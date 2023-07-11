package main

import (
	"fmt"
	"sort"
	"strings"
	"text/template"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
)

func genFile(gen *protogen.Plugin, file *protogen.File) error {
	if len(file.Services) == 0 {
		return nil
	}

	g := gen.NewGeneratedFile(file.GeneratedFilenamePrefix+".cobra.pb.go", file.GoImportPath)
	g.P("// Code generated by protoc-gen-cobra. DO NOT EDIT.")
	g.P()
	g.P("package ", file.GoPackageName)
	g.P()

	for _, srv := range file.Services {
		if err := genService(g, srv); err != nil {
			return err
		}
	}

	return nil
}

var (
	serviceTemplateCode = `
func {{.GoName}}ClientCommand(options ...client.Option) *cobra.Command {
	cfg := client.NewConfig(options...)
	cmd := &cobra.Command{
		Use: cfg.CommandNamer("{{.GoName}}"),
		Short: "{{.GoName}} service client",
		Long: {{.Comments.Leading | cleanComments | printf "%q"}},{{if .Desc.Options.GetDeprecated}}
		Deprecated: "deprecated",{{end}}
	}
	cfg.BindFlags(cmd.PersistentFlags())
	cmd.AddCommand({{range .Methods}}
		_{{$.GoName}}{{.GoName}}Command(cfg),{{end}}
	)
	return cmd
}
`
	serviceTemplate = template.Must(template.New("service").
			Funcs(template.FuncMap{"cleanComments": cleanComments}).
			Parse(serviceTemplateCode))
	serviceImports = []protogen.GoImportPath{
		"github.com/NathanBaulch/protoc-gen-cobra/client",
		"github.com/spf13/cobra",
	}
)

func genService(g *protogen.GeneratedFile, service *protogen.Service) error {
	for _, imp := range serviceImports {
		g.QualifiedGoIdent(protogen.GoIdent{GoImportPath: imp})
	}
	if err := serviceTemplate.Execute(g, service); err != nil {
		return err
	}

	for _, mth := range service.Methods {
		if err := genMethod(g, mth); err != nil {
			return err
		}
	}

	return nil
}

var (
	methodTemplateCode = `
func _{{.Parent.GoName}}{{.GoName}}Command(cfg *client.Config) *cobra.Command {
	req := &{{.InputType}}{}

	cmd := &cobra.Command{
		Use: cfg.CommandNamer("{{.GoName}}"),
		Short: "{{.GoName}} RPC client",
		Long: {{.Comments.Leading | cleanComments | printf "%q"}},{{if .Desc.Options.GetDeprecated}}
		Deprecated: "deprecated",{{end}}
		RunE: func(cmd *cobra.Command, args []string) error {
			if cfg.UseEnvVars {
				if err := flag.SetFlagsFromEnv(cmd.Parent().PersistentFlags(), true, cfg.EnvVarNamer, cfg.EnvVarPrefix, "{{.Parent.GoName}}"); err != nil {
					return err
				}
				if err := flag.SetFlagsFromEnv(cmd.PersistentFlags(), false, cfg.EnvVarNamer, cfg.EnvVarPrefix, "{{.Parent.GoName}}", "{{.GoName}}"); err != nil {
					return err
				}
			}
			return client.RoundTrip(cmd.Context(), cfg, func(cc grpc.ClientConnInterface, in iocodec.Decoder, out iocodec.Encoder) error {
				cli := New{{.Parent.GoName}}Client(cc)
				v := &{{.InputType}}{}
	{{if .Desc.IsStreamingClient}}
				stm, err := cli.{{.GoName}}(cmd.Context())
				if err != nil {
					return err
				}
				for {
					if err := in(v); err != nil {
						if err == io.EOF {
							_ = stm.CloseSend()
							break
						}
						return err
					}
					proto.Merge(v, req)
					if err = stm.Send(v); err != nil {
						return err
					}
				}
	{{else}}
				if err := in(v); err != nil {
					return err
				}
				proto.Merge(v, req)
		{{if .Desc.IsStreamingServer}}
				stm, err := cli.{{.GoName}}(cmd.Context(), v)
		{{else}}
				res, err := cli.{{.GoName}}(cmd.Context(), v)
		{{end}}
				if err != nil {
					return err
				}
	{{end}}
	{{if .Desc.IsStreamingServer}}
				for {
					res, err := stm.Recv()
					if err != nil {
						if err == io.EOF {
							break
						}
						return err
					}
					if err = out(res); err != nil {
						return err
					}
				}
				return nil
	{{else}}
		{{if .Desc.IsStreamingClient}}
				res, err := stm.CloseAndRecv()
				if err != nil {
					return err
				}
		{{end}}
				return out(res)
	{{end}}
			})
		},
	}

	{{.InputCode}}

	return cmd
}
`
	methodTemplate = template.Must(template.New("method").
			Funcs(template.FuncMap{"cleanComments": cleanComments}).
			Parse(methodTemplateCode))
	methodImports = []protogen.GoImportPath{
		"github.com/NathanBaulch/protoc-gen-cobra/client",
		"github.com/NathanBaulch/protoc-gen-cobra/flag",
		"github.com/NathanBaulch/protoc-gen-cobra/iocodec",
		"github.com/spf13/cobra",
		"google.golang.org/grpc",
		"google.golang.org/protobuf/proto",
	}
)

func genMethod(g *protogen.GeneratedFile, method *protogen.Method) error {
	for _, imp := range methodImports {
		g.QualifiedGoIdent(protogen.GoIdent{GoImportPath: imp})
	}
	if method.Desc.IsStreamingClient() || method.Desc.IsStreamingServer() {
		g.QualifiedGoIdent(protogen.GoIdent{GoImportPath: "io"})
	}

	code := walkMessage(g, method.Input, nil, nil, false, make(map[protogen.GoIdent]bool))
	data := struct {
		*protogen.Method
		InputType string
		InputCode string
	}{method, g.QualifiedGoIdent(method.Input.GoIdent), strings.Join(code, "\n")}
	return methodTemplate.Execute(g, data)
}

var (
	basicTypes = map[protoreflect.Kind]struct{ Type, Parse, Value, Slice, Pointer, Default string }{
		protoreflect.BoolKind:   {"bool", "ParseBoolE", "BoolVar", "BoolSliceVar", "BoolPointerVar", "false"},
		protoreflect.Int32Kind:  {"int32", "ParseInt32E", "Int32Var", "Int32SliceVar", "Int32PointerVar", "0"},
		protoreflect.Uint32Kind: {"uint32", "ParseUint32E", "Uint32Var", "Uint32SliceVar", "Uint32PointerVar", "0"},
		protoreflect.Int64Kind:  {"int64", "ParseInt64E", "Int64Var", "Int64SliceVar", "Int64PointerVar", "0"},
		protoreflect.Uint64Kind: {"uint64", "ParseUint64E", "Uint64Var", "Uint64SliceVar", "Uint64PointerVar", "0"},
		protoreflect.FloatKind:  {"float32", "ParseFloat32E", "Float32Var", "Float32SliceVar", "Float32PointerVar", "0"},
		protoreflect.DoubleKind: {"float64", "ParseFloat64E", "Float64Var", "Float64SliceVar", "Float64PointerVar", "0"},
		protoreflect.StringKind: {"string", "ParseStringE", "StringVar", "StringSliceVar", "StringPointerVar", `""`},
		protoreflect.BytesKind:  {"bytesBase64", "ParseBytesBase64E", "BytesBase64Var", "BytesBase64SliceVar", "", "nil"},
	}
	wrappersPkg  = protogen.GoImportPath("google.golang.org/protobuf/types/known/wrapperspb")
	timestampPkg = protogen.GoImportPath("google.golang.org/protobuf/types/known/timestamppb")
	durationPkg  = protogen.GoImportPath("google.golang.org/protobuf/types/known/durationpb")
	knownTypes   = map[protogen.GoIdent]struct{ Type, Parse, Value, Slice string }{
		timestampPkg.Ident("Timestamp"):  {"timestamp", "ParseTimestampE", "TimestampVar", "TimestampSliceVar"},
		durationPkg.Ident("Duration"):    {"duration", "ParseDurationE", "DurationVar", "DurationSliceVar"},
		wrappersPkg.Ident("DoubleValue"): {"float64", "ParseDoubleWrapperE", "DoubleWrapperVar", "DoubleWrapperSliceVar"},
		wrappersPkg.Ident("FloatValue"):  {"float32", "ParseFloatWrapperE", "FloatWrapperVar", "FloatWrapperSliceVar"},
		wrappersPkg.Ident("Int64Value"):  {"int64", "ParseInt64WrapperE", "Int64WrapperVar", "Int64WrapperSliceVar"},
		wrappersPkg.Ident("UInt64Value"): {"uint64", "ParseUInt64WrapperE", "UInt64WrapperVar", "UInt64WrapperSliceVar"},
		wrappersPkg.Ident("Int32Value"):  {"int32", "ParseInt32WrapperE", "Int32WrapperVar", "Int32WrapperSliceVar"},
		wrappersPkg.Ident("UInt32Value"): {"uint32", "ParseUInt32WrapperE", "UInt32WrapperVar", "UInt32WrapperSliceVar"},
		wrappersPkg.Ident("BoolValue"):   {"bool", "ParseBoolWrapperE", "BoolWrapperVar", "BoolWrapperSliceVar"},
		wrappersPkg.Ident("StringValue"): {"string", "ParseStringWrapperE", "StringWrapperVar", "StringWrapperSliceVar"},
		wrappersPkg.Ident("BytesValue"):  {"bytesBase64", "ParseBytesBase64WrapperE", "BytesBase64WrapperVar", "BytesBase64WrapperSliceVar"},
	}
)

func walkMessage(g *protogen.GeneratedFile, msg *protogen.Message, path, postSetCode []string, deprecated bool, visited map[protogen.GoIdent]bool) []string {
	var code []string
	sort.Slice(msg.Fields, func(i, j int) bool { return msg.Fields[i].Desc.Index() < msg.Fields[j].Desc.Index() })
	for _, fld := range msg.Fields {
		code = append(code, walkField(g, fld, path, postSetCode, deprecated, visited)...)
	}
	return code
}

func walkField(g *protogen.GeneratedFile, fld *protogen.Field, path, postSetCode []string, deprecated bool, visited map[protogen.GoIdent]bool) []string {
	target := "req"
	if len(path) > 0 {
		target = "_" + strings.Join(path, "_")
	}
	path = append(path, fld.GoName)
	flagName := fmt.Sprintf("cfg.FlagNamer(%q)", strings.Join(path, " "))
	varName := "_" + strings.Join(path, "_")
	oneof := fld.Oneof != nil && !fld.Oneof.Desc.IsSynthetic()

	if f := flagFormat(g, fld); f != "" {
		var code []string
		if oneof {
			postSetCode = append(postSetCode, fmt.Sprintf("%s.%s = %s", target, fld.Oneof.GoName, varName))
			target = varName
			code = append(code, fmt.Sprintf("%s := &%s{}", target, g.QualifiedGoIdent(fld.GoIdent)))
		}
		ptr := fmt.Sprintf("&%s.%s", target, fld.GoName)
		code = append(code, fmt.Sprintf(f, ptr, flagName, cleanComments(fld.Comments.Leading)))
		if len(postSetCode) > 0 {
			code = append(code, fmt.Sprintf("flag.WithPostSetHook(cmd.PersistentFlags(), %s, func() { %s })", flagName, strings.Join(postSetCode, ";")))
		}
		if deprecated || fld.Desc.Options().(*descriptorpb.FieldOptions).GetDeprecated() {
			code = append(code, fmt.Sprintf("_ = cmd.PersistentFlags().MarkDeprecated(%s, \"deprecated\")", flagName))
		}
		return code
	} else if normalizeKind(fld.Desc.Kind()) == protoreflect.MessageKind {
		if fld.Desc.IsList() {
			// message list not supported
		} else if fld.Desc.IsMap() {
			// limited map support
		} else if visited[fld.Parent.GoIdent] = true; visited[fld.Message.GoIdent] {
			// cycle detected
		} else {
			if oneof {
				postSetCode = append(postSetCode, fmt.Sprintf("%s.%s = &%s{%s: %s}", target, fld.Oneof.GoName, g.QualifiedGoIdent(fld.GoIdent), fld.GoName, varName))
			} else {
				postSetCode = append(postSetCode, fmt.Sprintf("%s.%s = %s", target, fld.GoName, varName))
			}
			subVisited := make(map[protogen.GoIdent]bool, len(visited))
			for k, v := range visited {
				subVisited[k] = v
			}
			if subCode := walkMessage(g, fld.Message, path, postSetCode, deprecated, subVisited); len(subCode) > 0 {
				code := []string{fmt.Sprintf("%s := &%s{}", varName, g.QualifiedGoIdent(fld.Message.GoIdent))}
				if oneof {
					code = append(code, fmt.Sprintf("cmd.PersistentFlags().Bool(%s, false, \"\")", flagName))
					code = append(code, fmt.Sprintf("flag.WithPostSetHook(cmd.PersistentFlags(), %s, func() { %s })", flagName, strings.Join(postSetCode, ";")))
				}
				return append(code, subCode...)
			}
		}
	}

	return nil
}

func flagFormat(g *protogen.GeneratedFile, fld *protogen.Field) string {
	k := normalizeKind(fld.Desc.Kind())

	if bt, ok := basicTypes[k]; ok {
		if fld.Desc.IsList() {
			switch k {
			case protoreflect.Uint32Kind, protoreflect.Uint64Kind, protoreflect.BytesKind:
				return fmt.Sprintf("flag.%s(cmd.PersistentFlags(), %%s, %%s, %%q)", bt.Slice)
			default:
				return fmt.Sprintf("cmd.PersistentFlags().%s(%%s, %%s, nil, %%q)", bt.Slice)
			}
		} else if k == protoreflect.BytesKind {
			return fmt.Sprintf("flag.%s(cmd.PersistentFlags(), %%s, %%s, %%q)", bt.Value)
		} else if (fld.Oneof == nil && fld.Desc.HasPresence()) || (fld.Oneof != nil && fld.Oneof.Desc.IsSynthetic()) {
			return fmt.Sprintf("flag.%s(cmd.PersistentFlags(), %%s, %%s, %%q)", bt.Pointer)
		} else {
			return fmt.Sprintf("cmd.PersistentFlags().%s(%%s, %%s, %s, %%q)", bt.Value, bt.Default)
		}
	}

	switch k {
	case protoreflect.EnumKind:
		if fld.Desc.IsList() {
			return "flag.EnumSliceVar(cmd.PersistentFlags(), %s, %s, %q)"
		} else if fld.Desc.HasPresence() {
			return "flag.EnumPointerVar(cmd.PersistentFlags(), %s, %s, %q)"
		} else {
			return "flag.EnumVar(cmd.PersistentFlags(), %s, %s, %q)"
		}
	case protoreflect.MessageKind:
		if kt, ok := knownTypes[fld.Message.GoIdent]; ok {
			if fld.Desc.IsList() {
				return fmt.Sprintf("flag.%s(cmd.PersistentFlags(), %%s, %%s, %%q)", kt.Slice)
			} else {
				return fmt.Sprintf("flag.%s(cmd.PersistentFlags(), %%s, %%s, %%q)", kt.Value)
			}
		} else if fld.Desc.IsList() {
			return fmt.Sprintf("flag.SliceVar(cmd.PersistentFlags(), flag.ParseMessageE[*%s], %%s, %%s, %%q)", g.QualifiedGoIdent(fld.Message.GoIdent))
		} else if fld.Desc.IsMap() {
			kk := normalizeKind(fld.Desc.MapKey().Kind())
			vk := normalizeKind(fld.Desc.MapValue().Kind())
			if kk == protoreflect.StringKind {
				switch vk {
				case protoreflect.StringKind:
					return "cmd.PersistentFlags().StringToStringVar(%s, %s, nil, %q)"
				case protoreflect.Int64Kind:
					return "cmd.PersistentFlags().StringToInt64Var(%s, %s, nil, %q)"
				}
			}

			if bt, ok := basicTypes[kk]; ok {
				keyParser := "flag." + bt.Parse
				valParser := ""
				switch vk {
				case protoreflect.EnumKind:
					valParser = fmt.Sprintf("flag.ParseEnumE[%s]", g.QualifiedGoIdent(fld.Message.Fields[1].Enum.GoIdent))
				case protoreflect.MessageKind:
					if kt, ok := knownTypes[fld.Message.Fields[1].Message.GoIdent]; ok {
						valParser = "flag." + kt.Parse
					} else {
						valParser = fmt.Sprintf("flag.ParseMessageE[*%s]", g.QualifiedGoIdent(fld.Message.Fields[1].Message.GoIdent))
					}
				default:
					if bt, ok := basicTypes[vk]; ok {
						valParser = "flag." + bt.Parse
					}
				}
				if valParser != "" {
					return fmt.Sprintf("flag.MapVar(cmd.PersistentFlags(), %s, %s, %%s, %%s, %%q)", keyParser, valParser)
				}
			}
		}
	}

	return ""
}

func normalizeKind(kind protoreflect.Kind) protoreflect.Kind {
	switch kind {
	case protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
		return protoreflect.Int32Kind
	case protoreflect.Fixed32Kind:
		return protoreflect.Uint32Kind
	case protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
		return protoreflect.Int64Kind
	case protoreflect.Fixed64Kind:
		return protoreflect.Uint64Kind
	case protoreflect.GroupKind:
		return protoreflect.MessageKind
	default:
		return kind
	}
}

func cleanComments(comments protogen.Comments) string {
	return strings.TrimSpace(string(comments))
}

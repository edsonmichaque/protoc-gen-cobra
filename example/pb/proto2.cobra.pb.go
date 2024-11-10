// Code generated by protoc-gen-cobra. DO NOT EDIT.

package pb

import (
	client "github.com/edsonmichaque/protoc-gen-cobra/client"
	flag "github.com/edsonmichaque/protoc-gen-cobra/flag"
	iocodec "github.com/edsonmichaque/protoc-gen-cobra/iocodec"
	cobra "github.com/spf13/cobra"
	grpc "google.golang.org/grpc"
	proto "google.golang.org/protobuf/proto"
)

func Proto2ClientCommand(options ...client.Option) *cobra.Command {
	cfg := client.NewConfig(options...)
	cmd := &cobra.Command{
		Use:   cfg.CommandNamer("Proto2"),
		Short: "Proto2 service client",
		Long:  "",
	}
	cfg.BindFlags(cmd.PersistentFlags())
	cmd.AddCommand(
		_Proto2EchoCommand(cfg),
	)
	return cmd
}

func _Proto2EchoCommand(cfg *client.Config) *cobra.Command {
	req := &Sound2{}

	cmd := &cobra.Command{
		Use:   cfg.CommandNamer("Echo"),
		Short: "Echo RPC client",
		Long:  "",
		RunE: func(cmd *cobra.Command, args []string) error {
			if cfg.UseEnvVars {
				if err := flag.SetFlagsFromEnv(cmd.Parent().PersistentFlags(), true, cfg.EnvVarNamer, cfg.EnvVarPrefix, "Proto2"); err != nil {
					return err
				}
				if err := flag.SetFlagsFromEnv(cmd.PersistentFlags(), false, cfg.EnvVarNamer, cfg.EnvVarPrefix, "Proto2", "Echo"); err != nil {
					return err
				}
			}
			return client.RoundTrip(cmd.Context(), cfg, func(cc grpc.ClientConnInterface, in iocodec.Decoder, out iocodec.Encoder) error {
				cli := NewProto2Client(cc)
				v := &Sound2{}

				if err := in(v); err != nil {
					return err
				}
				proto.Merge(v, req)

				res, err := cli.Echo(cmd.Context(), v)

				if err != nil {
					return err
				}

				return out(res)

			})
		},
	}

	flag.Float64PointerVar(cmd.PersistentFlags(), &req.Double, cfg.FlagNamer("Double"), "")
	flag.Float32PointerVar(cmd.PersistentFlags(), &req.Float, cfg.FlagNamer("Float"), "")
	flag.Int32PointerVar(cmd.PersistentFlags(), &req.Int32, cfg.FlagNamer("Int32"), "")
	flag.Int64PointerVar(cmd.PersistentFlags(), &req.Int64, cfg.FlagNamer("Int64"), "")
	flag.Uint32PointerVar(cmd.PersistentFlags(), &req.Uint32, cfg.FlagNamer("Uint32"), "")
	flag.Uint64PointerVar(cmd.PersistentFlags(), &req.Uint64, cfg.FlagNamer("Uint64"), "")
	flag.Int32PointerVar(cmd.PersistentFlags(), &req.Sint32, cfg.FlagNamer("Sint32"), "")
	flag.Int64PointerVar(cmd.PersistentFlags(), &req.Sint64, cfg.FlagNamer("Sint64"), "")
	flag.Uint32PointerVar(cmd.PersistentFlags(), &req.Fixed32, cfg.FlagNamer("Fixed32"), "")
	flag.Uint64PointerVar(cmd.PersistentFlags(), &req.Fixed64, cfg.FlagNamer("Fixed64"), "")
	flag.Int32PointerVar(cmd.PersistentFlags(), &req.Sfixed32, cfg.FlagNamer("Sfixed32"), "")
	flag.Int64PointerVar(cmd.PersistentFlags(), &req.Sfixed64, cfg.FlagNamer("Sfixed64"), "")
	flag.BoolPointerVar(cmd.PersistentFlags(), &req.Bool, cfg.FlagNamer("Bool"), "")
	flag.StringPointerVar(cmd.PersistentFlags(), &req.String_, cfg.FlagNamer("String_"), "")
	flag.BytesBase64Var(cmd.PersistentFlags(), &req.Bytes, cfg.FlagNamer("Bytes"), "")
	cmd.PersistentFlags().Float64SliceVar(&req.ListDouble, cfg.FlagNamer("ListDouble"), nil, "")
	cmd.PersistentFlags().Float32SliceVar(&req.ListFloat, cfg.FlagNamer("ListFloat"), nil, "")
	cmd.PersistentFlags().Int32SliceVar(&req.ListInt32, cfg.FlagNamer("ListInt32"), nil, "")
	cmd.PersistentFlags().Int64SliceVar(&req.ListInt64, cfg.FlagNamer("ListInt64"), nil, "")
	flag.Uint32SliceVar(cmd.PersistentFlags(), &req.ListUint32, cfg.FlagNamer("ListUint32"), "")
	flag.Uint64SliceVar(cmd.PersistentFlags(), &req.ListUint64, cfg.FlagNamer("ListUint64"), "")
	cmd.PersistentFlags().Int32SliceVar(&req.ListSint32, cfg.FlagNamer("ListSint32"), nil, "")
	cmd.PersistentFlags().Int64SliceVar(&req.ListSint64, cfg.FlagNamer("ListSint64"), nil, "")
	flag.Uint32SliceVar(cmd.PersistentFlags(), &req.ListFixed32, cfg.FlagNamer("ListFixed32"), "")
	flag.Uint64SliceVar(cmd.PersistentFlags(), &req.ListFixed64, cfg.FlagNamer("ListFixed64"), "")
	cmd.PersistentFlags().Int32SliceVar(&req.ListSfixed32, cfg.FlagNamer("ListSfixed32"), nil, "")
	cmd.PersistentFlags().Int64SliceVar(&req.ListSfixed64, cfg.FlagNamer("ListSfixed64"), nil, "")
	cmd.PersistentFlags().BoolSliceVar(&req.ListBool, cfg.FlagNamer("ListBool"), nil, "")
	cmd.PersistentFlags().StringSliceVar(&req.ListString, cfg.FlagNamer("ListString"), nil, "")
	flag.BytesBase64SliceVar(cmd.PersistentFlags(), &req.ListBytes, cfg.FlagNamer("ListBytes"), "")
	cmd.PersistentFlags().StringToInt64Var(&req.MapStringInt64, cfg.FlagNamer("MapStringInt64"), nil, "")
	cmd.PersistentFlags().StringToStringVar(&req.MapStringString, cfg.FlagNamer("MapStringString"), nil, "")
	flag.EnumPointerVar(cmd.PersistentFlags(), &req.Enum, cfg.FlagNamer("Enum"), "")
	flag.EnumSliceVar(cmd.PersistentFlags(), &req.ListEnum, cfg.FlagNamer("ListEnum"), "")
	flag.TimestampVar(cmd.PersistentFlags(), &req.Timestamp, cfg.FlagNamer("Timestamp"), "")
	flag.DurationVar(cmd.PersistentFlags(), &req.Duration, cfg.FlagNamer("Duration"), "")
	flag.BoolWrapperVar(cmd.PersistentFlags(), &req.WrapperBool, cfg.FlagNamer("WrapperBool"), "")
	flag.BytesBase64WrapperVar(cmd.PersistentFlags(), &req.WrapperBytes, cfg.FlagNamer("WrapperBytes"), "")
	flag.DoubleWrapperVar(cmd.PersistentFlags(), &req.WrapperDouble, cfg.FlagNamer("WrapperDouble"), "")
	flag.FloatWrapperVar(cmd.PersistentFlags(), &req.WrapperFloat, cfg.FlagNamer("WrapperFloat"), "")
	flag.Int32WrapperVar(cmd.PersistentFlags(), &req.WrapperInt32, cfg.FlagNamer("WrapperInt32"), "")
	flag.Int64WrapperVar(cmd.PersistentFlags(), &req.WrapperInt64, cfg.FlagNamer("WrapperInt64"), "")
	flag.StringWrapperVar(cmd.PersistentFlags(), &req.WrapperString, cfg.FlagNamer("WrapperString"), "")
	flag.UInt32WrapperVar(cmd.PersistentFlags(), &req.WrapperUint32, cfg.FlagNamer("WrapperUint32"), "")
	flag.UInt64WrapperVar(cmd.PersistentFlags(), &req.WrapperUint64, cfg.FlagNamer("WrapperUint64"), "")
	flag.TimestampSliceVar(cmd.PersistentFlags(), &req.ListTimestamp, cfg.FlagNamer("ListTimestamp"), "")
	flag.DurationSliceVar(cmd.PersistentFlags(), &req.ListDuration, cfg.FlagNamer("ListDuration"), "")
	flag.BoolWrapperSliceVar(cmd.PersistentFlags(), &req.ListWrapperBool, cfg.FlagNamer("ListWrapperBool"), "")
	flag.BytesBase64WrapperSliceVar(cmd.PersistentFlags(), &req.ListWrapperBytes, cfg.FlagNamer("ListWrapperBytes"), "")
	flag.DoubleWrapperSliceVar(cmd.PersistentFlags(), &req.ListWrapperDouble, cfg.FlagNamer("ListWrapperDouble"), "")
	flag.FloatWrapperSliceVar(cmd.PersistentFlags(), &req.ListWrapperFloat, cfg.FlagNamer("ListWrapperFloat"), "")
	flag.Int32WrapperSliceVar(cmd.PersistentFlags(), &req.ListWrapperInt32, cfg.FlagNamer("ListWrapperInt32"), "")
	flag.Int64WrapperSliceVar(cmd.PersistentFlags(), &req.ListWrapperInt64, cfg.FlagNamer("ListWrapperInt64"), "")
	flag.StringWrapperSliceVar(cmd.PersistentFlags(), &req.ListWrapperString, cfg.FlagNamer("ListWrapperString"), "")
	flag.UInt32WrapperSliceVar(cmd.PersistentFlags(), &req.ListWrapperUint32, cfg.FlagNamer("ListWrapperUint32"), "")
	flag.UInt64WrapperSliceVar(cmd.PersistentFlags(), &req.ListWrapperUint64, cfg.FlagNamer("ListWrapperUint64"), "")

	return cmd
}

package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"

	"github.com/pulumi/pulumi-terraform/sdk/v5/go/state"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {

		conf := config.New(ctx, "")
		bucket := conf.Require("bucket")
		prefix := conf.Require("prefix")
		encryptionKey := conf.Require("encryptionKey")

		state, err := state.NewRemoteStateReference(ctx, "gcsstate", &state.GcsStateArgs{
			Bucket: pulumi.String(bucket),
			Prefix: pulumi.Sprintf("%s/terraform.tfstate", prefix),
			EncryptionKey: pulumi.String(encryptionKey),
		})
		if err != nil {
			return err
		}

		ctx.Export("test", state.Outputs)

		return nil
	})
}

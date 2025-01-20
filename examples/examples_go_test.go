// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.
//go:build go || all
// +build go all

package examples

import (
	"path"
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
)

func TestGoLocal013(t *testing.T) {
	test := getGoBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: path.Join(getCwd(t), "localstate-go"),
			Config: map[string]string{
				"statefile": "terraform.0-13-0.tfstate",
			},
		})

	integration.ProgramTest(t, &test)
}

func TestGoGcsPlainText(t *testing.T) {
	test := getGoBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: path.Join(getCwd(t), "gcsstate-go"),
			Config: map[string]string{
				"bucket": "terraform-to-g0wn0",
				"prefix": "pulumi-terraform-provider/plain-text",
			},
		})

	integration.ProgramTest(t, &test)
}

func TestGoGcsEncrypted(t *testing.T) {
	test := getGoBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: path.Join(getCwd(t), "gcsstate-go"),
			Config: map[string]string{
				"bucket": "terraform-to-g0wn0",
				"encryptionKey": "1hh7l8LeVu4fPjEk8BfpkGDX+N/ZmLsQRvFT+uSMAJE=",
				"prefix": "pulumi-terraform-provider/encrypted",
			},
		})

	integration.ProgramTest(t, &test)
}


// func TestGoLocal012(t *testing.T) {
// 	test := getGoBaseOptions(t).
// 		With(integration.ProgramTestOptions{
// 			Dir: path.Join(getCwd(t), "localstate-go"),
// 			Config: map[string]string{
// 				"statefile": "terraform.0-12-24.tfstate",
// 			},
// 		})

// 	integration.ProgramTest(t, &test)
// }

// func TestGoS3013(t *testing.T) {
// 	test := getGoBaseOptions(t).
// 		With(integration.ProgramTestOptions{
// 			Dir: path.Join(getCwd(t), "s3state-go"),
// 			Config: map[string]string{
// 				"bucketName": "pulumi-terraform-remote-state-testing",
// 				"key":        "0-11-state",
// 				"region":     "us-west-2",
// 			},
// 		})

// 	integration.ProgramTest(t, &test)
// }

// func TestGoS3012(t *testing.T) {
// 	test := getGoBaseOptions(t).
// 		With(integration.ProgramTestOptions{
// 			Dir: path.Join(getCwd(t), "s3state-go"),
// 			Config: map[string]string{
// 				"bucketName": "pulumi-terraform-remote-state-testing",
// 				"key":        "0-12-state",
// 				"region":     "us-west-2",
// 			},
// 		})

// 	integration.ProgramTest(t, &test)
// }

// func TestGoRemoteBackend(t *testing.T) {
// 	t.Skip("TODO: https://github.com/pulumi/pulumi-terraform/issues/730")
// 	test := getGoBaseOptions(t).
// 		With(integration.ProgramTestOptions{
// 			Dir: path.Join(getCwd(t), "remote-backend-go"),
// 			Config: map[string]string{
// 				"organization":  getRemoteBackendOrganization(t),
// 				"workspaceName": "dev",
// 			},
// 			Secrets: map[string]string{
// 				"tfeToken": getRemoteBackendToken(t),
// 			},
// 		})

// 	integration.ProgramTest(t, &test)
// }

func getGoBaseOptions(t *testing.T) integration.ProgramTestOptions {
	base := getBaseOptions()
	baseGo := base.With(integration.ProgramTestOptions{
		RunUpdateTest: false,
		Dependencies: []string{
			"github.com/kmochejski/pulumi-terraform/sdk/v3/go",
		},
	})

	return baseGo
}

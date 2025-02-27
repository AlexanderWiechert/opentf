// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package main

import (
	localexec "github.com/placeholderplaceholderplaceholder/opentf/internal/builtin/provisioners/local-exec"
	"github.com/placeholderplaceholderplaceholder/opentf/internal/grpcwrap"
	"github.com/placeholderplaceholderplaceholder/opentf/internal/plugin"
	"github.com/placeholderplaceholderplaceholder/opentf/internal/tfplugin5"
)

func main() {
	// Provide a binary version of the internal terraform provider for testing
	plugin.Serve(&plugin.ServeOpts{
		GRPCProvisionerFunc: func() tfplugin5.ProvisionerServer {
			return grpcwrap.Provisioner(localexec.New())
		},
	})
}

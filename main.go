package main

import (
	"context"

	"github.com/cappyzawa/terraform-provider-concourse/concourse"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	tf5server "github.com/hashicorp/terraform-plugin-go/tfprotov5/server"
	tfmux "github.com/hashicorp/terraform-plugin-mux"
)

// ProviderPATH describes path of this provider in registry
const ProviderPATH = "registry.terraform.io/cappyzawa/concourse"

func main() {
	ctx := context.Background()
	muxed, err := tfmux.NewSchemaServerFactory(ctx, concourse.Provider().GRPCProvider)
	if err != nil {
		panic(err)
	}

	if err := tf5server.Serve(ProviderPATH, func() tfprotov5.ProviderServer {
		return muxed.Server()
	}); err != nil {
		panic(err)
	}
}

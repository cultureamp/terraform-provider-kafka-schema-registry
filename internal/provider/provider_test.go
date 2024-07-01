package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	tfprotov6 "github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

const (
	testAccProviderVersion = "test-version"
	testAccProviderType    = "schemaregistry"
)

var testAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
	"schemaregistry": providerserver.NewProtocol6WithError(New(testAccProviderVersion)()),
}

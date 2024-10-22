package main

import (
    "github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
    "github.com/AvinBajelani/terraform-provider-dns-soeldnerconsult/soeldnerconsult"
)

func main() {
    plugin.Serve(&plugin.ServeOpts{
        ProviderFunc: soeldnerconsult.Provider,
    })
}

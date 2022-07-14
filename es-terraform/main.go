package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() terraform.ResourceProvider {
			return Provider()
		},
	})

	// user := map[string]string{}
	// user["name"] = "employee2"
	// buff, _ := json.Marshal(user)
	// err := es_ops.Create(2, "employee", string(buff))
	// if err != nil {
	// 	panic(err)
	// }
}

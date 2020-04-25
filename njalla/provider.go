package njalla

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"token": {
				Type:	schema.TypeString,
				Required: true,
				DefaultFunc: schema.EnvDefaultFunc("NJALLA_TOKEN", ""),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"njalla_domain_record": resourceDomainRecord(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	return NjallaClient{
		d.Get("token").(string),
		"https://njal.la/api/1/",
	}, nil
}

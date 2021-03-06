package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataWLANGroup() *schema.Resource {
	return &schema.Resource{
		Description: `
unifi_wlan_group data source can be used to retrieve the ID for a WLAN group by name.
`,

		Read: dataWLANGroupRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Description: "The name of the WLAN group to look up.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "Default",
			},
		},
	}
}

func dataWLANGroupRead(d *schema.ResourceData, meta interface{}) error {
	c := meta.(*client)

	name := d.Get("name").(string)

	groups, err := c.c.ListWLANGroup(context.TODO(), c.site)
	if err != nil {
		return err
	}
	for _, g := range groups {
		if g.Name == name {
			d.SetId(g.ID)
			return nil
		}
	}

	return fmt.Errorf("WLAN group not found with name %s", name)
}

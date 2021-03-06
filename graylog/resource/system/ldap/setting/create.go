package setting

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func create(d *schema.ResourceData, m interface{}) error {
	if err := update(d, m); err != nil {
		return err
	}
	d.SetId(ldapSettingID)
	return nil
}

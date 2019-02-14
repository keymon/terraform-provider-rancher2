package rancher2

import (
	"github.com/hashicorp/terraform/helper/schema"
	managementClient "github.com/rancher/types/client/management/v3"
)

//Schemas

func privateRegistriesFields() map[string]*schema.Schema {
	s := map[string]*schema.Schema{
		"is_default": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"password": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"url": {
			Type:     schema.TypeString,
			Required: true,
		},
		"user": {
			Type:     schema.TypeString,
			Optional: true,
		},
	}
	return s
}

// Flatteners

func flattenPrivateRegistries(p []managementClient.PrivateRegistry) ([]interface{}, error) {
	out := []interface{}{}

	for _, in := range p {
		obj := make(map[string]interface{})
		obj["is_default"] = in.IsDefault

		if len(in.Password) > 0 {
			obj["password"] = in.Password
		}

		if len(in.URL) > 0 {
			obj["url"] = in.URL
		}

		if len(in.User) > 0 {
			obj["user"] = in.User
		}

		out = append(out, obj)
	}

	return out, nil
}

// Expanders

func expandPrivateRegistries(p []interface{}) ([]managementClient.PrivateRegistry, error) {
	out := []managementClient.PrivateRegistry{}
	if len(p) == 0 || p[0] == nil {
		return out, nil
	}

	for i := range p {
		in := p[i].(map[string]interface{})
		obj := managementClient.PrivateRegistry{}

		if v, ok := in["is_default"].(bool); ok {
			obj.IsDefault = v
		}

		if v, ok := in["password"].(string); ok && len(v) > 0 {
			obj.Password = v
		}

		if v, ok := in["url"].(string); ok && len(v) > 0 {
			obj.URL = v
		}

		if v, ok := in["user"].(string); ok && len(v) > 0 {
			obj.User = v
		}
		out = append(out, obj)
	}

	return out, nil
}

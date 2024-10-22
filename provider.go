package soeldnerconsult

import (
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider returns a schema.Provider.
func Provider() *schema.Provider {
    return &schema.Provider{
        ResourcesMap: map[string]*schema.Resource{
            "soeldnerconsult_dns_record": resourceDnsRecord(),
        },
    }
}

// resourceDnsRecord نمونه‌ای از تعریف منبع DNS است.
func resourceDnsRecord() *schema.Resource {
    return &schema.Resource{
        // ویژگی‌ها و متدهای منبع DNS را اینجا تعریف کنید
    }
}

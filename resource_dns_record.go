package soeldnerconsult

import (
    "context"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
)

func resourceDNSRecord() *schema.Resource {
    return &schema.Resource{
        CreateContext: resourceDNSRecordCreate,
        ReadContext:   resourceDNSRecordRead,
        UpdateContext: resourceDNSRecordUpdate,
        DeleteContext: resourceDNSRecordDelete,

        Schema: map[string]*schema.Schema{
            "name": {
                Type:     schema.TypeString,
                Required: true,
            },
            "type": {
                Type:     schema.TypeString,
                Required: true,
            },
            "value": {
                Type:     schema.TypeString,
                Required: true,
            },
            "ttl": {
                Type:     schema.TypeInt,
                Optional: true,
                Default: 3600,
            },
        },
    }
}

func resourceDNSRecordCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
    // منطق ایجاد رکورد DNS
    name := d.Get("name").(string)
    recordType := d.Get("type").(string)
    value := d.Get("value").(string)
    ttl := d.Get("ttl").(int)

    // فرض کنید اینجا یک API را برای ایجاد رکورد تماس می‌زنیم
    // ...

    // قرار دادن شناسه رکورد در داده‌های منابع
    d.SetId(name + "_" + recordType)

    return resourceDNSRecordRead(ctx, d, m)
}

func resourceDNSRecordRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
    // منطق خواندن رکورد DNS
    return diag.Diagnostics{}
}

func resourceDNSRecordUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
    // منطق به‌روزرسانی رکورد DNS
    return diag.Diagnostics{}
}

func resourceDNSRecordDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
    // منطق حذف رکورد DNS
    return diag.Diagnostics{}
}

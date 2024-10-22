terraform {
  required_providers {
    soeldnerconsult = {
      source  = "avinbajelani/dns-soeldnerconsult"  # آدرس مناسب
      version = "0.1.0"  # نسخه مناسب را در اینجا قرار دهید
    }
  }
}

provider "soeldnerconsult" {
  # تنظیمات مربوط به provider
}

resource "soeldnerconsult_dns_record" "example_a_record" {
  name  = "sclabs.com"
  type  = "A"
  value = "110.137.37.178"
  ttl   = 3600
}

resource "soeldnerconsult_dns_record" "example_cname_record" {
  type  = "CNAME"
  name  = "www.sclabs.com"
  value = "sclabs.com"
  ttl   = 3600
}


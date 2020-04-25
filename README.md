# Terraform Njalla Provider

## Provider

Terraform provider for [njal.la](https://njal.la) domain records using the njalla API.

The API gives access to a lot more resources, so far just domain records is implemented

### Example Usage
```hcl-terraform
provider "njalla" { 
  token = "de65d2cec2b202c9a37089ad4ac9b81e"
}
```

### Argument Reference

* `token` - (Required) your token in njalla or via env variable `NJALLA_TOKEN`, can be generated [here](https://njal.la/settings/api/)

## domain_record

This resource can create/edit your domain records

### Example Usage

```hcl-terraform
resource "njalla_domain_record" "this" {
  domain = "example.com"
  name = "subdomain"
  content = "127.0.0.1"
  type = "A"
  ttl = 10800
}
```

### Argument Reference

* `domain` (Required) the domain you want to add a record to
* `name` (Required) name of the domain record
* `content` (Required) content of the record
* `type` (Optional, default `A`) the type of the record
* `ttl` (Optional, default `10800`) the Time To Live for the record 

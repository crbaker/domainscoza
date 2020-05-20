resource "domainscoza_record" "navigator_dns" {
  key      = "123456789"
  tld      = "wethinkcode"
  sld      = "co.za"
  content  = "SAM"
  name     = "navigator"
  ttl      = 11
  priority = 31
  type     = "A"
}

resource "domainscoza_record" "zipkin_dns" {
  depends_on = [domainscoza_record.navigator_dns]
  key      = "123456789"
  tld      = "wethinkcode"
  sld      = "co.za"
  content  = "BOB"
  name     = "zipkin"
  ttl      = 666
  priority = 777
  type     = "A"
}

resource "domainscoza_record" "gitlab_dns" {
  depends_on = [domainscoza_record.zipkin_dns]
  key      = "123456789"
  tld      = "wethinkcode"
  sld      = "co.za"
  content  = "MIKE"
  name     = "somewhere"
  ttl      = 33300
  priority = 10
  type     = "A"
}
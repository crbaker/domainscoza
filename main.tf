resource "domainscoza_record" "some_record" {
  points_to = "12345"
  host = "some-other-host"
  ttl = 300
  priority = 100
  type = "A"
}

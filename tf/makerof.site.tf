resource "digitalocean_domain" "default" {
  name = "makerof.site"
  ip_address = "${digitalocean_droplet.haproxy-www.ipv4_address}"
}

resource "digitalocean_record" "CNAME-www" {
  domain = "${digitalocean_domain.default.name}"
  type = "CNAME"
  name = "www"
  value = "@"
}

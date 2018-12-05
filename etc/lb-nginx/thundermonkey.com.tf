resource "digitalocean_domain" "thundermonkey" {
  name = "thundermonkey.com"
  ip_address = "${digitalocean_droplet.haproxy-www.ipv4_address}"
}

resource "digitalocean_record" "tm-www" {
  domain = "${digitalocean_domain.thundermonkey.name}"
  type = "CNAME"
  name = "www"
  value = "@"
}

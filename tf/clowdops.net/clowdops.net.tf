resource "digitalocean_domain" "clowdops" {
  name = "clowdops.net"
  ip_address = "${digitalocean_droplet.dev.ipv4_address}"
}

resource "digitalocean_record" "clowdops-www" {
  domain = "${digitalocean_domain.clowdops.name}"
  type = "CNAME"
  name = "www"
  value = "@"
}

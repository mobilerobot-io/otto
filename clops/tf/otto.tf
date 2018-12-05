resource "digitalocean_droplet" "otto" {
  name = "otto"
  image = "ubuntu-18-04-x64"
  region = "sfo2"
  size = "s-1vcpu-1gb"

  private_networking = true
  ssh_keys = [
    "${var.ssh_fingerprint}"
  ]

  connection {
    user = "root"
    type = "ssh"
    private_key = "${file(var.pvt_key)}"
    timeout = "2m"
  }

  provisioner "local-exec" {
    command = "echo ${digitalocean_droplet.otto.ipv4_address} > /srv/inventory/clops/hosts-otto"
  }

  provisioner "remote-exec" {
    inline = [
      "export PATH=$PATH:/usr/bin",
      "sudo apt-get update",
      "sudo apt-get -y install python"
    ]
  }
}

output "oTToIP" {
  value = "${digitalocean_droplet.otto.ipv4_address}"
}

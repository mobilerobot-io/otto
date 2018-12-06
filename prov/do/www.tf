resource "digitalocean_droplet" "w01" {
  name = "w01"
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

  ## TODO - put
  provisioner "local-exec" {
    command = "echo ${digitalocean_droplet.w01.ipv4_address} >>  ${var.hostsfile}"
  }

  provisioner "remote-exec" {
    inline = [
      "export PATH=$PATH:/usr/bin",
      "sudo apt-get update",
      "sudo apt-get -y install python"
      "sudo adduser --system --shell /bin/bash --ingroup sudoers"
    ]
  }
}

output "w01" {
  value = "${digitalocean_droplet.w01.ipv4_address}"
}

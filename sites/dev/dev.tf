resource "digitalocean_droplet" "dev" {
  name = "dev"
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

  provisioner "remote-exec" {
    inline = [
      "export PATH=$PATH:/usr/bin",
      "sudo apt-get update",
      "sudo apt-get -y install python"
    ]
  }
}

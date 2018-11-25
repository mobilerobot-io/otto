resource "digitalocean_droplet" "www-2" {
  image = "ubuntu-14-04-x64"
  name = "www-2"
  region = "sfo2"
  size = "512mb"
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
      "sudo apt-get -y install nginx"
    ]
  }
}

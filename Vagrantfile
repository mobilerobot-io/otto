# -*- mode: ruby -*-
# vi: set ft=ruby :

# docs: https://docs.vagrantup.com
# search for boxes at: https://vagrantcloud.com/search

Vagrant.configure("2") do |config|

  # go with ub18 server for now
  config.vm.box = "bento/ubuntu-18.04"
  config.ssh.insert_key = false

  # our local name is loca
  config.vm.hostname = "loca.local"
  config.vm.post_up_message = "Run 'vagrant ssh' and do what it says "
  config.vm.network "public_network"
  # config.vm.network "forwarded_port", guest: 80, host: 1001


  # Make sure the local repo is there
  # config.vm.synced_folder "config", "/srv/config"

  # virtualbox is the "provider"
  config.vm.provider "virtualbox" do |vb|
    vb.memory = "1024"  # make this smaller for production
    vb.linked_clone = true
  end

  # Ansible
  config.vm.provision "ansible" do |ansible|
    ansible.playbook = "config/vagrant.yml"
  end

  # otto is our application
  config.vm.define "o01" do |app|
    app.vm.hostname = "o01.local"
    app.vm.network :private_network, ip: "10.24.13.2"
  end

  # otto is our application
  config.vm.define "w01" do |app|
    app.vm.hostname = "w01.local"
    app.vm.network :private_network, ip: "10.24.13.12"
  end
end

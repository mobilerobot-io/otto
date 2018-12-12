providers = do vagrant
provider = vagrant

ap	= /usr/local/bin/ansible-playbook
inv	= /srv/inventory/${provider}

status:
	$(ap) -i $(inv) -status --list-hosts --list-tags --list-tasks

config:
	$(ap) -i $(inv) init.yml

otto:
	$(ap) -i $(inv) otto.yml

vagrant:
	$(ap) -i ../inv/vagrant

clean:
	rm -rf *~

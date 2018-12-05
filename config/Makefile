## Maybe make these variables
ap	= /usr/local/bin/ansible-playbook
inv	= ../inv/do/

# providers = do vagrant
# providers = vagrant
providers = do

status:
	$(ap) -i $(inv) -status --list-hosts --list-tags --list-tasks

config:
	$(ap) -i $(inv) init.yml

otto:
	$(ap) -i $(inv) otto.yml

vagrant:
	$(ap) -i ../inv/vagrant

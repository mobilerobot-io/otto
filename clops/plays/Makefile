## Maybe make these variables
ap	= /usr/local/bin/ansible-playbook
inv	= /srv/inventory/clowdops/hosts

status:
	$(ap) -i $(inv) -status site.yml --list-hosts --list-tags --list-tasks

config:
	$(ap) -i $(inv) site.yml


## Maybe make these variables
ap	= /usr/local/bin/ansible-playbook
inv	= /srv/inventory/clowdops

status:
	$(ap) -i $(inv) -status site.yml --list-hosts --list-tags --list-tasks


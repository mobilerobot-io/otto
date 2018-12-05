## Maybe make these variables
ap	= /usr/local/bin/ansible-playbook
inv	= ../inv/clowd/

status:
	$(ap) -i $(inv) -status --list-hosts --list-tags --list-tasks

config:
	$(ap) -i $(inv) playbook.yml


plays: /srv/otto-local
	cp -r plays /srv/otto-local

up: plays
	vagrant up

provision: plays
	vagrant provision

destroy:
	vagrant destroy

clean:
	go clean
	rm -rf *~ \#*
	rm -rf tmp

status:
	vagrant status

keys:
	mkdir tmp; cd tmp && ssh-keygen -q -f ./id_rsa

.PHONY: plays

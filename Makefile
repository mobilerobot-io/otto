bindir = /usr/local/bin

vag = $(bindir)/vagrant 
ans = $(bindir)/ansible
tf  = $(bindir)/terraform
hugo= $(bindir)/hugo

provider=do

status: 
	@echo "Vagrant hosts status..."
	@echo "-----------------------"
	@vagrant status | grep virtualbox
	@echo "Digital Ocean Droplets..."
	@echo "-----------------------"
	@doctl compute droplet list | awk '{ print $2 }'

prov:
	make -C prov $(provider)

up:
	$(vag) up

destroy:
	$(vag) $@

clean:
	rm -rf *~

.PHONY: up prov status destroy clean

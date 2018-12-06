bindir = /usr/local/bin

vag = $(bindir)/vagrant 
ans = $(bindir)/ansible
tf  = $(bindir)/terraform
hugo= $(bindir)/hugo

providers = do vagrant

status: 
	@echo "Vagrant hosts status..."
	@echo "-----------------------"
	@vagrant status | grep virtualbox


up:
	$(vag) up

destroy:
	$(vag) $@

clean:
	rm -rf *~

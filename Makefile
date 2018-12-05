bindir = /usr/local/bin

vag = $(bindir)/vagrant 
ans = $(bindir)/ansible
tf  = $(bindir)/terraform
hugo= $(bindir)/hugo

providers = do vagrant

status: 
	$(print-status) $(providers)

images:
	@echo TODO make images

prov:
	$(MAKE) -C $< $@

config:
	$(MAKE) -C $@ $< 

test:
	@echo TODO make test

stage:
	@echo TODO make stage

prod:
	@echo TODO make prod

destroy:
	$(MAKE) -C prov $<

.PHONY: images status prov

up:
	vagrant up

otto:
	$(MAKE) -C otto $@


define print-status = 
$(MAKE) -C images $@
$(MAKE) -C prov $@
$(MAKE) -C config $@
endef


dirs = src loca infra sites

bindir = /usr/local/bin

vag = $(bindir)/vagrant 
ans = $(bindir)/ansible
tf  = $(bindir)/tf

define print-status = 
$(MAKE) -C images $@
$(MAKE) -C prov $@
$(MAKE) -C config $@
endef

status:
	$(print-status)

images:
	$(MAKE) -C images $@

prov:
	$(MAKE) -C prov $<

config:
	$(MAKE) -C config

test:
	$(MAKE) -C test

stage:
	$(MAKE) -C stage

prod:
	$(MAKE) -C prod
	
destroy:
	$(MAKE) -C destory

.PHONY: images status prov

up:
	vagrant up

destroy:
	vagrant destroy

otto:
	$(MAKE) -C otto $@

provision: otto
	$(MAKE) -C loca $@	
	$(MAKE) -C infra $@
	$(MAKE) -C sites $@

config:
	$(MAKE) -C loca $@	
	$(MAKE) -C infra $@
	$(MAKE) -C sites $@
	vagrant provision

destroy:
	$(MAKE) -C loca $@	
	$(MAKE) -C infra $@
	$(MAKE) -C sites $@

clean:
	rm -rf *~ 

.PHONY: otto up provision status
status:
	vagrant status


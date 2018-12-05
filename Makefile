dirs = src loca infra sites

bindir = /usr/local/bin

vag = $(bindir)/vagrant 
ans = $(bindir)/ansible

status:
	$(MAKE) -C loca $@
	$(MAKE) -C infra $@
	$(MAKE) -C sites $@

local:
	$(vag) up

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

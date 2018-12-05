dirs = src loca infra sites

bindir = /usr/local/bin

vag = $(bindir)/vagrant 
ans = $(bindir)/ansible

status:
	make -C loca $@
	make -C infra $@
	make -C sites $@

local:
	$(vag) up

otto:
	make -C otto $@

provision: otto
	make -C loca $@	
	make -C infra $@
	make -C sites $@

config:
	make -C loca $@	
	make -C infra $@
	make -C sites $@
	vagrant provision

destroy:
	make -C loca $@	
	make -C infra $@
	make -C sites $@

clean:
	rm -rf *~ 

.PHONY: otto up provision status

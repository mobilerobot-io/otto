dirs = src infra sites

bidndir = /usr/local/bin

vag = $(bindir)/vagrant 
ans = $(bindir)/ansible

otto:
	make -C otto $@

provision: otto
	make -C infra $@
	make -C sites $@

config:
	make -C infra $@
	make -C sites $@
	vagrant provision

status:
	vagrant status
	make -C infra $@
	make -C sites $@

destroy:
	vagrant destroy
	make -C infra $@
	make -C sites $@

clean:
	rm -rf *~ 

.PHONY: otto up provision status

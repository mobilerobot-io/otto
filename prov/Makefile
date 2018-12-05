tf = /usr/local/bin/terraform
tfv = ./tfvars

status:
	$(tf) show

up: init plan apply

init plan apply destroy:
	$(tfv) $@

clean:
	rm *~

.PHONY: up init plan apply

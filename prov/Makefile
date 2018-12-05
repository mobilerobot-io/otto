tf = /usr/local/bin/terraform
tfv = ./tfvars

status:
	@echo ""
	@echo "provisions"
	@echo "------------------------------"
	@$(tf) show

up: init plan apply

init plan:
	$(tfv) $@ -input=false

apply destroy:
	$(tfv) $@ -input=false -auto-approve

clean:
	rm *~

check:
	doctl compute droplet list

.PHONY: up init plan apply

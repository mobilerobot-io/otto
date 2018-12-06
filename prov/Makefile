tf = /usr/local/bin/terraform
tfv = ./tfvars

# provider=$(provider)

status:
	@echo ""
	@echo " > " $(provider) $@
	@echo "------------------------------"
	@$(tf) show

up: init plan apply

init plan:
	$(tfv) $@ -input=false $(provider)

apply destroy:
	$(tfv) $@ -input=false -auto-approve

clean:
	rm *~

check:
	doctl compute droplet list

.PHONY: up init plan apply

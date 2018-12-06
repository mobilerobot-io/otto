tf = /usr/local/bin/terraform
tfv = ./tfvars

# provider=$(provider)

status:
	@echo ""
	@echo "Provider: " $(provider) $@
	@echo "--------------------------"
	@$(tf) show
	@echo ""
	@echo "Digital Ocean: "
	@echo "--------------------------"
	@doctl compute droplet list

up: init plan apply

init plan:
	$(tfv) $@ -input=false $(provider)

apply destroy:
	$(tfv) $@ -input=false -auto-approve $(provider)

clean:
	rm *~

check:
	doctl compute droplet list

.PHONY: up init plan apply

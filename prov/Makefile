allcmds = "get init plan apply"

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

.terraform:
	@echo "terraform must download required plugins first"
	$(tf) get

$(allcmds): .terraform
	@echo ".terraform directory must exist first ..."

up: $(allcmds)
	@echo "Bringing up $(topo) on $(provider)..."
	
apply destroy: init
	$(tfv) $@ -input=false -auto-approve $(provider)

clean:
	rm *~

check:
	doctl compute droplet list

.PHONY: up init plan apply

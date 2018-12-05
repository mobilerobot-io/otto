tf = /usr/local/bin/terraform

status:
	@echo "TODO - move this into its own module"

init:
	$(tf) init

plan:
	$(tf) plan

apply:
	$(tf) apply

destroy:
	$(tf) plan -destroy -out=terraform.tfplan
	$(tf) apply terraform.tfplan

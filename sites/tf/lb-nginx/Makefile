tf = ./tf

init:
	$(tf) init

plan:
	$(tf) plan

apply:
	$(tf) apply

destroy:
	$(tf) plan -destroy -out=terraform.tfplan
	$(tf) apply terraform.tfplan

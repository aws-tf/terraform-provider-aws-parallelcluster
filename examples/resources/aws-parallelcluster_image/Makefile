setup:
	@stat terraform.tfvars > /dev/null 2>&1 || cp terraform.tfvars.EXAMPLE terraform.tfvars
	@echo "[INFO] Add your settings to terraform.tfvars"

deploy:
	terraform init
	terraform plan -out tfplan
	terraform apply tfplan

clean:
	rm -rf .terraform
	rm -rf terraform.tfstate
	rm -rf terraform.tfstate.backup
	rm -rf .terraform.tfstate.lock.info
	rm -rf .terraform.lock.hcl
	rm -rf tfplan

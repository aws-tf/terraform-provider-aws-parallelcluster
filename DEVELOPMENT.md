# DEVELOPMENT

## Requirements
* go 1.21.0+: https://go.dev/
* Terraform 1.8.0+: https://developer.hashicorp.com/Terraform/install?product_intent=Terraform
* pre-commit 3.6.0+: https://pre-commit.com/
* git-secrets 1.3.0+: https://github.com/awslabs/git-secrets


## Setup
```
pre-commit install
git secrets --install
git secrets --register-aws
```

## Developing the Provider

`./internal/provider/openapi` contains a generated ParallelCluster go client.
It should be regenerated when a new release of ParallelCluster introduces changes to the OpenAPI spec file.
It can be generated with:

```shell
openapi-generator generate -g go -i /path/to/aws-parallelcluster/api/spec/openapi/ParallelCluster.openapi.yaml -o internal/provider/openapi -p 'structPrefix=true,withAWSV4Signature=true,enumClassPrefix=true'
```

The ParallelCluster spec file can be retrieved from https://github.com/aws/aws-parallelcluster/tree/develop/api/spec/openapi.

See https://openapi-generator.tech/docs/installation/ for openapi generator installation instructions.

After the go client is regenerated unit and End-to-End tests should be run (See below).

### Building

`make build` will build the provider for windows, macOS, and linux. An arm version will also be built for macOS. Binaries will be placed in the `./build` folder. See below for specific make targets.

- `darwin_arm64` Arm version for macOS.
- `darwin_amd64` x86_64 version for macOS.
- `windows_amd64`x86_64 version for windows.
- `linux_amd64` x86_64 version for linux.

### Installing

The provider can be installed locally for testing with the `install` target. The `VERSION` environment is used to determine the install path. To increment the version for local installs you could use the following:

`make install VERSION=1.0.0`

In order to use the local provider, a configuration block like the following may be used.

```Terraform
terraform {
  required_providers {
    aws-parallelcluster = {
      source  = "terraform.local/local/aws-parallelcluster"
      version = "1.0.0"
    }
  }
}
```

The hyphen in the version denotes a prerelease and will not be picked up by the version filter unless explicitly defined.

## Documentation
To generate or update documentation, run `go generate`. Documentation will be generated with Terraform-plugin-docs.

Terraform-plugin-docs will perform the following actions:

* Copy all the templates and static files to a temporary directory
* Build (`go build`) a temporary binary of the provider source code
* Collect schema information using `Terraform providers schema -json`
* Generate a default provider template file, if missing (**index.md**)
* Generate resource template files, if missing
* Generate data source template files, if missing
* Generate function template files, if missing (Requires Terraform v1.8.0+)
* Copy all non-template files to the output website directory
* Process all the remaining templates to generate files for the output website directory

Notable directories are in the `examples` folder. Examples in the documentation will be taken from the `examples/resources`, `examples/data-sources`, and `examples/provider` folders.

See https://github.com/hashicorp/Terraform-plugin-docs for more information.

## Testing

Workflows exist to lint go and terraform code, check code coverage, and run tests.

To check go code before creating a PR golangci-lint may be used. See for installation instructions:  https://golangci-lint.run/usage/install/#local-installation. Run with the following command: `golangci-lint run`.

To check terraform before creating a PR, the terraform fmt command may be used. Run with the following command: `terraform fmt -recursive ./examples`. Note, by default terraform fmt will fix files, if you only want to run a check add the `-check` flag.

Running `make test` is sufficient to run unit tests and end-to-end tests using the default aws profile in us-east-1. The following environment variables can customize the end-to-end tests.

- `TEST_REGION` The region used for deploying resources and query data sources. Defaults to "us-east-1".
- `TEST_ENDPOINT` The endpoint of the ParallelCluster API. If unset will be autodetected.
- `TEST_ROLE` The role used for deploying resources and query data sources. If unset the users default role will be used.
- `TEST_USE_USER_ROLE` Whether or not to use the user role exported by the ParallelCluster cloudformation stack. Either set this to true, supply a `role_arn` or set neither. If unset the default role will be used.
- `TEST_PCAPI_STACK_NAME` The ParallelCluster cloudformation stack name containing the endpoint and role outputs. If unset the API called "ParallelCluster" will be used.
- `TEST_CLUSTER_NAME` The name to use for the ParallelCluster created during end-to-end tests. Defaults to "test-cluster".

AWS configuration environment variables can be used as well. Such as `AWS_PROFILE`, `AWS_DEFAULT_PROFILE`, `AWS_SECRET_ACCESS_KEY`, `AWS_ACCESS_KEY_ID`.

The sdk test framework also has configuration variables. These control some of the behavior of End-to-End tests (called acceptance tests in the documentation) including Terraform configuration and log output. See https://developer.hashicorp.com/Terraform/plugin/sdkv2/testing/acceptance-tests#environment-variables for more information.

In order to run the full suite of End-to-End tests,  Run `make test_end2end`.

In order to run unit tests, run `make test_unit`. To create a coverage profile the following may be used: `make test_unit TESTARGS="-coverprofile=cover.out"`. To read the coverage profile use: `go tool cover -html=cover.out`.

To run end-to-end tests and unit tests, run `make test`.

*Note:* End-to-end tests create real resources, and often cost money to run.

Additional information: https://developer.hashicorp.com/Terraform/plugin/sdkv2/testing

## Releasing

Releases are created using `goreleaser`. A github action exists at `./github/workflows/release.yml` that will create a new release when new tags are pushed to the repo. Tags must match the following filter: `v*`.

The goreleaser configuration exists at `.goreleaser.yml`. The configuration is copied from the official terraform scaffholding.

For more information see:

https://github.com/hashicorp/terraform-provider-scaffolding-framework/blob/main/.goreleaser.yml
https://goreleaser.com/intro/

This Github workflow expects the following secrets:
- `GPG_SECRET_KEY` This secret GPG should be exported with the `--armor` flag and will need to match the public key pushed to the terraform registry (see below)
- `PASSPHRASE` The passphrase of the secret GPG key

NOTE: A version is of the form `vX.Y.Z`, if a `-` delimiter exists (ie. `v1.2.3-alpha`) terraform will consider it a prerelease and will not upgrade to it unless its explicitly defined in the configuration.

See https://developer.hashicorp.com/terraform/registry/providers/publishing#creating-a-github-release for more information.

## Publishing to the terraform registry

Releases are published automatically when a new release is created. This is done with a webhook that is created during the setup processes on the terraform registry. The public gpg key used during setup should match the private key used for releases. Otherwise terraform will throw signature errors during init.

See https://developer.hashicorp.com/terraform/registry/providers/publishing#publishing-to-the-registry for more information.

## Generating attribution document
Follow the below steps to generate/update the attribution document in `THIRD-PARTY-LICENSES.txt`.

1. Install [go-licenses](https://github.com/google/go-licenses)
```
go install -v github.com/google/go-licenses@latest
```

2. Execute the below command from repository root dir
```
go-licenses report . --template tools/licensing/go-licenses.tpl > THIRD-PARTY-LICENSES.txt 2> /tmp/go-licenses-errors.txt
```


## Logging
Enable logging by setting envars:
```
export TF_LOG_PATH="tf.log"
export TF_LOG="TRACE"
export TF_LOG_PROVIDER="TRACE"
```

Execute the provider in debug mode by setting `Debug: true` in `internal/provider/openapi/configuration.go`:

```
func NewConfiguration() *Configuration {
	cfg := &Configuration{
	    ...
		Debug:            true,
		...
	}
	return cfg
}
```

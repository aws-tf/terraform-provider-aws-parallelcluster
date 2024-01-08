# DEVELOPMENT

## Requirements
* go 1.21.5+: https://go.dev/
* terraform 1.6.6+: https://developer.hashicorp.com/terraform/install?product_intent=terraform
* pre-commit 3.6.0+: https://pre-commit.com/
* git-secrets 1.3.0+: https://github.com/awslabs/git-secrets


## Setup
```
pre-commit install
git secrets --install
git secrets --register-aws
```

## Developing the Provider

`./internal/provider/openapi` contains a generated parallelcluster go client. It can be generated with:

```shell
openapi-generator-cli generate -g go -i path-to-openapi-pcluster-spec -o internal/provider/openapi -p 'structPrefix=true,withAWSV4Signature=true,enumClassPrefix=true'
```

## Documentation
To generate or update documentation, run `go generate`.


## Testing
`TEST_REGION` and the `TEST_CLUSTER_NAME` environment variables must be set (ie. `export TEST_REGION=us-east-1 TEST_CLUSTER_NAME=test-cluster`).

In order to run the full suite of Acceptance tests,  Run `make testacc`.

In order to run unit tests, run `make testunit`.

To run acceptance tests and unit tests, run `make test`.

*Note:* Acceptance tests create real resources, and often cost money to run.

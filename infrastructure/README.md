## Setup GitHub and AWS credentials using Cloudformation
To access AWS resources from a GitHub workflow you need to create new IAM resources and use a specific action 
that will retrieve temporary credentials to access your account.

To create the resources needed by the workflow action you must deploy the `./github-env.yml` to CloudFormation 
with the default parameters. The stack name is not significant (example: GitHubEnv).

The stack will create, among other resources, the following roles, whose ARN is returned as stack outputs:
- the `ProviderE2ETestExecutionRole` with the permissions needed to run the end-to-end tests workflow for the provider, 
  its ARN should be put in the GitHub secret named `ACTION_E2E_TESTS_ROLE` in the provider repository.
- the `ModuleE2ETestExecutionRole` with the permissions needed to run the end-to-end tests workflow for the provider,
  its ARN should be put in the GitHub secret named `ACTION_E2E_TESTS_ROLE` in the module repository.

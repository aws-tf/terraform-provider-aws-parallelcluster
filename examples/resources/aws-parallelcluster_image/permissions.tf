resource "aws_iam_role" "cleanup_lambda_role" {
  name               = "CleanupLambdaRole-${random_id.suffix.id}"
  description        = "Role to be attached to the ParallelCluster cleanup lambda function for building image"
  path               = "/parallelcluster/"
  assume_role_policy = data.aws_iam_policy_document.assume_role_lambda_policy.json
}

data "aws_iam_policy_document" "assume_role_lambda_policy" {
  statement {
    actions = ["sts:AssumeRole"]
    principals {
      type        = "Service"
      identifiers = ["lambda.amazonaws.com"]
    }
  }
}

resource "aws_iam_policy" "cleanup_lambda_policy" {
  name        = "CleanupLambdaPolicy-${random_id.suffix.id}"
  description = "IAM policy for the ParallelCluster cleanup lambda function for building image"
  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect   = "Allow"
        Action   = ["iam:DetachRolePolicy", "iam:DeleteRole", "iam:DeleteRolePolicy"]
        Resource = "arn:${local.partition}:iam::${local.account_id}:role/parallelcluster/*"
      },
      {
        Effect   = "Allow"
        Action   = ["iam:DeleteInstanceProfile", "iam:RemoveRoleFromInstanceProfile"]
        Resource = "arn:${local.partition}:iam::${local.account_id}:instance-profile/parallelcluster/*"
      },
      {
        Effect   = "Allow"
        Action   = "imagebuilder:DeleteInfrastructureConfiguration"
        Resource = "arn:${local.partition}:imagebuilder:${local.region}:${local.account_id}:infrastructure-configuration/parallelclusterimage-*"
      },
      {
        Effect   = "Allow"
        Action   = ["imagebuilder:DeleteComponent"]
        Resource = "arn:${local.partition}:imagebuilder:${local.region}:${local.account_id}:component/parallelclusterimage-*/*"
      },
      {
        Effect   = "Allow"
        Action   = "imagebuilder:DeleteImageRecipe"
        Resource = "arn:${local.partition}:imagebuilder:${local.region}:${local.account_id}:image-recipe/parallelclusterimage-*/*"
      },
      {
        Effect   = "Allow"
        Action   = "imagebuilder:DeleteDistributionConfiguration"
        Resource = "arn:${local.partition}:imagebuilder:${local.region}:${local.account_id}:distribution-configuration/parallelclusterimage-*"
      },
      {
        Effect   = "Allow"
        Action   = ["imagebuilder:DeleteImage", "imagebuilder:GetImage", "imagebuilder:CancelImageCreation"]
        Resource = "arn:${local.partition}:imagebuilder:${local.region}:${local.account_id}:image/parallelclusterimage-*/*"
      },
      {
        Effect   = "Allow"
        Action   = "cloudformation:DeleteStack"
        Resource = "arn:${local.partition}:cloudformation:${local.region}:${local.account_id}:stack/*/*"
      },
      {
        Effect   = "Allow"
        Action   = "ec2:CreateTags"
        Resource = "arn:${local.partition}:ec2:${local.region}::image/*"
      },
      {
        Effect   = "Allow"
        Action   = "tag:TagResources"
        Resource = "*"
      },
      {
        Effect   = "Allow"
        Action   = ["lambda:DeleteFunction", "lambda:RemovePermission"]
        Resource = "arn:${local.partition}:lambda:${local.region}:${local.account_id}:function:ParallelClusterImage-*"
      },
      {
        Effect   = "Allow"
        Action   = "logs:DeleteLogGroup"
        Resource = "arn:${local.partition}:logs:${local.region}:${local.account_id}:log-group:/aws/lambda/ParallelClusterImage-*:*"
      },
      {
        Effect   = "Allow"
        Action   = ["SNS:GetTopicAttributes", "SNS:DeleteTopic", "SNS:GetSubscriptionAttributes", "SNS:Unsubscribe"]
        Resource = "arn:${local.partition}:sns:${local.region}:${local.account_id}:ParallelClusterImage-*"
      }
    ]
  })
}

resource "aws_iam_policy_attachment" "cleanup_lambda_role_attachment_1" {
  name       = "CleanupLambdaRolePolicyAttachment1-${random_id.suffix.id}"
  roles      = [aws_iam_role.cleanup_lambda_role.name]
  policy_arn = "arn:${local.partition}:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
}

resource "aws_iam_policy_attachment" "cleanup_lambda_role_attachment_2" {
  name       = "CleanupLambdaRolePolicyAttachment2-${random_id.suffix.id}"
  roles      = [aws_iam_role.cleanup_lambda_role.name]
  policy_arn = aws_iam_policy.cleanup_lambda_policy.arn
}
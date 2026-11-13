#!/usr/bin/env bash
#
# Deletes the resources that an end-to-end test run left behind, matching them
# by the suffix that identifies the run. The suffix is mandatory.
#
# Usage: cleanup-e2e-leftovers.sh <name-suffix> [--dry-run]

set -o pipefail

SUFFIX="${1:-}"
DRY_RUN="${2:-}"

# On stderr, to keep the stdout of the functions below free of log lines.
function info() { echo "[cleanup] $*" >&2; }

if [[ ! "${SUFFIX}" =~ ^-[A-Za-z0-9][A-Za-z0-9-]*$ ]]; then
  info "ERROR: expected a run suffix such as '-gh123-1-0', got '${SUFFIX}'. Refusing to run."
  exit 1
fi

function run() {
  if [[ "${DRY_RUN}" == "--dry-run" ]]; then
    info "WOULD RUN: $*"
  else
    info "RUN: $*"
    "$@" || info "WARNING: command failed: $*"
  fi
}

info "Cleaning up leftovers of the run identified by the suffix '${SUFFIX}'."

### CloudFormation stacks

# ParallelCluster names the stack of a cluster after the cluster, and the stack
# of an image build after the image, so matching the suffix covers both without
# assuming how the tests name them.
stacks="$(aws cloudformation list-stacks \
  --stack-status-filter CREATE_COMPLETE CREATE_FAILED CREATE_IN_PROGRESS \
  ROLLBACK_COMPLETE ROLLBACK_FAILED UPDATE_COMPLETE UPDATE_FAILED \
  UPDATE_ROLLBACK_COMPLETE UPDATE_ROLLBACK_FAILED DELETE_FAILED \
  --query "StackSummaries[?ends_with(StackName, '${SUFFIX}')].StackName" \
  --output text 2>/dev/null)"

deleted_stacks=()
for stack in ${stacks}; do
  run aws cloudformation delete-stack --stack-name "${stack}"
  deleted_stacks+=("${stack}")
done

# Every deletion is already in flight, so these waits overlap. They are what
# surfaces a stack that failed to delete.
if [[ "${DRY_RUN}" != "--dry-run" ]]; then
  for stack in "${deleted_stacks[@]}"; do
    info "Waiting for stack ${stack} to be deleted."
    aws cloudformation wait stack-delete-complete --stack-name "${stack}" ||
      info "WARNING: stack ${stack} was not deleted. Check what it still owns manually."
  done
fi

### AMIs built by the image test, and their snapshots

# Deleting the image stack removes the EC2 Image Builder resources, but not the
# AMI it produced, so deregister it and reclaim the snapshots it holds.
ami_ids="$(aws ec2 describe-images --owners self \
  --filters "Name=name,Values=*${SUFFIX}" \
  --query 'Images[].ImageId' --output text 2>/dev/null)"

if [[ -z "${ami_ids}" ]]; then
  info "No AMI whose name ends with ${SUFFIX} to clean up."
fi

for ami_id in ${ami_ids}; do
  # Read the snapshots while the AMI still exists: deregistering it is the
  # point at which the mapping to them is lost.
  snapshot_ids="$(aws ec2 describe-images --image-ids "${ami_id}" \
    --query 'Images[].BlockDeviceMappings[].Ebs.SnapshotId' --output text 2>/dev/null)"

  run aws ec2 deregister-image --image-id "${ami_id}"
  for snapshot_id in ${snapshot_ids}; do
    run aws ec2 delete-snapshot --snapshot-id "${snapshot_id}"
  done
done

### IAM roles and policies of the image test

# The image test grants the ParallelCluster cleanup lambda its own role, which
# survives a build whose teardown fails and then accumulates against the IAM
# quotas. Its policy has to be detached before either can be deleted.
for role in $(aws iam list-roles --path-prefix /parallelcluster/ \
  --query "Roles[?ends_with(RoleName, '${SUFFIX}')].RoleName" \
  --output text 2>/dev/null); do
  for attached in $(aws iam list-attached-role-policies --role-name "${role}" \
    --query 'AttachedPolicies[].PolicyArn' --output text 2>/dev/null); do
    run aws iam detach-role-policy --role-name "${role}" --policy-arn "${attached}"
  done
  run aws iam delete-role --role-name "${role}"
done

for policy_arn in $(aws iam list-policies --scope Local --path-prefix /parallelcluster/ \
  --query "Policies[?ends_with(PolicyName, '${SUFFIX}')].Arn" \
  --output text 2>/dev/null); do
  run aws iam delete-policy --policy-arn "${policy_arn}"
done

info "Done"

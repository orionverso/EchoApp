#!/usr/bin/env bash
##You must trust dev account in order to deploy to prod
if [[ $# -ge 3 ]]; then
    export CDK_DEV_ACCOUNT=$1
    export CDK_PROD_ACCOUNT=$2
    export CDK_PROD_REGION=$3
    shift; shift;

cdk bootstrap --trust $1 \
    --cloudformation-execution-policies arn:aws:iam::aws:policy/AdministratorAccess \
    aws://$2/$3 \
    --profile workerprod
    exit $?
else
    echo 1>&2 "Provide development region and account as first two args."
    echo 1>&2 "You must use --profile ROLE_WITH_PROD_PERMISION"
    echo 1>&2 "Additional args are passed through to cdk deploy."
    exit 1
fi

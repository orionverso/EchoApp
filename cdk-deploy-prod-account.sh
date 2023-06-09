#!/usr/bin/env bash
if [[ $# -ge 3 ]]; then
    export CDK_PROD_REGION=$1
    export CDK_PROD_ACCOUNT=$2
    shift; shift;
    cdk deploy "$@"
    exit $?
else
    echo 1>&2 "Provide development region and account as first two args."
    echo 1>&2 "You must use --profile ROLE_WITH_PROD_PERMISION"
    echo 1>&2 "Additional args are passed through to cdk deploy."
    exit 1
fi


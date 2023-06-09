#!/usr/bin/env bash
if [[ $# -ge 3 ]]; then
    export CDK_DEV_REGION=$1
    export CDK_DEV_ACCOUNT=$2
    shift; shift;
    cdk deploy "$@"
    exit $?
else
    echo 1>&2 "Provide development region and account as first two args."
    echo 1>&2 "You must use --profile ROLE_WITH_DEV_PERMISION"
    echo 1>&2 "Additional args are passed through to cdk deploy."
    exit 1
fi


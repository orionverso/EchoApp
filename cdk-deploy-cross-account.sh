#!/usr/bin/env bash
if [[ $# -ge 4 ]]; then
    export CDK_DEV_REGION=$1
    export CDK_DEV_ACCOUNT=$2
    export CDK_PROD_REGION=$3
    export CDK_PROD_ACCOUNT=$4
    shift; shift;shift;shift;
    cdk deploy "$@"
    exit $?
else
    echo 1>&2 "Provide development region and account as first two args."
    echo 1>&2 "Provide production region and account as second two args."
    echo 1>&2 "Additional args are passed through to cdk deploy."
    exit 1
fi

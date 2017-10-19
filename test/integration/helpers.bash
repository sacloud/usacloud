#!/bin/bash

function echo_to_log {
    echo "$BATS_TEST_NAME
----------
$output
----------

"   >> ${BATS_LOG}
}

function teardown {
    echo_to_log
}

function errecho {
    >&2 echo "$@"
}

function only_if_env {
    if [[ ${!1} != "$2" ]]; then
        errecho "This test requires the $1 environment variable to be set to $2. Skipping..."
        skip
    fi
}

function require_env {
    if [[ -z ${!1} ]]; then
        errecho "This test requires the $1 environment variable to be set in order to run."
        exit 1
    fi
}

#unset SAKURACLOUD_ACCESS_TOKEN SAKURACLOUD_ACCESS_TOKEN_SECRET SAKURACLOUD_ZONE USACLOUD_PROFILE

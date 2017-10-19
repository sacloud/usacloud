#!/bin/bash

set -e

# Usage: ./run-bats.sh [subtest]

function quiet_run () {
    if [[ "$VERBOSE" == "1" ]]; then
        "$@"
    else
        "$@" &>/dev/null
    fi
}

function cleanup_resources() {
    if [[ "$1" == "ALL" ]]; then
        echo "[TODO] cleanup_resources is not implements yet"
        # quiet_run run_usacloud_bin rm -f $MACHINE_NAME
    fi
}

function cleanup_config_store() {
    if [[ -d "$USACLOUD_PROFILE_DIR" ]]; then
        rm -r "$USACLOUD_PROFILE_DIR"
    fi
}


function run_usacloud() {
    "$USACLOUD_BIN_NAME" "$@"
}

function run_bats() {
    for bats_file in $(find "$1" -name \*.bats); do
        echo "=> $bats_file"
        # BATS returns non-zero to indicate the tests have failed, we shouldn't
        # necessarily bail in this case, so that's the reason for the e toggle.
        set +e
        bats "$bats_file"
        if [[ $? -ne 0 ]]; then
            EXIT_STATUS=1
        fi
        set -e
        echo
    done
}

function check_required_env() {
    local envs=(SAKURACLOUD_ACCESS_TOKEN SAKURACLOUD_ACCESS_TOKEN_SECRET SAKURACLOUD_ZONE)
    for e in ${envs[@]}; do
        if [[ -z "${!e}" ]]; then
            echo "This test requires the $e environment variable to be set in order to run."
            exit 1
        fi
    done
}
# Set this ourselves in case bats call fails
EXIT_STATUS=0
export BATS_FILE="$1"

if [[ -z "$BATS_FILE" ]]; then
    echo "You must specify a bats test to run."
    exit 1
fi

if [[ ! -e "$BATS_FILE" ]]; then
    echo "Requested bats file or directory not found: $BATS_FILE"
    exit 1
fi

export BASE_TEST_DIR=$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )
export USACLOUD_PROFILE_DIR="/tmp/usacloud_test_profile"
export PROJECT_ROOT="$BASE_TEST_DIR/../.."
export USACLOUD_BIN_NAME="usacloud"
export BATS_LOG="$PROJECT_ROOT/bats.log"

# Local builded binary (./bin/) takes precedence
export PATH="$PROJECT_ROOT"/bin:$PATH

# This function gets used in the integration tests, so export it.
export -f run_usacloud

> "$BATS_LOG"

check_required_env

cleanup_resources "ALL"
cleanup_config_store

mkdir -p "${USACLOUD_PROFILE_DIR}"

run_bats "$BATS_FILE"

cleanup_resources "ALL"
cleanup_config_store

exit ${EXIT_STATUS}
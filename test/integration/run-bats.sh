#!/bin/bash
# Copyright 2017-2020 The Usacloud Authors
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.


set -e

# Usage: ./run-bats.sh [subtest]

source $(cd $(dirname $0); pwd)/helpers.bash

function run_bats() {
    bats_files=("$1")
    if [ -d "$1" ]; then
        bats_files=$(find "$1" -name \*.bats)
    fi

    for bats_file in ${bats_files[@]}; do
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
export BATS_LOG="$PROJECT_ROOT/bats.log"
export TMP_PASSWORD=$(cat /dev/urandom | base64 | tr -dc 'a-zA-Z0-9' | fold -w 16 | head -n 1)

# Local builded binary (./bin/) takes precedence
export PATH="$PROJECT_ROOT"/bin:$PATH

# This function gets used in the integration tests, so export it.

> "$BATS_LOG"

check_required_env

cleanup_resources "ALL"
cleanup_config_store

mkdir -p "${USACLOUD_PROFILE_DIR}"

run_bats "$BATS_FILE"

cleanup_resources "ALL"
cleanup_config_store

exit ${EXIT_STATUS}
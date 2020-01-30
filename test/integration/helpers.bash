#
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
#
#!/bin/bash

# define environment variables
export USACLOUD_BIN_NAME="bin/usacloud"
export TEST_TARGET_NAME="usacloud-integration-test"
export TEST_TARGET_TAG="usacloud-integration-test"

ALL_RESOURCE_TYPES=(dns gslb simple-monitor icon license ssh-key startup-script server private-host database load-balancer nfs vpc-router archive auto-backup disk iso-image internet packet-filter switch bridge)
#ALL_RESOURCE_TYPES=(internet)
GLOBAL_RESOURCE_TYPES=(dns gslb simple-monitor icon license ssh-key startup-script bridge)
ZONES=(is1a is1b tk1a tk1v)
if [ -n "${USACLOUD_ZONES}" ]; then
    ZONES=(`echo ${USACLOUD_ZONES} | tr ',' ' '`)
fi

MK_ISO_CMD=""
if type mkisofs > /dev/null 2>&1; then
    MK_ISO_CMD="mkisofs -R -V config-2 "
elif type genisoimage > /dev/null 2>&1; then
    MK_ISO_CMD="genisoimage -R -V config-2 "
else
    MK_ISO_CMD="hdiutil makehybrid -iso -joliet -default-volume-name config-2 "
fi
export MK_ISO_CMD

function usacloud_run() {
    usacloud "$@" 2>/dev/null
}
function usacloud_run_with_stderr() {
    usacloud "$@"
}

export -f usacloud_run

function quiet_run () {
    if [[ "$VERBOSE" == "1" ]]; then
        echo "$@"
        "$@"
    else
        "$@" &>/dev/null
    fi
}

function cleanup_resources() {
    echo "========================================"
    echo -n "Cleanup resources..."
    if [[ "$SKIP_CLEANUP" == "1" ]]; then
        echo "skipped."
        echo "========================================"
        echo
        return
    fi
    for resource_type in ${ALL_RESOURCE_TYPES[@]}; do
        for zone in ${ZONES[@]}; do

            if [ "$1" == "ALL" ] || [ "$1" == "$resource_type" ]; then
                IDs=()
                SCOPE_OPTION=""
                # has --scope parameter in ls command?
                if [ $(usacloud_run --zone $zone $resource_type ls -h | fgrep -e '--scope' | wc -l) -ne 0 ]; then
                    SCOPE_OPTION=" --scope user "
                fi

                # has --name parameter in ls command?
                if [ $(usacloud_run --zone $zone $resource_type ls -h | fgrep -e '--name' | wc -l) -ne 0 ]; then
                    IDs+=(`usacloud_run --zone $zone $resource_type ls -q $SCOPE_OPTION --name "$TEST_TARGET_NAME" 2>/dev/null`)
                fi
                # has --tags parameter in ls command?
                if [ $(usacloud_run --zone $zone $resource_type ls -h | fgrep -e '--tags' | wc -l) -ne 0 ]; then
                    IDs+=(`usacloud_run --zone $zone $resource_type ls -q $SCOPE_OPTION --tags "$TEST_TARGET_TAG" 2>/dev/null`)
                fi
                IDs=(`echo ${IDs[*]} | tr ' ' '\n' | sort | uniq`)

                for id in ${IDs[@]}; do
                     case "$resource_type" in
                         "internet" )
                             ipv6nets=$(usacloud_run --zone $zone internet read --out json $id | jq ".[].Switch.IPv6Nets[] | .ID")
                             if [ -n "$(echo $ipv6nets)" ]; then
                                 for ipv6net_id in $(echo $ipv6nets); do
                                    # ipv6 prt record registered?
                                    ipv6addrs=$(usacloud_run --zone $zone ipv6 list --ipv6net-id $ipv6net_id --out json | jq -r ".[].IPv6Addr")
                                    if [ -n "$(echo $ipv6addrs)" ]; then
                                        for ipv6addr in $(echo $ipv6addrs); do
                                            # delete ipv6 prt record
                                            quiet_run usacloud_run --zone $zone ipv6 ptr-delete -y $ipv6addr
                                        done
                                    fi
                                 done
                                 # disable ipv6
                                 quiet_run usacloud_run --zone $zone internet ipv6-disable -y $id
                             fi

                             subnets=$(usacloud_run --zone $zone internet read --out json $id | jq ".[].Switch.Subnets[] | .ID")
                             if [ -n "$(echo $subnets)" ]; then
                                 i=0
                                 for subnet_id in $(echo $subnets); do
                                    if [ $i -gt 0 ]; then
                                        # delete subnets
                                        quiet_run usacloud_run --zone $zone internet subnet-delete -y --subnet-id $subnet_id $id
                                    fi
                                    i=$( expr $i + 1 )
                                 done
                             fi

                             ;;
                         "switch" )
                             # is connected to bridge?
                             bridge=$(usacloud_run --zone $zone switch read --out json $id | jq ".[].Bridge.ID")
                             if [ "$bridge" != "null" ]; then
                                 # disconnect from bridge
                                 quiet_run usacloud_run --zone $zone switch bridge-disconnect -y $id
                             fi
                             ;;
                     esac

                     FORCE_OPTION=""
                     # has --force parameter in rm command?
                     if [ $(usacloud_run --zone $zone $resource_type rm -h | fgrep -e '--force' | wc -l) -ne 0 ]; then
                         FORCE_OPTION=" --force "
                     fi

                     quiet_run usacloud_run --zone $zone $resource_type rm -y $FORCE_OPTION $id 2>/dev/null
                done

            # quiet_run usacloud_run_bin rm -f $MACHINE_NAME
            fi

            # グローバルリソースの場合はbreak
            if [[ " ${GLOBAL_RESOURCE_TYPES[@]} " =~ " ${resource_type} " ]]; then
                break
            fi
        done
    done
    echo "done."
}

function cleanup_config_store() {
    if [[ -d "$USACLOUD_PROFILE_DIR" ]]; then
        rm -r "$USACLOUD_PROFILE_DIR"
    fi
}

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

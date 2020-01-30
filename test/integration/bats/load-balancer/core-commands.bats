#!/usr/bin/env bats
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


load ${BASE_TEST_DIR}/helpers.bash

resource_name=${TEST_TARGET_NAME}01

@test "Usacloud: should can create load-balancer" {
  # create Switch
  run usacloud_run switch create --name "$resource_name" -y -q

  [ -n "${output}" ]
  [ ${status} -eq 0 ]

  switch_id=$(echo ${output})

  run usacloud_run load-balancer create -y -q \
          --name "$resource_name" \
          --switch-id "$switch_id" \
          --ipaddress1 "192.168.100.101" \
          --nw-mask-len 24

  [ -n "${output}" ]
  [ ${status} -eq 0 ]

}

@test "Usacloud: should can CRUD VIP" {

  lb_id=$(usacloud_run load-balancer read -q "$resource_name")

  # run vip-add
  run usacloud_run load-balancer vip-add --vip 192.168.100.1 --port 80 -y "$resource_name"
  [ "${status}" -eq 0 ]

  run usacloud_run load-balancer vip-add --vip 192.168.100.1 --port 443 -y "$resource_name"
  [ "${status}" -eq 0 ]

  # run vip-update
  # should error(duplicate ip:port)
  run usacloud_run load-balancer vip-update --index 2 --port 80 -y "$resource_name"
  [ "${status}" -eq 1 ]

  run usacloud_run load-balancer vip-update --index 1 --vip 192.168.100.2 -y "$resource_name"
  [ "${status}" -eq 0 ]

  # run vip-del
  run usacloud_run load-balancer vip-delete --index 1 -y "$resource_name"
  [ "${status}" -eq 0 ]
}

@test "Usacloud: should can delete load-balancer" {

  run usacloud_run load-balancer rm -f -y "${resource_name}"

  [ -n "${output}" ]
  [ "${status}" -eq 0 ]

  run usacloud_run switch rm -y "${resource_name}"

  [ -n "${output}" ]
  [ "${status}" -eq 0 ]

}

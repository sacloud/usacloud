#!/usr/bin/env bats
#
# Copyright 2017-2019 The Usacloud Authors
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

load ${BASE_TEST_DIR}/helpers.bash

password="$TMP_PASSWORD"
resource_name=${TEST_TARGET_NAME}01

@test "Usacloud: should can create database(MariaDB)" {
  # create Switch
  run usacloud_run switch create --name "$resource_name" -y -q

  [ -n "${output}" ]
  [ ${status} -eq 0 ]

  switch_id=$(echo ${output})

  run usacloud_run database create -y -q \
          --name "$resource_name" \
          --switch-id "$switch_id" \
          --ipaddress1 "192.168.100.101" \
          --nw-mask-len 24 \
          --default-route "192.168.100.1" \
          --database mariadb \
          --username "UsacloudTest" \
          --password "${password}"

  [ -n "${output}" ]
  [ ${status} -eq 0 ]

}

@test "Usacloud: should can create database(PostgreSQL)" {
  # create Switch
  run usacloud_run switch create --name "$resource_name" -y -q

  [ -n "${output}" ]
  [ ${status} -eq 0 ]

  switch_id=$(echo ${output})

  run usacloud_run database create -y -q \
          --name "$resource_name" \
          --switch-id "$switch_id" \
          --ipaddress1 "192.168.100.101" \
          --nw-mask-len 24 \
          --default-route "192.168.100.1" \
          --database postgresql \
          --username "UsacloudTest" \
          --password "${password}"

  [ -n "${output}" ]
  [ ${status} -eq 0 ]

}

#!/usr/bin/env bats
# Copyright 2017-2020 The Usackoud Authors
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

setup(){
  quiet_run usacloud_run server build -y --disk-mode diskless --name Test11 --tags Test11 --tags $TEST_TARGET_TAG
  quiet_run usacloud_run server build -y --disk-mode diskless --name Test12 --tags Test12 --tags $TEST_TARGET_TAG
  quiet_run usacloud_run server build -y --disk-mode diskless --name Test21 --tags Test21 --tags $TEST_TARGET_TAG
  quiet_run usacloud_run server build -y --disk-mode diskless --name Test31 --tags Test31 --tags $TEST_TARGET_TAG
}

teardown(){
  quiet_run usacloud_run server rm -y -f --selector $TEST_TARGET_TAG
}

@test "Usacloud: should be able to can list servers" {
  # no filter
  run usacloud_run server list -q --tags $TEST_TARGET_TAG
  [ -n "${output}" ]
  [ ${#lines[*]} -eq 4 ]
  [ ${status} -eq 0 ]

  run usacloud_run server list -q --name "MissingResourceName" --tags $TEST_TARGET_TAG
  [ -z "${output}" ]
  [ "${#lines[*]}" -eq 0 ]
  [ "${status}" -eq 0 ]

  # filter by name
  run usacloud_run server list -q --name "Test1" --tags $TEST_TARGET_TAG
  [ "${#lines[*]}" -eq 2 ]

  run usacloud_run server list -q --name "Test2" --tags $TEST_TARGET_TAG
  [ "${#lines[*]}" -eq 1 ]

  run usacloud_run server list -q --name "Test3" --tags $TEST_TARGET_TAG
  [ "${#lines[*]}" -eq 1 ]

  run usacloud_run server list -q --name "2" --tags $TEST_TARGET_TAG
  [ "${#lines[*]}" -eq 2 ]

  run usacloud_run server list -q --name "3" --tags $TEST_TARGET_TAG
  [ "${#lines[*]}" -eq 1 ]

  # filter by tags
  run usacloud_run server list -q --tags "Test11" --tags $TEST_TARGET_TAG
  [ "${#lines[*]}" -eq 1 ]

}




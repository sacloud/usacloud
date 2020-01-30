#!/usr/bin/env bats
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

load ${BASE_TEST_DIR}/helpers.bash

@test "Usacloud: should show auth status with default sub-command" {
  run usacloud_run auth-status 2>/dev/null
  [ -n "${output}" ]
  [ "${status}" -eq 0 ]
}

@test "Usacloud: should show auth status using valid api-key" {
  run usacloud_run auth-status show 2>/dev/null
  [ -n "${output}" ]
  [ "${status}" -eq 0 ]
}

@test "Usacloud: should not show auth status using invalid api-key" {
  run usacloud_run --token "invalid" --secret "invalid" auth-status show
  [ -z "${output}" ]
  [ "${status}" -eq 1 ]
}
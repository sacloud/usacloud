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

DUMMY_RAW_FILE="integration-test-dummy.raw"
ARCHIVE_NAME=${TEST_TARGET_NAME}01

function setup() {
  dd if=/dev/zero of="${DUMMY_RAW_FILE}" bs=1024 count=1024
}

function teardown() {
  rm -f "${DUMMY_RAW_FILE}"
}

@test "Usacloud: should can create and upload archive" {
  run usacloud_run archive create -y --size 20 --archive-file "${DUMMY_RAW_FILE}" --name "${ARCHIVE_NAME}"

  [ -n "${output}" ]
  [ "${status}" -eq 0 ]
}

@test "Usacloud: should can read archive" {
  run usacloud_run archive read --out json "${ARCHIVE_NAME}"

  [ -n "${output}" ]
  [ "${status}" -eq 0 ]

  # parse JSON
  res=$(echo ${output} | jq ".[]")
  [ "$(echo ${res} | jq ".Availability")" == '"available"'  ]

  id=$(echo ${res} | jq ".ID")
  originID=$(echo ${res} | jq ".OriginalArchive.ID")
  [ "$id" -eq "$originID" ]

  [ "$(echo ${res} | jq ".SizeMB")" -eq 20480 ]
}

@test "Usacloud: should can delete archive" {
  run usacloud_run archive delete -y "${ARCHIVE_NAME}"

  [ -n "${output}" ]
  [ "${status}" -eq 0 ]
}
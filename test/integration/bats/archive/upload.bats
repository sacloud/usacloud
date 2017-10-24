#!/usr/bin/env bats

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
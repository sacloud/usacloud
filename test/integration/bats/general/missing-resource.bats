#!/usr/bin/env bats

load ${BASE_TEST_DIR}/helpers.bash

BAD_RESOURCE="foobar"

@test "Usacloud: should not allow run with bad resource type" {
  run usacloud_run $BAD_RESOURCE
  [ "${status}" -eq 1 ]
}
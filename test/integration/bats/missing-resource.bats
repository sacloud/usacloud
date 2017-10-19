#!/usr/bin/env bats

load ${BASE_TEST_DIR}/helpers.bash

BAD_RESOURCE="foobar"

@test "usacloud: Should not allow run with bad resource type" {
  run run_usacloud $BAD_RESOURCE
  [[ ${status} -eq 1 ]]
}
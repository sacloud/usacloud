#!/usr/bin/env bats

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
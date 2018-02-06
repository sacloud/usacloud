#!/usr/bin/env bats

load ${BASE_TEST_DIR}/helpers.bash

password="$TMP_PASSWORD"
hostname=${TEST_TARGET_NAME}01

@test "Usacloud: should build server with using Windows 2016 RDS" {

  run usacloud_run_with_stderr server build -y -q \
          --os-type windows2016-rds \
          --password "$password" \
          --disk-size 100 \
          --name "$hostname"

  [ -n "${output}" ]
  [ ${status} -eq 1 ]
}


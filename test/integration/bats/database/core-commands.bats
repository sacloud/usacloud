#!/usr/bin/env bats

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

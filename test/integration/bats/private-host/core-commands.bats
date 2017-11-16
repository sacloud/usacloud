#!/usr/bin/env bats

load ${BASE_TEST_DIR}/helpers.bash

hostname=${TEST_TARGET_NAME}01

@test "Usacloud: should can create private-host" {

  # build server with CentOS pablic archive and minimum options
  run usacloud_run private-host create -y -q \
          --name "$hostname" \
          --tags "tags1" \
          --description "description" \

  [ -n "${output}" ]
  [ ${status} -eq 0 ]

}

@test "Usacloud: should can add server" {

  # create server(boot after created)
  run usacloud_run server build --name "$hostname" --disk-mode diskless -y -q

  [ -n "${output}" ]
  [ ${status} -eq 0 ]

  serverID=$(echo ${output})

  # run server-add(should error)
  run usacloud_run private-host server-add --server-id "$serverID" -y "$hostname"
  [ "${status}" -eq 1 ]

  # shutdown server
  run usacloud_run server shutdown-force -y "$serverID"
  run usacloud_run server wait-for-down "$serverID"

  # run server-add(should ok)
  run usacloud_run private-host server-add --server-id "$serverID" -y "$hostname"
  [ "${status}" -eq 0 ]

  # delete private-host(should error)
  run usacloud_run private-host rm -y "${hostname}"
  [ "${status}" -eq 1 ]

  # boot server
  run usacloud_run server boot -y "$serverID"

  # run server-del(should error)
  run usacloud_run private-host server-delete --server-id "$serverID" -y "$hostname"
  [ "${status}" -eq 1 ]

  # shutdown server
  run usacloud_run server shutdown-force -y "$serverID"

  # run server-del(should ok)
  run usacloud_run private-host server-delete --server-id "$serverID" -y "$hostname"
  [ "${status}" -eq 0 ]

  # cleanup
  run usacloud_run server rm -f -y "$serverID"
}

@test "Usacloud: should can delete private-host" {

  # build server with CentOS pablic archive and minimum options
  run usacloud_run private-host rm -y "${hostname}"

  [ -n "${output}" ]
  [ ${status} -eq 0 ]

}
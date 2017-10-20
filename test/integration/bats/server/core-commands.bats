#!/usr/bin/env bats

load ${BASE_TEST_DIR}/helpers.bash

password="$TMP_PASSWORD"
hostname=${TEST_TARGET_NAME}01

@test "Usacloud: should build server with using CentOS and minimum options" {

  # build server with CentOS pablic archive and minimum options
  run usacloud_run server build -y -q \
          --os-type centos \
          --password "$password" \
          --name "$hostname" \
          --hostname "$hostname"

  [ -n "${output}" ]
  [ ${status} -eq 0 ]

}

@test "Usacloud: should read server JSON with CentOS and minimum options" {

  # read server
  run usacloud_run server read --out json $hostname

  [ -n "${output}" ]
  [ ${status} -eq 0 ]

  # parse JSON
  res=$(echo ${output} | jq ".[]")

  [ "$(echo ${res} | jq ".ServerPlan.CPU")" -eq 1  ]
  [ "$(echo ${res} | jq ".ServerPlan.MemoryMB")" -eq 1024  ]
  [ "$(echo ${res} | jq ".Disks | length")" -eq 1  ]
  [ "$(echo ${res} | jq ".Disks[].Plan.ID ")" -eq 4  ] # hdd=2 / ssd=4
  [ "$(echo ${res} | jq ".Disks[].Connection ")" == '"virtio"'  ] # virtio or ide
  [ "$(echo ${res} | jq ".Disks[].SizeMB ")" -eq 20480  ]

  # check source_archive_id
  disk_id="$(echo ${res} | jq ".Disks[].ID")"
  centos_archive_id=$(usacloud_run archive read -q --selector "distro-centos" --selector "current-stable")
  source_archive_id=$(usacloud_run disk read --out json "$disk_id" | jq ".[].SourceArchive.ID")
  [ "$centos_archive_id" == "$source_archive_id" ]

  source_disk_id=$(usacloud_run disk read --out json "$centos_archive_id" | jq ".[].SourceDisk.ID")
  [ -z "$source_disk_id" ]

}

@test "Usacloud: should can delete server with server name" {
  run usacloud_run server rm -f -y $hostname

  [ -n "${output}" ]
  [ ${status} -eq 0 ]

}
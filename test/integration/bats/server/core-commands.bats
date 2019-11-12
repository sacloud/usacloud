#!/usr/bin/env bats
# Copyright 2017-2019 The Usackoud Authors
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


load ${BASE_TEST_DIR}/helpers.bash

password="$TMP_PASSWORD"
hostname=${TEST_TARGET_NAME}01

@test "Usacloud: should be able to build server with using CentOS and minimum options" {

  # build server with CentOS pablic archive and minimum options
  run usacloud_run server build -y -q \
          --os-type centos \
          --password "$password" \
          --name "$hostname" \
          --hostname "$hostname"

  [ -n "${output}" ]
  [ ${status} -eq 0 ]

}

@test "Usacloud: should be able to read server JSON with CentOS and minimum options" {

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

@test "Usacloud: should be able to delete server with server name" {
  run usacloud_run server rm -f -y $hostname

  [ -n "${output}" ]
  [ ${status} -eq 0 ]

}
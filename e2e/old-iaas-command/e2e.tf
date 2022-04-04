terraform {
  required_providers {
    sakuracloud = {
      source  = "sacloud/sakuracloud"
      version = "2.16.2"
    }
  }
}

locals {
  zones = ["is1a", "is1b", "tk1a", "tk1b"]
}

resource "sakuracloud_server" "server" {
  count = length(local.zones)

  zone = local.zones[count.index]
  name = format("usacloud-e2e-old-iaas-command%02d", count.index)

  force_shutdown = true
}
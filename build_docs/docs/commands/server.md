# server

---

サーバーの操作を行います。

# コマンド

* 基本的な操作
    - [`list`](#list) - 一覧表示
    - [`build`](#build) - 作成
    - [`read`](#read) - 詳細表示
    - [`update`](#update) - 更新
    - [`delete`](#delete) - 削除
    - [`plan-change`](#plan-change) - プラン変更
* 電源操作
    - [`boot`](#boot) - 起動
    - [`shutdown`](#shutdown) - シャットダウン(graceful)
    - [`shutdown-force`](#shutdown-force) - シャットダウン(force)
    - [`reset`](#reset) - リセット
    - [`wait-for-boot`](#wait-for-boot) - 起動するまで待機
    - [`wait-for-down`](#wait-for-down) - 終了するまで待機
* SSH/SCP/VNC
    - [`ssh`](#ssh) - SSH接続
    - [`ssh-exec`](#ssh-exec) - SSH接続(コマンド実行)
    - [`scp`](#scp) - SCPでのファイル送受信
    - [`vnc`](#vnc) - VNCクライアント起動
    - [`vnc-info`](#vnc-info) - VNC接続情報表示
* ディスク管理
    - [`disk-info`](#disk-info1) - 接続されているディスクを一覧表示
    - [`disk-connect`](#disk-connect) - ディスクの接続
    - [`disk-disconnect`](#disk-disconnect) - ディスクの切断
* ネットワーク(NIC)管理
    - [`interface-info`](#interface-info) - 接続されているNICを一覧表示
    - [`interface-add-for-internet`](#interface-add-for-internet) - インターネット(共有セグメント)に接続するNICを追加
    - [`interface-add-for-router`](#interface-add-for-router) - ルーターに接続するNICを追加
    - [`interface-add-for-switch`](#interface-add-for-switch) - スイッチに接続するNICを追加
    - [`interface-add-disconnected`](#interface-add-disconnected) - 上流ネットワークに接続していないNICを追加
* ISOイメージ(CD-ROM)
    - [`iso-info`](#iso-info) - 挿入されているISOイメージ(CD-ROM)の詳細表示
    - [`iso-insert`](#iso-insert) - ISOイメージ(CD-ROM)の挿入
    - [`iso-eject`](#iso-eject) - ISOイメージ(CD-ROM)の排出
* 監視/モニタリング
    - [`monitor-cpu`](#monitor-cpu) - アクティビティモニタ(CPU)取得
    - [`monitor-nic`](#monitor-nic) - アクティビティモニタ(NIC)取得
    - [`monitor-disk`](#monitor-disk) - アクティビティモニタ(ディスク)取得

# == 基本的な操作 ==

---

## list

```bash
USAGE:
   usacloud server list [command options] [arguments...]

OPTIONS:
   --from value  set offset (default: 0)
   --id value    set filter by id(s)
   --max value   set limit (default: 0)
   --name value  set filter by name(s)
   --sort value  set field(s) for sort
   --help, -h    show help (default: false)
```

---

## build

```bash
USAGE:
   usacloud server build [command options] [arguments...]

OPTIONS:
 === For server-plan options ===
   --core value    [Required] set CPU core count (default: 1)
   --memory value  [Required] set memory size(GB) (default: 1)
   
 === For disk options ===
   --disk-mode value          [Required] disk create mode[create/connect/diskless] (default: "create")
   --os-type value            set source OS type
   --disk-plan value          set disk plan('hdd' or 'ssd') (default: "ssd")
   --disk-connection value    set disk connection('virtio' or 'ide') (default: "virtio")
   --disk-size value          set disk size(GB) (default: 20)
   --source-archive-id value  set source disk ID (default: 0)
   --source-disk-id value     set source disk ID (default: 0)
   --distant-from value       set distant from disk IDs
   --disk-id value            set connect disk ID (default: 0)
   
 === For ISO image options ===
   --iso-image-id value  set iso-image ID (default: 0)
   
 === For network options ===
   --network-mode value      [Required] network connection mode[shared/switch/disconnect/none] (default: "shared")
   --use-nic-virtio          use virtio on nic (default: true)
   --packet-filter-id value  set packet filter ID (default: 0)
   --switch-id value         set connect switch ID (default: 0)
   
 === For edit-disk options ===
   --hostname value                            set hostname
   --password value                            set password
   --disable-password-auth, --disable-pw-auth  disable password auth on SSH (default: false)
   
 === For edit-disk(network settings) options ===
   --ipaddress value, --ip value                set ipaddress
   --nw-masklen value, --network-masklen value  set ipaddress  prefix (default: 24)
   --default-route value, --gateway value       set default gateway
   
 === For edit-disk(startup-script) options ===
   --startup-scripts value, --notes value        set startup script(s)
   --startup-script-ids value, --note-ids value  set startup script ID(s)
   --startup-scripts-ephemeral                   set startup script persist mode (default: true)
   
 === For edit-disk(ssh-key) options ===
   --ssh-key-mode value                ssh-key mode[none/id/generate/upload]
   --ssh-key-name value                set ssh-key name
   --ssh-key-ids value                 set ssh-key ID(s)
   --ssh-key-pass-phrase value         set ssh-key pass phrase
   --ssh-key-description value         set ssh-key description
   --ssh-key-private-key-output value  set ssh-key privatekey output path
   --ssh-key-public-keys value         set ssh-key public key
   --ssh-key-public-key-files value    set ssh-key public key file
   --ssh-key-ephemeral                 set ssh-key persist mode (default: true)
   
 === For server-info options ===
   --name value                       [Required] set resource display name
   --description value, --desc value  set resource description
   --tags value                       set resource tags
   --icon-id value                    set Icon ID (default: 0)
   
 === Common options ===
   --assumeyes, -y              assume that the answer to any question which would be asked is yes (default: false)
   --us-keyboard                use us-keyboard (default: false)
   --disable-boot-after-create  boot after create (default: false)
   --help, -h                   show help (default: false)
```

---

## read

```bash
USAGE:
   usacloud server read [command options] <ID or Name(allow multiple target)>

OPTIONS:
 === Common options ===
   --help, -h  show help (default: false)
```

---

## update

```bash
USAGE:
   usacloud server update [command options] <ID or Name(allow multiple target)>

OPTIONS:
 === Common options ===
   --assumeyes, -y                    assume that the answer to any question which would be asked is yes (default: false)
   --description value, --desc value  set resource description
   --icon-id value                    set Icon ID (default: 0)
   --name value                       set resource display name
   --tags value                       set resource tags
   --help, -h                         show help (default: false)
```

---

## delete

```bash
USAGE:
   usacloud server delete [command options] <ID or Name(allow multiple target)>

OPTIONS:
 === Common options ===
   --assumeyes, -y  assume that the answer to any question which would be asked is yes (default: false)
   --force, -f      forced-shutdown flag if server is running (default: false)
   --with-disk      delete connected disks with server (default: true)
   --help, -h       show help (default: false)
```

---

## plan-change

```bash
USAGE:
   usacloud server plan-change [command options] <ID or Name(allow multiple target)>

OPTIONS:
 === Common options ===
   --assumeyes, -y  assume that the answer to any question which would be asked is yes (default: false)
   --core value     [Required] set CPU core count (default: 0)
   --memory value   [Required] set memory size(GB) (default: 0)
   --help, -h       show help (default: false)
```

---

# == 電源操作 ==

---

## boot

```bash
USAGE:
   usacloud server boot [command options] <ID or Name(allow multiple target)>

OPTIONS:
   --assumeyes, -y  assume that the answer to any question which would be asked is yes (default: false)
   --help, -h       show help (default: false)
```

---

## shutdown

```bash
USAGE:
   usacloud server shutdown [command options] <ID or Name(allow multiple target)>

OPTIONS:
   --assumeyes, -y  assume that the answer to any question which would be asked is yes (default: false)
   --help, -h       show help (default: false)
```

---

## shutdown-force

```bash
USAGE:
   usacloud server shutdown-force [command options] <ID or Name(allow multiple target)>

OPTIONS:
   --assumeyes, -y  assume that the answer to any question which would be asked is yes (default: false)
   --help, -h       show help (default: false)
```

---

## reset

```bash
USAGE:
   usacloud server reset [command options] <ID or Name(allow multiple target)>

OPTIONS:
   --assumeyes, -y  assume that the answer to any question which would be asked is yes (default: false)
   --help, -h       show help (default: false)
```

---

## wait-for-boot

```bash
USAGE:
   usacloud server wait-for-boot [command options] <ID or Name(allow multiple target)>

OPTIONS:
   --help, -h  show help (default: false)
```

---

## wait-for-down

---

# == SSH/SCP/VNC ==

```bash
USAGE:
   usacloud server wait-for-down [command options] <ID or Name(allow multiple target)>

OPTIONS:
   --help, -h  show help (default: false)
```

---

## ssh

```bash
USAGE:
   usacloud server ssh [command options] <ID or Name(only single target)>

OPTIONS:
   --key value, -i value   private-key file path
   --password value        password(or private-key pass phrase) [$SAKURACLOUD_SSH_PASSWORD]
   --port value, -p value  [Required] port (default: 22)
   --quiet, -q             disable information messages (default: false)
   --user value, -l value  user name
   --help, -h              show help (default: false)
```

---

## ssh-exec

```bash
USAGE:
   usacloud server ssh-exec [command options] <ID or Name(only single target)>

OPTIONS:
   --key value, -i value   private-key file path
   --password value        password(or private-key pass phrase) [$SAKURACLOUD_SSH_PASSWORD]
   --port value, -p value  [Required] port (default: 22)
   --quiet, -q             disable information messages (default: false)
   --user value, -l value  user name
   --help, -h              show help (default: false)
```

---

## scp

```bash
USAGE:
   usacloud server scp [command options] [ServerID:]<FROM> [ServerID:]<TO>

OPTIONS:
   --assumeyes, -y         assume that the answer to any question which would be asked is yes (default: false)
   --key value, -i value   private-key file path
   --password value        password(or private-key pass phrase) [$SAKURACLOUD_SSH_PASSWORD]
   --port value, -p value  [Required] port (default: 22)
   --quiet, -q             disable information messages (default: false)
   --recursive, -r         set recursive copy flag (default: false)
   --user value, -l value  user name
   --help, -h              show help (default: false)
```

---

## vnc

```bash
USAGE:
   usacloud server vnc [command options] <ID or Name(allow multiple target)>

OPTIONS:
   --assumeyes, -y  assume that the answer to any question which would be asked is yes (default: false)
   --help, -h       show help (default: false)
```

---

## vnc-info

```bash
USAGE:
   usacloud server vnc-info [command options] <ID or Name(allow multiple target)>

OPTIONS:
 === Common options ===
   --help, -h  show help (default: false)
```

---

# == ディスク管理 ===

---

## disk-info

```bash
USAGE:
   usacloud server disk-info [command options] <ID or Name(only single target)>

OPTIONS:
 === Common options ===
   --help, -h  show help (default: false)
```

---

## disk-connect

```bash
USAGE:
   usacloud server disk-connect [command options] <ID or Name(only single target)>

OPTIONS:
   --assumeyes, -y  assume that the answer to any question which would be asked is yes (default: false)
   --disk-id value  [Required] set target disk ID (default: 0)
   --help, -h       show help (default: false)
```

---

## disk-disconnect

```bash
USAGE:
   usacloud server disk-disconnect [command options] <ID or Name(only single target)>

OPTIONS:
   --assumeyes, -y  assume that the answer to any question which would be asked is yes (default: false)
   --disk-id value  [Required] set target disk ID (default: 0)
   --help, -h       show help (default: false)
```

---

# == ネットワーク(NIC)管理 ==

---

## interface-info

```bash
USAGE:
   usacloud server interface-info [command options] <ID or Name(only single target)>

OPTIONS:
 === Common options ===
   --help, -h  show help (default: false)
```

---

## interface-add-for-internet

```bash
USAGE:
   usacloud server interface-add-for-internet [command options] <ID or Name(only single target)>

OPTIONS:
   --assumeyes, -y      assume that the answer to any question which would be asked is yes (default: false)
   --without-disk-edit  set skip edit-disk flag. if true, don't call DiskEdit API after interface added (default: false)
   --help, -h           show help (default: false)
```

---

## interface-add-for-router

```bash
USAGE:
   usacloud server interface-add-for-router [command options] <ID or Name(only single target)>

OPTIONS:
   --assumeyes, -y                              assume that the answer to any question which would be asked is yes (default: false)
   --default-route value, --gateway value       set default gateway
   --ipaddress value, --ip value                set ipaddress
   --nw-masklen value, --network-masklen value  set ipaddress  prefix (default: 24)
   --switch-id value                            [Required] set connect switch(connected to router) ID (default: 0)
   --without-disk-edit                          set skip edit-disk flag. if true, don't call DiskEdit API after interface added (default: false)
   --help, -h                                   show help (default: false)
```

---

## interface-add-for-switch

```bash
USAGE:
   usacloud server interface-add-for-switch [command options] <ID or Name(only single target)>

OPTIONS:
   --assumeyes, -y                              assume that the answer to any question which would be asked is yes (default: false)
   --default-route value, --gateway value       set default gateway
   --ipaddress value, --ip value                set ipaddress
   --nw-masklen value, --network-masklen value  set ipaddress  prefix (default: 24)
   --switch-id value                            [Required] set connect switch ID (default: 0)
   --without-disk-edit                          set skip edit-disk flag. if true, don't call DiskEdit API after interface added (default: false)
   --help, -h                                   show help (default: false)
```

---

## interface-add-disconnected

```bash
USAGE:
   usacloud server interface-add-disconnected [command options] <ID or Name(only single target)>

OPTIONS:
   --assumeyes, -y  assume that the answer to any question which would be asked is yes (default: false)
   --help, -h       show help (default: false)
```

---

# == ISOイメージ(CD-ROM) ==

---

## iso-info

```bash
USAGE:
   usacloud server iso-info [command options] <ID or Name(only single target)>

OPTIONS:
 === Common options ===
   --help, -h  show help (default: false)
```

---

## iso-insert

```bash
USAGE:
   usacloud server iso-insert [command options] <ID or Name(only single target)>

OPTIONS:
   --assumeyes, -y                    assume that the answer to any question which would be asked is yes (default: false)
   --description value, --desc value  set resource description
   --icon-id value                    set Icon ID (default: 0)
   --iso-file value                   set iso image file
   --iso-image-id value               set iso-image ID (default: 0)
   --name value                       set resource display name
   --size value                       set iso size(GB) (default: 5)
   --tags value                       set resource tags
   --help, -h                         show help (default: false)
```

---

## iso-eject

```bash
USAGE:
   usacloud server iso-eject [command options] <ID or Name(only single target)>

OPTIONS:
   --assumeyes, -y  assume that the answer to any question which would be asked is yes (default: false)
   --help, -h       show help (default: false)
```

---

# == 監視/モニタリング ==

---

## monitor-cpu

```bash
USAGE:
   usacloud server monitor-cpu [command options] <ID or Name(allow multiple target)>

OPTIONS:
 === Common options ===
   --end value         set end-time
   --key-format value  [Required] set monitoring value key-format (default: "sakuracloud.{{.ID}}.cpu")
   --start value       set start-time
   --help, -h          show help (default: false)
```

---

## monitor-nic

```bash
USAGE:
   usacloud server monitor-nic [command options] <ID or Name(allow multiple target)>

OPTIONS:
 === Common options ===
   --end value         set end-time
   --index value       target index(es)
   --key-format value  [Required] set monitoring value key-format (default: "sakuracloud.{{.ID}}.nic.{{.Index}}")
   --start value       set start-time
   --help, -h          show help (default: false)
```

---

## monitor-disk

```bash
USAGE:
   usacloud server monitor-disk [command options] <ID or Name(allow multiple target)>

OPTIONS:
 === Common options ===
   --end value         set end-time
   --index value       target index(es)
   --key-format value  [Required] set monitoring value key-format (default: "sakuracloud.{{.ID}}.disk.{{.Index}}")
   --start value       set start-time
   --help, -h          show help (default: false)
```


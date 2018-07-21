# Usacloud

![usacloud_logo_h.png](build_docs/docs/images/usacloud_logo_h.png)

[`usacloud`](https://github.com/sacloud/usacloud) is a CLI client for [Sakura Cloud](http://cloud.sakura.ad.jp/index.html).

[![Build Status](https://travis-ci.org/sacloud/usacloud.svg?branch=master)](https://travis-ci.org/sacloud/usacloud)
[![Build status](https://ci.appveyor.com/api/projects/status/pt8g7b9ht3t8ohox?svg=true)](https://ci.appveyor.com/project/sacloud-bot/usacloud)
[![Slack](https://slack.usacloud.jp/badge.svg)](https://slack.usacloud.jp/)

## Major features

- Following the latest function of Sakura Cloud. Billing Information, Object Storage, Web Accelerator, etc. can also be used.
- Support for cross-platform version　(Windows / macOS / Linux). It can also be executed with ARM.
- Single binary implemented in Go language. Installation method is only binary copy (yum / apt / brew is also supported).
- Supports SSH / SCP / SFTP / VNC etc. only with "binary".

## Install

### macOS(`brew`) / Linux(`apt` or `yum` or `brew`) / bash on Windows(Ubuntu)

    curl -fsSL https://releases.usacloud.jp/usacloud/repos/install.sh | bash

### Windows(`chocolatey`)

    choco install usacloud

> The [usacloud package of chocolatey](https://chocolatey.org/packages/usacloud) is maintained by @223n.

### Windows(Other than `chocolatey`) / For other environments

Please downloaded the binary file from the following link.

After expanding the file, please it in an arbitrary folder.

(It is convenient to set PATH.)

- Windows 64bit : [https://releases.usacloud.jp/usacloud/repos/windows/usacloud_windows-amd64.zip](https://releases.usacloud.jp/usacloud/repos/windows/usacloud_windows-amd64.zip)
- Windows 32bit : [https://releases.usacloud.jp/usacloud/repos/windows/usacloud_windows-386.zip](https://releases.usacloud.jp/usacloud/repos/windows/usacloud_windows-386.zip)
- Other : [https://github.com/sacloud/usacloud/releases/latest/](https://github.com/sacloud/usacloud/releases/latest/)

If `bash_completion` is available, you can use it with `usacloud` by the following commands.

```bash
curl -s -L https://releases.usacloud.jp/usacloud/contrib/completion/bash/usacloud >> ~/.bashrc
```

> ※To active `bash_completion`, execute it and do re-login.

## Initial setting

Set the API key using `usacloud config` command.

(The API key setting is saved as a file in `~/.usacloud/default/config.json`.)

```bash
    $ usacloud config

    Setting SakuraCloud API Token => 
    	Enter token: [ENTER YOUR_API_TOKEN]

    Setting SakuraCloud API Secret => 
    	Enter secret: [ENTER YOUR_API_SECRET]
    	
    Setting SakuraCloud Zone => 
    	Enter zone[is1a/is1b/tk1a/tk1v](default:tk1a): [ENTER ZONE]
   
    Written your settings to ~/.usacloud_config
```

You can confirm the API key setting with `usacloud config --show` command.

```bash
   $ usacloud config show
   
   token  = [YOUR_API_TOKEN]
   secret = [YOUR_API_SECRET]
   zone   = [YOUR_ZONE]
   
```

Note: The API key can also be set using environment variables.

```bash
   $ export SAKURACLOUD_ACCESS_TOKEN=[YOUR_API_TOKEN]
   $ export SAKURACLOUD_ACCESS_TOKEN_SECRET=[YOUR_API_SECRET]
   $ export SAKURACLOUD_ZONE=tk1v
```

For details, refer to [usacloud documentation](https://sacloud.github.io/usacloud/).
   
### usage

Refer to [usacloud documentation:Basic usage](https://sacloud.github.io/usacloud/basic_usage/).

```bash
NAME:
   usacloud - CLI client for SakuraCloud

USAGE:
   usacloud [global options] resource command [command options] [arguments...]

VERSION:
   NN.NN.NN, build xxxxxx

COMMANDS:
   config, profile                  A manage command of APIKey settings
   auth-status                      A manage commands of AuthStatus
   private-host                     A manage commands of PrivateHost
   server                           A manage commands of Server
   archive                          A manage commands of Archive
   auto-backup                      A manage commands of AutoBackup
   disk                             A manage commands of Disk
   iso-image                        A manage commands of ISOImage
   bridge                           A manage commands of Bridge
   interface                        A manage commands of Interface
   internet                         A manage commands of Internet
   ipv4                             A manage commands of IPv4
   ipv6                             A manage commands of IPv6
   packet-filter                    A manage commands of PacketFilter
   switch                           A manage commands of Switch
   database                         A manage commands of Database
   load-balancer                    A manage commands of LoadBalancer
   nfs                              A manage commands of NFS
   vpc-router                       A manage commands of VPCRouter
   dns                              A manage commands of DNS
   gslb                             A manage commands of GSLB
   simple-monitor                   A manage commands of SimpleMonitor
   icon                             A manage commands of Icon
   license                          A manage commands of License
   ssh-key                          A manage commands of SSHKey
   startup-script, note             A manage commands of StartupScript
   bill                             A manage commands of Bill
   object-storage, ojs              A manage commands of ObjectStorage
   web-accel                        A manage commands of WebAccel
   price, public-price              A manage commands of Price
   product-disk, disk-plan          A manage commands of ProductDisk
   product-internet, internet-plan  A manage commands of ProductInternet
   product-license, license-info    A manage commands of ProductLicense
   product-server, server-plan      A manage commands of ProductServer
   region                           A manage commands of Region
   zone                             A manage commands of Zone
   summary                          Show summary of resource usage

GLOBAL OPTIONS:
   --token value                    API Token of SakuraCloud (default: none) [$SAKURACLOUD_ACCESS_TOKEN]
   --secret value                   API Secret of SakuraCloud (default: none) [$SAKURACLOUD_ACCESS_TOKEN_SECRET]
   --zone value                     Target zone of SakuraCloud (default: tk1a) [$SAKURACLOUD_ZONE]
   --config value, --profile value  Config(Profile) name [$USACLOUD_PROFILE]
   --help, -h                       show help (default: false)
   --version, -v                    print the version (default: false)

COPYRIGHT:
   Copyright (C) 2017 Kazumichi Yamamoto.
```


#### Examples: List display / search

```bash
    # Show all
    $ usacloud switch ls 
   
    # List items that include `example` in their names.
    $ usacloud switch ls --name example
    
    # Specify sorting conditions(Ascending by Name, Descending by ID)
    $ usacloud switch ls --sort Name --sort -ID
    
    # Specify Limit/Offset(Display up to 5 items, display from the second item)
    $ usacloud switch ls --max 5 --from 2
    
```

#### Examples: CRUD operation

```bash
    # Create
    $ usacloud switch create --name "Example" --desc "description" --tags "Tag1" --tags "Tag2"
     
    # Read
    $ usacloud switch read <ID or Name>
   
    # Update
    $ usacloud switch update --name "Example-update" <ID or Name>
    
    # Delete
    $ usacloud switch delete <ID or Name>
    
```

#### Examples: Create server

```bash
    # Build a server with CentOS installed.
    $ usacloud server build \
             --name server01 \               # Server name
             --os-type centos \              # OS type (Specify a public archive)
             --hostname server01 \           # Host name
             --password "$YOUR_PASSWORD" \   # Administrator password
             --ssh-key-mode generate \       # SSH public key (Generate on the cloud)
             --disable-pw-auth               # Disable password and challenge response authentication when connecting to SSH

    # generated private-key is saved to ~/.ssh/sacloud_pkey_[ServerID]
```

#### Examples: Server operation(Power supply related)

```bash
    # Start up
    $ usacloud server boot <ID or Name>
    
    # Shutdown(graceful)
    $ usacloud server shutdown <ID or Name>
    
    # Shutdown(force)
    $ usacloud server shutdown-force <ID or Name>
    
    # Reset(hard)
    $ usacloud server reset <ID or Name>
    
```

#### Examples: Connect to server(SSH/SCP/VNC/Remote desktop)

```bash
    # Connect to the server via SSH
    # By default, if `~/.ssh/sacloud_pkey_[Server ID]` file exists, it will be used as a secret key.(It can also be specified with the -i option.)
    $ usacloud server ssh <ID or Name> 
    
    # Execute arbitrary command after SSH connection to server.(It can also be executed from the Windows command prompt.)
    $ usacloud server ssh-exec <ID or Name>  cat /etc/passwd
    
    # Upload/Download with SCP
    $ usacloud server scp local-file.txt [ID or Name]:/home/ubuntu/remote-file.txt # From local to remote
    $ usacloud server scp [ID or Name]:/home/ubuntu/remote-file.txt local-file.txt # From remote to local
    $ usacloud server scp -r local-dir [ID or Name]:/home/ubuntu/remote-dir        # Process the directory recursively.
   
    # Connect to the server with VNC using OS default VNC client.
    # (For Windows, you need to associate the appropriate VNC client with the `.vnc` extension.)
    $ usacloud server vnc <ID or Name>

    # Use Remote Desktop to connect to the server using OS default RDP client.
    $ usacloud server remote-desktop <ID or Name>
    # or
    $ usacloud server rdp <ID or Name>
```


#### Examples: Upload/Download by FTPS(Archive/ISO image)

```bash
    # Upload ISO image
    $ usacloud iso-image create --name example --iso-file example.iso
    
    # Download archive (You can download only my archives.)
    $ usacloud archive download --file-destination example.img <ID or Name> 
    
```

#### Examples: Billing related

```bash
    # List of billing information
    $ usacloud bill list
    
    # Output billing CSV file output
    $ usacloud bill csv [BillID]

``` 

#### Examples: Operation of object storage

```bash
    # Set API key for object storage (per bucket)
    $ export SACLOUD_OJS_ACCESS_KEY_ID="[YOUR_BUCKET_ACCESS_KEY]"
    $ export SACLOUD_OJS_SECRET_ACCESS_KEY="[YOUR_BUCKET_SECRET_KEY]"
    
    # List of object storage
    $ usacloud object-storage ls 
    $ usacloud object-storage ls dir1/dir2

    # Download object
    $ usacloud object-storage get remote.txt           # To standard output
    $ usacloud object-storage get remote.txt local.txt # To local file
    $ usacloud object-storage get -r remote/ local/    # Process directories recursively

    # Upload object
    $ usacloud object-storage put local.txt remote.txt 
    $ usacloud object-storage put local.txt dir1/dir2/remote.txt
    $ usacloud object-storage put -r local/ remote/    # Process directories recursively
    
    # Delete object
    $ usacloud object-storage del remote.txt
    $ usacloud object-storage del -r remote/           # Process directories recursively
```

#### Examples: Web accelerator

```bash
    # Delete cache on web accelerator
    $ usacloud web-accel purge https://example.com https://foobar.com

```

#### Examples: Definition of output

```bash
    # Table format (Default)
    $ usacloud switch ls

    # JSON format
    $ usacloud switch ls --output-type json

    # CSV/TSV format
    $ usacloud switch ls --output-type csv # or tsv

    # Specify the column to output (It can be specified when outputting in CSV / TSV format.)
    $ usacloud switch ls --output-type tsv --col ID --col Name
   
    # Output ID or key only
    $ usacloud swtich ls -q # or --quiet
    
    # Use custom golang templates to output
    $ usacloud switch ls --format "ID is '{{.ID}}', Name is '{{.Name}}'"
    ID is '123456789012', Name is 'example'
```

#### Examples: Definition of output (For monitoring tools)

```bash

    # target resource ID = 123456789012

    # for munin
    $ usacloud internet monitor --format "target.value {{.In}}" 123456789012
    
    # for zabbix_sender(zabbix_hostname=router01 , item_key=packet.in)
    $ usacloud internet monitor --format "router01 packet.in {{.UnixTime}} {{.In}}" 123456789012 \
         | zabbix_sender -z your.zabbix.hostname -p 10051 -T -i -
    
    # for sensu/mackerel
    $ OUTPUT_FORMAT=`echo -e "{{.Key}}\t{{.In}}\t{{.UnixTime}}"`
    $ usacloud internet monitor --format "$OUTPUT_FORMAT" 123456789012
```

### Examples: Use multiple API keys (Profile function)

```bash
    # List
    $ usacloud config list

    # Create profile (Interactive format)
    $ usacloud config edit your-profile-name1

    # Create profile (Non-interactive format)
    $ usacloud config edit --zone "is1a" --token "token" --secret "secret" your-profile-name1
    
    # Show profile contents
    $ usacloud config show your-profile-name1

    # Display the name of the currently selected profile
    $ usacloud config current

    # Switch profile
    $ usacloud config use your-profile-name1

    # Delete profile
    $ usacloud config delete your-profile-name1
```

## Develop

#### Build

    $ make build

#### Build (Build on Docker / Build for cross platform.)
    
    $ make docker-build
        
#### Test
    
    $ make test
    
#### Integration test

To execute the integration test, you need to set the following environment variables.

- `SAKURACLOUD_ACCESS_TOKEN`
- `SAKURACLOUD_ACCESS_TOKEN_SECRET`
- `SAKURACLOUD_ZONE`

```bash
# When running on the local machine (bats/jq required)
$ make integration-test
# If you want to run only individual tests (You can also test by directory or file.)
$ test/integration/run_bats.sh test/integration/bats/[Target directory or target file name]

# When running on Docker (No need for bats/jq)
$ make docker-integration-test
```

If you want to add a test, see [test/integration/README.md](test/integration/README.md).

#### Generate the source of each command

```bash
    $ make gen
    $ # or
    $ make gen-force
```
    
#### Add a new resource / command

Create a definition file under `define`.
Execute the `make gen-force` command to generate the source.

#### Document

The document uses Github Pages. (It is under the "docs" directory of the master branch.)

You can generate static files with the `mkdocs` command.

**For document PR, please modify only under "build_docs" directory.**

**Please do not change under "docs" directory.**

**The "docs" directory is updated in bulk at the time of release.**

```bash
    # Start document preview server (Preview is possible with `http://localhost/`.)
    make serve-docs
    
    # Document validation (textlint)
    make lint-docs
```

## License

 `usacloud` Copyright (C) 2017-2018 Kazumichi Yamamoto.

  This project is published under [Apache 2.0 License](LICENSE.txt).
  
## Author

  * [Usacloud Authors](AUTHORS)

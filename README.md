# Usacloud

[![Build Status](https://travis-ci.org/sacloud/usacloud.svg?branch=master)](https://travis-ci.org/sacloud/usacloud)

## Overview

:rabbit:[`usacloud`](https://github.com/sacloud/usacloud) is New CLI client for the [SakuraCloud](http://cloud.sakura.ad.jp/index.html).  
That means **"Unofficial + [`sacloud`](https://github.com/sakura-internet/node-sacloud)"**.  

## Key Features

- Cover the latest features of the SakuraCloud
- Cross Platform(Windows/macOS/Linux) support(on Linux version, supports ARM!)
- Includes Upload of [Archive/ISO-image](http://cloud.sakura.ad.jp/specification/server-disk/#server-disk-content03) by FTPS
- Includes [WebAccelerator](http://cloud.sakura.ad.jp/specification/web-accelerator/) support
- Includes [ObjectStorage](http://cloud.sakura.ad.jp/specification/object-storage/) support
- Includes [Billing API](http://cloud-news.sakura.ad.jp/billapi/) support

## Install

### RHEL / CentOS

    curl -fsSL https://usacloud.b.sakurastorage.jp/repos/setup-yum.sh | sh

### Ubuntu / debian / bash on Windows

    curl -fsSL https://usacloud.b.sakurastorage.jp/repos/setup-apt.sh | sh

### Others

#### Using docker

    alias usacloud="docker run -it --rm sacloud/usacloud" 

#### Using docker with [`whalebrew`](https://github.com/bfirsh/whalebrew)
    whalebrew install sacloud/usacloud

#### Manual install

Download binary file from [Release page](https://github.com/sacloud/usacloud/releases/latest). 

**[OPTION]** Enable bash-completion

    $ eval "`curl -s -L https://usacloud.b.sakurastorage.jp/contrib/completion/bash/usacloud`"
    # or
    $ curl -s -L https://usacloud.b.sakurastorage.jp/contrib/completion/bash/usacloud >> ~/.bashrc

## Setting

Set API key and secret by `usacloud config` command.  
(Your API key settings will write to `~/.usacloud_config`)

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

If you want to confirm settings, use `usacloud config --show`.

```bash
   $ usacloud config --show
   
   token=[YOUR_API_TOKEN]
   secret=[YOUR_API_SECRET]
   zone=[YOUR_ZONE]
   
```

Note: API key and secret can be also set using environment variables.

```bash
   $ export SAKURACLOUD_ACCESS_TOKEN=[YOUR_API_TOKEN]
   $ export SAKURACLOUD_ACCESS_TOKEN_SECRET=[YOUR_API_SECRET]
   $ export SAKURACLOUD_ZONE=tk1v
```
   
### Usage

```bash
NAME:
   usacloud - Unofficial 'sacloud' - CLI client of the SakuraCloud

USAGE:
   usacloud [global options] command [command options] [arguments...]

VERSION:
   0.0.5, build xxxxxxx

COMMANDS:
   config                           A manage command of APIKey settings
   server                           A manage commands of Server
   archive                          A manage commands of Archive
   auto-backup                      A manage commands of AutoBackup
   disk                             A manage commands of Disk
   iso-image                        A manage commands of ISOImage
   bridge                           A manage commands of Bridge
   interface                        A manage commands of Interface
   internet                         A manage commands of Internet
   packet-filter                    A manage commands of PacketFilter
   switch                           A manage commands of Switch
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
   help, h                          Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --token value   API Token of SakuraCloud (default: none) [$SAKURACLOUD_ACCESS_TOKEN]
   --secret value  API Secret of SakuraCloud (default: none) [$SAKURACLOUD_ACCESS_TOKEN_SECRET]
   --zone value    Target zone of SakuraCloud (default: tk1a) [$SAKURACLOUD_ZONE]
   --help, -h      show help (default: false)
   --version, -v   print the version (default: false)

COPYRIGHT:
   Copyright (C) 2017 Kazumichi Yamamoto.
```   


#### Examples: List/Find/Search resource 

```bash
    # list all switches
    $ usacloud switch ls 
   
    # list switches with name "example"
    $ usacloud switch ls --name example
    
    # list switches order by Name(asc) and ID(desc)
    $ usacloud switch ls --sort Name --sort -ID
    
    # list switches with limit=5 and offset=2
    $ usacloud switch ls --max 5 --from 2
    
```   

#### Examples: Basic CRUD

```bash
    # create switch
    $ usacloud switch create --name "Example" --desc "description" --tags "Tag1" --tags "Tag2"
     
    # show switch detail
    $ usacloud switch read <ID or Name>
   
    # update switch
    $ usacloud switch update --name "Example-update" <ID or Name>
    
    # delete switch
    $ usacloud switch delete <ID or Name>
    
```  

#### Examples: Create Server 

```bash
    # Build server from Public Archive(CentOS)
    $ usacloud server build \
             --name server01 \               # ServerName
             --os-type centos \              # OSType(use PublicArchive)
             --hostname server01 \           # for DiskEdit parameter
             --password "$YOUR_PASSWORD" \   # for DiskEdit parameter
             --ssh-key-mode generate \       # generate ssh-key
             --ssh-key-name key01 \          # generate ssh-key name
             --disable-pw-auth               # disable password auth

    # generated private-key is saved to ~/.ssh/sacloud_pkey_[ServerID]
```   

#### Examples: Manipulate Server

```bash
    # boot
    $ usacloud server boot <ID or Name>
    
    # shutdown(graceful)
    $ usacloud server shutdown <ID or Name>
    
    # shutdown(force)
    $ usacloud server shutdown-force <ID or Name>
    
    # reset(hard)
    $ usacloud server reset <ID or Name>
    
```

#### Examples: Connect to server(SSH/SCP/VNC)

```bash
    # connect to server by ssh using generated private-key(Not supported on Windows)
    $ usacloud server ssh <ID or Name>
    
    # exec command on SSH(no-pty, support Windows)
    $ usacloud server ssh-exec <ID or Name> cat /etc/passwd
    
    # upload/download by SSH(like scp)
    $ usacloud server scp local-file.txt [ServerID]:/home/ubuntu/remote-file.txt # local to remote
    $ usacloud server scp [ServerID]:/home/ubuntu/remote-file.txt local-file.txt # remote to local
    $ usacloud server scp -r local-dir [ServerID]:/home/ubuntu/remote-dir        # recursive
    
    # open VNC client using the OS's default application
    # In Windows, it is need to associate vnc client with .vnc extension
    $ usacloud server vnc <ID or Name>
```

#### Examples: Upload/Download iso-image or archive by FTPS

```bash
    # upload iso-image
    $ usacloud iso-image create --name example --iso-file example.iso
    
    # download archive(Only MyArchive can be downloaded)
    $ usacloud archive download --file-destination example.img <ID or Name>
    
```

#### Examples: Billing(download csv)

```bash
    # list bill
    $ usacloud bill list
    
    # download bill-detail by CSV
    $ usacloud bill csv [BillID]

``` 

#### Examples: Object Storage

```bash
    # set API key of Object Storage(per bucket)
    $ export SACLOUD_OJS_ACCESS_KEY_ID="[YOUR_BUCKET_ACCESS_KEY]"
    $ export SACLOUD_OJS_SECRET_ACCESS_KEY="[YOUR_BUCKET_SECRET_KEY]"
    
    # list objects
    $ usacloud object-storage ls 
    $ usacloud object-storage ls dir1/dir2

    # download object to local(download remote.txt)
    $ usacloud object-storage get remote.txt           # output to os.StdOut
    $ usacloud object-storage get remote.txt local.txt # save as local file
    $ usacloud object-storage get -r remote/ local/    # download recursive

    # upload object
    $ usacloud object-storage put local.txt remote.txt
    $ usacloud object-storage put local.txt dir1/dir2/remote.txt
    $ usacloud object-storage put -r local/ remote/    # upload recursive
    
    # delete object
    $ usacloud object-storage del remote.txt
    $ usacloud object-storage del -r remote/           # delete recursive

```   

#### Examples: Web Accelerator(purge cache)

```bash
    # purge cache on web-accel
    $ usacloud web-accel purge https://example.com https://foobar.com

```   

#### Examples: Output format

```bash
    # output table(default)
    $ usacloud switch ls

    # output JSON
    $ usacloud switch ls --output-type json

    # output csv/tsv
    $ usacloud switch ls --output-type csv # or tsv

    # output csv/tsv + include columns
    $ usacloud switch ls --output-type tsv --col ID --col Name
   
    # output QuietMode(output ID/Key only)
    $ usacloud swtich ls -q # or --quiet
    
    # output custom format(using text/template style template)
    $ usacloud switch ls --format "ID is '{{.ID}}', Name is '{{.Name}}'"
    ID is '123456789012', Name is 'example'
```

#### Examples: Output format for Monitoring

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


## Development

#### Build(includes src generate)

    $ make build
    
#### Generate each command source 

    $ make gen
    $ # or
    $ make gen-force
    
#### Add new resource or command

Edit under the `define` package.


## License

 `usacloud` Copyright (C) 2017 Kazumichi Yamamoto.

  This project is published under [Apache 2.0 License](LICENSE.txt).
  
## Author

  * Kazumichi Yamamoto ([@yamamoto-febc](https://github.com/yamamoto-febc))

# Usacloud

This project is work in progress!!

## Overview

:rabbit:[`usacloud`](https://github.com/sacloud/usacloud) is New CLI client for the [SakuraCloud](http://cloud.sakura.ad.jp/index.html).  
That means **"Unofficial + [`sacloud`](https://github.com/sakura-internet/node-sacloud)"**.  
Written by Go(golang) from scratch.

## Key Features

- Cover the latest features of the SakuraCloud
- Cross Platform(Windows/macOS/Linux) support(on Linux version, supports ARM!)
- Includes Upload of [Archive/ISO-image](http://cloud.sakura.ad.jp/specification/server-disk/#server-disk-content03) by FTPS
- Includes [WebAccelerator](http://cloud.sakura.ad.jp/specification/web-accelerator/) support
- Includes [ObjectStorage](http://cloud.sakura.ad.jp/specification/object-storage/) support
- Includes [Billing API](http://cloud-news.sakura.ad.jp/billapi/) support

## Install

   Download binary file from [Release page](https://github.com/sacloud/usacloud/releases/latest). 

#### Enable bash-completion

    $ eval "`curl -s -L https://usacloud.b.sakurastorage.jp/contrib/completion/bash/usacloud`"
    # or
    $ curl -s -L https://usacloud.b.sakurastorage.jp/contrib/completion/bash/usacloud >> ~/.bashrc

## Usage

### Setting API Keys

```bash
    # set API Key to environment variables
    $ export SAKURACLOUD_ACCESS_TOKEN=[YOUR_API_TOKEN]
    $ export SAKURACLOUD_ACCESS_TOKEN_SECRET=[YOUR_API_SECRET]

    # set default zone
    $ export SAKURACLOUD_DEFAULT_ZONE=tk1v
```
   
#### Usage

```bash
NAME:
   usacloud - Unofficial 'sacloud' - CLI client of the SakuraCloud

USAGE:
   usacloud [global options] command [command options] [arguments...]

VERSION:
   0.0.1, build xxxxxxx

COMMANDS:
     archive                          A manage commands of Archive
     auto-backup                      A manage commands of AutoBackup
     bill                             A manage commands of Bill
     bridge                           A manage commands of Bridge
     disk                             A manage commands of Disk
     dns                              A manage commands of DNS
     gslb                             A manage commands of GSLB
     icon                             A manage commands of Icon
     interface                        A manage commands of Interface
     internet                         A manage commands of Internet
     iso-image                        A manage commands of ISOImage
     license                          A manage commands of License
     object-storage, ojs              A manage commands of ObjectStorage
     packet-filter                    A manage commands of PacketFilter
     price, public-price              A manage commands of Price
     product-disk, disk-plan          A manage commands of ProductDisk
     product-internet, internet-plan  A manage commands of ProductInternet
     product-license, license-info    A manage commands of ProductLicense
     product-server, server-plan      A manage commands of ProductServer
     region                           A manage commands of Region
     server                           A manage commands of Server
     simple-monitor                   A manage commands of SimpleMonitor
     ssh-key                          A manage commands of SSHKey
     startup-script, note             A manage commands of StartupScript
     switch                           A manage commands of Switch
     web-accel                        A manage commands of WebAccel
     zone                             A manage commands of Zone
     help, h                          Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --format value, --output-format value                    Output format[table/json/csv/tsv] (default: table)
   --secret value, --sakuracloud-access-token-secret value  API Secret of SakuraCloud (default: none) [$SAKURACLOUD_ACCESS_TOKEN_SECRET]
   --token value, --sakuracloud-access-token value          API Token of SakuraCloud (default: none) [$SAKURACLOUD_ACCESS_TOKEN]
   --trace, --sakuracloud-trace-mode                        Flag of SakuraCloud debug-mode (default: false) [$SAKURACLOUD_TRACE_MODE]
   --zone value, --sakuracloud-default-zone value           Target zone of SakuraCloud (default: tk1a) [$SAKURACLOUD_ZONE]
   --help, -h                                               show help (default: false)
   --version, -v                                            print the version (default: false)

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
    
    # list switches with limit=5,offset=2
    $ usacloud switch ls --max 5 --from 2
    
```   

#### Examples: Basic CRUD

```bash
    # create switch
    $ usacloud switch create --name "Example" --desc "description" --tags "Tag1" --tags "Tag2"
     
    # show switch detail
    $ usacloud switch read [SwitchID]
   
    # update switch
    $ usacloud switch update --name "Example-update" [SwitchID]
    
    # delete switch
    $ usacloud switch delete [SwitchID]
    
```  

#### Examples: Create Server 

```bash
    # Build server from Public Archive(CentOS)
    $ usacloud server build \
             --name server01 \               # ServerName
             --os-type centos \              # OSType(use PublicArchive)
             --hostname server01 \           # for DiskEdit parameter
             --password "$YOUR_PASSWORD" \   # for DiskEdit parameter
             --ssh-key-mode generatea \      # generate ssh-key
             --ssh-key-name key01 \          # generate ssh-key name
             --disable-pw-auth               # disable password auth

    # generated private-key is saved to ~/.ssh/sacloud_pkey_[ServerID]
```   

#### Examples: Connect to server by SSH

```bash
    # connect to server by ssh using generated private-key(Not supported on Windows)
    $ usacloud server ssh [ServerID]
    
    # [not implemented yet]exec command by SSH
    $ usacloud server ssh-exec cat /etc/passwd
    
    # [not implemented yet] upload/download by SSH(like scp)
    $ usacloud server scp local-file.txt [ServerID]:/home/ubuntu/remote-file.txt # local to remote
    $ usacloud server scp [ServerID]:/home/ubuntu/remote-file.txt local-file.txt # remote to local
```

#### Examples: Upload/Download iso-image or archive by FTPS

```bash
    # upload iso-image
    $ usacloud iso-image create --name example --iso-file example.iso
    
    # download archive(Only MyArchive can be downloaded)
    $ usacloud archive download --file-destination example.img [MyArchiveID]
    
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
    $ usacloud get remote.txt           # output to os.StdOut
    $ usacloud get remote.txt local.txt # save as local file

    # upload object
    $ usacloud put local.txt remote.txt
    $ usacloud put local.txt dir1/dir2/remote.txt
    

```   

#### Examples: Web Accelerator(purge cache)

```bash
    # purge cache on web-accel
    $ usacloud web-accel purge https://example.com https://foobar.com

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

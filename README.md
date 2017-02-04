# Usacloud

This project is work in progress!!

## Overview

[`usacloud`](https://github.com/sacloud/usacloud) is New CLI client for the [SakuraCloud](http://cloud.sakura.ad.jp/index.html).  
That means "Unofficial + [`sacloud`](https://github.com/sakura-internet/node-sacloud)".  
Written by Go(golang) from scratch.

## Key Features

- Cover the latest features of SakuraCloud
- Cross Platform(Windows/macOS/Linux) support(on Linux version, supports ARM!)
- Includes Upload of [Archive/ISO-image](http://cloud.sakura.ad.jp/specification/server-disk/#server-disk-content03) by FTPS
- Includes [WebAccelerator](http://cloud.sakura.ad.jp/specification/web-accelerator/) support
- Includes [ObjectStorage](http://cloud.sakura.ad.jp/specification/object-storage/) support
- Includes [Billing API](http://cloud-news.sakura.ad.jp/billapi/) support

## Install

    TODO write about how to install

## Development

#### Build(includes src generate)

    $ make build
    
#### Generate each command source 

    $ make gen
    $ # or
    $ make gen-force
    
#### Add new resource or command

Edit under the `define` package.
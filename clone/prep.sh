#!/bin/bash
yum install git
useradd -m  happyleaf
passwd happyleaf
ssh-keygen -t rsa -b 4096 -C "witcxc@gmail.com"
git clone git@github.com:witcxc/funnycode.git
yum install go
go get github.com/shadowsocks/shadowsocks-go/cmd/shadowsocks-server
make all

nohup ./shadowsocks-server -c config.json > shadow_out.log 2>&1 &
#{
#    "server":"0.0.0.0",
#    "server_port":8389,
#    "local_port":1080,
#    "local_address":"127.0.0.1",
#    "password":"XXXX",
#    "method": "aes-256-ctr",
#    "timeout":600
#}

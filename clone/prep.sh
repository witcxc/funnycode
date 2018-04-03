#!/bin/bash
yum install git
ssh-keygen -t rsa -b 4096 -C "witcxc@gmail.com"

useradd -m  happyleaf
passwd happyleaf
visudo
# add
#username        ALL=(ALL)      ALL

git clone git@github.com:witcxc/funnycode.git
yum install go
go get github.com/shadowsocks/shadowsocks-go/cmd/shadowsocks-server
make all

nohup ./shadowsocks-server -c config.json > shadow_out.log 2>&1 &
sudo bash -c su
nice -n -20 nohup ./serv.bin -c config.json -d

#python
wget https://www.python.org/ftp/python/2.7.14/Python-2.7.14.tgz
tar xzf Python-2.7.14.tgz
cd Python-2.7.14
./configure --prefix=/home/happyleaf/python --enable-optimizations --enable-unicode
make
make install

#virtual env
curl -O https://pypi.python.org/packages/source/v/virtualenv/virtualenv-X.X.tar.gz
tar xvfz virtualenv-X.X.tar.gz
cd virtualenv-X.X
[sudo] python setup.py install

#vim
sudo yum install ncurses-devel
sudo yum install -y ruby ruby-devel lua lua-devel luajit \
    luajit-devel ctags git python python-devel \
    python3 python3-devel tcl-devel \
    perl perl-devel perl-ExtUtils-ParseXS \
    perl-ExtUtils-XSpp perl-ExtUtils-CBuilder \
    perl-ExtUtils-Embed

git clone https://github.com/vim/vim.git
cd vim/
export PATH=/home/happyleaf/python/bin/:$PATH
make distclean
./configure --with-features=huge \
            --enable-multibyte \
            --enable-rubyinterp=yes \
            --enable-pythoninterp=yes \
            --with-python-config-dir=/home/happyleaf/python/lib/python2.7/config \
            --enable-perlinterp=yes \
            --enable-luainterp=yes \
            --enable-cscope \
            --prefix=/usr/local
#            --enable-python3interp=yes \
#            --with-python3-config-dir=/usr/lib/python3.5/config \

#vim Makefile
make
sudo make install

#{
#    "server":"0.0.0.0",
#    "server_port":8389,
#    "local_port":1080,
#    "local_address":"127.0.0.1",
#    "password":"XXXX",
#    "method": "aes-256-ctr",
#    "timeout":600
#}

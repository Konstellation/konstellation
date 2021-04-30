#!/bin/sh

if test -n "$1"; then
    # need -R not -r to copy hidden files
    cp -R "$1/.knstld" /root
fi

mkdir -p /root/log
knstld start --rpc.laddr tcp://0.0.0.0:26657 --trace --log_level info

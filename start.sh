#! /bin/bash

set -e

cd /usr/bin

if [ ! -x crawl ];then
    chmod +x crawl

cd /home/crawl

/usr/bin/crawl -_ui=cmd -a_mode=0 -c_spider=0 -a_outtype=mysql -a_thread=10 -a_dockercap=30 -a_pause=5000 -a_proxyminute=0 -a_keyins="<crawl><golang>" -a_success=true -a_failure=true
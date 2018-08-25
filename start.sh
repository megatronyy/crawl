#! /bin/bash

set -e

cd /usr/bin

if [ ! -f crawl ];then
    echo "crawl progrom not exists"
    exit 0
fi

# 判断是否有执行权限
if [ ! -x crawl ];then
    # 添加执行权限
    chmod +x crawl
    
fi

# 跳转目录
cd /home/crawl
# 后台启动爬虫
nohup /usr/bin/crawl \
-_ui=cmd \
-a_mode=0 \
-c_spider=0 \
-a_outtype=mysql \
-a_thread=10 \
-a_dockercap=30 \
-a_pause=5000 \
-a_proxyminute=0 \
-a_keyins="<crawl><golang>" \
-a_success=true -a_failure=true &

echo "crawl start..."
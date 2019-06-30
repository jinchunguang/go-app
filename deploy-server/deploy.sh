#!/usr/bin/env bash

# 关闭原服务
kill -9 $(pgrep deploy-web)

# 进入工作目录
cd ~/work/web-app

# 下载代码
git pull https://github.com/jinchunguang/web-app.git

#
cd web-server

#
./deploy-web &


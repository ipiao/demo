#!/usr/bin/env bash

# 监听文件夹改动并自动编译重启
p=`pwd`
fun(){
    echo "开始监听文件夹 $p"
    if inotifywait -rq -e modify,create,move,delete *.go; then 
        echo "执行重启"
        restart
        sleep 5s
        fun
    fi
}

# 重启方法
restart(){
    go build
    if [ $? -eq 0 ] 
    then
        pkill fweb
        nohup ./fweb>>nohup.out &
        echo "执行状态:$?"
        if [ $? -eq 0 ]
        then
            echo "success"
        else
            echo "fail"
        fi
    else
        echo "代码还没写完哦！"
    fi
}

# 入口
fun
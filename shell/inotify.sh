#!/usr/bin/env bash

# 监听文件夹改动并自动编译重启
SRCDIR=`pwd`
dorestart(){
    sleep 2
    echo "restart $1"
    echo "do something"
    echo "restart $1 done"
}

inotifywait -mrd -o inotify.log --timefmt '%d/%m/%y %H:%M' --format '%T %Xe %w %f' \
 -e CLOSE_WRITE,CREATE --excludei *.* $SRCDIR/* | while read DATE TIME EVENT DIR FILE; do
RDIR=${DIR%/}              # 去除末尾/
dir=${RDIR##*/}            # 去除路径前缀
echo "notify $EVENT:$dir-$FILE" 
# if [ "CLOSE_WRITEXCLOSE" == "$EVENT" ] #CLOSE事件不止一次发生,并且暂时不知道怎么判断文件上传结束
if [ "EOF" == "$FILE" ]   # 通过EOF文件标记上传结束
then
   dorestart $dir
fi
done
 
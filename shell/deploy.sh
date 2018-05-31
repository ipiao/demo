#!/usr/bin/env bash

# source base.sh
user="work"
hosts=("test.mys4s.cn")
gitbranch="master_dev"
buildpath="/home/work/s4s_dev/build"

gitbranch="ykk"
srcpath="/home/work/s4s_dev/src/s4s/mservices/drive_school"

for host in ${hosts[*]};do
{
echo $host
# ssh
ssh $user@$host <<EOF
cd $srcpath
git checkout $gitbranch
git pull origin $gitbranch
cd $buildpath
./install.sh drive_school
ps -aux|grep drive_school|grep s4s_dev|grep -v grep|awk '{print \$2}'|xargs kill -9

EOF
if [ $? -eq 0 ]
then
    echo "布置到$host成功！"
else
    echo "布置到$host失败！"
fi
}&
done
wait

echo done

# 批量异步上传二进制文件到服务器，并重启服务器程序
# nohup $buildpath/bin/drive_school -c conf/drive_school.conf >/dev/null 2>&1 &
# supervisorctl restart drive_school_dev
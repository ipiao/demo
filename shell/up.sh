#!/usr/bin/env bash

user="root"
hosts=("192.168.1.99")
dest_path="/home/go/src/thinkdata.cc/factory/datasrv/cmd/fweb/"
dest_file="fweb_up"
src_file="fweb"

# 编译上传
go build
echo ${hosts[*]}
for host in ${hosts[*]};do
{
echo $host
scp $src_file $user@$host:$dest_path$dest_file
# ssh
ssh $user@$host <<EOF
cd $dest_path
sleep 1s
pkill $src_file
rm $src_file
mv $dest_file $src_file
nohup ./$src_file >>nohup.out 2>&1 &
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
# whois
一款跨平台的命令行whois查询工具
# 快速安装
go get -u github.com/xiaoqidun/whois
# 编译安装
```
git clone https://github.com/xiaoqidun/whois.git
cd whois
go build whois.go
```
# 手动安装
1. 根据系统架构下载为你编译好的[二进制文件](https://github.com/xiaoqidun/whois/releases)
2. 将下载好的二进制文件重命名为whois并保留后缀
3. 把whois文件移动到系统PATH环境变量中的目录下
4. windows外的系统需使用chmod命令赋予可执行权限
# 使用说明
```shell script
whois aite.xyz
```
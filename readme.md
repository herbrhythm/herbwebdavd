# Herbwebdavd 轻量级跨平台Webdav服务

## 简介

Herbwebdavd是一款轻量级的Webdav服务。

主要针对的使用场景为

* 虚拟机里宿主机与客户机交换文件
* 内网小范围传文件，作为文件中心使用
* 临时需要上传文件到服务/VPS上，通用http端口，不需要额外安装服务和调整防火墙

特色包括

* 跨平台，支持Linux和Windows系统
* 简单，默认情况下只需要设置目录和帐号密码
* 绿色，可以直接通过复制粘贴在不同的系统下运行
* 简单的权限系统，可以针对不同的目录给不同用户设置权限

## 使用方式

下载最新的Herbwebdavd,解压后执行 解压后文件的bin/herbwebdavd.exe\(Windows\)或bin/herbwebdavd\(Linux\)即可

## 设置目录

修改程序目录下的/config/system.toml文件，格式为Toml.内容为

```toml
#Folders config file
[Folders]
"workspace"="d:/workspace"
```

Folders小节中是网页路径和实际路径的键值对。

* 网页路径 必须是1-64位的大小写英文，数字，可用的符号为'_'和'-'
* 目录必须在系统内存在，而且程序有权限读取
* 实际访问路径为 http://\[HOST\]:\[PORT\]/\[Name\]/ 的格式

## 设置访问口令

口令配置文件是 程序目录的/appdata/appkey.static.toml

程序第一次运行时会自动创建，也可以从程序目录的/system/exampledata/appkey.static.toml复制到该位置

格式为toml，内容为

```toml
[[Apps]]
# ID used for http basicauth username
ID="testid"
# Owner must not be empty
Owner="test"
# Key used for http basicauth password
Key="testkey"
[Apps.Payloads]
# user roles,splite by ,
# For example root user roles 
# roles="root"
# Roles by folder:
# roles="folder:name=folder1,name=folder2,name=folder%202"
roles="root"
```
* 每一个Apps 小节为一个新的用户
* ID是登陆帐号
* Key是登陆密码
* Owner不能为空
* Apps.Payloads中的roles为用户对应的权限
* 权限由分号分割，格式为\[权限名\]:\[属性\]=\[属性值\]，\[属性2\]=\[属性值2\];权限名，属性，属性值可以使用 URLEncode转义，即%20代表空格
* root权限代表超级管理员权限，可以访问所有目录
* folder权限为针对目录设置的权限，通过name属性来指定有权限操作的目录。name为具体的网页前缀

## 默认端口与修改

默认端口为 全局4380端口

如果需要修改，可以复制 程序目录/system/defaultconfig/http.toml到 程序目录/config 下，再进行编辑

## 关于访问安全

处于设计目的，整个webdav服务只有帐号密码验证，没有失败次数限制等额外的防护，因此如果在公网中长期使用是存在一定安全隐患的。

如果需要在公网中使用，建议

* 随用随开，用完关闭
* 通过nginx/apache/caddy之类反代并绑定域名，而非通过端口访问
* 用户名和网页前缀不要设置的过于简单
* 密码设置的不要过短，易猜测

## 开机自动启动

如果需要开机自动启动，可以将herbwebdavd注册为服务

### 注册Windows服务

程序并不支持直接注册为Windows服务，如需要注册，请使用[nssm](https://nssm.cc/)注册服务

### 注册Linux服务

注册Linux服务建议使用systemd来进行操作

程序在 程序目录/system/exampledata/herbwebdavd.service 提供了一个 herbwebdavd.service 的模板
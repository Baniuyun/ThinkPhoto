# ThinkPhoto

## 项目介绍

想拍视频，一个web端的短视频应用。

Github 地址：https://github.com/Baniuyun/ThinkPhoto



### 项目架构



### 技术介绍

前端技术栈：

* npm 包管理工具

* 基于 JavaScript
* Vue3
* Vue Router
* Axios 
* ...

后端技术栈：

* Golang
* go-zero
* JWT
* Protocol Buffers
* etcd
* ZincSearch
* MySQL
* ...



### 目录结构

> 这里简单列出目录结构

代码仓库目录：

```bash
ThinkPhoto
├─go
├─js
├─README.md
├─go.mod
```

简略前端目录

```bash
js
├─public
├─src
├─api
├─assets
├─components
│  ├─channel
│  ├─icons
│  ├─index
│  ├─player
│  ├─search
│  ├─user
│  └─video
├─router
├─store
├─utils
└─views
```

简略后端目录

```
go
 ├─log
 ├─pkg
 │  ├─cors
 │  ├─encrypt
 │  ├─interceptor
 │  ├─jwt
 │  └─xcode
 ├─script
 └─server
 ├─follow
 ├─user
 ├─video
 ├─webapp
 └─Zinc
```



## 环境介绍

### 线上环境

服务器：腾讯云 2核 2G Ubuntu 22.04 LTS

对象存储：七牛云

### 开发环境

| 开发工具 | 说明                |
| -------- | ------------------- |
| WebStorm | 前端开发工具        |
| GoLand   | 后端开发工具        |
| Navicat  | MySQL 远程连接工具  |
| XShell   | 远程连接工具        |
| Apifox   | 接口调试 + 文档生成 |

| 开发环境         | 版本   |
| ---------------- | ------ |
| Golang           | 1.21.3 |
| ZincSearch       | 0.4.9  |
| etcd             | 3.5.10 |
| Protocol Buffers | 3.17.3 |
| MySQL            | 8.x    |



## 项目运行



### 本地运行

Windows10/11 环境

拉取项目前安装 Golang、Node、MySQL、Redis 环境：

Node 安装参考 [Nvm](https://nvm.uihtm.com/) 

Golang 安装参考 [官方文档](https://go.dev/doc/install)

goctl、protoc、go-zero 安装参考 [golang 安装 | go-zero Documentation](https://go-zero.dev/docs/tasks)

etcd 安装参考 [彻底搞懂 etcd 系列文章（二）：etcd 的多种安装姿势 - 知乎 (zhihu.com)](https://zhuanlan.zhihu.com/p/144056143)

Protocol Buffers 安装参考 [windows下安装和配置protobuf流程_protobuf 环境变量-CSDN博客](https://blog.csdn.net/bean_business/article/details/109688219)

ZincSearch 安装参考 [ZincSearch 新手安装教程 - 知乎 (zhihu.com)](https://zhuanlan.zhihu.com/p/657054683)

MySQL 安装参考 [Documentation](https://dev.mysql.com/doc/)

拉取项目到本地：

```bash
git clone https://github.com/Baniuyun/ThinkPhoto.git
```

前端项目运行：

```bash
# 1.进入前端项目根目录
cd ThinkPhoto/js
# 2.安装依赖
npm install
# 3.运行项目
npm run dev
```

后端项目运行：

```bash
# 1.进入后端项目根目录
cd ThinkPhoto/go

# 2.修改项目的配置文件
cd go/server
# 修改 folow/etc、user/api/etc、user/rpc/etc、video/api/etc、video/rpc/etc、webapp/etc、Zinc/etc 这些目录下的 .yaml为结尾的文件，改为自己的地址。

# 3.MySQL 导入 sql 文件(位于 ThinkPhoto/go/script 目录下)
# 在 Navicat 中运行这些脚本文件

# 4.配置 etcd，根据 yaml 文件修改的配置进行 etcd 的键值配置
# 打开 cmd 输入：
etcdctl put follow.rpc 127.0.0.1:8001
etcdctl put user.rpc 127.0.0.1:8002
etcdctl put video.rpc 127.0.0.1:8003
etcdctl put zinc.rpc 127.0.0.1:8004

# 5.配置 ZincSearch

# 6.运行项目
# 先进入根目录
go mod tidy

go run go/server/follow/follow.go
go run go/server/user/rpc/user.go
go run go/server/video/rpc/video.go
go run go/server/Zinc/zinc.go

go run go/server/user/api/userapi.go
go run go/server/video/api/videoapi.go
go run go/server/webapp/webapp.go

```




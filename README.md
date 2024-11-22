<h1 align="center">
  Douyin-Backend
</h1>


该项目基于开源前端项目，采用 Go 语言开发后端，使用 Gin 框架构建，结合 MySQL 进行数据存储，并通过 Redis 实现鉴权 Token 的缓存管理。
同时，集成 Gorse 提供推荐算法服务，基于 WebSocket 实现实时聊天功能。项目功能涵盖短视频的点赞、评论、收藏、分享以及用户关注和实时聊天和等核心互动场景。
后续将继续补充直播和其他功能。

<div>
<img width="150px" src='docs/imgs/vfcfs-95rgz.gif'/>
<img width="150px" src='docs/imgs/50nea-frbnj.gif'/>
<img width="150px" src='docs/imgs/xmg24-2nkbp.gif'/>
<img width="150px" src='docs/imgs/fxklv-5nafx.gif'/>
<img width="150px" src='docs/imgs/1mc6q-ywxs1.gif'/>
</div>

## 目录
- [在线访问](#在线访问)
- [部署及运行](#部署及运行)
- [用法说明](#用法说明)
- [技术交流](#技术交流)
- [注意事项](#注意事项)
- [许可证](#许可证)


## 在线访问

http://117.50.163.130:3000/login/password

使用下面用户名和密码登录体验：

phone: 19911220000 (非真实电话号码)
password: 00000000

## 部署及运行

```bash
# 1.克隆仓库
git clone https://github.com/GIS142857/douyin-backend.git

# 2.进入项目目录
cd douyin-backend

# 3.安装依赖
go mod tidy

# 4.数据库导入
mysql -u username -p db_douyin < database/db_douyin.sql

# 5.修改配置文件 config.yaml、gorm_v2.yaml

# 6.运行项目
go run cmd/web/main.go 

# 7.前端项目启动(参考前端项目中的方法)
```

## 用法说明
由于取消了前端的请求拦截以及将前端的很多代码做了修改，所以需要使用我修改过的前端代码进行部署，修改版的前端代码在这个 fork 的项目中：

https://github.com/GIS142857/douyin.git


## 技术交流
目前项目还需补充更多功能，欢迎提交 `PR`，非常感谢你对我们的支持！
技术交流可以联系我的邮箱 <a href="mailto:fridalongwayhk4@gmail.com">fridalongwayhk4@gmail.com</a>

## 注意事项
注意：本项目仅适用于学习和研究，不得用于商业使用。

## 许可证

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

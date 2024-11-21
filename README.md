# douyin-backend

该项目为开源前端项目的后端部分，使用Go语言编写，主要实现了短视频的点赞、评论、关注、分享等功能。
前端项目地址：https://github.com/zyronon/douyin.git

## 目录

- [背景](#背景)
- [部署及运行](#部署及运行)
- [用法](#用法)
- [技术交流](#技术交流)
- [注意事项](#注意事项)
- [许可证](#许可证)

## 背景

基于开源的前端项目，为其开发后端部分，并增加了一些接口，包括点赞、评论、关注、分享和聊天等功能。同时，接入了推荐算法，为不同用户提供个性化内容推荐。


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

## 用法
由于取消了前端的请求拦截以及将前端的部分代码做了修改，所以需要使用我修改过的前端代码进行部署，修改版的前端代码在这个fork的项目中：https://github.com/GIS142857/douyin.git。

目前项目已经部署上线，访问 http://117.50.163.130:3000/login/password 

使用下面用户名和密码登录体验：

phone: 19911220000 (非真实电话号码)
password: 00000000

## 技术交流
目前项目还需补充更多功能，欢迎提交 `PR`，非常感谢你对我们的支持！
技术交流可以联系我的邮箱 <a href="mailto:fridalongwayhk4@gmail.com">fridalongwayhk4@gmail.com</a>

## 注意事项
注意：本项目仅适用于学习和研究，不得用于商业使用。

## 许可证
[MIT](LICENSE)
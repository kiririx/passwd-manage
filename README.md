# passwd-manage
这是一个用于管理用户密码的web应用。

# 构成
前端：React+Antd+Typescript
后端：Go
数据库：MySQL或sqlite

# 功能
- 注册/登录
- 电脑端密码管理
- 移动端密码管理（待开发。。。）

# 环境变量
应用里很多配置项都采用的环境变量，比如数据库的连接地址、端口、用户名、密码等。

这样方便了在docker的部署，这种web应用也推荐在docker里部署。

- db: mysql/sqlite，数据库的类型，只有选择了mysql，后面的配置才有意义
- db_username: 数据库的用户名
- db_host: 数据库地址
- db_password: 数据库的密码
- db_port: 数据库的端口
- db_database: 数据库的名称

# 部署
如果想直接使用，可以拉取我的docker-image

```bash
docker pull kiririx/passwd:latest
```

接着使用docker run命令，启动应用。

```bash
docker run -p 10011:8080 -d --name passwd -e db=mysql -e db_username=root -e db_password=xxx -e db_host=your.mysql.host -e db_port=3306 -e db_database=passwd kiririx/passwd:latest

```
服务默认开启了8080端口进行监听，请自行改变端口映射。
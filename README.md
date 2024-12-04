# MongoCaseLearn
## mongo Docker
下载 `mongo` `Docker` 镜像
```shell
docker pull mongo:latest
```
`Docker` 启动本地 `mysql` 容器
```shell
docker run -d -p 27017:27017 --name example-mongo -e MONGODB_INITDB_ROOT_USERNAME=root -e MONGODB_INITDB_ROOT_PASSWORD=password  mongo:latest
```
进入 `Docker` 容器，通过用户名和密码操作数据库
```shell
docker exec -it example-mongo bash
```

## Download
[https://www.mongodb.com/try/download/community](https://www.mongodb.com/try/download/community)

勾选 `MongoDB Compass`工具，可视化数据库数据

## Preference
https://www.mongodb.com/zh-cn/docs/manual/tutorial/insert-documents/

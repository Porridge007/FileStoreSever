## 这是什么
这是使用go语言写的文件存储管理系统
## 关于需要手动安装的库
    go get github.com/go-sql-driver/mysql
    go get github.com/spf13/viper
    >go get viper安装失败：
    > 1.cd  ${gopath}/src/golang.org/x
    > 2.git clone https://github.com/golang/text.git
    > 3.go install -x  golang.org/x/text
    go get github.com/garyburd/redigo/redis
    go get github.com/json-iterator/go
1.将FileStoreServer移入$GOPATH<br>
2.在mysql中执行db/tables.sql
3.执行go run main.go

## 实现功能
- [x] 文件管理系统的基本增删改查
- [x] 数据持久化
- [x] 用户模块、session鉴权、资源隔离
- [ ] 分块上传


---
## What's this?
This is a File Store Server written in Golang.
## What's to go get?
    go get github.com/go-sql-driver/mysql.
    go get github.com/spf13/viper
    What if go get viper failed?
        > 1.cd  ${gopath}/src/golang.org/x
        > 2.git clone https://github.com/golang/text.git
        > 3.go install -x  golang.org/x/text
    go get github.com/garyburd/redigo/redis
    go get github.com/json-iterator/go
## How to start it?
1.Move Project File to your $Gopath<br>
2.Execute sql in doc/tables.sql in Mysql.
3.go run main.go

## Completed Functions
- [x] File Storage Sever basic CRUD.
- [ ] Data persistence
- [ ] Account System, Session Authentication, Resource isolation
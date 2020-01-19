## 这是什么
这是使用go语言写的文件存储管理系统
## 关于需要手动安装的库
    go get github.com/go-sql-driver/mysql
    go get github.com/spf13/viper
## 使用文档
1.将FileStoreServer移入$GOPATH<br>
2.执行go run main.go

## 实现功能
- [x] 文件管理系统的基本增删改查
- [ ] 数据持久化
- [ ] 用户模块、session鉴权、资源隔离

---
## What's this?
This is a File Store Server written in Golang.
## What's to go get?
    go get github.com/go-sql-driver/mysql.
    go get github.com/spf13/viper
## How to start it?
1.Move Project File to your $Gopath<br>
2.go run main.go

## Completed Functions
- [x] File Storage Sever basic CRUD.
- [ ] Data persistence
-[ ] Account System, Session Authentication, Resource isolation
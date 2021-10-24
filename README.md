
## go

> go version

> go env

## bee

> go get github.com/beego/bee

bee new <项目名>
bee api <项目名>
bee run
bee pack
bee version
bee generate

---

> bee api user-service

## go get 代理

> echo "export GOPROXY=https://goproxy.cn" >> ~/.profile && source ~/.profile

## 更新依赖

> go mod tidy

## 访问 

localhost:8080/v1/user 
GetAll

localhost:8080/v1/user/user1
Get id=user1

---

## orm

查看创建表的SQL
> go run main.go orm sqlall 

创建表
> go run main.go orm syncdb


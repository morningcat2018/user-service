
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

{
    "user1": {
        "Id": "user1",
        "Username": "晨猫",
        "Password": "123456",
        "Profile": {
            "Gender": "男",
            "Age": 28,
            "Address": "南京市雨花台区",
            "Email": "mengzhang1993@gmail.com"
        }
    }
}

---

## orm

> go run main.go orm sqlall 

> go run main.go orm syncdb


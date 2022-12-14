# go_oj

Front-end -- Vue、ElementUI

Back-end -- Gin、GORM

**Swagger**

Interface access address：[http://localhost:8080/swagger/index.html]()

```shell
// GetProblemList
// @Tags 公共方法
// @Summary 问题列表
// @Param page query int false "page"
// @Param size query int false "size"
// @Success 200 {string} json "{"code":"200","msg","","data":""}"
// @Router /problem-list [get]
```

Install jwt

```shell
go get github.com/dgrijalva/jwt-go
```

Install mail

```shell
go get github.com/jordan-wright/email
```

* [X] OJ
  * [X] User Module
    * [X] register
    * [X] login
    * [X] user detail
  * [X] Problem Module
    * [X] problem list、problem detail
    * [X] problem create
    * [X] problem modify
  * [X] Judge Module
    * [X] submit list
    * [X] submit and judge
  * [X] Rank Module
    * [X] rank list

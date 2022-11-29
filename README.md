# GoOJ

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


* [ ] OJ
  * [X] User Module
    * [X] register
    * [X] login
    * [X] user detail
  * [ ] Problem Module
    * [X] problem list、problem detail
    * [ ] problem create
    * [ ] problem modify
  * [ ] Judge Module
    * [X] submit list
    * [ ] submit and judge
  * [ ] Rank Module
    * [ ] rank list

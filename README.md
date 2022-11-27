# GoOJ

Front-end -- Vue、ElementUI

Back-end -- Gin、GORM


Swagger

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

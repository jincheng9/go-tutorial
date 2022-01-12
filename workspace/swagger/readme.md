# gin-swagger常见问题

## Fetch error Internal Server Error doc*.*json

需要导入`swag init`命令生成的docs包

```go
import _  module_name/docs
```



## gin-swagger支持multipart/form-data里带文件参数

Param Type: `formData`

Data Type: `file`

```markdown
// UpdateProfile godoc
// @Summary UpdateProfile updates user info
// @Description check username and token, then update user profile
// @Accept multipart/form-data
// @Produce json
// @param username path string true "username"
// @param token formData string true "token"
// @param nickname formData string false "nickname"
// @param image formData file false "image"
// @Success 200 {object} map[string]interface{}
// @Router /update/{username} [POST]
```

* https://github.com/swaggo/swag/blob/master/README.md

* https://swagger.io/docs/specification/2-0/file-upload/


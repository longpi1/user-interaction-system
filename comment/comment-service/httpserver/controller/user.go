package controller

import (
	"encoding/json"
	"net/http"

	"github.com/longpi1/user-interaction-system/comment-service/libary/constant"
	"github.com/longpi1/user-interaction-system/comment-service/model/dao/db/model"
	"github.com/longpi1/user-interaction-system/comment-service/model/service"

	"github.com/longpi1/gopkg/libary/utils"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var user model.User
	err := json.NewDecoder(c.Request.Body).Decode(&user)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"event":   constant.InvalidParam,
		})
		return
	}
	if err := utils.Validate.Struct(&user); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"event":   "输入不合法 " + err.Error(),
		})
		return
	}
	if err := service.Register(&user); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"event":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"event":   "",
	})
}

func Login(c *gin.Context) {
	var loginRequest model.LoginRequest
	err := json.NewDecoder(c.Request.Body).Decode(&loginRequest)
	if err != nil {
		utils.RespError(c, constant.InvalidParam)
		return
	}
	username := loginRequest.Username
	password := loginRequest.Password
	if username == "" || password == "" {
		utils.RespError(c, constant.InvalidParam)
		return
	}
	user := model.User{
		Username: username,
		Password: password,
	}
	err = service.Login(user)
	if err != nil {
		utils.RespError(c, err.Error())
		return
	}
	setupLogin(&user, c)
}

// setup session & cookies and then return user info
func setupLogin(user *model.User, c *gin.Context) {
	session := sessions.Default(c)
	session.Set("id", user.ID)
	session.Set("username", user.Username)
	session.Set("role", user.Role)
	session.Set("status", user.Status)
	err := session.Save()
	if err != nil {
		utils.RespError(c, "无法保存会话信息，请重试")
		return
	}
	cleanUser := model.User{
		Username:    user.Username,
		DisplayName: user.DisplayName,
		Role:        user.Role,
		Status:      user.Status,
	}
	utils.RespData(c, "登录成功", cleanUser)
}

func Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	err := session.Save()
	if err != nil {
		utils.RespError(c, err.Error())
		return
	}
	utils.RespSuccess(c, "退出登录成功")
}

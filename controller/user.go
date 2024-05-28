package controller

import (
	"bluebell/dao/mysql"
	"bluebell/logic"
	"bluebell/models"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// 处理注册请求的函数
func SignUpHandler(c *gin.Context) {
	//1.获取参数和参数校验
	p := new(models.ParamSignUp)
	if err := c.ShouldBindJSON(p); err != nil {
		//请求参数有误
		zap.L().Error("SignUp with invalid param", zap.Error(err)) //记录一个错误日志
		//判断err是不是validator类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			//非validator类型错误
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans))) //翻译错误
		return
	}
	//手动对函数进行详细的业务规则校验
	/*if len(p.Username) == 0 || len(p.Password) == 0 || len(p.RePassword) == 0 || p.RePassword != p.Password {
		zap.L().Error("SignUp with invalid param") //记录一个错误日志
		c.JSON(http.StatusOK, gin.H{
			"msg": "请求参数有误",
		})
		return
	}*/
	fmt.Println(p)
	//2.业务处理 （一般放在logic层）
	if err := logic.SignUp(p); err != nil {
		zap.L().Error("logic.SignUp failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserExist) {
			ResponseError(c, CodeUserExist)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}
	//3.返回响应
	ResponseSuccess(c, nil)
}

// 登录请求的函数
func LoginHandler(c *gin.Context) {
	//获取请求参数及参数校验
	p := new(models.ParamLogin)
	if err := c.ShouldBindJSON(p); err != nil {
		//请求参数有误
		zap.L().Error("Login with invalid param", zap.Error(err)) //记录一个错误日志
		//判断err是不是validator类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {

			//非validator类型错误
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	//业务逻辑处理
	user, err := logic.Login(p)
	if err != nil {
		zap.L().Error("login.Login failed", zap.String("username", p.Username), zap.Error(err))
		if errors.Is(err, mysql.ErrorUserExist) {
			ResponseError(c, CodeUserExist)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}
	//返回响应
	ResponseSuccess(c, gin.H{
		"user_id":   fmt.Sprintf("%d", user.UserID),
		"user_name": user.Username,
		"token":     user.Token,
	})
}

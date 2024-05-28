package controller

import "bluebell/models"

//专门用来存放接口文档用到的models
//接口文档返回的数据格式是一致的，但是具体的data类型不一样

type _ResponsePostList struct {
	Code    ResCode                 `json:"code"`    //业务响应状态码
	Message string                  `json:"message"` //提示信息
	Data    []*models.ApiPostDetail `json:"data"`    //数据
}

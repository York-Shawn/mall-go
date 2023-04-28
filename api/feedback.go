package api

import (
	"mall/constant"
	"mall/models/app"
	"mall/models/web"
	"mall/response"
	"mall/service"

	"github.com/gin-gonic/gin"
)

type WebFeedback struct {
	service.WebFeedBackService
}

type AppFeedback struct {
	service.AppFeedBackService
}

func GetWebFeedback() *WebFeedback {
	return &WebFeedback{}
}

func GetAppFeedback() *AppFeedback {
	return &AppFeedback{}
}

func (f *WebFeedback) SendFeedback(context *gin.Context) {
	var param web.FeedbackParam
	if err := context.ShouldBind(&param); err != nil {
		response.Failed(constant.ParamInvalid, context)
		return
	}
	if err := f.Send(param); err != nil {
		response.Failed(constant.SendFailed, context)
		return
	}
	response.Success(constant.SendSuccess, nil, context)
}

func (f *AppFeedback) SendFeedback(context *gin.Context) {
	var param app.FeedbackParam
	if err := context.ShouldBind(&param); err != nil {
		response.Failed(constant.ParamInvalid, context)
		return
	}
	if err := f.Send(param); err != nil {
		response.Failed(constant.SendFailed, context)
		return
	}
	response.Success(constant.SendSuccess, nil, context)
}

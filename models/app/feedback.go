package app

type FeedbackParam struct {
	Content string `form:"content"` // 反馈的内容
	Email   string `form:"email"`   // 反馈的邮箱
}

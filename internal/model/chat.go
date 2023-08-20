package model

// ChatMsg 消息结构体
type ChatMsg struct {
	Type string      `json:"type" v:"required#消息类型不能为空"`
	Data interface{} `json:"data" v:""`
	From string      `json:"name" v:""`
}

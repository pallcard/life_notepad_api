package model

// ChatMsg 消息结构体
type ChatMsg struct {
	Type       string      `json:"type" v:"required#消息类型不能为空"` // Single单聊 Group群聊
	SenderId   string      `json:"sender_id" v:""`             // 发送者
	ReceiverId string      `json:"receiver_id" v:""`           // 接收者
	CreateTime string      `json:"create_time" v:""`           // 创建时间
	Data       interface{} `json:"data" v:""`                  // 聊天数据，特别：链接数据以 [link] 开头
}

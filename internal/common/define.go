package common

// Res 返回数据结构
type Res struct {
	Code  int
	Error string
	Data  interface{}
}

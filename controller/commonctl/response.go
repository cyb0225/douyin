package commonctl
// 返回相应的基本信息，作为user接口的基本类型，嵌入到各个user接口的响应类中
type Response struct {
	// 返回状态码 0 表示成功， 其他表示失败
	Status_code int `json:"status_code"`
	// 返回状态描述
	Status_msg string `json:"status_msg"`
}

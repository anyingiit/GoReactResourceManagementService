package structs

type InvateClientMessage struct {
	ServerHost string `json:"server_host"`
	ServerPort int    `json:"server_port"`
	InvateCode string `json:"invate_code"` // 本身是无意义的字符，被服务器验证成功后会返回一个和服务器的通信凭证
}

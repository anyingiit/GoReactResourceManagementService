package structs

type InvateClientMessage struct {
	ServerIP   string
	ServerPort int
	InvateCode string // 本身是无意义的字符，被服务器验证成功后会返回一个和服务器的通信凭证
}

package email

import(
	"github.com/wjpxxx/letgo/lib"
)

//邮件消息
type Message struct {
	From        string       `json:"from"`//发送邮箱
	To          []string     `json:"to"`//接收邮箱
	Cc          []string     `json:"cc"`//抄送邮箱
	Bcc         []string     `json:"bcc"`//暗抄送邮箱
	Subject     string       `json:"subject"`//邮件主题
	Body        string       `json:"body"`//邮件内容
	PlainText   string       `json:"plainText"`//纯文本内容
	ContentType string       `json:"contentType"`//邮件内容编码
	Attachment  []Attachment  `json:"attachment"`//附件
}
//String
func(m *Message)String() string {
    return lib.ObjectToString(m)
}
//邮件配置
type EmailConfig struct{
	User        string       `json:"user"`//邮箱账号
	Password    string       `json:"password"`//邮箱密码
	Host        string       `json:"host"`//邮箱服务器地址
	Port        string       `json:"port"`//邮箱服务器端口号
}
//String
func(e *EmailConfig)String() string {
    return lib.ObjectToString(e)
}

//附件
type Attachment struct {
	Data 		[]byte  `json:"data"`//文件二进制
	Name        string  `json:"name"`//文件名
	ContentType string  `json:"contentType"`//附件编码类型,比如:multipart/related;boundary=GoBoundary
}

//String
func(e *Attachment)String() string {
    return lib.ObjectToString(e)
}
package email
import (
	"crypto/tls"
	"io/ioutil"
	"net/smtp"
	"strings"
	"net/http"
	"github.com/wjpxxx/letgo/lib"
	"github.com/wjpxxx/letgo/file"
	"github.com/wjpxxx/letgo/encry"
	"fmt"
)
/**
邮件类 基类
**/
type Email struct {
	config EmailConfig
	auth     smtp.Auth
	message Message
	boundary string
}
//初始化邮件配置
func WithConfig(config EmailConfig) Emailer{
	return &Email{
		config: config,
		auth:smtp.PlainAuth("", config.User, config.Password, config.Host),
		message:Message{
			From:config.User,
		},
		boundary:"LetGoBoundary",
	}
}
//发送邮件
func (e *Email)Send(to,subject,body string)error{
	return e.Sends([]string{to}, subject,body)
}

//发送邮件
func (e *Email)Sends(to []string,subject,body string)error{
	e.message.To=to
	e.message.Subject=subject
	e.message.Body=body
	addr:=e.config.Host
	if e.config.Port!=""{
		addr+=fmt.Sprintf(":%s",e.config.Port)
	}
	client, err := smtp.Dial(addr)
	if err != nil {
		return err
	}
	defer client.Close()
	if ok, _ := client.Extension("AUTH"); ok {
		if err := client.Auth(e.auth); err != nil {
			return err
		}
	}
	if err := client.Mail(e.message.From); err != nil {
		return err
	}
	for _, to := range e.message.To {
		if err := client.Rcpt(to); err != nil {
			return err
		}
	}
	w, err := client.Data()
	if err != nil {
		return err
	}
	content:=e.packageMessage()
	_, err = w.Write([]byte(content))
	if err != nil {
		return err
	}
	w.Close()
	client.Quit()
	return nil
}
//发送邮件
func (e *Email)SendTLSs(to []string,subject,body string)error{
	e.message.To=to
	e.message.Subject=subject
	e.message.Body=body
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName: e.config.Host,
	}
	addr:=e.config.Host
	if e.config.Port!=""{
		addr+=fmt.Sprintf(":%s",e.config.Port)
	}
	con, err := tls.Dial("tcp", addr, tlsconfig)
	if err != nil {
		return err
	}
	client, err := smtp.NewClient(con, e.config.Host)
	if err != nil {
		return err
	}
	defer client.Close()
	//fmt.Println(client.Extension("AUTH"))
	if ok, _ := client.Extension("AUTH"); ok {
		if err := client.Auth(e.auth); err != nil {
			return err
		}
	}
	if err := client.Mail(e.message.From); err != nil {
		return err
	}
	for _, to := range e.message.To {
		if err := client.Rcpt(to); err != nil {
			return err
		}
	}
	w, err := client.Data()
	if err != nil {
		return err
	}
	content:=e.packageMessage()
	//fmt.Println(content)
	_, err = w.Write([]byte(content))
	if err != nil {
		return err
	}
	w.Close()
	client.Quit()
	return nil
}
//发送邮件
func (e *Email)SendTLS(to,subject,body string)error{
	return e.SendTLSs([]string{to},subject,body)
}
//WithAppendByPath 添加附件
func (e *Email)WithAppendByPath(path string)Emailer{
	f, _ := ioutil.ReadFile(path)
	e.message.Attachment=append(e.message.Attachment,Attachment{
		Name:file.BaseName(path),
		ContentType:http.DetectContentType(f),
		Data:f,
	})
	return e
}
//WithAppendBytes 添加附件
func (e *Email)WithAppendBytes(name string,data []byte)Emailer{
	e.message.Attachment=append(e.message.Attachment,Attachment{
		Name:name,
		ContentType:http.DetectContentType(data),
		Data:data,
	})
	return e
}
//packageMessage 封装消息
func (e *Email)packageMessage()string{
	content:="From:"+e.message.From+"\r\n"
	content+="To:"+strings.Join(e.message.To, ";")+"\r\n"
	if len(e.message.Cc) != 0 {
		content+="Cc:"+strings.Join(e.message.Cc, ";")+"\r\n"
	}
	if len(e.message.Bcc) != 0 {
		content+="Bcc:"+strings.Join(e.message.Bcc, ";")+"\r\n"
	}
	content+="Subject:"+e.message.Subject+"\r\n"
	content+="Content-Type:multipart/mixed;boundary="+e.boundary+"\r\n"
	content+="Date:"+lib.Now()+"\r\n\r\n"
	if len(e.message.Attachment)!=0{
		for _, atr := range e.message.Attachment {
			atr.ContentType=strings.Replace(atr.ContentType,"text/plain","application/octet-stream",-1)
			content += "\r\n--" + e.boundary + "\r\n"
			content += "Content-Transfer-Encoding:base64\r\n"
			content += "Content-Type:" + atr.ContentType + ";name=\"" + atr.Name + "\"\r\n"
			content += "Content-ID: <" + atr.Name + "> \r\n\r\n"
			content += encry.Base64StdEncode(string(atr.Data))
		}
	}
	if e.message.PlainText != "" {
		content += "\r\n--" + e.boundary + "\r\n"
		content += "Content-Type:text/plain;charset=\"UTF-8\" \r\n"
		content += "\r\n" + e.message.PlainText
	}
	body := "<html><body>" + e.message.Body + "</body></html>"
	content += "\r\n--" + e.boundary + "\r\n"
	content += "Content-Transfer-Encoding:base64\r\n"
	content += "Content-Type: text/html;charset=\"UTF-8\" \r\n"
	content += "\r\n" + encry.Base64StdEncode(body)
	content += "\r\n--" + e.boundary + "--\r\n\r\n"
	return content
}
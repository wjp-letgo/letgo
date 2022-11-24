package email
import(
	"testing"
	"fmt"
)
func TestSend(t *testing.T){
	e:=WithConfig(EmailConfig{
		User:"234234234324@qq.com",
		Password:"23423423424",
		Host:"smtp.qq.com",
		Port: "465",
	}).WithAppendByPath("./email_test.go").SendTLS("23424234243@qq.com","nihao","<b>hello</b>")
	fmt.Println(e)
}
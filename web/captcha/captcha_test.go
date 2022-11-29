package captcha

import (
	"image/color"
	"testing"

	"github.com/wjp-letgo/letgo/file"
)

func TestCaptcha(t *testing.T){
	c:=NewCaptcha()
	c.AddFonts("STHUPO.TTF")
	c.AddColors(color.RGBA{255,0,255,255},color.RGBA{0,255,255,255},color.Black)
	c.SetSize(200,60)
	//c.AddBackColors(color.Black,color.Opaque)
	//c.Create(4,NUM)
	//c.Create(4,LCHAR)
	//c.SetDisturbLevel(HIGH)
	c.SetDisturbLevel(MEDIUM)
	img1,_:=c.Create(4,ALL)
	//img1.DrawLine(10,0,100,20,color.RGBA{255,0,255,255})
	img1.SaveImage("1.png")
	img2,_:=c.Create(4,ALL)
	file.PutContent("2",img2.Base64Encode())
	img2.SaveImage("2.png")
}
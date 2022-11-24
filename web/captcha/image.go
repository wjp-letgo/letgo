package captcha

import (
	"image"
	"image/color"
	"image/png"
	"path/filepath"
	"image/gif"
	"image/jpeg"
	"image/draw"
	"math"
	"os"
	"bytes"
	"encoding/base64"
	"fmt"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"github.com/wjpxxx/letgo/file"
)

//Image
type Image struct {
	*image.RGBA
}

//NewImage
func NewImage(width, height int) *Image {
	return &Image{image.NewRGBA(image.Rect(0,0,width,height))}
}

//DrawString
func (i *Image)DrawString(font *truetype.Font, color color.Color,text string,fontSize float64){
	ctx:=freetype.NewContext()
	ctx.SetDst(i)
	ctx.SetClip(i.Bounds())
	ctx.SetSrc(image.NewUniform(color))
	ctx.SetFontSize(fontSize)
	ctx.SetFont(font)
	pt:=freetype.Pt(0,int(-fontSize/6)+ctx.PointToFixed(fontSize).Ceil())
	ctx.DrawString(text,pt)
}

//SaveImage
func (i *Image)SaveImage(fullPath string){
	path:=file.DirName(fullPath)
	file.Mkdir(path)
	f,_:=os.OpenFile(fullPath, os.O_SYNC|os.O_RDWR|os.O_CREATE, 0666)
	defer f.Close()
	switch filepath.Ext(fullPath) {
		case ".jpg":
			jpeg.Encode(f,i,&jpeg.Options{Quality: 100})
		case ".jpeg":
			jpeg.Encode(f,i,&jpeg.Options{Quality: 100})
		case ".png":
			png.Encode(f,i)
		case ".gif":
			gif.Encode(f,i,&gif.Options{NumColors: 256})
	}
}

//Base64Encode
func (i *Image)Base64Encode()string{
	buf:=bytes.NewBuffer(nil)
	png.Encode(buf,i)
	return fmt.Sprintf("%s%s","data:image/png;base64,",base64.StdEncoding.EncodeToString(buf.Bytes()))
}


//Fill
func(i *Image)Fill(img image.Image){
	draw.Draw(i,i.Bounds(),img,image.ZP,draw.Over)
}
//DrawLine
func (i *Image)DrawLine(x1,y1,x2,y2 int,colr color.Color){
	dx, dy, flag := int(math.Abs(float64(x2-x1))),
		int(math.Abs(float64(y2-y1))),
		false
	if dy > dx {
		flag = true
		x1, y1 = y1, x1
		x2, y2 = y2, x2
		dx, dy = dy, dx
	}
	ix, iy := sign(x2-x1), sign(y2-y1)
	n2dy := dy * 2
	n2dydx := (dy - dx) * 2
	d := n2dy - dx
	for x1 != x2 {
		if d < 0 {
			d += n2dy
		} else {
			y1 += iy
			d += n2dydx
		}
		if flag {
			i.Set(y1, x1, colr)
		} else {
			i.Set(x1, y1, colr)
		}
		x1 += ix
	}
}
//sign
func sign(x int) int {
	if x > 0 {
		return 1
	}
	return -1
}
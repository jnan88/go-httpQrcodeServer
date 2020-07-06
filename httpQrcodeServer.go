package main

import (
qrcode "github.com/skip2/go-qrcode"
gin "github.com/gin-gonic/gin"
"log"
"net/http"
"net/url"
"os"
"strconv"
)
import "fmt"

func main() {
	router := gin.Default()
	router.GET("/favicon.ico", func(context *gin.Context) {
		context.Writer.WriteHeader(http.StatusNoContent)
		return
	})
	router.GET("/createQrcode", func(context *gin.Context) {
		var text string= context.Request.URL.Query().Get("text")
		encode := context.Request.URL.Query().Get("encode")
		size := context.DefaultQuery("size","256")
		imgSize ,err:= strconv.Atoi(size)
		if err != nil {
			imgSize=256
		}
		if imgSize>512 {
			imgSize=512
		}
		if "url"==encode {
			text,_= url.QueryUnescape(text)
		}
		log.Println("size="+strconv.Itoa(imgSize))
		log.Println("text="+text)
		imgButes,err := qrcode.Encode(text, qrcode.Medium,imgSize)
		if err != nil {
			fmt.Println("write error")
			context.Writer.WriteHeader(http.StatusNoContent)
			return
		}
		context.Writer.Write(imgButes)
	})
	var arg_num int = len(os.Args)
	fmt.Printf(" run args is %d\n",arg_num)
	var runPort string= ":9090"
	if arg_num>1{
		p,err:=strconv.Atoi(os.Args[1])
		if err==nil{
			runPort=":"+strconv.Itoa(p)
		}
	}
	// 指定地址和端口号
	fmt.Println("run on "+runPort)
	router.Run(runPort)
}


package paraseURL

import (
	"VideoSpider/public/error_operate"
	"regexp"
)

func ParasListUrl(urls []string,name string)  {
	for _,url := range urls{
		resContent, err := ParaseURL(url)
		if err != nil{
			panic(err)
		}
		ParasResponeData(resContent,name)
	}
}

func JudgeForRespond(content string){

	err:=regexp.MustCompile(`("errorcode":"(.*?)",)`)

	aimError := err.FindAllStringSubmatch(content,-1)

	if aimError[0][2] != ""{
		panic(error_operate.NoticeError{"访问的 url 格式有错误或者网址出错！"})
	}
}

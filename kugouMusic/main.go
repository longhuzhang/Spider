package main

//https://www.cnblogs.com/sam-uncle/archive/2019/05/27/10934064.html 参考博客内容

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
)

var regxStr = `Hash":"([^\"]+)"`
var regxAlbum = `"album_id":([0-9]+)`
var regexTitle = ` class="pc_temp_songname" title="([^"]+)" hidefocus="true"`

func main()  {
	resp,err := http.Get("https://www.kugou.com/yy/rank/home/1-8888.html?from=rank")
	defer resp.Body.Close()
	if err != nil{
		panic("访问榜单失败")
	}

	content,err := ioutil.ReadAll(resp.Body)
	myMaterial:=string(content)

	rex:= regexp.MustCompile(regxStr)
	result := rex.FindAllStringSubmatch(myMaterial,-1)
	fmt.Println(result)
	fmt.Println(len(result))

	rexAlbum := regexp.MustCompile(regxAlbum)
	albumResult := rexAlbum.FindAllStringSubmatch(myMaterial,-1)

	rexTitle := regexp.MustCompile(regexTitle)
	titleResult := rexTitle.FindAllStringSubmatch(myMaterial,-1)


	fmt.Println(albumResult)
	fmt.Println(len(albumResult))
	fmt.Println(len(titleResult))

	for key,hash := range result{
		fmt.Println("song name is",titleResult[key][1])
		url := "https://wwwapi.kugou.com/yy/index.php?r=play/getdata&callback=jQuery19107422002281373865_1562120909057&hash="+hash[1]+"&album_id="+albumResult[key][1]+"&dfid=4EANJt3L4s910PkKm92qgWDz&mid=ae151661a596a5c61c5ba739d283eac3&platid=4&_=1562120909058"
		//fmt.Println(url)
		RequestUrl(url,titleResult[key][1])
	}
}

var downUrl = `"play_url":"([^\"]+)"`

func RequestUrl(url,name string)  {
	fmt.Println(url)
	resp,err := http.Get(url)
	defer resp.Body.Close()
	if err != nil{
		panic(err)
	}

	result,err := ioutil.ReadAll(resp.Body)

	fmt.Println(string(result))

	regex:=regexp.MustCompile(downUrl)
	res:=regex.FindAllStringSubmatch(string(result),-1)
	fmt.Println(res)
	if len(res) == 0{
		return
	}
	aimUrl := strings.Replace(res[0][1],"\\","",-1)
	fmt.Println("aimUrl",aimUrl)

	fileResp,err := http.Get(aimUrl)
	defer fileResp.Body.Close()
	if err != nil{
		fmt.Println("request err",err)
		panic(err)
	}

	_, err = os.Stat("kugou")
	if err != nil{
		os.Mkdir("kugou",os.ModePerm)
	}

	file,err:=os.Create("kugou/"+name+".mp3")
	defer file.Close()
	if err != nil{
		panic(err)
	}

	io.Copy(file,fileResp.Body)
}

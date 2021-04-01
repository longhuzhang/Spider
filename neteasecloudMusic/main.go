package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
)


var regxStr = `<a href="/song\?id=([0-9]+)">([^<]+)</a>?`

func main() {

	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://music.163.com/discover/toplist?id=3778678", nil)
	if err != nil {
		fmt.Println("获取热歌榜失败！")
		return
	}
	//设置请求头，防止504
	//req.Header.Set("Referer", "https://music.163.com/")
	//req.Header.Set("Host", "music.163.com")
	//req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.100 Safari/537.36")
	//req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
	resp, err := client.Do(req)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("获取热歌榜失败！")
		return
	}
	content := string(body)

	doRegx := regexp.MustCompile(regxStr)
	fingResult := doRegx.FindAllStringSubmatch(content, -1)
	fmt.Println("find song result", len(fingResult))
	count := 0
	for _,song := range fingResult{
		if count== 10 {
			return
		}
		DownLoadSong(song[1],song[2])
		count++
	}

	fmt.Println("获取热歌榜成功！")
}


func DownLoadSong(id, name string)  {
	url := "http://music.163.com/song/media/outer/url?id="+id+".mp3"
	fmt.Println(url)
	resp,err  := http.Get(url)
	defer resp.Body.Close()
	if err != nil{
		panic(err)
	}

	os.Mkdir("song",os.ModeDir)
	file, err := os.Create("song/"+name+".mp3")
	defer file.Close()
	if err != nil{
		panic(err)
	}
	io.Copy(file,resp.Body)
}


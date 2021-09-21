package util

import (
	"Spider/public"
	"crypto/md5"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"strings"
)

//根据视频名字获取
func SearchVideoByName(name string) (string, error) {
	url := strings.Replace(public.SearchUrl, "{keyword}", name, -1)
	url = strings.Replace(url, "{pagesize}", "1", -1)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	return string(data), err
}

//根据hash获取视频下载地址
func ObtainVideoDownUrl(hash string) (string, error) {
	data := []byte(strings.ToUpper(hash) + "kugoumvcloud")
	val := md5.Sum(data)
	key := fmt.Sprintf("%x", val)
	url := strings.Replace(public.DownloadUrl, "{hash}", hash, -1)
	url = strings.Replace(url, "{key}", key, -1)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	return string(content), err
}

//视频文件下载
func DownVideo(url, id, title, actor string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	dir := fmt.Sprintf("./download/%s", id)
	_, err = os.Stat(dir)
	if err != nil && !os.IsExist(err) {
		err = os.MkdirAll(dir, os.ModePerm)
	}

	name := dir + "/" + title + "-" + actor + ".mp4"
	file, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = io.Copy(file, resp.Body)
	return err
}

func SendError(errMessage chan public.ErrorMessage, err error) {
	_, file, line, _ := runtime.Caller(1)
	errMessage <- public.ErrorMessage{
		File: file,
		Line: line,
		Err:  err,
	}

}

func ParaseUnicode(source string) string {
	sUnicodev := strings.Split(source, `\u`)
	aim := ""
	for _, v := range sUnicodev {
		if v == "" {
			continue
		}
		fmt.Println("v", v)
		temp, err := strconv.ParseInt(v, 16, 32)
		if err != nil {
			panic(err)
		}
		aim += fmt.Sprintf("%c", temp)
	}
	return aim
}

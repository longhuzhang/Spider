package paraseURL

import (
	"VideoSpider/public"
	"VideoSpider/public/error_operate"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Data struct {
	Url         string `json:"url"`
	Extname     string `json:"extname"`
	Duration    int    `json:"duration"`
	FileSize    int    `json:"filesize"`
	ThumbUrl    string `json:"thumb_url"`
	TransStatus int    `json:"trans_status"`
}

type Respone struct {
	Status  int    `json:"status"`
	Errcode int    `json:"errcode"`
	Error   string `json:"error_operate"`
	Data    Data   `json:"data"`
}

func ParseId(id string) {
	url := "http://visitor.fanxing.kugou.com/VServices/Video.OfflineVideoService.getVideoList/" + id + "-1-0-10/"
	resContent, err := ParaseURL(url)
	if err != nil {
		panic(error_operate.NoticeError{"ParaseURL error"})
	}

	JudgeForRespond(resContent)

	//idRegex := regexp.MustCompile(`getVideoList(.*?)`)
	//myId := idRegex.FindAllStringSubmatch(resContent, -1)
	countMatch := regexp.MustCompile(`("count":(.*?)})`)
	getCount := countMatch.FindAllStringSubmatch(resContent, -1)

	fmt.Println(getCount, len(getCount[0]))
	//fmt.Println(myId,len(myId[0]))
	count, err := strconv.Atoi(getCount[0][2])
	if err != nil {
		panic(err)
	}

	ParasResponeData(resContent, id)

	pageNum := count/10 + 1
	for i := 2; i < pageNum; i++ {
		page := strconv.Itoa(i)
		url = "http://visitor.fanxing.kugou.com/VServices/Video.OfflineVideoService.getVideoList/" + id + "-" + page + "-0-10/"
		fmt.Println(url)
		aim, err := ParaseURL(url)
		if err != nil {
			panic(err)
		}
		ParasResponeData(aim, id)
	}

}

func ParaseURL(url string) (string, error) {
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	defer func() {
		res.Body.Close()
	}()
	resContent := string(b)
	return resContent, nil
}

func ParasResponeData(resContent string, id string) {
	public.Cond.L.Lock()
	title := regexp.MustCompile(`("title":"(.*?)")`)
	actor := regexp.MustCompile(`("actor":"(.*?)")`)
	hashl := regexp.MustCompile(`("hashValue":"(.*?)")`)
	hashList := hashl.FindAllStringSubmatch(resContent, -1)
	titleList := title.FindAllStringSubmatch(resContent, -1)
	actorList := actor.FindAllStringSubmatch(resContent, -1)

	var myTitle public.MyTitle
	var myActor public.MyActor

	fmt.Println("the title is :", len(titleList), len(hashList))
	fmt.Println("actor is:", actorList, len(actorList))
	for key, v := range hashList {

		select {
		case i := <-public.MyChannel:
			if i == 1 {
				return
			} else if i == 2 {
				public.Cond.Wait()
			}
		default:
		}

		theHash := v[2]
		if strings.Contains(v[2], "|") {
			split := strings.Split(v[2], "|")
			theHash = split[0]
		}

		//md5加密
		hashValue := "hash" + theHash + "pid6kugou_video"
		a := md5.Sum([]byte(hashValue))
		value := fmt.Sprintf("%x", a)

		fmt.Println(v[2])

		url := "http://tracker.v.kugou.com/video/query?pid=6&sign=" + value + "&hash=" + theHash

		json.Unmarshal([]byte("{"+titleList[key][0]+"}"), &myTitle)
		json.Unmarshal([]byte("{"+actorList[key][0]+"}"), &myActor)

		getVideoAndSave(GetVedioPlayAddress(url), id, myTitle.Title, myActor.Actor)
		fmt.Println("url:", url)
	}
	public.Cond.L.Unlock()
}

func GetVedioPlayAddress(utl string) string {
	res, err := http.Get(utl)
	if err != nil {
		panic(err)
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	defer func() {
		res.Body.Close()
	}()
	var resContent Respone
	fmt.Println(string(b))
	err = json.Unmarshal(b, &resContent)
	if err != nil {
		panic(err)
	}
	return resContent.Data.Url
}

//获取视频并保存到本地
func getVideoAndSave(url, id, title, actor string) {
	//url := "http://fx.v.kugou.com/G127/M07/0B/17/vw0DAFx6mdiAP82TAXAgmN91YYw320.mp4"
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	path := fmt.Sprintf("download/%s", id)
	_, err = os.Stat(path)
	fmt.Println(path, err)
	if err != nil {
		os.MkdirAll(path, os.ModePerm)
	}
	myName := title + "_" + actor + ".mp4"
	fmt.Println("myName", myName)
	name := fmt.Sprintf("download/%s/%s", id, myName)
	//如果文件已经存在就跳过
	_, err = os.Stat(name)
	if err == nil {
		return
	}

	f, err := os.Create(name)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	io.Copy(f, res.Body)
}

package parase_url

import (
	"Spider/public"
	"Spider/util"
	"errors"
	"io/ioutil"
	"net/http"
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

type Response struct {
	Status  int    `json:"status"`
	Errcode int    `json:"errcode"`
	Error   string `json:"error_operate"`
	Data    Data   `json:"data"`
}

//
//func ParseId(id string) {
//	url := "http://visitor.fanxing.kugou.com/VServices/Video.OfflineVideoService.getVideoList/" + id + "-1-0-10/"
//	resContent, err := ParaseURL(url)
//	if err != nil {
//		panic(error_operate.NoticeError{"ParaseURL error"})
//	}
//
//	JudgeForRespond(resContent)
//
//	//idRegex := regexp.MustCompile(`getVideoList(.*?)`)
//	//myId := idRegex.FindAllStringSubmatch(resContent, -1)
//	countMatch := regexp.MustCompile(`("count":(.*?)})`)
//	getCount := countMatch.FindAllStringSubmatch(resContent, -1)
//
//	fmt.Println(getCount, len(getCount[0]))
//	//fmt.Println(myId,len(myId[0]))
//	count, err := strconv.Atoi(getCount[0][2])
//	if err != nil {
//		panic(err)
//	}
//
//	ParasResponeData(resContent, id)
//
//	pageNum := count/10 + 1
//	for i := 2; i < pageNum; i++ {
//		page := strconv.Itoa(i)
//		url = "http://visitor.fanxing.kugou.com/VServices/Video.OfflineVideoService.getVideoList/" + id + "-" + page + "-0-10/"
//		fmt.Println(url)
//		aim, err := ParaseURL(url)
//		if err != nil {
//			panic(err)
//		}
//		ParasResponeData(aim, id)
//	}
//
//}

func ParasURL(url string) (string, error) {
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}
	if res.StatusCode != http.StatusOK {
		return "", errors.New("系统错误，错误码" + strconv.Itoa(res.StatusCode))
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

func ParasResponseData(resContent string, id string, stop, proceed chan struct{}, errChan chan public.ErrorMessage) {
	public.Cond.L.Lock()
	defer public.Cond.L.Unlock()
	title := regexp.MustCompile(`("title":"(.*?)")`)
	actor := regexp.MustCompile(`("actor":"(.*?)")`)
	hashl := regexp.MustCompile(`("MvHash":"(.*?)")`)
	download := regexp.MustCompile(`("downurl":"(.*?)")`)
	titleList := title.FindAllStringSubmatch(resContent, -1)
	actorList := actor.FindAllStringSubmatch(resContent, -1)

	//fmt.Println("actor is:", actorList, len(actorList))
	for key, _ := range titleList {
		select {
		case <-stop:
			proceed <- struct{}{}
		default:
		}
		videoData, err := util.SearchVideoByName(titleList[key][0])
		if err != nil {
			util.SendError(errChan, err)
			continue
		}
		videoHash := hashl.FindAllStringSubmatch(videoData, -1)
		if len(videoHash) < 1 {
			util.SendError(errChan, errors.New("未匹配到视频哈希"))
			continue
		}

		//log.Println("video hash:",videoHash)
		videoInfo, err := util.ObtainVideoDownUrl(videoHash[0][2])
		if err != nil {
			util.SendError(errChan, err)
			continue
		}
		//log.Println("video info : ",videoInfo)
		videoDown := download.FindAllStringSubmatch(videoInfo, -1)
		if len(videoDown) < 1 {
			util.SendError(errChan, err)
			continue
		}

		downurl := strings.Replace(videoDown[0][2], `\`, "", -1)
		//go func(key int) {
		err = util.DownVideo(downurl, id, titleList[key][2], actorList[key][2])
		if err != nil {
			util.SendError(errChan, err)
		}
		util.SendError(errChan, errors.New("成功下载视频："+titleList[key][2]))
		//}(key)
	}
}

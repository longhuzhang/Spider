package client

import (
	"Spider/parase_url"
	"Spider/public"
	"strconv"
)

type Work struct {
	prevFilePath string
	path         string
}

//
//func (mw *Work) getUserIdAndOperate() {
//	defer func() {
//		if err := recover(); err != nil {
//			errMsg := ""
//			switch errType := reflect.TypeOf(err); errType {
//			case reflect.TypeOf(error_operate.NoticeError{}):
//				errMsg = err.(error_operate.NoticeError).Error()
//			default:
//				errMsg = "未知错误"
//			}
//			userId.SetText(errMsg)
//		}
//	}()
//	users := userId.Text()
//
//	if users == "" {
//		//userId.SetText("开始运行")
//		return
//	}
//
//	if strings.Contains(users, "shoudaoshuju") {
//		panic(error_operate.NoticeError{"传入玩家编号出错"})
//	}
//
//	//userId.SetText("")
//
//	runStopBtn.SetText("暂停")
//
//	userIds := strings.Split(users, ",")
//
//	fmt.Println(userIds)
//	for _, userId := range userIds {
//		go urlWork(userId)
//	}
//}
//
//func urlWork(MyUserId string) {
//	defer func() {
//		if err := recover(); err != nil {
//			errMsg := ""
//			switch errType := reflect.TypeOf(err); errType {
//			case reflect.TypeOf(error_operate.NoticeError{}):
//				errMsg = err.(error_operate.NoticeError).Error()
//			default:
//				errMsg = "未知错误"
//			}
//			userId.SetText(errMsg)
//		}
//	}()
//	parase_url.ParseId(MyUserId)
//}

func (mw *Work) dynamicList(stop, proceed chan struct{}, errMessage chan public.ErrorMessage) {
	urls := make([]string, 0)
	for i := 0; i < 10; i++ {
		strNum := strconv.Itoa(i)
		url := "https://fx.service.kugou.com/VServices/Video.OfflineVideoRankService.getDynamicRank/" + strNum + "/?jsonpcallback=jsonphttpsfxservicekugoucomVServicesVideoOfflineVideoRankServicegetDynamicRank" + strNum + "jsonpcallback"
		urls = append(urls, url)
	}
	go parase_url.ParasListUrl(urls, "动态榜单", stop, proceed, errMessage)
}

func (mw *Work) soaringList(stop, proceed chan struct{}, errMessage chan public.ErrorMessage) {
	urls := make([]string, 0)
	for i := 0; i < 10; i++ {
		strNum := strconv.Itoa(i)
		url := "https://fx.service.kugou.com/VServices/Video.OfflineVideoService.getWeeklyRisingRank/" + strNum + "/?jsonpcallback=jsonphttpsfxservicekugoucomVServicesVideoOfflineVideoServicegetWeeklyRisingRank" + strNum + "jsonpcallback"
		//https://fx.service.kugou.com/VServices/Video.OfflineVideoService.getWeeklyRisingRank/0/?jsonpcallback=jsonphttpsfxservicekugoucomVServicesVideoOfflineVideoServicegetWeeklyRisingRank0jsonpcallback
		urls = append(urls, url)
	}
	go parase_url.ParasListUrl(urls, "飙升榜单", stop, proceed, errMessage)
}

func (mw *Work) dayList(strTime string, stop, proceed chan struct{}, errMessage chan public.ErrorMessage) {
	urls := make([]string, 0)
	for i := 0; i < 2; i++ {
		strNum := strconv.Itoa(i)
		url := "https://fx.service.kugou.com/VServices/Video.OfflineVideoService.getDailyRank/" + strNum + "-" + strTime + "/?jsonpcallback=jsonphttpsfxservicekugoucomVServicesVideoOfflineVideoServicegetDailyRank" + strNum + strTime + "jsonpcallback"
		urls = append(urls, url)
	}
	go parase_url.ParasListUrl(urls, "日榜", stop, proceed, errMessage)
}

func (mw *Work) mouthList(strTime string, stop, proceed chan struct{}, errMessage chan public.ErrorMessage) {
	urls := make([]string, 0)
	for i := 0; i < 2; i++ {
		strNum := strconv.Itoa(i)
		url := "https://fx.service.kugou.com/VServices/Video.OfflineVideoService.getMonthRank/" + strNum + "-" + strTime + "/?jsonpcallback=jsonphttpsfxservicekugoucomVServicesVideoOfflineVideoServicegetMonthRank" + strNum + strTime + "jsonpcallback"
		//https://fx.service.kugou.com/VServices/Video.OfflineVideoService.getMonthRank/1/?jsonpcallback=jsonphttpsfxservicekugoucomVServicesVideoOfflineVideoServicegetMonthRank1jsonpcallback
		urls = append(urls, url)
	}
	go parase_url.ParasListUrl(urls, "月榜", stop, proceed, errMessage)
}

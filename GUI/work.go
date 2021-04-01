package GUI

import (
	"VideoSpider/paraseURL"
	"VideoSpider/public"
	"VideoSpider/public/error_operate"
	"fmt"
	"github.com/lxn/walk"
	"reflect"
	"strconv"
	"strings"
)

type MyMainWindow struct {
	*walk.MainWindow
	prevFilePath string
	path         string
}

func (mw *MyMainWindow) getUserIdAndOperate() {
	defer func() {
		if err := recover(); err != nil {
			errMsg := ""
			switch errType := reflect.TypeOf(err); errType {
			case reflect.TypeOf(error_operate.NoticeError{}):
				errMsg = err.(error_operate.NoticeError).Error()
			default:
				errMsg = "未知错误"
			}
			userId.SetText(errMsg)
		}
	}()
	users := userId.Text()

	if users == "" {
		//userId.SetText("开始运行")
		return
	}

	if strings.Contains(users, "shoudaoshuju") {
		panic(error_operate.NoticeError{"传入玩家编号出错"})
	}

	//userId.SetText("")

	runStopBtn.SetText("暂停")

	userIds := strings.Split(users, ",")

	fmt.Println(userIds)
	for _, userId := range userIds {
		go urlWork(userId)
	}
}

func urlWork(MyUserId string) {
	defer func() {
		if err := recover(); err != nil {
			errMsg := ""
			switch errType := reflect.TypeOf(err); errType {
			case reflect.TypeOf(error_operate.NoticeError{}):
				errMsg = err.(error_operate.NoticeError).Error()
			default:
				errMsg = "未知错误"
			}
			userId.SetText(errMsg)
		}
	}()
	paraseURL.ParseId(MyUserId)
}

func (mw *MyMainWindow) dynamicList() {
	urls := make([]string, 0)
	for i := 0; i < 4; i++ {
		strNum := strconv.Itoa(i)
		url := "https://fx.service.kugou.com/VServices/Video.OfflineVideoRankService.getDynamicRank/" + strNum + "/?jsonpcallback=jsonphttpsfxservicekugoucomVServicesVideoOfflineVideoRankServicegetDynamicRank" + strNum + "jsonpcallback"
		urls = append(urls, url)
	}
	go paraseURL.ParasListUrl(urls, "动态榜单")
}

func (mw *MyMainWindow) soaringList() {
	urls := make([]string, 0)
	for i := 0; i < 2; i++ {
		strNum := strconv.Itoa(i)
		url := "https://fx.service.kugou.com/VServices/Video.OfflineVideoService.getWeeklyRisingRank/" + strNum + "/?jsonpcallback=jsonphttpsfxservicekugoucomVServicesVideoOfflineVideoServicegetWeeklyRisingRank" + strNum + "jsonpcallback"
		//https://fx.service.kugou.com/VServices/Video.OfflineVideoService.getWeeklyRisingRank/0/?jsonpcallback=jsonphttpsfxservicekugoucomVServicesVideoOfflineVideoServicegetWeeklyRisingRank0jsonpcallback
		urls = append(urls, url)
	}
	go paraseURL.ParasListUrl(urls, "飙升榜单")
}

func (mw *MyMainWindow) dayList() {
	urls := make([]string, 0)
	strTime := timeText.Text()
	timeText.SetText(fmt.Sprintf("爬取日期为：%s的视频", strTime))
	for i := 0; i < 2; i++ {
		strNum := strconv.Itoa(i)
		url := "https://fx.service.kugou.com/VServices/Video.OfflineVideoService.getDailyRank/" + strNum + "-" + strTime + "/?jsonpcallback=jsonphttpsfxservicekugoucomVServicesVideoOfflineVideoServicegetDailyRank" + strNum + strTime + "jsonpcallback"
		urls = append(urls, url)
	}
	go func() {
		defer func() {
			if err := recover(); err != nil {
				errMsg := ""
				switch errType := reflect.TypeOf(err); errType {
				case reflect.TypeOf(error_operate.NoticeError{}):
					errMsg = err.(error_operate.NoticeError).Error()
				default:
					errMsg = "未知错误"
				}
				timeText.SetText(errMsg)
			}
		}()
		paraseURL.ParasListUrl(urls, "日榜")
	}()
}

func (mw *MyMainWindow) mouthList() {
	strTime := timeText.Text()
	urls := make([]string, 0)
	for i := 0; i < 2; i++ {
		strNum := strconv.Itoa(i)
		url := "https://fx.service.kugou.com/VServices/Video.OfflineVideoService.getMonthRank/" + strNum + "-" + strTime + "/?jsonpcallback=jsonphttpsfxservicekugoucomVServicesVideoOfflineVideoServicegetMonthRank" + strNum + strTime + "jsonpcallback"
		urls = append(urls, url)
	}
	go paraseURL.ParasListUrl(urls, "月榜")
}

func (mw *MyMainWindow) offlinePauseRecover() {
	//users := userId.Text()
	public.MyChannel <- 1
	pauseRecoverBtn.SetText("退出成功")

}
func (mw *MyMainWindow)offlineStop()  {
	str := pauseStopBtn.Text()
	if str == "暂停" {
		public.MyChannel <-2
		pauseStopBtn.SetText("继续")
	}else {
		public.Cond.Signal()
		pauseStopBtn.SetText("暂停")
	}
}

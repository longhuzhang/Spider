package public

import "sync"

var MyChannel = make(chan int)
var Locker = new(sync.Mutex)
var Cond = sync.NewCond(Locker)

var SearchUrl = `https://mvsearch.kugou.com/mv_search?keyword={keyword}&page=1&pagesize={pagesize}&platform=WebFilter`
var DownloadUrl = `http://trackermv.kugou.com/interface/index/cmd=100&hash={hash}&key={key}&pid=6&ext=mp4&ismp3=0`

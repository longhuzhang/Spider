package public

import "sync"

var MyChannel = make( chan int)
var Locker = new(sync.Mutex)
var Cond  = sync.NewCond(Locker)

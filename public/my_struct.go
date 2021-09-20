package public

import "context"

type MyContext struct {
	context.Context
	UserList    []string
	wokeList    []string
}

type MyActor struct {
	Actor  string`json:"actor"`
}

type MyTitle struct {
	Title string`json:"title"`
}


type UrlRespond struct {
	ServerTime int      `json:"servertime"`
	Callback   []string `json:"callback"`
	data       string   `json:"data"`
	Status     int      `json:"status"`
	ErrorCode  string   `json:"errorcode"`
	ErrorNo    int      `json:"errorno"`
	Code       int      `json:"code"`
	Msg        string   `json:"msg"`
	Time       int      `json:"time"`
	Times      int      `json:"times"`
}
package parase_url

import (
	"Spider/public"
	"Spider/util"
)

func ParasListUrl(urls []string, name string, stop, proceed chan struct{}, errMessage chan public.ErrorMessage) {
	for _, url := range urls {
		resContent, err := ParasURL(url)
		if err != nil {
			util.SendError(errMessage, err)
		}
		ParasResponseData(resContent, name, stop, proceed, errMessage)
	}
}

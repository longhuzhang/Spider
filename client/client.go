package client

import (
	"Spider/public"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"log"
	"os"
	"strconv"
)

var Stop chan struct{}
var Proceed chan struct{}
var DownState int

var ErrorMessage chan public.ErrorMessage

func InitClient() {
	myApp := app.New()
	window := myApp.NewWindow("Spider")
	window.Resize(fyne.NewSize(500, 400))
	window.CenterOnScreen()

	inputText := widget.NewMultiLineEntry()
	inputText.SetPlaceHolder("Enter code...")
	inputText.Resize(fyne.NewSize(500, 200))

	startButton := widget.NewButton("start", func() {
		if DownState == 1 {
			<-Proceed
			DownState = 0
		}
	})
	stopButton := widget.NewButton("stop", func() {
		if DownState == 0 {
			Stop <- struct{}{}
			DownState = 1
		}
	})
	quitButton := widget.NewButton("quit", func() {
		os.Exit(0)
	})

	logText := widget.NewEntry()
	logText.MultiLine = true
	logText.Resize(fyne.NewSize(500, 200))

	work := &Work{}
	Stop = make(chan struct{})
	Proceed = make(chan struct{})
	ErrorMessage = make(chan public.ErrorMessage, 20)
	dynamicButton := widget.NewButton("dynamic rank", func() {
		work.dynamicList(Stop, Proceed, ErrorMessage)
	})
	soarButton := widget.NewButton("soar rank", func() {
		work.soaringList(Stop, Proceed, ErrorMessage)
	})
	dayButton := widget.NewButton("day rank", func() {
		day := inputText.Text
		work.dayList(day, Stop, Proceed, ErrorMessage)
	})
	weekButton := widget.NewButton("mouth rank", func() {
		mouth := inputText.Text
		work.mouthList(mouth, Stop, Proceed, ErrorMessage)
	})

	var message public.ErrorMessage
	go func() {
		for {
			message = <-ErrorMessage
			log.Println("协程获取的信息：", message)
			logText.SetText(logText.Text + "\n" + message.File + ":" + strconv.Itoa(message.Line) + ":" + message.Err.Error())
		}
	}()

	operateContain := container.NewGridWithColumns(4, dynamicButton, soarButton, dayButton, weekButton)

	buttonContain := container.NewGridWithColumns(3, startButton, stopButton, quitButton)
	content := container.NewVBox(operateContain, inputText, buttonContain, logText)

	window.SetContent(content)
	window.ShowAndRun()
}

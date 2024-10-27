package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"strconv"
	"time"
)

type Timer struct {
	sec  uint8
	min  uint8
	hour uint8
}

func (t *Timer) AddSec() {
	t.sec++

	if t.sec == 60 {
		t.sec = 0
		t.min++
	}
	if t.min == 60 {
		t.min = 0
		t.hour++
	}
	if t.hour == 24 {
		t.hour = 0
	}
}

func (t *Timer) Reset() {
	t.sec = 0
	t.min = 0
	t.hour = 0
}

func main() {
	timer := Timer{sec: 0, min: 0, hour: 0}

	timerApp := app.New()
	timerWindow := timerApp.NewWindow("Timer")

	var clockEnable = false
	clock := widget.NewLabel("")
	clock.TextStyle = fyne.TextStyle{Bold: true}
	clock.Alignment = fyne.TextAlignCenter

	timerString := fmt.Sprintf("%02d:%02d:%02d", timer.hour, timer.min, timer.sec)
	clock.SetText(timerString)

	convertedHours := widget.NewLabel("0.00")
	convertedHours.Alignment = fyne.TextAlignCenter
	convertedHoursFloat := 0.0

	clockEnableButton := widget.NewButton("Start", func() {
		clockEnable = true
	})
	clockDisableButton := widget.NewButton("Stop", func() {
		clockEnable = false
	})
	clockResetButton := widget.NewButton("Reset", func() {
		clockEnable = false
		timer.Reset()
		convertedHoursFloat = 0.0
		convertedHours.SetText(strconv.FormatFloat(convertedHoursFloat, 'f', 2, 64))
	})

	go func() {
		for {
			timerString = fmt.Sprintf("%02d:%02d:%02d", timer.hour, timer.min, timer.sec)
			clock.SetText(timerString)
			if clockEnable {
				time.Sleep(1 * time.Second)

				if clockEnable {
					timer.AddSec()
					convertedHoursFloat = float64(timer.hour) + float64(timer.min)/60 + float64(timer.sec)/3600
					convertedHours.SetText(strconv.FormatFloat(convertedHoursFloat, 'f', 2, 64))
				}
			}
		}
	}()

	timerButtonBox := container.NewHBox(
		clockEnableButton,
		clockDisableButton,
		clockResetButton)

	timerBox := container.NewVBox(
		clock,
		convertedHours,
		timerButtonBox)

	timerWindow.SetContent(timerBox)

	timerWindow.ShowAndRun()
}

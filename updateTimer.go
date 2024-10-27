package main

import (
	"fyne.io/fyne/v2/widget"
	"time"
)

func updateTimer(clock *widget.Label) {
	clock.SetText(time.Now().Format("15:04:05"))
}

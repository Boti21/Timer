package main

import "log"

/*
	func toggleClock(clockEnable *atomic.Bool) {
		clockEnable.Store(!clockEnable.Load())
		log.Println("clockEnable:", clockEnable.Load())
	}
*/
func toggleClock(clockEnable *bool) {
	*clockEnable = !*clockEnable
	log.Println("clockEnable:", *clockEnable)
}

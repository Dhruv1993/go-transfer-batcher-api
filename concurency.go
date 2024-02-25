package app

import "log"

func RunInBackground(f func()) {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Panic occurred in background goroutine: %v", r)
			}
		}()
		f()
	}()
}

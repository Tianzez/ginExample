package utils

import (
	"runtime"
	"runtime/debug"

	log "github.com/cihub/seelog"
)

func Async(f func()) {
	go func() {
		defer func() {
			if eface := recover(); eface != nil {
				log.Errorf("recover called, eface is %v", eface)
				if _, ok := eface.(runtime.Error); ok {
					log.Errorf("Internal error stack: %s", debug.Stack())
				}
			}
		}()
		f()
	}()
}

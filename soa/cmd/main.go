// cmd/runall/main.go
package main

import (
	"log"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	start := func(name string, run func()) {
		wg.Add(1)
		go func() {
			defer wg.Done()
			log.Printf("starting %s…", name)
			run()
		}()
	}

	start("service-a", servicea.Run)
	start("service-b", serviceb.Run)
	start("watermark", watermark.Run)

	wg.Wait()
}

package parent

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"sync"
	"time"
)

func Run(children int, bytesPerSecond uint64) {
	var wg sync.WaitGroup
	wg.Add(children)
	for i := 0; i < children; i++ {
		i := i
		go func() {
			defer wg.Done()
			runChild(i, bytesPerSecond)
		}()
	}
	go ps()
	wg.Wait()
}

func ps() {
	t := time.NewTicker(2 * time.Second)
	defer t.Stop()

	for range t.C {
		if err := run("ps", "axo", "pid,rss,args", "-H"); err != nil {
			log.Printf("ps: %v", err)
			return
		}
	}
}

func runChild(childNum int, bytesPerSecond uint64) {
	err := run("go", "run", ".",
		"-child-number", fmt.Sprint(childNum),
		"-rate", fmt.Sprint(bytesPerSecond))
	if err != nil {
		log.Printf("child %d: %v", childNum, err)
	}
}

func run(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Start(); err != nil {
		return err
	}
	return cmd.Wait()
}

package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"regexp"

	"github.com/ActiveState/tail"
)

var follow = false
const version = "v0.1.0"

func main() {
	matchers := processArgs()

	appLog := os.Getenv("BOXEN_APPLICATION_LOG")
	if appLog == "" {
		appLog = os.Getenv("APPLICATION_LOG")
	}
	if appLog == "" {
		appLog = "/var/log/application.log"
	}

	t, err := tail.TailFile(appLog, tail.Config{
		Follow: follow,
		ReOpen: follow,
		MustExist: true,
	})

	if err != nil {
		log.Fatal(err.Error())
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	go func(t *tail.Tail) {
		sig := <-ch
		fmt.Println("received", sig)
		t.Stop()
	}(t)

	for line := range t.Lines {
		matched := true

		for _, matcher := range matchers {
			if matched == false {
				break
			} else {
				matched, err = regexp.MatchString(matcher, line.Text)
				if err != nil {
					log.Fatal(err.Error())
				}
			}
		}

		if matched == true {
			fmt.Println(line.Text)
		}
	}
}

func processArgs() (matchers []string) {
	for _, arg := range os.Args[1:] {
		switch arg {
		case "--version", "-v":
			fmt.Println("log-me", version)
			os.Exit(0)
		case "--tail", "-t":
			follow = true
		default:
			matchers = append(matchers, massageMatcher(arg))
		}
	}
	return
}

func massageMatcher(matcher string) (massaged string) {
	massager, err := regexp.Compile("\\*")
	if err != nil {
		log.Fatal(err.Error())
	}

	massaged = massager.ReplaceAllLiteralString(matcher, ".*")
	return
}
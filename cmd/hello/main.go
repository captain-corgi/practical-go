package main

import (
	"fmt"
	"strings"
)

var (
	urls = []string{
		"",
	}
	watching = "https://dave.cheney.net/practical-go"
	done     = []string{
		"",
		"https://dave.cheney.net/paste/gopherchina-2019-testing-talk.pdf?fbclid=IwAR2C-fGo3ztJuxETInCPMPgVXzInsKLBKycbeEHb30yMPEsdVCm7qjy6Gfk",
		"https://dave.cheney.net/practical-go/presentations/qcon-china.html#_concurrency",
	}
)

func main() {
	fmt.Printf("Watch list:\t %+v\n", strings.Join(urls, "\n\t"))
	fmt.Printf("Watching:\n \t%s\n", watching)
	fmt.Printf("Finished:\t %+v\n", strings.Join(done, "\n\t"))
}

package main

/*
read zimbra logs retrieving data that is interesting for further analysis
*/
import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

var zlog *string = flag.String("log", "/var/log/zimbra.log", "location of zimbra log file")
var v *bool = flag.Bool("v", false, "verbosity true or false")

func usage() {
	flag.PrintDefaults()
	os.Exit(-1)
}

func spinner(p int) int {
	if p >= 4 {
		return 0
	}
	p++
	return p
}

func saslUserName(l *[]byte, v *bool) {
	if *v == true {
		fmt.Println(string(*l))
	}
}

func nonDel(l *[]byte, v *bool) {
	if *v == true {
		fmt.Println(string(*l))
	}
}

func authFail(l *[]byte, v *bool) {
	if *v == true {
		fmt.Println(string(*l))
	}
}

func statDefer(l *[]byte, v *bool) {
	if *v == true {
		fmt.Println(string(*l))
	}
}

func main() {
	s := 0
	sSlices := []string{"-", "\\", "|", "/", "|"}
	flag.Parse()
	fi, err := os.Open(*zlog)
	defer fi.Close()
	if err != nil {
		log.Fatal(err)
	}
	bfi := bufio.NewReader(fi)
	for {
		line, _, err := bfi.ReadLine()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		switch {
		case strings.Contains(string(line), "sasl_username="):
			saslUserName(&line, v)
		case strings.Contains(string(line), "non-delivery notification:"):
			nonDel(&line, v)
		case strings.Contains(string(line), "authentication failed for"):
			authFail(&line, v)
		case strings.Contains(string(line), "status=deferred"):
			statDefer(&line, v)
		default:
			s = spinner(s)
			fmt.Printf(sSlices[s])
			fmt.Printf("\b")
		}

	}
}

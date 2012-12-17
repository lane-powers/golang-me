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

func usage() {
	flag.PrintDefaults()
	os.Exit(-1)
}

func main() {
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
			fmt.Println(string(line))
		case strings.Contains(string(line), "non-delivery notification:"):
			fmt.Println(string(line))
		default:
			fmt.Printf(".\b")
		}

	}
}

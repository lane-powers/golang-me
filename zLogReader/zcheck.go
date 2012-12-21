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
	"regexp"
	"sort"
	"strings"
)

var zlog *string = flag.String("log", "/var/log/zimbra.log", "location of zimbra log file")
var minCount *int = flag.Int("min", 5, "minimum entry counts for a record to be reported")
var regCleaner = regexp.MustCompile("[><,=\\]]|sasl_username|(.*\\[|)")

func usage() {
	flag.PrintDefaults()
	os.Exit(-1)
}

func spinner(p *int) {
	if *p >= 3 {
		*p = 0
	}
	*p++
}

func saslUserName(l *string, mySaslUser map[string]int, mySaslSource map[string]int) {
	z := strings.Fields(*l)
	addy := regCleaner.ReplaceAllString(z[8], "")
	src := regCleaner.ReplaceAllString(z[6], "")
	aval, aok := mySaslUser[addy]
	// track it by sending email
	if aok {
		mySaslUser[addy] = int(aval) + 1
	} else {
		mySaslUser[addy] = 1
	}
	// now track it by source IP
	sval, sok := mySaslSource[src]
	if sok {
		mySaslSource[src] = int(sval) + 1
	} else {
		mySaslSource[src] = 1
	}
}

func authFail(l *string, myAuthFail map[string]int) {
	z := strings.Fields(*l)
	addy := regCleaner.ReplaceAllString(z[6], "")
	val, ok := myAuthFail[addy]
	if ok {
		myAuthFail[addy] = int(val) + 1
	} else {
		myAuthFail[addy] = 1
	}
}

func statDefer(l *string, myDefList map[string]int) {
	z := strings.Fields(*l)
	addy := regCleaner.ReplaceAllString(z[6], "")
	val, ok := myDefList[addy]
	if ok {
		myDefList[addy] = int(val) + 1
	} else {
		myDefList[addy] = 1
	}
}

func sortPrint(myMinCount *int, inMap map[string]int) {
	var i int
	myList := make([]int, 0, 10000)
	for _, v := range inMap {
		if v >= *myMinCount {
			myList = append(myList, v)
		}
	}
	sort.Ints(myList)
	for len(myList) > 0 {
		i, myList = myList[len(myList)-1], myList[:len(myList)-1]
		for k, v := range inMap {
			if v == i {
				fmt.Printf("  %-65s\t:%d\n", k, v)
				delete(inMap, k)
			}
		}
	}
}

func main() {
	s := 0
	sSlices := []string{"|", "/", "_", "\\"}
	defList := make(map[string]int)
	saslUser := make(map[string]int)
	saslSource := make(map[string]int)
	failList := make(map[string]int)
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
		sline := string(line)
		switch {
		case strings.Contains(sline, "sasl_username="):
			saslUserName(&sline, saslUser, saslSource)
		case strings.Contains(sline, "authentication failed for"):
			if strings.Contains(sline, "auth_zimbra:") {
				authFail(&sline, failList)
			}
		case strings.Contains(sline, "status=deferred"):
			statDefer(&sline, defList)
		default:
			spinner(&s)
			fmt.Printf("%s", sSlices[s])
			fmt.Printf("\b")
		}
	}

	fmt.Printf("Deferred connections >= %d:\n", *minCount)
	sortPrint(minCount, defList)
	fmt.Printf("SASL user auths per username >= %d:\n", *minCount)
	sortPrint(minCount, saslUser)
	fmt.Printf("SASL user auths per source IP >= %d:\n", *minCount)
	sortPrint(minCount, saslSource)
	fmt.Printf("Failed connections per username >= %d\n", *minCount)
	sortPrint(minCount, failList)
}

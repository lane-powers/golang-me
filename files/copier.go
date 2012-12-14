package main

import (
	"bufio"
	"io"
	"os"
)

func main() {
	fi, err := os.Open("in")
	// bail if error
	if err != nil {
		panic(err)
	}
	defer fi.Close() // beautiful go... close the fi when your all done
	in := bufio.NewReader(fi)

	fo, err := os.Create("out") // Create instead of open since we want to truncate/new
	if err != nil {
		panic(err)
	}
	defer fo.Close() // seriously, I love this language
	out := bufio.NewWriter(fo)

	buf := make([]byte, 1024) // now that I did the slices stuff, this makes since
	for {
		n, err := in.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		} // if we have an error besides EOF, bail
		if n == 0 {
			break
		} // abort when we don't have data

		o, err := out.Write(buf[:n])
		if err != nil {
			panic(err)
		} else if o != n {
			panic("error in writing, mismatch bytes in and bytes out")
		}

	}
	err = out.Flush()
	if err != nil {
		panic(err)
	}
}

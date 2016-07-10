package main

import (
	"flag"
	"fmt"
	"strconv"
	"time"
)

func main() {
	printHumanReadable := flag.Bool("h", false, "Prints human readable date")

	flag.Parse()

	t := time.Now()

	if len(flag.Args()) == 1 {
		millis, err := strconv.ParseInt(flag.Arg(0), 10, 64)
		if err != nil {
			panic(err)
		}
		t = time.Unix(0, millis*1000000)
	}

	fmt.Printf("Epoch millis:\t%v\n", t.UnixNano()/1000000)

	if *printHumanReadable {
		utc := t.In(time.UTC)
		fmt.Printf("UTC:\t%v\n", formatTime(&utc))
		fmt.Printf("Local:\t%v\n", formatTime(&t))
	}

}

func formatTime(t *time.Time) string {
	return t.Format("Mon Jan _2 2006 03:04:05.000 PM")
}

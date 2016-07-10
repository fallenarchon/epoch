package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	printHumanReadable := flag.Bool("h", false, "Prints human readable date")

	flag.Parse()

	t := time.Now()

	info, err := os.Stdin.Stat()
	checkErr(err)

	//If data is being piped in, parse millis from that
	if info.Size() > 0 {
		t = convertStringMillisToTime(fetchInputFromPipe())
	} else if len(flag.Args()) == 1 {
		t = convertStringMillisToTime(flag.Arg(0))
	}

	fmt.Printf("Epoch millis:\t%v\n", t.UnixNano()/1000000)

	if *printHumanReadable {
		utc := t.In(time.UTC)
		fmt.Printf("UTC:\t%v\n", formatTime(&utc))
		fmt.Printf("Local:\t%v\n", formatTime(&t))
	}

}

//Currently only fetches the first token in the pipe.  Future work could spit out data for each token
func fetchInputFromPipe() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func convertStringMillisToTime(inputMillis string) time.Time {
	millis, err := strconv.ParseInt(inputMillis, 10, 64)
	checkErr(err)
	return time.Unix(0, millis*1000000)
}

func formatTime(t *time.Time) string {
	return t.Format("Mon Jan _2 2006 03:04:05.000 PM")
}

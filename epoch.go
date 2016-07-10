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
	stripHumanReadable := flag.Bool("s", false, "Strips human readable date info")
	date := flag.String("d", "", "Supply date to find epoch milliseconds <MM/DD/YYYY hh:mm:ss>")

	flag.Parse()

	var err error
	t := time.Now()

	if *date != "" {
		t, err = time.Parse("01/02/2006 15:04:05", *date)
		if err != nil {
			t, err = time.Parse("01/02/2006 15:04:05.000", *date)
		}
		checkErr(err)
	} else {
		info, err := os.Stdin.Stat()
		checkErr(err)

		//If data is being piped in, parse millis from that
		if info.Size() > 0 {
			t = convertStringMillisToTime(fetchInputFromPipe())
		} else if len(flag.Args()) == 1 {
			t = convertStringMillisToTime(flag.Arg(0))
		}
	}

	if *stripHumanReadable {
		fmt.Printf("%v\n", t.UnixNano()/1000000)
	} else {
		fmt.Printf("Epoch millis:\t%v\n", t.UnixNano()/1000000)
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

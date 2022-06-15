package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/micmonay/keybd_event"
)

func main() {
	k := flag.Int("k", 0x5B+0xFFF, "key") // LWIN
	t := flag.Int("t", 3600, "seconds")
	linux := flag.Bool("l", false, "linux")
	delay := flag.Duration("d", 0, "delay")
	quiet := flag.Bool("q", false, "quiet")
	noDelay := flag.Bool("n", false, "no delay")
	flag.Parse()

	if *quiet {
		log.SetOutput(io.Discard)
		log.SetFlags(0)
	}

	if *delay == 0 && !*noDelay {
		var s string
		var err error
		fmt.Print("delay: ")
		fmt.Scanf("%s", &s)
		*delay, err = time.ParseDuration(s)
		if err != nil && len(s) > 0 {
			log.Print(err)
		}
	}

	if *delay != 0 {
		log.Printf("Delay for %.2f seconds.", delay.Seconds())
		time.Sleep(*delay)
	}

	kb, err := keybd_event.NewKeyBonding()
	if err != nil {
		log.Fatal(err)
	}

	if *linux {
		time.Sleep(2 * time.Second) // https://github.com/micmonay/keybd_event/issues/25
	}

	kb.SetKeys(*k)

	defer kb.Release()

	log.Printf("Pressing %d for %d seconds...\n", *k, *t)

	for i := 0; i < *t; i++ {
		if err := kb.Press(); err != nil {
			log.Print(err)
		}
		time.Sleep(500 * time.Millisecond)
		if err := kb.Release(); err != nil {
			log.Print(err)
		}
		time.Sleep(500 * time.Millisecond)
	}
}

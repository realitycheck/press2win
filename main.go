package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/micmonay/keybd_event"
)

func main() {
	k := flag.Int("k", 0x5B+0xFFF, "key") // LWIN
	t := flag.Int("t", 3600, "seconds")
	linux := flag.Bool("l", false, "linux")
	flag.Parse()

	fmt.Printf("Pressing %d for %d seconds...\n", *k, *t)

	kb, err := keybd_event.NewKeyBonding()
	if err != nil {
		panic(err)
	}

	if *linux {
		time.Sleep(2 * time.Second) // https://github.com/micmonay/keybd_event/issues/25
	}

	kb.SetKeys(*k)

	defer kb.Release()

	for i := 0; i < *t; i++ {
		if err := kb.Press(); err != nil {
			println(err)
		}
		time.Sleep(500 * time.Millisecond)
		if err := kb.Release(); err != nil {
			println(err)
		}
		time.Sleep(500 * time.Millisecond)
	}
}

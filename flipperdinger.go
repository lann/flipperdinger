package main

import (
	"fmt"
	"log"
	"os"
	"github.com/lann/mpris2"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("usage: %s <cmd>")
	}

	conn, err := mpris2.Connect()
	if err != nil {
		log.Fatal(err)
	}

	mp, err := conn.GetFirstMediaPlayer()
	if err != nil {
		log.Fatal(err)
	}

	switch cmd := os.Args[1]; cmd {
	case "play":
		err = mp.Play()
	case "pause":
		err = mp.Pause()
	case "playpause":
		err = mp.PlayPause()
	case "stop":
		err = mp.Stop()
	case "prev":
		err = mp.Previous()
	case "next":
		err = mp.Next()
	default:
		err = fmt.Errorf("Unknown cmd %s", cmd)
	}
	if err != nil {
		log.Fatal(err)
	}
	
	return 
}
package main

import (
	"fmt"
	"log"
	"os"
	"sort"
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
		
    case "metadata":
		data, err := mp.Metadata()
		if err == nil {
			if len(os.Args) > 2 {
				for _, k := range os.Args[2:] {
					fmt.Println(data[k])
				}
			} else {
				keys := make([]string, len(data))
				i := 0
				maxLen := 0
				for k := range data {
					keys[i] = k
					i++
					if len(k) > maxLen {
						maxLen = len(k)
					}
				}

				sort.Strings(keys)

				for _, k := range keys {
					fmt.Printf("%-*s  %v\n", maxLen, k, data[k])
				}
			}
		}
		
	default:
		err = fmt.Errorf("Unknown cmd %s", cmd)
	}
	if err != nil {
		log.Fatal(err)
	}
	
	return 
}
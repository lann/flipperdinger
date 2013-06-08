package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"github.com/lann/mpris2"
)

func main() {
	log.SetFlags(0)
	
	if len(os.Args) < 2 {
		log.Fatalf("usage: %s <cmd> [param...]", os.Args[0])
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

	case "seek":
		if len(os.Args) != 3 {
			log.Fatalf("usage: %s seek <offset>", os.Args[0])
		}

		var offset int64
		_, err = fmt.Sscan(os.Args[2], &offset)
		if err == nil {
			err = mp.Seek(offset)
		}

	case "open":
		if len(os.Args) != 3 {
			log.Fatalf("usage: %s open <uri>", os.Args[0])
		}

		err = mp.OpenUri(os.Args[2])

	case "identity":
		var identity string
		identity, err = mp.Identity()
		if err == nil {
			fmt.Println(identity)
		}

	case "desktop":
		var name string
		name, err = mp.DesktopEntry()
		if len(name) > 0 {
			fmt.Printf("/usr/share/applications/%s.desktop\n", name)
		}
		
	case "status":
		var status string
		status, err = mp.PlaybackStatus()
		if err == nil {
			fmt.Println(status)
		}
		
	case "pos":
		var pos int64 
		pos, err = mp.Position()
		if err == nil {
			fmt.Println(pos)
		}
		
    case "metadata":
		var data mpris2.Metadata
		data, err = mp.Metadata()
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
		log.Fatalf("Unknown cmd %s", cmd)
	}
	
	if err != nil {
		log.Fatal(err)
	}
	
	return 
}
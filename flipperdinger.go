package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"github.com/lann/mpris2"
)

func main() {
	log.SetFlags(0)

	var playerName string
	flag.StringVar(&playerName, "player", "", "media player to control")
	flag.Parse()
	
	cmd := flag.Arg(0)
	if cmd == "" {
		log.Fatalf("Usage: %s [<options>] <cmd> [<params>]", os.Args[0])
	}

	conn, err := mpris2.Connect()
	if err != nil {
		log.Fatal(err)
	}

	if cmd == "list" {
		names, err := conn.ListMediaPlayers()
		if err != nil {
			log.Fatal(err)
		}

		for _, name := range names {
			fmt.Println(name)
		}
		os.Exit(0)
	}
	
	var mp *mpris2.MediaPlayer
	if playerName == "" {
		mp, err = conn.GetFirstMediaPlayer()
	} else {
		mp = conn.GetMediaPlayer(playerName)
	}
	
	if err != nil {
		log.Fatal(err)
	}

	switch cmd {
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
		if flag.NArg() != 2 {
			log.Fatalf("Usage: %s seek <offset>", os.Args[0])
		}

		var offset int64
		_, err = fmt.Sscan(flag.Arg(1), &offset)
		if err == nil {
			err = mp.Seek(offset)
		}

	case "open":
		if flag.NArg() != 2 {
			log.Fatalf("usage: %s open <uri>", os.Args[0])
		}

		err = mp.OpenUri(flag.Arg(1))

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
			if flag.NArg() > 1 {
				for _, k := range flag.Args()[1:] {
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
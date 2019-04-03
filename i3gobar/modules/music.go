package modules

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/fsnotify/fsnotify"

	"github.com/zwzn/i3/i3gobar/bar"
	"github.com/zwzn/i3/i3gobar/icon"
)

type Song struct {
	Title    string `json:"title"`
	Artist   string `json:"artist"`
	Album    string `json:"album"`
	AlbumArt string `json:"albumArt"`
}

type Rating struct {
	Liked    bool `json:"liked"`
	Disliked bool `json:"disliked"`
}

type Time struct {
	Current int `json:"current"`
	Total   int `json:"total"`
}

type GPMDP struct {
	Playing    bool   `json:"playing"`
	Song       Song   `json:"song"`
	Rating     Rating `json:"rating"`
	Time       Time   `json:"time"`
	SongLyrics string `json:"songLyrics"`
	Shuffle    string `json:"shuffle"`
	Repeat     string `json:"repeat"`
	Volume     int    `json:"volume"`
}

var changeListeners []func(*GPMDP, error)
var changeListenersMtx sync.RWMutex

func sendChange(info *GPMDP, err error) {
	changeListenersMtx.Lock()
	for _, listener := range changeListeners {
		listener(info, err)
	}
	changeListenersMtx.Unlock()
}

func onChange(cb func(*GPMDP, error)) error {
	changeListenersMtx.RLock()
	changeListeners = append(changeListeners, cb)
	changeListenersMtx.RUnlock()

	w, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}
	go func() {
		w.Add(GPMDPPath())
		defer w.Close()
		for {
			select {
			case <-time.Tick(time.Second):
				info, err := CurrentSong()
				sendChange(info, err)

			case <-w.Events:
				info, err := CurrentSong()
				sendChange(info, err)

			case err := <-w.Errors:
				sendChange(nil, err)
			}
		}
	}()
	return nil
}
func onSong() chan Song {
	songs := make(chan Song)
	oldSong := Song{}
	onChange(func(info *GPMDP, err error) {
		if err != nil {
			return
		}
		if oldSong != info.Song {
			oldSong = info.Song
			songs <- info.Song
		}
	})
	return songs
}
func onPlaying() chan bool {
	playings := make(chan bool)
	oldPlaying := false
	onChange(func(info *GPMDP, err error) {
		if err != nil {
			return
		}
		if oldPlaying != info.Playing {
			oldPlaying = info.Playing
			playings <- info.Playing
		}
	})
	return playings
}

func GPMDPPath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	return filepath.Join(home, ".config/Google Play Music Desktop Player/json_store/playback.json")
}

func CurrentSong() (*GPMDP, error) {
	b, err := ioutil.ReadFile(GPMDPPath())
	if err != nil {
		return nil, err
	}
	info := &GPMDP{}
	err = json.Unmarshal(b, info)
	if err != nil {
		return nil, err
	}
	return info, nil
}

func MustCurrentSong() *GPMDP {
	info, err := CurrentSong()
	if err != nil {
		panic(err)
	}
	return info
}

func Music() *bar.Block {
	cs := make(chan string)
	info := MustCurrentSong()
	song := info.Song
	playing := info.Playing
	go func() {
		for {
			select {
			case song = <-onSong():
			case playing = <-onPlaying():
			}
			ico := icon.PauseCircle
			if playing {
				ico = icon.PlayCircle
			}
			if (song == Song{}) {
				continue
			}
			cs <- fmt.Sprintf("%s %s - %s", ico, song.Artist, song.Title)
		}
	}()
	return bar.NewBlock(cs)
}

package modules

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/davecgh/go-spew/spew"
	"github.com/zwzn/i3/i3gobar/bar"
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

func errBlock(err error) *bar.Block {
	cs := make(chan string)
	go func() { cs <- err.Error() }()
	return bar.NewBlock(cs)
}

func CurrentSong() (*GPMDP, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	b, err := ioutil.ReadFile(filepath.Join(home, ".config/Google Play Music Desktop Player/json_store/playback.json"))
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

func Music() *bar.Block {
	info, err := CurrentSong()
	if err != nil {
		return errBlock(err)
	}
	cs := make(chan string)
	spew.Dump(info)
	os.Exit(1)
	return bar.NewBlock(cs)
}

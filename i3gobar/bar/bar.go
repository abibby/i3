package bar

import (
	"encoding/json"
	"flag"
	"fmt"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/microcosm-cc/bluemonday"
)

type Block struct {
	name    string
	channel chan string
}
type barSection struct {
	id     int
	Text   string `json:"full_text"`
	Markup string `json:"markup,omitempty"`
}

type Bar []*barSection

var (
	consoleMode bool
)

func init() {
	flag.BoolVar(&consoleMode, "console", false, "runs the bar in console mode for testing")
}

func (b *Bar) Update(id int, text string) {
	for _, bs := range *b {
		if bs.id == id {
			bs.Text = text
			return
		}
	}
}
func (b *Bar) Print() {
	sort.Slice(*b, func(i, j int) bool {
		return (*b)[i].id < (*b)[j].id
	})
	if consoleMode {

		parts := []string{}
		for _, sec := range *b {
			if sec.Text != "" {
				parts = append(parts, bluemonday.StrictPolicy().Sanitize(sec.Text))
			}
		}
		fmt.Printf("%s\r", strings.Join(parts, " | "))

	} else {
		bJSON, err := json.Marshal(b)
		if err == nil {
			fmt.Printf(",%s\n", bJSON)
		}
	}
}

var bar = Bar{}

func Run(blocks ...chan string) {
	if !flag.Parsed() {
		flag.Parse()
	}
	wg := sync.WaitGroup{}
	for id := range blocks {
		bar = append(bar, &barSection{
			id:     id,
			Text:   "",
			Markup: "pango",
		})
		wg.Add(1)

	}
	fmt.Printf("{\"version\":1,\"click_events\":true}\n[[]\n")
	bar.Print()
	for id, cs := range blocks {
		go func(id int, cs chan string) {
			for text := range cs {
				parts := strings.Split(strings.TrimRight(text, "\n"), "\n")

				bar.Update(id, parts[len(parts)-1])
				bar.Print()
			}
			wg.Done()
		}(id, cs)
	}
	wg.Wait()
}

func Schedule(cb func() string, every time.Duration) chan string {
	cs := make(chan string)
	go func() {
		cs <- cb()
		for range time.Tick(every) {
			cs <- cb()
		}
	}()
	return cs
}

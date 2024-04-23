package bar

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"sort"
	"strings"
	"sync"

	"github.com/microcosm-cc/bluemonday"
)

type BarSection struct {
	id       int
	Name     string `json:"Name"`
	Instance string `json:"Instance"`
	Text     string `json:"full_text"`
	Markup   string `json:"markup,omitempty"`
}

type Bar []*BarSection

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

func Run(blocks ...*Block) {
	if !flag.Parsed() {
		flag.Parse()
	}

	wg := &sync.WaitGroup{}

	wg.Add(1)
	go readClicks(wg, blocks)

	for id := range blocks {
		bar = append(bar, &BarSection{
			id:       id,
			Instance: fmt.Sprintf("%d", id),
			Text:     "",
			Markup:   "pango",
		})
		wg.Add(1)
	}

	fmt.Printf("{\"version\":1,\"click_events\":true}\n[[]\n")
	bar.Print()
	for id, cs := range blocks {
		go runBlock(wg, id, cs)
	}
	wg.Wait()
}

func readClicks(wg *sync.WaitGroup, blocks []*Block) {
	defer wg.Done()
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := strings.Trim(scanner.Text(), ",")
		if line == "[" {
			continue
		}
		click := &Click{}
		err := json.Unmarshal([]byte(line), click)
		if err != nil {
			continue
		}
		for id, block := range blocks {
			if block.onClick != nil {
				if fmt.Sprint(id) == click.Instance {
					go block.onClick(*click)
				}
			}
		}
	}
}

func runBlock(wg *sync.WaitGroup, id int, cs *Block) {
	defer wg.Done()
	for text := range cs.text {
		parts := strings.Split(strings.TrimRight(text, "\n"), "\n")

		bar.Update(id, parts[len(parts)-1])
		bar.Print()
	}
}

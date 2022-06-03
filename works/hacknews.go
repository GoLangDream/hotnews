package works

import (
	"github.com/peterhellberg/hn"
)

func init() {
	cronTask.AddFunc("@hourly", syncHotNews)
}

func syncHotNews() {
	hnClient := hn.DefaultClient

	ids, _ := hnClient.TopStories()

	var stories = []*hn.Item{}

	for _, id := range ids[:10] {
		item, _ := hnClient.Item(id)

		stories = append(stories, item)
	}

}

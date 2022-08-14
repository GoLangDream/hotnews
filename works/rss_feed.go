package works

import "hot_news/service/rss"

func init() {
	for name, source := range rss.Sources {
		syncRss(name, source.Url, source.NeedTranslate)
	}
}

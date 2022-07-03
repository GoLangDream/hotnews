package works

func init() {
	syncRss("producthunt", "https://www.producthunt.com/feed")
	syncRss("dev.to", "https://dev.to/feed")
	syncRss("golang_weekly", "https://golangweekly.com/rss/")
	syncRss("ruby_weekly", "https://rubyweekly.com/rss/")
	syncRss("kotlin_weekly", "https://us12.campaign-archive.com/feed?u=f39692e245b94f7fb693b6d82&id=93b2272cb6")
	syncRss("database_weekly", "https://dbweekly.com/rss/")
	syncRss("elixir_weekly", "https://elixirstatus.com/rss")
	syncRss("stackoverflow", "https://stackoverflow.com/feeds")
	syncRss("oschina", "https://rsshub.app/oschina/news", false)
	syncRss("掘金", "https://rsshub.app/juejin/trending/all/weekly", false)
	syncRss("Go语言爱好者周刊", "https://rsshub.app/go-weekly", false)
	syncRss("安全内参", "https://rsshub.app/secrss/category/%E4%BA%A7%E4%B8%9A%E8%B6%8B%E5%8A%BF", false)
	syncRss("微博热搜", "https://rsshub.app/weibo/search/hot", false)
	syncRss("看雪论坛", "https://rsshub.app/pediy/topic/all/latest", false)
	syncRss("InfoQ中文", "https://rsshub.app/infoq/recommend", false)
}

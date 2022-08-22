package rss

type Source struct {
	Url           string
	NeedTranslate bool
}

var Sources = map[string]Source{
	"producthunt":     {"https://www.producthunt.com/feed", true},
	"dev.to":          {"https://dev.to/feed", true},
	"golang_weekly":   {"https://golangweekly.com/rss/", true},
	"ruby_weekly":     {"https://rubyweekly.com/rss/", true},
	"kotlin_weekly":   {"https://us12.campaign-archive.com/feed?u=f39692e245b94f7fb693b6d82&id=93b2272cb6", true},
	"database_weekly": {"https://dbweekly.com/rss/", true},
	"elixir_weekly":   {"https://elixirstatus.com/rss", true},
	"cncf":            {"https://rsshub.app/cncf", true},
	"ribbonfarm":      {"https://www.ribbonfarm.com/feed/", true},

	"oschina":      {"https://rsshub.app/oschina/news", false},
	"掘金":           {"https://rsshub.app/juejin/trending/all/weekly", false},
	"安全内参":         {"https://rsshub.app/secrss/category/%E4%BA%A7%E4%B8%9A%E8%B6%8B%E5%8A%BF", false},
	"微博热搜":         {"https://rsshub.app/weibo/search/hot", false},
	"看雪论坛":         {"https://rsshub.app/pediy/topic/all/latest", false},
	"InfoQ中文":      {"https://rsshub.app/infoq/recommend", false},
	"dbaplus":      {"https://rsshub.app/dbaplus", false},
	"软件活动":         {"https://rsshub.app/dbaplus/activity", false},
	"Dockone":      {"https://rsshub.app/dockone/weekly", false},
	"GitChat":      {"https://rsshub.app/gitchat/newest", false},
	"HelloGitHub":  {"https://rsshub.app/hellogithub/article", false},
	"NOSEC.org":    {"https://rsshub.app/nosec/hole", false},
	"0day漏洞":       {"https://rsshub.app/project-zero-issues", false},
	"segmentfault": {"https://rsshub.app/segmentfault/channel/frontend", false},
	"美团技术团队":       {"https://rsshub.app/meituan/tech/home", false},
	"知乎日报":         {"https://rsshub.app/zhihu/daily", false},
	"知乎热榜":         {"https://rsshub.app/zhihu/hotlist", false},
	"知乎想法":         {"https://rsshub.app/zhihu/pin/daily", false},
	"果壳网":          {"https://rsshub.app/guokr/scientific", false},
	"后续":           {"https://rsshub.app/houxu/events", false},
	//"汇通网":          {"https://rsshub.app/fx678/kx", false},
	"人人都是产品经理":  {"https://rsshub.app/woshipm/popular", false},
	"推酷":        {"https://rsshub.app/tuicool/mags/tech", false},
	"少数派":       {"https://rsshub.app/sspai/index", false},
	"深潮":        {"https://rsshub.app/techflow520/newsflash", false},
	"探物":        {"https://rsshub.app/tanwu/products", false},
	"瘾科技":       {"https://rsshub.app/engadget/chinese", false},
	"IT之家":      {"https://rsshub.app/ithome/ranking/24h", false},
	"豆瓣-一周口碑电影": {"https://rsshub.app/douban/movie/weekly", false},
	"豆瓣-热门书":    {"https://rsshub.app/douban/book/rank", false},
	"happyxiao": {"https://happyxiao.com/feed/", false},
	"decohack":  {"https://www.decohack.com/feed", false},
	"100x":      {"https://100x.today/feed/", false},
}

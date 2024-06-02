package goal_ten_api

type TenAPI interface {
	// 热点趋势

	BaiduHots() ([]Hot, error)
	DouyinHots() ([]Hot, error)
	WeiboHots() ([]Hot, error)
	ZhihuHots() ([]Hot, error)
	BiliHots() ([]Hot, error)

	//ToutiaoHots 头条热搜
	ToutiaoHots() ([]Hot, error)

	// ToutiaoHotNews 头条热点新闻
	ToutiaoHotNews() ([]Hot, error)
}

package constants

var Settings = map[string][]Setting{
    "settings": []Setting{
        {
            Title: "博客",
            Url:   "https://xiaoxuan6.github.io",
        },
        {
            Title: "假数据生成器",
            Url:   "https://xiaoxuan6.github.io/faker",
        },
        {
            Title: "Free VIP视频解析",
            Url:   "https://xiaoxuan6.github.io/free-vip-video",
        },
        {
            Title: "chat gpt 在线免费网站",
            Url:   "https://xiaoxuan6.github.io/chatgpt-server",
        },
    },
}

type Setting struct {
    Title string `json:"title"`
    Url   string `json:"url"`
}

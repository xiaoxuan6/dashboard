package common

const HOLIDAY = "https://timor.tech/api/holiday/tts/next"
const HOLIDAY_INFO = "http://timor.tech/api/holiday/info/%s"

const (
    CollectExpiration = 30 // 分钟
    GoTags            = "go-package"
    PackageUrl        = "https://ghproxy.com/https://github.com/xiaoxuan6/go-package-example/blob/main/README.md"

    ImageUrl     = "https://api.airandomimage.top/api/open/picture?current=%s"
    BingUrl      = "http://cn.bing.com"
    BingImageUrl = BingUrl + "/HPImageArchive.aspx?idx=%s&n=1"
)

var (
    GoLanguage  = []string{"Go", "go", "Golang", "golang"}
    PhpLanguage = []string{"PHP", "php"}

    Communes = []string{"https://www.v2ex.com/index.xml", "https://learnku.com/feed"}
    TAGS     = []string{"README", "action", "api", "chat", "docker", "dockerfile", "email", "git", "go", "heiliao", "js", "linux", "logo", "makefile", "ocr", "pic", "send", "tool", "vip", GoTags}

    // PHP docs
    PHP = map[string]string{
        "Guzzle 中文文档":   "https://guzzle-cn.readthedocs.io/zh_CN/latest/quickstart.html",
        "Laravel excel": "https://docs.laravel-excel.com/2.1/import/injection.html",
        "Laravel Api":   "https://laravel.com/api/8.x/Illuminate.html",
        "Deployer":      "https://deployer.org/docs/6.x/installation",
        "PHP 语言设计模式":    "https://refactoringguru.cn/design-patterns",
    }

    GO = map[string]string{
        "Go 语言设计模式":         "https://www.topgoer.cn/docs/golang-design-pattern/Singleton",
        "Go Example":        "https://gobyexample.com",
        "goreleaser":        "https://llever.com/goreleaser-zh/",
        "gin":               "https://gin-gonic.com/docs/introduction/",
        "GORM":              "https://gorm.io/zh_CN/docs/",
        "beego":             "https://git-books.github.io/books/beego/",
        "pre-commit-golang": "https://github.com/dnephin/pre-commit-golang",
    }

    Python = map[string]string{
        "Python 语法": "https://www.nowcoder.com/tutorial/10005/f9a1fa805b604d0dbddcb8835286d8cb",
    }

    Other = map[string]string{
        "Swoole":               "https://wiki.swoole.com/#/environment?id=安装准备",
        "Elasticsearch":        "https://www.elastic.co/guide/cn/elasticsearch/php/current/_index_management_operations.html",
        "Elasticsearch 中文文档":   "https://doc.codingdict.com/elasticsearch/74",
        "RabbitMQ 中文文档－PHP版":   "https://rabbitmq.shujuwajue.com/ying-yong-jiao-cheng/php-ban",
        "pre-commit":           "https://pre-commit.com/",
        "gRPC":                 "https://grpc.io/docs/languages/",
        "Tampermonkey - 篡改猴":   "https://www.tampermonkey.net/documentation.php?locale=zh",
        "Greasy Fork - 油猴":     "https://greasyfork.org/zh-CN/help/meta-keys",
        "Tampermonkey油猴脚本中文文档": "https://www.itblogcn.com/article/2233.html",
        "Caddy":                "https://caddyserver.com/docs/",
        "Caddy - 中文文档":         "https://dengxiaolong.com/caddy/zh/",
        "Makefile":             "https://www.zhaixue.cc/makefile/makefile-intro.html",
        "Vercel":               "https://vercel.com/docs/projects/project-configuration#routes",
    }

    Html = map[string]string{
        "Typescript":   "https://wangdoc.com/typescript",
        "Axios - 中文文档": "http://axios-js.com/zh-cn/docs/index.html",
        "Axios":        "https://axios-http.com/zh/docs/intro",
        "Notiflix":     "https://notiflix.github.io/notify",
        "NProgress":    "https://github.com/rstacruz/nprogress",
    }

    Docs = map[string]map[string]string{
        "PHP":    PHP,
        "前端":     Html,
        "Go":     GO,
        "Python": Python,
        "Other":  Other,
        "Reference": map[string]string{
            "Quick":                     "https://wangchujiang.com/reference/index.html",
            "在 Git 项目中增加 pre-commit 校验": "https://jasonkayzk.github.io/2021/10/10/%E5%9C%A8Git%E9%A1%B9%E7%9B%AE%E4%B8%AD%E5%A2%9E%E5%8A%A0pre-commit%E6%A0%A1%E9%AA%8C/",
            "phpstorm 与 xdebug 的调试配置":   "https://www.cnblogs.com/eminer/p/17247178.html",
        },
    }
)

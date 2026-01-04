package handlers

import (
	"bytes"
	"dashboard/util"
	"fmt"
	"github.com/OwO-Network/gdeeplx"
	"github.com/abadojack/whatlanggo"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
	db "github.com/xiaoxuan6/go-package-db"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

var (
	CollectHandler = new(collectHandler)
)

type collectHandler struct {
}

func (c collectHandler) Put(ctx *gin.Context) {
	b, _ := ioutil.ReadAll(ctx.Request.Body)
	decodedString, err := url.QueryUnescape(string(b))
	if err != nil {
		util.Fail(ctx, err.Error())
		return
	}

	logrus.Info("decodedString", decodedString)
	values, _ := url.ParseQuery(decodedString)
	auth := values.Get("auth")
	uri := values.Get("url")
	description := values.Get("description")
	language := values.Get("language")

	ok, msg := checkParams(auth, uri)
	if !ok {
		util.Fail(ctx, msg)
		return
	}

	if isExists(uri) {
		util.Fail(ctx, "url exists")
		return
	}

	description = descriptionDo(strings.TrimSpace(description))
	language = languageDo(uri, strings.TrimSpace(language))

	var body bytes.Buffer
	body.WriteString(fmt.Sprintf(`{"event_type": "push", "client_payload": {"url": "%s", "description":"%s", "demo_url":"", "language": "%s"}}`, uri, description, language))

	r, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("https://api.github.com/repos/%s/go-package-example/dispatches", auth), &body)
	r.Header.Set("Accept", "application/vnd.github+json")
	r.Header.Set("Authorization", fmt.Sprintf("token %s", os.Getenv("GITHUB_TOKEN")))
	r.Header.Set("X-GitHub-Api-Version", "2022-11-28")

	res, _ := http.DefaultClient.Do(r)
	if res.StatusCode != 204 {
		logrus.Error("dispatch error", body.String())
		util.Fail(ctx, "dispatch error")
		return
	}

	util.Success(ctx)
}

func checkParams(auth, uri string) (bool, string) {
	if auth == "" || uri == "" {
		return false, "auth or url not empty"
	}

	if auth != os.Getenv("GITHUB_OWNER") {
		return false, "auth error"
	}

	u, _ := url.Parse(uri)
	if u.Host != "github.com" {
		return false, "url error"
	}

	return true, ""
}

func isExists(uri string) bool {
	db.Init(
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	defer db.Close()
	db.AutoMigrate()

	uri = strings.ReplaceAll(uri, "https://", "")
	return db.DB.Where("url = ?", uri).First(&db.Collect{}).RowsAffected > 0
}

func (c collectHandler) IsExists(ctx *gin.Context) {
	b, _ := ioutil.ReadAll(ctx.Request.Body)
	decodedString, err := url.QueryUnescape(string(b))
	if err != nil {
		util.Fail(ctx, err.Error())
		return
	}

	values, _ := url.ParseQuery(decodedString)
	auth := values.Get("auth")
	uri := values.Get("url")
	ok, msg := checkParams(auth, uri)
	if !ok {
		util.Fail(ctx, msg)
		return
	}

	util.SuccessWithData(ctx, isExists(uri))

}

func descriptionDo(description string) string {
	if len(description) < 1 {
		return ""
	}

	info := whatlanggo.Detect(description)
	lang := info.Lang.String()
	if lang != "" && lang != "Mandarin" {
		result, err := gdeeplx.Translate(description, "", "zh", 0)
		if err == nil {
			res := result.(map[string]interface{})
			description = strings.TrimSpace(res["data"].(string))
		}
	}

	return description
}

func languageDo(uri, language string) string {
	if len(language) > 10 {
		language = ""
	}

	if len(language) > 0 {
		return language
	}

	u, _ := url.Parse(uri)
	u.Path = strings.TrimPrefix(u.Path, "/")
	paths := strings.Split(u.Path, "/")

	res, err := util.Get(fmt.Sprintf("https://api.github.com/repos/%s/%s", paths[0], paths[1]))
	if err != nil {
		return "Other"
	}

	language = gjson.Get(res, "language").String()
	if len(language) < 1 {
		return "Other"
	}

	return language
}

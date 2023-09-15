let keyword;

function init() {
    keyword = localStorage.getItem("keyword")
    if (keyword === undefined || keyword == null) {
        window.location.href = "/search"
        return
    }

    if (token === undefined || token == null) {
        window.location.href = "/login"
    }

    search_do()
}

init()

function search_do() {
    NProgress.start();
    post("search_do", {
        token: token,
        email: email,
        keyword: keyword,
    }, function (response) {
        let data = response.data
        if (data.status != 200) {
            NProgress.done();
            error(data.msg);
            setTimeout(function () {
                localStorage.removeItem("keyword")
                window.location.href = "/search"
            }, 1000)
            return
        }

        setContent(data.data)
    }, function (err) {
        NProgress.done();
        error(`请求失败：${err}`)
    })
}

function setContent(data) {
    console.log(data)
}
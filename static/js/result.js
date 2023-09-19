let keyword;

function init() {
    keyword = localStorage.getItem("keyword")
    if (keyword === undefined || keyword == null) {
        redirect("/search", 0)
        return
    }

    token = localStorage.getItem("token")
    if (token === undefined || token == null) {
        redirect("/login", 0)
        return
    }

    email = localStorage.getItem("email")
    search_do()
}

init()

function search_do() {
    NProgress.start();
    postWithHeader("search_do", {
        keyword: keyword,
    }, function (response) {
        let data = response.data
        if (data.status === 401) {
            NProgress.done();
            error(data.msg)
            localStorage.removeItem("token")
            redirect("/login", 0)
            return
        }
        localStorage.removeItem('keyword')
        setContent(data.data)
        NProgress.done();
    }, function (err) {
        NProgress.done();
        localStorage.removeItem("keyword")
        error(`请求失败：${err}`)
    })
}

function setContent(data) {
    console.log(data)
    document.getElementById('key').innerHTML = data.keyword
    document.getElementById('date').innerHTML = data.date
}

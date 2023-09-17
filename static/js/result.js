let keyword;

function init() {
    keyword = localStorage.getItem("keyword")
    if (keyword === undefined || keyword == null) {
        window.location.href = "/search"
        return
    }

    token = localStorage.getItem("token")
    if (token === undefined || token == null) {
        window.location.href = "/login"
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

        if (data.status !== 200 && data.status === 401) {
            NProgress.done();
            error(data.msg);

            // todo:: 设置 token
            return;
        }

        if (data.status !== 200 && data.status !== 401) {
            NProgress.done();
            error(data.msg);
            setTimeout(function () {
                localStorage.removeItem("keyword")
                window.location.href = "/search"
            }, 1000)
            return
        }

        setContent(data.data)
        NProgress.done();
    }, function (err) {
        NProgress.done();
        error(`请求失败：${err}`)
    })
}

function setContent(data) {
    $('.loading').hide()
    let dataDiv = document.getElementById("data")
    dataDiv.style.display = "";
    dataDiv.innerHTML = data.keyword
}

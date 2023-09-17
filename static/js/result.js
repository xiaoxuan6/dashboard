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
    // dataDiv.innerHTML = data.keyword

    console.log(data)
}

function init() {
    token = localStorage.getItem("token")
    email = localStorage.getItem("email")
    if (token === undefined || email === undefined || email == null || token == null) {
        error("未登录！");
        setTimeout(function () {
            window.location.href = "/login"
        }, 2000)
        return
    }
}

init()

$(document).ready(function () {
    $('#keyword').keypress(function (event) {
        if (event.which === 13) {
            search()
        }
    });
});

function search() {
    let keyword = document.getElementById("keyword").value
    if (!keyword) {
        error("请输入搜索问题")
        return;
    }

    localStorage.setItem("keyword", keyword)
    window.location.href = "/result"
}

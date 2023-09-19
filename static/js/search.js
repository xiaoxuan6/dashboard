$(document).ready(function () {
    $('#keyword').keypress(function (event) {
        if (event.which === 13) {
            search()
        }
    });

    function init() {
        document.getElementById("right_button").style.display = "";

        token = localStorage.getItem("token")
        email = localStorage.getItem("email")
        if (token === undefined || email === undefined || email == null || token == null) {
            error("未登录！");
            redirect("/login", 0)
        }

        keyword = localStorage.getItem("keyword")
        if (keyword !== undefined && keyword != null) {
            redirect("/result", 0)
        }
    }

    init()
});

function search() {
    let keyword = document.getElementById("keyword").value
    if (!keyword) {
        error("请输入搜索问题")
        return;
    }

    localStorage.setItem("keyword", keyword)
    redirect("/result", 0)
}

function load() {
    error("功能待开发中……")
}

function logout() {
    localStorage.removeItem("token")
    Notiflix.Notify.success("退出成功")
    redirect("/login", 1000)
}

$(document).ready(function () {
    $('#keyword').keypress(function (event) {
        if (event.which === 13) {
            search()
        }
    });

    function init() {
        token = localStorage.getItem("token")
        email = localStorage.getItem("email")
        if (token === undefined || email === undefined || email == null || token == null) {
            error("未登录！");
            redirect("/login", 0)
        }
    }

    init()
});

function search() {
    let keyword = document.getElementById("keyword").value
    if (!keyword) {
        error("请输入关键字")
        return;
    }
    alert(keyword)

    // todo:搜索
}

function logout() {
    localStorage.removeItem("token")
    Notiflix.Notify.success("退出成功")
    redirect("/login", 1000)
}

function search() {
    let token = localStorage.getItem("token")
    let email = localStorage.getItem("email")
    let passwd = localStorage.getItem("password")
    if (!token || !email || !passwd) {
        Notiflix.Notify.failure("未登录！");
        setTimeout(function () {
            window.location.href = "/login"
        }, 1000)
        return
    }

    let keyword = document.getElementById("keyword").value
    if (!keyword) {
        error("请输入搜索问题")
        return;
    }

    localStorage.setItem("keyword", keyword)
    // window.location.href = "/result"
}

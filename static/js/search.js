init()

function init() {
    let token = localStorage.getItem("token")
    let email = localStorage.getItem("email")
    console.log("token", token)
    console.log("email", email)
    
    if (!token || !email) {
        Notiflix.Notify.failure("未登录！");
        setTimeout(function () {
            window.location.href = "/login"
        }, 2000)
        return
    }
}

function search() {
    let keyword = document.getElementById("keyword").value
    if (!keyword) {
        error("请输入搜索问题")
        return;
    }

    localStorage.setItem("keyword", keyword)
    // window.location.href = "/result"
}

init()

let token = localStorage.getItem("token")
let email = localStorage.getItem("email")

function init() {
    if (token === undefined || email === undefined) {
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

    post("check_token", {
        token: token,
        email: email,
    }, function (response) {
        let data = response.data;
        if (data.status != 200) {
            Notiflix.Notify.warning(data.data.msg);
            setTimeout(function () {
                window.location.href = "/search"
            }, 1000)
        }
    }, function (error) {

    })

    localStorage.setItem("keyword", keyword)
    // window.location.href = "/result"
}

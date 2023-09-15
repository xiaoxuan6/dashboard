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

    post("check_token", {
        token: token,
        email: email,
    }, function (response) {
        let data = response.data;
        if (data.status != 200) {
            warning(data.msg);
            setTimeout(function () {
                localStorage.removeItem("token")
                window.location.href = "/login"
            }, 1000)
        }
    }, function (error) {

    })

    localStorage.setItem("keyword", keyword)
    window.location.href = "/result"
}

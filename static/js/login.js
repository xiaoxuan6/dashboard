$(function () {
    $('#password').focus(function () {
        // 密码框获得焦点，追加样式.password
        $('#owl').addClass('password');
    }).blur(function () {
        // 密码框失去焦点，移除样式.password
        $('#owl').removeClass('password');
    })
})

function login() {
    let email = document.getElementById("email")
    let passwd = document.getElementById("password")

    if (!email) {
        error("账号不能为空！")
        return
    }
    if (!passwd) {
        error("密码不能为空！");
        return
    }

    NProgress.start()
    get("login", success(), response())
}

function success() {
    return function (response) {
        NProgress.done();
        let data = response.data;
        if (data.status == 200) {
            Notiflix.Notify.success("登录成功！")
            localStorage.setItem("token", data.token)
            setTimeout(function () {
                window.location.href = "/search"
            }, 1000)
        } else {
            error("账号或密码错误！");
        }
    }
}

function response() {
    return function (error) {
        NProgress.done();
        error(`请求失败: ${error}`);
    }
}

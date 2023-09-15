$(function () {
    $('#password').focus(function () {
        // 密码框获得焦点，追加样式.password
        $('#owl').addClass('password');
    }).blur(function () {
        // 密码框失去焦点，移除样式.password
        $('#owl').removeClass('password');
    })
})

$(document).ready(function () {
    $('#password').keypress(function (event) {
        if (event.which === 13) {
            login()
        }
    });
});

init()

function init() {
    let token = localStorage.getItem("token")
    let email = localStorage.getItem("email")
    if (token === undefined || email === undefined || token == null) {
        return
    }

    post("check_token", {
        token: token,
        email: email,
    }, function (response) {
        let data = response.data;
        if (data.status == 200) {
            setTimeout(function () {
                window.location.href = "/search"
            }, 1000)
        } else {
            warning(data.msg);
        }
    }, function (error) {

    })
}

function login() {
    let email = document.getElementById("email").value
    let passwd = document.getElementById("password").value

    if (!email) {
        error("账号不能为空！")
        return
    }
    if (!passwd) {
        error("密码不能为空！");
        return
    }

    NProgress.start()
    post("login_do", {
        email: email,
        passwd: passwd
    }, success(), response())
}

function success() {
    return function (response) {
        NProgress.done();
        let data = response.data;
        if (data.status == 200) {
            Notiflix.Notify.success("登录成功！")
            localStorage.setItem("token", data.data.token)
            localStorage.setItem("email", data.data.email)

            setTimeout(function () {
                window.location.href = "/search"
            }, 2000)
        } else {
            error("账号或密码错误！");
        }
    }
}

function response() {
    return function (err) {
        NProgress.done();
        error(`请求失败: ${err}`);
    }
}

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
    NProgress.start()
    const now_time = new Date().getHours() + new Date().getMinutes() + new Date().getSeconds();
    axios.get("/api?action=login&time=" + now_time)
        .then(function (response) {
            NProgress.done();
            let data = response.data;
            if (data.status == 200) {
                Notiflix.Notify.success("登录成功！")
                localStorage.setItem("token", data.token)
                setTimeout(function () {
                    window.location.href = "/search"
                }, 1000)
            } else {
                Notiflix.Notify.failure("登录失败！");
            }
        })
        .catch(function (error) {
            NProgress.done();
            Notiflix.Notify.failure(`请求失败: ${error}`);
        })
}

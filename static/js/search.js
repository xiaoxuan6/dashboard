function search() {
    let token = localStorage.getItem("token")
    if (token.length < 1) {
        Notiflix.Notify.failure("未登录！", () => {
            window.location.href = "/login"
        });
        return
    }

    // todo::验证token

    const now_time = new Date().getHours() + new Date().getMinutes() + new Date().getSeconds();
    axios.get("/api?action=search&time=" + now_time)
        .then(function (response) {
            // let data = response.data;
            // if (data.status == 200) {
            //     append(data.data.settings)
            // } else {
            //     Notiflix.Notify.failure("请求失败: " + data.msg);
            // }
        })
        .catch(function (error) {
            Notiflix.Notify.failure(`请求失败: ${error}`);
        })
}

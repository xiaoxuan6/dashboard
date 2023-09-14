function search() {
    let token = localStorage.getItem("token")
    if (!token) {
        Notiflix.Notify.failure("未登录！", () => {
            window.location.href = "/login"
        });
        return
    }

    let keyword = document.getElementById("keyword")
    if (!keyword) {
        error("请输入搜索问题")
        return;
    }

    post("search", {keyword: keyword, token: token}, success(), response())
}

function success() {
    return function (response) {
        // let data = response.data;
        // if (data.status == 200) {
        //     append(data.data.settings)
        // } else {
        //     Notiflix.Notify.failure("请求失败: " + data.msg);
        // }
    }
}

function response() {
    return function (error) {
        Notiflix.Notify.failure(`请求失败: ${error}`);
    }
}

let keyword
$(document).ready(function () {
    $('.keyword').keypress(function (event) {
        if (event.which === 13) {
            keyword = $(this).val()
            if (!keyword) {
                error("请输入关键字")
                return;
            }

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

        let target = localStorage.getItem('disable')
        if (target === "true") {
            document.getElementById('pre').style.display = "none"
            document.getElementById('next').style.display = "block"
        }
    }

    init()
});

function onSearch(than) {
    keyword = $(than).prev('input').val()
    if (!keyword) {
        error("请输入关键字")
        return;
    }
    search()
}

function search() {
    NProgress.start()
    postWithHeader('search_do', {'keyword': keyword}, function (response) {
        NProgress.done();
        let data = response.data;
        if (data.status === 200) {
            console.log("data", data.data.)
        } else {
            error(data.msg);
        }
    }, function (err) {
        NProgress.done();
        error(`请求失败: ${err}`);
    })

    document.getElementById('keyword').value = keyword

    let target = localStorage.getItem('disable')
    if (!target) {
        document.getElementById('pre').style.display = "none"
        document.getElementById('next').style.display = "block"
        localStorage.setItem('disable', "true")
    }
}

function logout() {
    localStorage.removeItem("token")
    Notiflix.Notify.success("退出成功")
    redirect("/login", 1000)
}

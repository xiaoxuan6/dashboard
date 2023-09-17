function prompt() {
    layer.prompt({title: '输入 token 口令，并确认', formType: 1}, function (pass, index) {
        layer.close(index);
        postWithHeader("github_token", {
            "token": pass
        }, function (response) {
            let data = response.data
            if (data.status === 200) {
                console.log('data.data', data.data)
                redirect()
            } else {
                error(response.msg)
                setTimeout(function () {
                    prompt()
                }, 2000)
            }
        }, function (err) {
            error(`请求失败：${err}`)
            setTimeout(function () {
                prompt()
            }, 2000)
        })
    });
}

init()

function init() {
    const now_time = new Date().getHours() + "hrs" + new Date().getMinutes() + "min";
    axios.get("/api?action=index&time=" + now_time)
        .then(function (response) {
            let data = response.data;
            if (data.status == 200) {
                append(data.data)
            } else {
                Notiflix.Notify.failure("请求失败: " + data.msg);
            }
        })
        .catch(function (error) {
            Notiflix.Notify.failure(`请求失败: ${error}`);
        })
}

function append(item) {
    document.getElementById('content').innerHTML = '';
    for (let i = 0; i < item; i++) {
        const li = document.createElement('li');
        let innerHtml = ""
        innerHtml = "<div>" + item[i]["title"] + "</div>" +
            "<div><a href=\"" + item[i]["url"] + "\">" + item[i]["title"] + "</a></div>"

        li.innerHTML = innerHtml;
        document.getElementById('content').appendChild(li);
    }
}

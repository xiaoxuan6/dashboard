init()

function init() {

    get("index", function (response) {
            let data = response.data;
            if (data.status == 200) {
                append(data.data.settings)
            } else {
                error("请求失败: " + data.msg)
            }
        },
        function (err) {
            error(`请求失败: ${err}`);
        }
    )
}

function append(item) {
    document.getElementById('content').innerHTML = '';
    for (let i = 0; i < item.length; i++) {
        const li = document.createElement('li');
        let innerHtml = ""
        innerHtml = "<div>" + item[i].title + "</div>";

        if (isURL(item[i].url)) {
            innerHtml = innerHtml + "<div><a href=\"" + item[i].url + "\" style='color: white' target=\"_blank\">" + item[i]["title"] + "</a></div>"
        } else {
            innerHtml = innerHtml + "<div><a href=\"" + item[i].url + "\" style='color: white'>" + item[i]["title"] + "</a></div>"
        }

        li.innerHTML = innerHtml;
        document.getElementById('content').appendChild(li);
    }
}

function isURL(str) {
    const pattern = new RegExp('^https?://[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}(/\\S*)?$');
    return pattern.test(str);
}

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
    const container = document.getElementById("container");

    let innerHtml;
    for (let i = 0; i < item.length; i++) {

        if (isURL(item[i].url)) {
            innerHtml = "<a href=\"" + item[i].url + "\" target=\"_blank\">" + item[i]["title"] + "</a>"
        } else {
            innerHtml = "<a href=\"" + item[i].url + "\">" + item[i]["title"] + "</a>"
        }

        const websiteElement = document.createElement("div");
        websiteElement.className = "website";
        websiteElement.innerHTML = innerHtml;
        container.appendChild(websiteElement);
    }
}

function isURL(str) {
    const pattern = new RegExp('^https?://[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}(/\\S*)?$');
    return pattern.test(str);
}

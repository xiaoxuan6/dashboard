load().then(res => {
    let html = ''
    for (let i = 0; i < res.length; i++) {
        html += '<div class="col-md-4">\n' +
            '            <a data-name="' + res[i].url + '" onclick="redirects()" class="block block-link-hover2 ribbon ribbon-modern ribbon-success"\n' +
            '               target="_blank">\n' +
            '                <div class="ribbon-box font-w600">状态：正常</div>\n' +
            '                <div class="block-content">\n' +
            '                    <div class="h4 push-5">' + res[i].title + '</div>\n' +
            '                    <p class="text-muted">' + res[i].desc + '</p></div>\n' +
            '            </a>\n' +
            '        </div>'
    }

    document.getElementById('common-packages').innerHTML = html
});

async function load() {
    let data = await (
        await fetch(`/apis/docs`)
    ).json()

    if (data && data.status === 200) {
        return data.data.apis
    }

    return {}
}

function redirects() {
    let target = $(this).data("name")
    localStorage.setItem("docs", target)
    window.location.href = "/apis/docs"
}

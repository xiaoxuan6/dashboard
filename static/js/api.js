load().then(res => {
    let html = ''
    for (let i = 0; i < res.length; i++) {
        html += '<div className="col-md-4">\n' +
            '            <a href="' + res.url + '" className="block block-link-hover2 ribbon ribbon-modern ribbon-success"\n' +
            '               target="_blank">\n' +
            '                <div className="ribbon-box font-w600">状态：正常</div>\n' +
            '                <div className="block-content">\n' +
            '                    <div className="h4 push-5">' + res.title + '</div>\n' +
            '                    <p className="text-muted">' + res.desc + '</p></div>\n' +
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

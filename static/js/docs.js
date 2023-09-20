load().then(r => {
    let url = document.getElementById('url')
    url.innerHTML = r.url
    url.setAttribute('data-clipboard-text', r.url)

    let method = document.getElementById('method')
    method.innerHTML = r.method
    method.setAttribute('data-clipboard-text', r.method)

    let url_demo = document.getElementById('url_demo')
    url_demo.innerHTML = r.url_demo
    url_demo.setAttribute('data-clipboard-text', r.url_demo)

    let params = r.params
    let params_html = ''
    for (let i = 0; i < params.length; i++) {
        params_html += '<tr>\n' +
            '<td>' + params[i].name + '</td>\n' +
            '<td>' + params[i].require + '</td>\n' +
            '<td>' + params[i].type + '</td>\n' +
            '<td>' + params[i].desc + '</td>\n' +
            '</tr>'
    }
    document.getElementById('params').innerHTML = params_html

    let response = r.response
    let response_html = ''
    for (let j = 0; j < response.length; j++) {
        response_html += '<tr>\n' +
            '<td>' + response[j].name + '</td>\n' +
            '<td>' + response[j].type + '</td>\n' +
            '<td>' + response[j].desc + '</td>\n' +
            '</tr>'
    }
    document.getElementById('response').innerHTML = response_html

    let codes = r.codes
    let code_html = ''
    for (let n = 0; n < codes.length; n++) {
        code_html += ' <tr>\n' +
            '<td>' + codes[n].code + '</td>\n' +
            '<td>' + codes[n].desc + '</td>\n' +
            '</tr>\n'
    }
    document.getElementById('codes').innerHTML = code_html
})

async function load() {
    let target = localStorage.getItem("docs")
    const item = await (
        await fetch(`/apis/docs${target}`)
    ).json()

    if (item && item.status === 200) {
        return item.data
    }

    return {}
}

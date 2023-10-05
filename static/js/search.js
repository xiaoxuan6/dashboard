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

    $(document).on('click', '.tag-item', function () {
        let than = $(this);
        let val = than.find('input[type="radio"]').data('val');
        than.find('input[type="radio"]').prop('checked', true);

        let data = localStorage.getItem('search_posts');
        data = JSON.parse(data);
        document.getElementById('accordion').innerHTML = '';
        if (val === '') {
            appendContent(data);
        } else {
            let posts = [];
            for (let i = 0; i < data.length; i++) {
                if (data[i].tag === val) {
                    posts.push(data[i]);
                }
            }
            appendContent(posts);
        }
    });
});

function onSearch(than) {
    keyword = $(than).prev('input').val()
    if (!keyword) {
        error("请输入关键字")
        return;
    }
    search()
}

function appendContent(data) {
    let html = ''
    for (let i = 0; i < data.length; i++) {
        html += '<div class="card mb-1 rounded-0 shadow border-0 card-PDF">\n' +
            '    <div id="will-updates-also-be-free">\n' +
            '        <div class="card-body font-secondary text-color">\n' +
            '            <h5>\n' +
            '                <a class="resource-link" target="_blank" rel="noopener noreferrer nofollow"\n' +
            '                   href="' + data[i].url + '">' + data[i].title + '</a>\n' +
            '            </h5>\n' +
            '            <p>\n' +
            '            </p>\n' +
            '            <p style="margin-bottom: 0">\n' +
            '                <span>' + data[i].tag + '</span>\n' +
            '            </p>\n' +
            '        </div>\n' +
            '    </div>\n' +
            '</div>'
    }
    document.getElementById('accordion').innerHTML = html
}

function append(data) {
    let posts = data.posts
    $('#tags').html('');
    $('#accordion').html('');
    if (posts !== null) {
        appendContent(posts)

        localStorage.setItem('search_posts', JSON.stringify(posts));
        let tags = data.tags
        let tags_html = '<span class="tag-item tag-item-all" style="margin-right: 20px">\n' +
            '   <input type="radio" class="tag" name="tag" data-val=""/>\n' +
            '   <label>全部</label>\n' +
            '</span>\n'
        $.each(tags, function (key, value) {
            tags_html += '<span class="tag-item" data-ext="DOC">\n' +
                '   <input type="radio" class="tag" name="tag" data-val="' + key + '"/>\n' +
                '   <label>' + key + '(' + value + ')</label>\n' +
                '</span>'
        })
        $('#tags').append(tags_html)
        $('.tag-item-all').find('input[type="radio"]').prop('checked', true);
    } else {
        error("没有找到相关资源")
    }
}

function search() {
    NProgress.start()
    postWithHeader('search_do', {'keyword': keyword}, function (response) {
        NProgress.done();
        let data = response.data;
        if (data.status === 200) {
            append(data.data)
        } else if (data.status === 401) {
            error(data.msg);
            localStorage.removeItem('token')
            redirect('/login', 1000)
        } else {
            error(data.msg);
        }
    }, function (err) {
        $('#tags').html('');
        $('#accordion').html('');
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
    redirect("/", 1000)
}

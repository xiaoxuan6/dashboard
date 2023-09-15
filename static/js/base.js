let token;
let email;

function error(msg) {
    Notiflix.Notify.failure(msg);
}

function warning(msg) {
    Notiflix.Notify.warning(msg);
}

const now_time = new Date().getHours() + "h" + new Date().getMinutes() + "i" + new Date().getSeconds() + "s";

function get(action, then, error) {
    axios.get("/api?action=" + action + "&time=" + now_time)
        .then(then)
        .catch(error)
}

const config = {
    headers: {
        'Content-Type': 'application/json',
        'Authorization': token,
        'email': email,
    },
}

function post(action, params, then, error) {
    axios.post("/api?action=" + action + "&time=" + now_time, params)
        .then(then)
        .catch(error);
}

function postWithHeader(action, params, then, error) {
    axios.post("/api?action=" + action + "&time=" + now_time, params, config)
        .then(then)
        .catch(error);
}

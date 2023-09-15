let token;
let email;

function error(msg) {
    Notiflix.Notify.failure(msg);
}

function warning(msg) {
    Notiflix.Notify.warning(msg);
}

const now_time = new Date().getHours() + new Date().getMinutes() + new Date().getSeconds();

function get(action, then, error) {
    axios.get("/api?action=" + action + "&time=" + now_time)
        .then(then)
        .catch(error)
}

function post(action, params, then, error) {
    axios.post("/api?action=" + action + "&time=" + now_time, params)
        .then(then)
        .catch(error);
}

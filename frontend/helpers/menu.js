window.onload = addMenu;

const buttons = ["Рейсы", "Аэропорты", "c", "d", "e"];
const functions = ["ShowAllFlights()", "ShowAirports()", "c", "d", "e"];
function addMenu() {
    main = document.getElementById('main')
    main.innerHTML = '';

    d = document.createElement('div');
    d.setAttribute("id", "navbar");

    for (let i = 0; i < buttons.length; i++) {
        a = document.createElement('a');
        a.setAttribute("onclick", functions[i]);
        a.innerText = buttons[i];
        d.appendChild(a);
    }

    let auth = getCookie('username');
    if (auth === undefined) {
        a = document.createElement('a');
        a.innerText = 'Login';
        a.setAttribute('id', 'login');
        a.setAttribute('onclick', 'addLoginBlock()');
        d.appendChild(a);
        a = document.createElement('a');
        a.innerText = 'Sign Up';
        a.setAttribute('id', 'signup');
        a.setAttribute('onclick', 'addSignUpBlock()');
        d.appendChild(a);
    } else {
        a = document.createElement('a');
        a.innerText = 'User :)';
        a.setAttribute('onclick', 'addUserBlock()');
        a.setAttribute("class", "user button");
        d.appendChild(a);
    }
    main.appendChild(d);
    d = document.createElement('div');
    d.setAttribute("id", "container");
    main.appendChild(d);
}

function getCookie(name) {
    if (document.cookie.split(';').filter((item) => item.trim().startsWith(name)).length) {
        return 'User';
    }
    return undefined;
}

function navigate(path) {
    window.history.pushState({}, path, window.location+path);
}
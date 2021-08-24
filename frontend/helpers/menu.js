window.onload = addMenu;

const buttons = ["a", "b", "c", "d", "e"];
document.cookie = "username=John; path=/; expires=Tue, 19 Jan 2038 03:14:07 GMT"
function addMenu() {
    main = document.getElementById('main')
    main.innerHTML = '';

    d = document.createElement('div');
    d.setAttribute("id", "navbar");

    for (let i = 0; i < buttons.length; i++) {
        a = document.createElement('a');
        a.setAttribute("href", buttons[i]);
        a.innerText = buttons[i];
        d.appendChild(a);
    }

    let auth = getCookie('username');
    if (auth === undefined) {
        a = document.createElement('a');
        a.innerText = 'Login';
        d.appendChild(a);
        a = document.createElement('a');
        a.innerText = 'Sign Up';
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

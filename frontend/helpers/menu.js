window.onload = addMenu;


const buttons = ["Рейсы", "Аэропорты", "Билеты"];
const functions = ["ShowAllFlights()", "ShowAllAirports()", ""];
async function addMenu() {
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


        let admin = await CheckAdmin();
        if (admin === true) {
            a = document.createElement('a');
            a.innerText = 'Админка';
            a.setAttribute('id', 'admin');
            a.setAttribute('onclick', 'addAdminBlock()');
            d.appendChild(a);
        }


    let check = await isAuthed();
    if (check === false) {
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
    window.history.pushState(path, path, path);
}

window.onpopstate = function (event) {
    state = event.state === null? 'state' : event.state;
    render(state);
};

function render(state) {
    console.log(state);
    switch (state) {
        case '/user':
            addUserBlock();
            break;
        case '/airports':
            ShowAllAirports();
            break;
        case '/admin':
            addAdminBlock();
            break;
        case '/tickets':
            GetAll();
            break;
        default:
            ShowAllFlights();
            break;
    }
}

function addError(text) {
    err = document.getElementById('error');
    if (err === undefined || err === null) {
        d = document.createElement('div');
        d.innerText = text;
        d.setAttribute("id", "error");
        main = document.getElementById('container');
        main.prepend(d);
    } else {
        err.innerText = text;
    }
}
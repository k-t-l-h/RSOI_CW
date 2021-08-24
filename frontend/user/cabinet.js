
function addUserBlock() {
    main = document.getElementById('container')
    main.innerHTML = '';

    d = document.createElement('div');
    d.setAttribute("class", "ticket-item");

    d.innerText = 'User';
    main.appendChild(d);
}


function addLoginBlock() {
    main = document.getElementById('container')
    main.innerHTML = '';

    d = document.createElement('div');
    d.setAttribute("class", "username");
    inputfield = document.createElement('input');
    inputfield.setAttribute("type", "text");
    inputfield.setAttribute("id", "username");
    inputfield.setAttribute("class", "user-input");
    inputfield.setAttribute("placeholder", "username");
    d.appendChild(inputfield);
    main.appendChild(d);

    d = document.createElement('div');
    d.setAttribute("class", "password");
    inputfield = document.createElement('input');
    inputfield.setAttribute("type", "password");
    inputfield.setAttribute("id", "password");
    inputfield.setAttribute("class", "pass-input");
    inputfield.setAttribute("placeholder", "password");
    d.appendChild(inputfield);
    main.appendChild(d);

    d = document.createElement('button');
    d.setAttribute("class", "signin-button");
    d.setAttribute('onclick', 'login()');
    d.innerText = 'Войти';
    main.appendChild(d);
}


function addSignUpBlock() {
    main = document.getElementById('container')
    main.innerHTML = '';

    d = document.createElement('div');
    d.setAttribute("class", "ticket-item");

    d.innerText = 'User';
    main.appendChild(d);
}

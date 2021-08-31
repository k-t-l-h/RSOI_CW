async function seeAllUsers() {
    const url = 'http://127.0.0.1:8010/api/v1/users';

    let response = await fetch(url, {
        method: 'GET',
        credentials: 'include',
        headers: {
            'Access-Control-Allow-Origin': 'http://127.0.0.1:8887/'
        },
    });
    if (response.ok) {

        main = document.getElementById('container')
        main.innerHTML = '';

        let airports = await response.json();
        for (let i = 0; i < airports.length; i++) {
            d = document.createElement('div');
            d.setAttribute("id", airports[i].id);
            d.setAttribute("class", "ticket-item");

            link = document.createElement('p');
            link.innerText = airports[i].id;
            d.appendChild(link);

            link = document.createElement('p');
            link.innerText = airports[i].login;
            d.appendChild(link);

            link = document.createElement('p');
            link.innerText = airports[i].role;
            d.appendChild(link);

            main.appendChild(d);
        }

    } else {
        console.log("Ошибка HTTP: " + response.status);
    }
}

async function createUserBlock() {
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
    d.setAttribute("class", "username");
    inputfield = document.createElement('input');
    inputfield.setAttribute("type", "text");
    inputfield.setAttribute("id", "role");
    inputfield.setAttribute("class", "user-input");
    inputfield.setAttribute("placeholder", "user role");
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
    d.setAttribute('onclick', 'signup()');
    d.innerText = 'Создать';
    main.appendChild(d);
}

async function getStatistics() {
    console.log('Stats')
}

async function getFillingStatistics() {
    console.log('Stats 2')
}

async function addFlightBlock() {
    main = document.getElementById('container')
    main.innerHTML = '';

    //TODO: добавить id
    d = document.createElement('div');

    inputfield = document.createElement('input');
    inputfield.setAttribute("type", "text");
    inputfield.setAttribute("id", "from");
    inputfield.setAttribute("class", "user-input");
    inputfield.setAttribute("placeholder", "Откуда");
    d.appendChild(inputfield);


    inputfield = document.createElement('input');
    inputfield.setAttribute("type", "text");
    inputfield.setAttribute("id", "to");
    inputfield.setAttribute("class", "user-input");
    inputfield.setAttribute("placeholder", "Куда");
    d.appendChild(inputfield);

    inputfield = document.createElement('input');
    inputfield.setAttribute("type", "date");
    inputfield.setAttribute("id", "date");
    inputfield.setAttribute("placeholder", "Когда");
    d.appendChild(inputfield);

    inputfield = document.createElement('input');
    inputfield.setAttribute("type", "text");
    inputfield.setAttribute("id", "tickets");
    inputfield.setAttribute("placeholder", "Сколько билетов");
    d.appendChild(inputfield);

    main.appendChild(d);
}

async function addFlightPatchBlock() {
    main = document.getElementById('container')
    main.innerHTML = '';

    d = document.createElement('div');

    inputfield = document.createElement('input');
    inputfield.setAttribute("type", "text");
    inputfield.setAttribute("id", "flight-id");
    inputfield.setAttribute("class", "user-input");
    inputfield.setAttribute("placeholder", "Откуда");
    d.appendChild(inputfield);

    b = document.createElement('button');
    b.setAttribute("class", "signin-button");
    b.setAttribute('onclick', 'ShowOneFlightPatch()');
    b.innerText = 'Найти';
    d.appendChild(b);
    main.appendChild(d);
}

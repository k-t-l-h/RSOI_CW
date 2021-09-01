async function CheckAdmin() {
    return true;
    const url = 'http://127.0.0.1:8010/api/v1/admin';
    let token = localStorage.getItem('token');

    let response = await fetch(url, {
        method: 'POST',
        credentials: 'same-origin',
        headers: {
            'Access-Control-Allow-Origin': 'http://127.0.0.1:8887/'
        },
        body: {token: token}
    });
    return response.ok;

}

function addAdminBlock() {
    navigate('/admin');
    main = document.getElementById('container')
    main.innerHTML = '';

    //6 кнопок
    d = document.createElement('div');
    d.innerText = 'Показать пользователей';
    d.setAttribute("class", "airport-items");
    d.setAttribute('onclick', 'seeAllUsers()');
    main.appendChild(d);

    d = document.createElement('div');
    d.innerText = 'Добавить пользователя';
    d.setAttribute("class", "airport-items");
    d.setAttribute('onclick', 'createUserBlock()');
    main.appendChild(d);

    d = document.createElement('div');
    d.innerText = 'Добавить рейс';
    d.setAttribute("class", "airport-items");
    d.setAttribute('onclick', 'addFlightBlock()');
    main.appendChild(d);

    d = document.createElement('div');
    d.innerText = 'Изменить рейс';
    d.setAttribute("class", "airport-items");
    d.setAttribute('onclick', 'addFlightPatchBlock()');
    main.appendChild(d);

    d = document.createElement('div');
    d.innerText = 'Посмотреть статистику';
    d.setAttribute("class", "airport-items");
    d.setAttribute('onclick', '');
    main.appendChild(d);

    d = document.createElement('div');
    d.innerText = 'Посмотреть статистику 2';
    d.setAttribute("class", "airport-items");
    d.setAttribute('onclick', '');
    main.appendChild(d);
}

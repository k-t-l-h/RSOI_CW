//узнать количество бонусов
async function GetMyBonus() {
    const url = 'http://127.0.0.1:8050/api/v1/miles';

    user = localStorage.getItem("token");

    let response = await fetch(url, {
        method: 'GET',
        credentials: 'include',
        headers: {
            'Authorization': `Bearer ${user}`,
            'Access-Control-Allow-Origin': 'http://127.0.0.1:8887'
        },
    });

    if (response.ok) {
        bonus = await response.json();
        main = document.getElementById('tickets-holder')
        if (main === null) {
            main = document.getElementById('container')
            d = document.createElement('div');
            d.setAttribute("id", "tickets-holder");
            main.appendChild(d);
            main = d;
        } else {
            main.innerHTML = '';
        }
        d = document.createElement('div');
        d.setAttribute("class", "ticket-item");
        id = document.createElement('p');
        id.innerText = 'Доступно бонусов: ' + bonus.balance;

        d.appendChild(id);
        main.appendChild(d);
    } else {
        addError("Бонусная система сейчас недоступна")
    }
}

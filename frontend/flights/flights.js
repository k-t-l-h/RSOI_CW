async function ShowAllFlights() {
    navigate('/tickets');
    const url = 'http://127.0.0.1:8030/api/v1/flights';

    let response = await fetch(url, {
        method: 'GET',
        credentials: 'same-origin',
        headers: {
            'Access-Control-Allow-Origin': '*'
        },
    });
    if (response.ok) {
        main = document.getElementById('container')
        main.innerHTML = '';
        let body = await response.json();
        for (let i = 0; i < body.length; i++) {
            d = document.createElement('div');
            d.setAttribute("id", body[i].id);
            d.setAttribute("class", "flight-item");

            a = document.createElement('a');
            a.innerText = body[i].from_city;
            a.setAttribute("id", body[i].from);
            a.setAttribute('onclick', 'ShowOneFlight(${body[i].from}))');
            d.appendChild(a);

            a = document.createElement('a');
            a.innerText = body[i].to_city;
            a.setAttribute("id", body[i].to);
            a.setAttribute('onclick', 'ShowOneFlight(${body[i].to}))');
            d.appendChild(a);

            main.appendChild(d);
        }

    } else {
        console.log("Ошибка Билетов HTTP: " + response.status);
    }

}

async function ShowOneFlight(id) {

}

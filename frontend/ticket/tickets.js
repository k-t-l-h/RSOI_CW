function Buy() {
    let from = document.getElementById("username").value;
    let to = document.getElementById("password").value;
    let date = document.getElementById("password-2").value;

}

async function GetAll() {
    navigate('/tickets');
    const url = 'http://127.0.0.1:8040/api/v1/tickets';

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
        for (let i = 0; i < tickets.all.length; i++) {
            d = document.createElement('div');
            d.setAttribute("id", tickets.all[i].id);
            d.setAttribute("class", "ticket-item");


            //link.addEventListener('onclick', ShowAirport.bind(airports.all[i].id), false);
            id = document.createElement('p');
            id.innerText = tickets.all[i].id;
            namef = document.createElement('p');
            namef.innerText = tickets.all[i].flight;
            city = document.createElement('p');
            city.innerText = tickets.all[i].city;

            link = document.createElement('button');
            link.innerText = 'Return';
            link.setAttribute('onclick', 'Return(' + d.id + ')');
            if (tickets.all[i].date < Date()) {
                link.disabled = false;
            }


            d.appendChild(id);
            d.appendChild(namef);
            d.appendChild(city);
            d.appendChild(link);
            main.appendChild(d);
        }

    } else {
        console.log("Ошибка Билетов HTTP: " + response.status);
    }


}

function Return(id) {
    main = document.getElementById(id);
    main.childNodes[main.childElementCount-1].disabled = true;
    console.log('Returned');
}

async function GetMyTickets(id) {
    const url = 'http://127.0.0.1:8040/api/v1/tickets/'+id.toString();

    let response = await fetch(url, {
        method: 'GET',
        credentials: 'same-origin',
        headers: {
            'Access-Control-Allow-Origin': '*'
        },
    });
    if (response.ok) {
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
        let tickets = await response.json();
        for (let i = 0; i < tickets.length; i++) {
            d = document.createElement('div');
            d.setAttribute("id", tickets[i].id);
            d.setAttribute("class", "ticket-item");


            //link.addEventListener('onclick', ShowAirport.bind(airports.all[i].id), false);
            id = document.createElement('p');
            id.innerText = tickets[i].ticket_uuid;
            namef = document.createElement('p');
            namef.innerText = tickets[i].flight_uuid;

            link = document.createElement('div');
            link.innerText = 'Return';
            link.setAttribute('onclick', 'Return(' + d.id + ')');
            link.setAttribute("class", "btn");
            if (tickets[i].date < Date()) {
                link.disabled = false;
            }


            d.appendChild(id);
            d.appendChild(namef);
            d.appendChild(link);
            main.appendChild(d);
        }

    } else {
        console.log("Ошибка Билетов HTTP: " + response.status);
    }


}
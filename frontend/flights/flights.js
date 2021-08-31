async function ShowAllFlights() {
    main = document.getElementById('container')
    main.innerHTML = '';

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
        console.log(body);
        for (let i = 0; i < body.length; i++) {
            d = document.createElement('div');
            d.setAttribute("id", body[i].id);
            d.setAttribute("class", "flight-item");

            a = document.createElement('p');
            a.innerText = body[i].from_city;
            a.setAttribute("id", body[i].from);
            a.setAttribute('onclick', 'ShowOneFlight("' +body[i].from + '")')
            d.appendChild(a);

            a = document.createElement('p');
            a.innerText = body[i].to_city;
            a.setAttribute("id", body[i].to);
            a.setAttribute('onclick', 'ShowOneFlight("' +body[i].to + '")');
            d.appendChild(a);

            a = document.createElement('p');
            a.innerText = body[i].date;
            d.appendChild(a);

            main.appendChild(d);
        }

    } else {
        console.log("Ошибка Билетов HTTP: " + response.status);
    }

}

async function ShowOneFlight(id) {
    console.log(id)
}

async function ShowOneFlightPatch() {
    let id = document.getElementById("flight-id").value;

    const url = 'http://127.0.0.1:8030/api/v1/flights/'+id;

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

        d = document.createElement('div');
        d.setAttribute("class", "username");
        inputfield = document.createElement('input');
        inputfield.setAttribute("type", "text");
        inputfield.setAttribute("id", "username");
        inputfield.setAttribute("class", "user-input");
        inputfield.setAttribute("value", body.id);
        inputfield.setAttribute("placeholder", "username");
        d.appendChild(inputfield);

        inputfield = document.createElement('input');
        inputfield.setAttribute("type", "text");
        inputfield.setAttribute("id", "username");
        inputfield.setAttribute("class", "user-input");
        inputfield.setAttribute("value", body.from);
        inputfield.setAttribute("placeholder", "username");
        d.appendChild(inputfield);


        inputfield = document.createElement('input');
        inputfield.setAttribute("type", "text");
        inputfield.setAttribute("id", "username");
        inputfield.setAttribute("class", "user-input");
        inputfield.setAttribute("value", body.to);
        inputfield.setAttribute("placeholder", "username");
        d.appendChild(inputfield);

        inputfield = document.createElement('input');
        inputfield.setAttribute("type", "date");
        inputfield.setAttribute("id", "username");
        inputfield.setAttribute("class", "user-input");
        inputfield.setAttribute("value", body.date);
        inputfield.setAttribute("placeholder", "username");
        d.appendChild(inputfield);


        main.appendChild(d);



    } else {
        console.log("Ошибка Билетов HTTP: " + response.status);
    }

}

let airports = { all: [
    {id: 1, name: "a", city: "b"},
        {id: 1, name: "a", city: "b"}]
}

window.onload = ShowAllAirports;

function ShowAllAirports() {
    main = document.getElementById('container')
    main.innerHTML = '';

    for (let i = 0; i < airports.all.length; i++) {
        d = document.createElement('div');
        d.setAttribute("id", airports.all[i].id);
        d.setAttribute("class", "ticket-item");

        link = document.createElement('button');
        link.innerText = airports.all[i].id;
        link.setAttribute('onclick', 'ShowAirport('+d.id+')');
        //link.addEventListener('onclick', ShowAirport.bind(airports.all[i].id), false);
        namef = document.createElement('p');
        namef.innerText = airports.all[i].name;
        city = document.createElement('p');
        city.innerText = airports.all[i].city;

        d.appendChild(link);
        d.appendChild(namef);
        d.appendChild(city);
        main.appendChild(d);
    }


}

async function ShowAirport(id) {
    console.log("Show");
    main = document.getElementById('container')
    main.innerHTML = '';

    url = "http:/localhost/airports/" + id;
    let response = await fetch(url, {
        method: 'GET',
        credentials: 'same-origin',
        headers: {
            'Content-Type': 'application/json',
        },
    });

    if (response.ok) {
        let airport = await response.json();

        d = document.createElement('div');
        d.setAttribute("id", airport.id);
        d.setAttribute("class", "ticket-item");

        link = document.createElement('a');
        link.innerText = airport.id;
        link.href = airport.id;
        namef = document.createElement('p');
        namef.innerText = airport.name;
        city = document.createElement('p');
        city.innerText = airport.city;

        d.appendChild(link);
        d.appendChild(namef);
        d.appendChild(city);
        main.appendChild(d);

    } else {
        d = document.createElement('div');
        sorry = document.createElement('p');
        sorry.innerText = 'Аэропорт не найден';
        d.appendChild(sorry);
        main.appendChild(d);
        console.log("Ошибка HTTP: " + response.status);
    }

}

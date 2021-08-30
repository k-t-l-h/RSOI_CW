async function ShowAllAirports() {

    navigate('/airports');
    const url = 'http://127.0.0.1:8020/api/v1/airports';

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

        let airports = await response.json();
        for (let i = 0; i < airports.length; i++) {
            d = document.createElement('div');
            d.setAttribute("id", airports[i].id);
            d.setAttribute("class", "airport-items");

            link = document.createElement('a');
            link.innerText = airports[i].id;
            link.setAttribute('onclick', 'ShowAirport("'+airports[i].id.toString()+'")');
            namef = document.createElement('p');
            namef.innerText = airports[i].name;

            d.appendChild(link);
            d.appendChild(namef);
            main.appendChild(d);
        }
    }


}

async function ShowAirport(id) {

    navigate('/airports');
    const url = 'http://127.0.0.1:8020/api/v1/airports/' + id;

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

        let airports = await response.json();
            d = document.createElement('div');
            d.setAttribute("id", airports.id);
            d.setAttribute("class", "airport-item");

            link = document.createElement('a');
            link.innerText = airports.id;
            namef = document.createElement('p');
            namef.innerText = airports.name;
            city = document.createElement('p');
            city.innerText = airports.city;
            description = document.createElement('p');
            description.innerText = airports.description;

            d.appendChild(link);
            d.appendChild(namef);
            d.appendChild(city);
            d.appendChild(description);
            main.appendChild(d);
    }

}

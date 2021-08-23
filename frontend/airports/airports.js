let airports = { all: [
    {id: 1, name: "a", city: "b"},
        {id: 1, name: "a", city: "b"}]
}

window.onload = ShowAirports;

function ShowAirports() {
    main = document.getElementById('container')
    main.innerHTML = '';

    for (let i = 0; i < airports.all.length; i++) {
        d = document.createElement('div');
        d.setAttribute("id", airports.all[i].id);
        d.setAttribute("class", "ticket-item");

        link = document.createElement('a');
        link.innerText = airports.all[i].id;
        link.href = airports.all[i].id;
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
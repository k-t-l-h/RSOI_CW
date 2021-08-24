let tickets = { all: [
        {id: 1, flight: "a", city: "b", date: Date()},
        {id: 2, flight: "a", city: "b", date: Date()}],
}

function Buy() {
    let from = document.getElementById("username").value;
    let to = document.getElementById("password").value;
    let date = document.getElementById("password-2").value;

}


window.onload = GetAll;
function GetAll() {

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
        link.setAttribute('onclick', 'Return('+d.id+')');
        if (tickets.all[i].date < Date()) {
            link.disabled = false;
        }


        d.appendChild(id);
        d.appendChild(namef);
        d.appendChild(city);
        d.appendChild(link);
        main.appendChild(d);
    }
}

function Return(id) {
    main = document.getElementById(id);
    main.childNodes[main.childElementCount-1].disabled = true;
    console.log('Returned');
}

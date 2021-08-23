let flights = { all: [
    {id: 1, from: "a", to: "b", available: 0},
        {id: 2, from: "c", to: "d", available: 0},
        {id: 3, from: "e", to: "f", available: 0}]
}


function ShowFlights() {
    console.log("Click");
    main = document.getElementById('container')
    main.innerHTML = '';

    for (let i = 0; i < flights.all.length; i++) {
        d = document.createElement('div');
        d.setAttribute("id", flights.all[i].id);
        d.setAttribute("class", "ticket-item");

        link = document.createElement('a');
        link.innerText = flights.all[i].id;
        link.href = flights.all[i].id;
        from = document.createElement('p');
        from.innerText = flights.all[i].from;
        to = document.createElement('p');
        to.innerText = flights.all[i].to;

        tickets = document.createElement('p');
        tickets.innerText = flights.all[i].to;

        d.appendChild(link);
        d.appendChild(from);
        d.appendChild(to);
        main.appendChild(d);
    }


}
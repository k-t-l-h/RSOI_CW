
function addUserBlock() {
    main = document.getElementById('container')
    main.innerHTML = '';

    d = document.createElement('div');
    d.setAttribute("class", "ticket-item");

    d.innerText = 'User';
    main.appendChild(d);
}

async function CheckAdmin() {
    const url = 'http://127.0.0.1:8010/api/v1/admin';

    let response = await fetch(url, {
        method: 'POST',
        credentials: 'same-origin',
        headers: {
            'Access-Control-Allow-Origin': '*'
        },
    });
    return response.ok;
}

function addAdminBlock() {
    main = document.getElementById('container')
    main.innerHTML = '';
    main.innerText = 'ADMIN PAGE';

    //6 кнопок
}

async function login() {
   let user = document.getElementById("username").value;
   let user_pass = document.getElementById("password").value;
   const url = 'http://127.0.0.1:8010/api/v1/auth';

   let auth = user +':' + user_pass;
   let response = await fetch(url, {
      method: 'POST',
      credentials: 'include',
      headers: {
         'Authorization': `Basic ${(btoa(auth))}`,
         'Access-Control-Allow-Origin': 'http://127.0.0.1:8887/'
      },
   });
   if (response.ok) {
      console.log(document.cookie)
      response.headers.forEach(function(val, key) { console.log(key + ' -> ' + val); });
      console.log(document.cookie)
      //document.cookie = response.headers.get('Set-Cookie');
      b1 = document.getElementById('login');
      b1.parentNode.removeChild(b1);
      b2 = document.getElementById('container');
      b2.parentNode.removeChild(b2);

      d = document.getElementById('navbar');
      a.innerText = 'User :)';
      a.setAttribute('onclick', 'addUserBlock()');
      a.setAttribute("class", "user button");
      d.appendChild(a);

      } else {
        console.log("Ошибка HTTP: " + response.status);
   }
}

async function signup() {
   let user = document.getElementById("username").value;
   let user_pass = document.getElementById("password").value;
   let user_pass_2 = document.getElementById("password-2").value;

   if (user_pass !== user_pass_2) {
      let validityState = document.getElementById("password").validity;
      console.log(validityState);
      document.getElementById("password").setCustomValidity('Пароли не совпадают');
      document.getElementById("password").reportValidity();
      console.log(document.getElementById("password").validity);
   } else {
      const url = 'http://127.0.0.1:8010/api/v1/auth';
      const data = {login: user, password: user_pass};

      let response = await fetch(url, {
         method: 'POST',
         credentials: 'same-origin',
         headers: {
            'Content-Type': 'application/json',
         },
         body: JSON.stringify(data)
      });
      if (response.ok) {
         let user = await response.json();
         //TODO: проставить тег isAuth
      } else {
         //TODO: вывести ошибку
         console.log("Ошибка HTTP: " + response.status);
      }
   }
}

function utf8_to_b64(str) {
   return window.btoa(unescape(encodeURIComponent(str)));
}

function b64_to_utf8(str) {
   return decodeURIComponent(escape(window.atob(str)));
}

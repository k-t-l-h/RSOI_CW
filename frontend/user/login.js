alert(document.cookie);

async function login() {
   let user = document.getElementById("username").value;
   let user_pass = document.getElementById("password").value;
   const url = 'http://127.0.0.1:8000/login';
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
      const url = 'http://127.0.0.1:8000/signup';
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

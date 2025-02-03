function submit() {
  let email = document.getElementById("inp-email").value;
  let password = document.getElementById("inp-password").value;
  if (email == "" || password == "") {
    alert("One or more input fields are empty.");
    return;
  }
  let auth = {
    email: email,
    password: password,
  };
  fetch("/api/login", {
    method: "POST",
    body: JSON.stringify(auth),
    headers: {
      "Content-type": "application/json; charset=UTF-8",
    },
  }).then((response) => {
    if (response.status != 200) {
      response.json().then((json) => {
        alert(json.error);
      });
    }
    response.json().then((json) => {
      localStorage.setItem("relay-auth-token", json.token);
      window.location.href = "/app";
    });
  });
}

function register() {
  window.location.href = "/register";
}

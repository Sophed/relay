function submit() {
  let username = document.getElementById("inp-username").value;
  let email = document.getElementById("inp-email").value;
  let password = document.getElementById("inp-password").value;
  if (username == "" || email == "" || password == "") {
    alert("One or more input fields are empty.");
    return;
  }
  let auth = {
    username: username,
    email: email,
    password: password,
  };
  fetch("/api/register", {
    method: "POST",
    body: JSON.stringify(auth),
    headers: {
      "Content-type": "application/json; charset=UTF-8",
    },
  }).then((response) => {
    if (response.status != 200) {
      alert(response.statusText);
      return;
    }
    response.json().then((json) => {
      window.location.href = "/app";
    });
  });
}

function login() {
  window.location.href = "/login";
}

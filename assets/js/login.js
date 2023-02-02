
login = document.getElementById("loginbtn");
signup = document.getElementById("signupbtn");
login.addEventListener("click", () => {
    fetch("/login", {
        method : 'POST',
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json'
          },
        body : JSON.stringify( {
            'email' : document.getElementById("loginemail").value,
            'password' : document.getElementById("loginpass").value,
        })
    })
    .then(async response => {
        if(response.status == 200) {
            window.location.href = "/todo";
        } else {
            let body = await response.json();
            if(body.error) {
                console.error(body.error);
                document.getElementById('error').innerHTML=body.error;
            }
            // var str = JSON.stringify(response.json());
            // document.write(str)
        }
        
    })
    .catch(error => {
        console.error(error);
    })
});

signup.addEventListener("click", () => {
    fetch("/signup", {
        method : 'POST',
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json'
          },
        body : JSON.stringify( {
            'username' : document.getElementById("signupname").value,
            'email' : document.getElementById("signupemail").value,
            'password' : document.getElementById("signuppass").value
        })
    })
    .then(response => {
        if(response.status == 200) {
            window.location.href = "/todo";
        } else {
            var str = JSON.stringify(response.json());
            document.write(str)
        }
        
    })
});

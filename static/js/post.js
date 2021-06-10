/*
--------------------------------------------------
--------------------TextArea----------------------
--------------------------------------------------
*/

var AreaPost = document.getElementById("AreaPost");
var btnCom = document.getElementById("btn-comment");

btnCom.addEventListener("click", () => {
  if(getComputedStyle(AreaPost).display != "none"){
    AreaPost.style.display = "none";
  } else {
    AreaPost.style.display = "flex";
  }
});


/*
---------------------------------------------------------------
--------------------Cookies Part-------------------------------
---------------------------------------------------------------


--------------------------------------------------
------------------Username------------------------
-----------Connexion/Deconnexion------------------
--------------------------------------------------
*/
function checkCookies(cookies) {
  var btnDeconnexion = document.getElementById("Deconnexion");
  var btnUser = document.getElementById("UserName");
  var btnConnexion = document.getElementById("Connexion");
  var btnInscription = document.getElementById("Inscription");
  var cookies = document.cookie;
  console.log("La liste des cookies :", cookies)

  if (cookies == "") {
    btnConnexion.style.display = "block";
    btnInscription.style.display = "block";
    btnDeconnexion.style.display = "none";
    btnUser.style.display = "none";
  } else {
    btnConnexion.style.display = "none";
    btnInscription.style.display = "none";
    btnDeconnexion.style.display = "block";
    btnUser.style.display = "block";
  }
}

window.onload = checkCookies();


/*
--------------------------------------------------
---------------Admin/Modo button------------------
--------------------------------------------------
*/

// function checkCookiesRole() {
//   var btnSupPost = document.getElementById("btn-sup");
//   // var btnSupCom = document.getElementById("btn-supcom");
//   var cookieValue = document.cookie
//   .split('; ')
//   .find(row => row.startsWith('Role='))
//   .split('=')[1];
  
//   if (cookieValue == undefined) {
//     if (cookieValue != "admin") {
//       console.log(cookieValue)
//       btnSupPost.style.display = "none";
//       // btnSupCom.style.display = "flex";
      
//     } else {
//       btnSupPost.style.display = "flex";
//       // console.log("il y a des cookies");
//       // btnSupCom.style.display = "none";
//     }
//   } else {
//     btnSupPost.style.display = "none";
//     console.log("pas de cookies ici");
//   }
  
        
// }

// window.onload = checkCookiesRole();

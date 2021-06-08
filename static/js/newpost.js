/*
--------------------------------------------------
------------------Username------------------------
-----------Connexion/Deconnexion------------------
--------------------------------------------------
*/
function checkCookies() {
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
  
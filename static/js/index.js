/*
--------------------------------------------------
------------------BouttonFiltre------------------
--------------------------------------------------
*/
var btn = document.getElementById("btnfilter");
var filter = document.getElementById("filtrer");

btn.addEventListener("click", () => {
  if(getComputedStyle(filter).display != "none"){
    filter.style.display = "none";
  } else {
    filter.style.display = "block";
  }
});


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


/*
--------------------------------------------------
------------------RandomColor---------------------
--------------------------------------------------
*/

function random_bg_color() {
  for (let id = 1; id < 11 ;id++) {
    var x = Math.floor(Math.random() * 256);
    var y = Math.floor(Math.random() * 256);
    var z = Math.floor(Math.random() * 256);
    var bgColor = "rgb(" + x + ", " + y + ", " + z + ")";
    console.log("Couleur aléatoire : ", bgColor);

    document.getElementById(id).style.backgroundColor = bgColor;
  }
}

window.onload = random_bg_color();
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

function checkCookiesRole() {
  // var btnSupPost = document.getElementById("suppriPost");
  // var btnSupCom = document.getElementById("suppriCom");
  const cookies2 = document.cookie;
  // console.log("La liste des cookies :", cookies2)
  const element = [] 
  element.push(cookies2);

  console.log(element)

  for (const i of element) {

    console.log(i);

    
  }


  
    
    

    
}

  // if (cookies == "") {
  //   btnSupPost.style.display = "none";
  //   btnSupCom.style.display = "none";
    
  // } else {
  //   console.log("il y a des cookies");
  //   btnSupPost.style.display = "flex";
  //   btnSupCom.style.display = "flex";
  // }


window.onload = checkCookiesRole();

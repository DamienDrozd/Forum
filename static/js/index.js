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
------------------Username Avatar-----------------
-----------------------Cookies--------------------
--------------------------------------------------
*/
var btn2 = document.querySelector("btnlogin");

function getCookies() {
  const allcookies = document.cookie
  console.log(allcookies)
  var cut = document.cookie.split(';');
  console.log(cut);

  for (var i = 1; i < cut.length-2; i++) {
    var c = cut[i]
    var user = c.split("=");
    var test = ""
    test += user[1]
    console.log(test)
    // console.log(user[1])
    // document.write(user[1])
    // var nom = document.querySelector('UserName') ;
    // nom.innerHTML = user[1]
    // document.querySelector('UserName').innerHTML = test;
  }
  

  return null
  
}

// getCookies()


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
    console.log(bgColor);

    document.getElementById(id).style.backgroundColor = bgColor;
  }
}

window.onload = random_bg_color();
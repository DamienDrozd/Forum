var btn = document.getElementById("btnfilter");
var filter = document.getElementById("filtrer");

btn.addEventListener("click", () => {
    if(getComputedStyle(filter).display != "none"){
      filter.style.display = "none";
    } else {
      filter.style.display = "block";
    }
});
var modal = document.getElementsByClassName("modal")[0];
var btn = document.getElementsByClassName("bron")[0];
var close = document.getElementsByClassName("btn-close")[0];

btn.onclick = function() {
    modal.style.display = "flex";
}

close.onclick = function() {
    modal.style.display = "none";
}
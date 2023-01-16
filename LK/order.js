var onDelete = document.getElementsByClassName("delete");
var onViuew = document.getElementsByClassName("viuew");
var modal = document.getElementsByClassName("modal")[0];

for (var i = 0; i < onDelete.length; i++) {
    onDelete[i].onclick = function() {
        fetch("http://192.168.31.104:8080/api/orders/" + this.id, {
            method: "DELETE",
            headers: {
                "Content-Type": "application/json"
            },
        });

        this.parentNode.parentNode.parentNode.remove();
    };
}

for (var i = 0; i < onViuew.length; i++) {
    onViuew[i].onclick = function() {
        var guid = modal.getElementsByClassName("guid-name")[0];
        var routr = modal.getElementsByClassName("route-name")[0];
        var price = modal.getElementsByClassName("price")[0];
        var date = modal.getElementsByClassName("date-input")[0];
        var time = modal.getElementsByClassName("time-input")[0];
        var persons = modal.getElementsByClassName("peoples-input")[0];
        var timeCount = modal.getElementsByClassName("count-time-input")[0];
        var op1 = modal.getElementsByClassName("opt-1")[0];
        var op2 = modal.getElementsByClassName("opt-2")[0];
        var price = modal.getElementsByClassName("price")[0];

        // atter
        console.log(5);
        date.readOnly = true;
        time.readOnly = true;
        persons.readOnly = true;
        timeCount.readOnly = true;
        op1.onclick = function () { return false; };
        op2.onclick = function () { return false; };

        guid.innerHTML = this.parentNode.getElementsByClassName("guid-name")[0].innerHTML.replaceAll("+", " ");
        routr.innerHTML = this.parentNode.getElementsByClassName("route-name")[0].innerHTML;
        price.innerHTML = this.parentNode.getElementsByClassName("price")[0].innerHTML;
        date.value = this.parentNode.getElementsByClassName("date")[0].innerHTML;
        time.value = this.parentNode.getElementsByClassName("time")[0].innerHTML;
        persons.value = this.parentNode.getElementsByClassName("persons")[0].innerHTML;
        timeCount.value = this.parentNode.getElementsByClassName("duration")[0].innerHTML;
        op1.checked = (this.parentNode.getElementsByClassName("optionFirst")[0].innerHTML === 'true');
        op2.checked = (this.parentNode.getElementsByClassName("optionSecond")[0].innerHTML === 'true');
        modal.style.display = "flex";
    };

}
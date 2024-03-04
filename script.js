document.addEventListener("DOMContentLoaded", function() {
    var getTimeBtn = document.getElementById("get-time-btn");
    var timeDisplay = document.getElementById("time-display");

    getTimeBtn.addEventListener("click", function() {
        fetch("/time")
            .then(response => response.text())
            .then(data => {
                timeDisplay.innerText = data;
            })
            .catch(error => {
                console.error("Error:", error);
            });
    });
});

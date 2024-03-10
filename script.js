document.getElementById("get-time-btn").addEventListener("click", function() {
    fetch('/time')
        .then(response => response.text())
        .then(data => {
            document.getElementById("time-display").innerHTML = data;
        })
        .catch(error => {
            console.error('Error:', error);
        });
});

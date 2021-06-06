//executes when the button is clicked! in Generals

document.getElementById("checkAvailabilitybutton").addEventListener("click", function () {
    const html = `
<form action="" method="POST" novalidate id="check-availability-form">
    <div class="row">
        <div class="row mb-3" id="reservation-date-modal">
            <div class="col-6">
                <label for="start_date" class="form-label">Starting Date</label>
                <input disabled placeholder="yyyy-mm-dd" type="text" class="form-control" id="start_date" name="start"
                    aria-describedby="startEmailHelp" required>
                <div id="emailHelp" class="form-text">Enter your starting date for stay.</div>
            </div>
            <div class="col-6">
                <label for="end_date" class="form-label">Ending Date</label>
                <input disabled placeholder="yyyy-mm-dd" type="text" class="form-control" id="end_date" name="end"
                    aria-describedby="endEmailHelp" required>
                <div id="emailHelp" class="form-text">Enter your ending date of stay.</div>
            </div>
        </div>
    </div>
    <hr />
</form>
`;

    attention.custom({
        text: html,
        title: "Choose your dates.",
        callback: function (result) {
            const form = document.getElementById("check-availability-form");
            const formData = new FormData(form);
            formData.append("csrf_token", "{{.CSRFToken}}");

            fetch("/search-availability-json", {
                method: "POST",
                body: formData,
            })
                .then(res => res.json())
                .then(data => console.log(data))
        }
    });
})
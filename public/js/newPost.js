const submitButton = document.getElementById("submitButton");
const SUBMISSION_TIMEOUT = 5000 // in ms

/**
 *  builds the post submission, with headers, and correct json body
 */
function submitPost(event, form) {
    event.preventDefault();
    submitButton.disabled = true;
    setTimeout(()=> submitButton.disabled = false, SUBMISSION_TIMEOUT);
    // get the fields from the form
    const title = document.getElementById("titleField").value;
    const postBody = document.getElementById("bodyField").value.split("\n");
    console.log(JSON.stringify({"title": title, "body": postBody}))
    // build the headers
    $.ajax({
        type: "POST",
        url: "/CreatePost",
        dataType: "json",
        contentType: "application/json; charset=utf-8",
        data: JSON.stringify({"Title": title, "Body": postBody}),
        success: (e) => alert("post was successful"),
    })
}

if(submitButton) {
    submitButton.addEventListener("click", (e) => {
        submitPost(e, this);
    })

}
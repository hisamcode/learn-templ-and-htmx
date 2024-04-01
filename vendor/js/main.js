htmx.logAll();
function SA(elt, config) {
    Swal.fire(config)
        .then((result) => {
            if (result.isConfirmed) {
                // htmx.trigger(this, 'confirmed')
                // htmx.trigger(elt, 'confirmed')
                elt.dispatchEvent(new Event('confirmed'));
                // htmx.trigger(elt, 'custom')
            }
        })
}


// document.addEventListener("htmx:confirm", function (e) {
//     console.log("oke")
//     e.preventDefault()
//     Swal.fire({
//         title: "Proceed?",
//         text: `I ask you... ${e.detail.question}`
//     }).then(function (result) {
//         if (result.isConfirmed) e.detail.issueRequest(true) // use true to skip window.confirm
//     })
// })
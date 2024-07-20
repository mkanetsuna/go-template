const evtSource = new EventSource("/reload");
evtSource.onmessage = function(event) {
    if (event.data === "reload") {
        window.location.reload();
    }
};
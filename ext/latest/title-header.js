htmx.defineExtension("title-header", {
    onEvent: function (name, evt) {
        if (name !== "htmx:afterRequest") return;
        const hxTitle = evt.detail.xhr.getResponseHeader("Hx-Title");
        if (hxTitle) window.document.title = hxTitle;
    },
});

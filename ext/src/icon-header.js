htmx.defineExtension("icon-header", {
    onEvent: function (name, evt) {
        if (name !== "htmx:afterRequest") return;
        const hxIcon = evt.detail.xhr.getResponseHeader("Hx-Icon");
        if (hxIcon) {
            var link = document.querySelector("link[rel='icon']");
            if (!link) {
                link = document.createElement("link");
                link.rel = "icon";
                document.head.appendChild(link);
            }
            link.href = hxIcon;
        }
    },
});

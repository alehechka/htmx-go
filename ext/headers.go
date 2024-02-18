package htmx_ext

// htmx custom extension headers
const (
	// HxTitle is associated with a custom htmx extension, "title-header", that intercepts this response header and updates the document's title to the new value when present.
	HxTitle string = "Hx-Title"
	// HxIcon is associated with a custom htmx extension, "icon-header", that intercepts this response header and updates the document's icon link element when present.
	HxIcon string = "Hx-Icon"
)

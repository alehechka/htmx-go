# Custom htmx extensions

This library includes a few custom `htmx` extensions that are not officially listed in htmx's [extension page](https://htmx.org/extensions). The extensions are written in raw JavaScript to make usage and maintenance of them simple. Additionally, these files will just be hosted via GitHub and cached with their CDNs. 

## Interfacing with Go

For extensions such as `title-header` and `icon-header` that introduce a new response header, these constants will be defined under [/ext](../ext/) and can be imported as `htmx_ext`. 

## Available Extensions and Usage

### The `title-header` Extension

This extension adds client support for a new `Hx-Title` response header that when present will update the browser tab's title element with the content provided. If the header is provided but with no content, the title will not be updated.

#### Install

```html
<script defer src="https://raw.githubusercontent.com/alehechka/htmx-go/v0.2.0/ext/title-header.js"></script>
```

#### Usage

```html
<body hx-ext="title-header">
 ...
</body>
```

### The `icon-header` Extension

This extension adds client support for a new `Hx-Icon` response header that when present will update the browser tab's `<link rel="icon" />` element's `href` attribute. This allows dynamic browser icons on HTTP responses. If the header is provided but with no content, the `href` will not be updated. If a `<link rel="icon">` element does not exist in the DOM, one will be created with the provided `href` from the response header.

```html
<script defer src="https://raw.githubusercontent.com/alehechka/htmx-go/v0.2.0/ext/icon-header.js"></script>
```

#### Usage

```html
<body hx-ext="icon-header">
 ...
</body>
```

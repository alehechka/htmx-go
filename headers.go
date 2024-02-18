package htmx

// htmx Header Reference: https://htmx.org/reference/#headers

// htmx Request Headers: https://htmx.org/reference/#request_headers
const (
	// HxBoosted is a request header that indicates that the request is via an element using hx-boost
	HxBoosted string = "Hx-Boosted"
	// HxCurrentURL is a request header that represents the current URL of the browser
	HxCurrentURL string = "Hx-Current-URL"
	// HxHistoryRestoreRequest is a request header that is "true" if the request is for history restoration after a miss in the local history cache
	HxHistoryRestoreRequest string = "Hx-History-Restore-Request"
	// HxPrompt is a request header that represents the user response to an hx-prompt
	HxPrompt string = "Hx-Prompt"
	// HxRequest is a request header that is always "true" if request is sent via htmx
	HxRequest string = "Hx-Request"
	// HxTarget is a request header that represents the id of the target element if it exists
	HxTarget string = "Hx-Target"
	// HxTriggerName is a request header that represents the name of the triggered element if it exists
	HxTriggerName string = "Hx-Trigger-Name"
)

// htmx Response Headers: https://htmx.org/reference/#response_headers
const (
	// HxLocation is a response header that allows you to do a client-side redirect that does not do a full page reload, see the documentation for more info: https://htmx.org/headers/hx-location/
	HxLocation string = "Hx-Location"
	// HxPushUrl is a response header that pushes a new url into the history stack, see the documentation for more info: https://htmx.org/headers/hx-push-url/
	HxPushUrl string = "Hx-Push-Url"
	// HxRedirect is a response header that can be used to do a client-side redirect to a new location
	HxRedirect string = "Hx-Redirect"
	// HxRefresh is a response header that if set to “true” the client side will do a a full refresh of the page
	HxRefresh string = "Hx-Refresh"
	// HxReplaceUrl is a response header that replaces the current URL in the location bar, see the documentation for more info: https://htmx.org/headers/hx-replace-url/
	HxReplaceUrl string = "Hx-Replace-Url"
	// HxReswap is a response header that allows you to specify how the response will be swapped. See hx-swap for possible values
	HxReswap string = "Hx-Reswap"
	// HxRetarget is a response header that represents a CSS selector that updates the target of the content update to a different element on the page
	HxRetarget string = "Hx-Retarget"
	// HxReselect is a response header that represents a CSS selector that allows you to choose which part of the response is used to be swapped in. Overrides an existing hx-select on the triggering element
	HxReselect string = "Hx-Reselect"
	// HxTriggerAfterSettle is a response header that allows you to trigger client side events, see the documentation for more info: https://htmx.org/headers/hx-trigger/
	HxTriggerAfterSettle string = "Hx-Trigger-After-Settle"
	// HxTriggerAfterSwap is a response header that allows you to trigger client side events, see the documentation for more info: https://htmx.org/headers/hx-trigger/
	HxTriggerAfterSwap string = "Hx-Trigger-After-Swap"
)

// htmx Request/Response Headers
const (
	// HxTrigger can be used as both a request and response header.
	//	- Request Header: represents the id of the triggered element if it exists
	//	- Response Header: allows you to trigger client side events, see the documentation for more info: https://htmx.org/headers/hx-trigger/
	HxTrigger string = "Hx-Trigger"
)

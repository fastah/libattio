package libattio

var (
	// ContextAPIKey can be used as a key name for context.WithValue()'s second parameter,
	// allowing you to pass an API key to functions via the context.
	ContextAPIKey = contextKey("apiKey")
)

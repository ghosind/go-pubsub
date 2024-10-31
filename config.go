package pubsub

// Config is the universal configuration for the pubsub client.
type Config struct {
	// URL is the URL of the broker.
	Address string
	// Username is the username to use for authentication.
	Username string
	// Password is the password to use for authentication.
	Password string
	// Logger is the logger to use for logging.
	Logger Logger
}

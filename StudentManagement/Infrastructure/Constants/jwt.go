package constants

const (
	// TODO: Store this in an .env file.
	signingKey        string = ""
	Issuer            string = "hadhri_backend"
	Audience          string = "hadhri_frontend"
	ExpirationMinutes int    = 15
)

var Jwtkey []byte = []byte(signingKey)

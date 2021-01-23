package config

// GetURLBase method to expose API Internal URL
func GetURLBase() string {
	return "/"
}

// IsLocalScope return true if environment is locally, false otherwise
func IsLocalScope() bool {
	return true
}

// GetCallerScopes return caller-scopes header value, for consume Insurance Purchase api
func GetCallerScopes() string {
	return ""
}

func GetRoutePrefix() string {
	return ""
}

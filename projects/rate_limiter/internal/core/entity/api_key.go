package core_entity

type ApiKey struct {
	ApiKey      string
	RateLimiter *RateLimiter
}

func NewApiKey(apiKey string, rateLimiter *RateLimiter) *ApiKey {
	return &ApiKey{
		ApiKey:      apiKey,
		RateLimiter: rateLimiter,
	}
}

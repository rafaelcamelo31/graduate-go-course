package core_entity

type Ip struct {
	Ip          string
	RateLimiter *RateLimiter
}

func NewIp(ip string, rateLimiter *RateLimiter) *Ip {
	return &Ip{
		Ip:          ip,
		RateLimiter: rateLimiter,
	}
}

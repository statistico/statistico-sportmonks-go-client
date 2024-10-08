package sportmonks

type ResponseDetails struct {
	Pagination   *Pagination    `json:"pagination"`
	Subscription []Subscription `json:"subscription"`
	RateLimit    RateLimit      `json:"rate_limit"`
	TimeZone     string         `json:"timezone"`
}

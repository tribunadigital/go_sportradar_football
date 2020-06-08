package go_sportradar_football

type HttpLog struct {
	Request    string
	Response   string
	CodeStatus string

	QuotaAllotted string
	QuotaCurrent  string
	QpsAllotted   string
	QpsCurrent    string
}

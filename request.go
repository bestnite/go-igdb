package igdb

import (
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"
)

var client *resty.Client

func init() {
	client = resty.New()
	client.SetRetryCount(3).SetRetryWaitTime(3 * time.Second).AddRetryCondition(
		func(r *resty.Response, err error) bool {
			return err != nil || r.StatusCode() == http.StatusTooManyRequests
		},
	)
}

func request() *resty.Request {
	return client.R().SetLogger(disableLogger{}).SetHeader("Accept-Charset", "utf-8").SetHeader("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:133.0) Gecko/20100101 Firefox/133.0")
}

type disableLogger struct{}

func (d disableLogger) Errorf(string, ...interface{}) {}
func (d disableLogger) Warnf(string, ...interface{})  {}
func (d disableLogger) Debugf(string, ...interface{}) {}

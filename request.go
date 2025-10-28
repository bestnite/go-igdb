package igdb

import (
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"
)

type SilentLogger struct{}

func (s SilentLogger) Errorf(string, ...any) {}
func (s SilentLogger) Warnf(string, ...any)  {}
func (s SilentLogger) Debugf(string, ...any) {}

func NewRestyClient() *resty.Client {
	return resty.New().SetRetryCount(10).SetRetryWaitTime(3 * time.Second).AddRetryCondition(
		func(r *resty.Response, err error) bool {
			if err != nil || r.StatusCode() == http.StatusTooManyRequests {
				return true
			}
			return false
		},
	)
}

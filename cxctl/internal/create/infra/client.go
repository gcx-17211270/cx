package infra

import (
	"crypto/tls"
	"net/http"
	"time"
)

func NewClient() *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
		Timeout: 1 * time.Minute, // 设置1分钟的超时时间
	}
}

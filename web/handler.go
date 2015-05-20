package web

import (
	"net/http"

	"net"

	"github.com/swanwish/go-helper/logs"
)

func MakeLogEnabledHandler(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		logRequest(r)
		fn(rw, r)
	}
}

func logRequest(r *http.Request) {
	// get client ip address
	ip, _, _ := net.SplitHostPort(r.RemoteAddr)

	logs.Debugf("Access url %v from ip: %s", r.URL, ip)
}

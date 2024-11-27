package rest

import (
	"net/http"
	"strings"
)

func allowedMethods(method ...string) func(h http.Handler) http.Handler {
	allowedMethods := make(map[string]struct{}, len(method))
	for _, m := range method {
		allowedMethods[strings.ToUpper(strings.TrimSpace(m))] = struct{}{}
	}
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if _, ok := allowedMethods[r.Method]; ok {
				next.ServeHTTP(w, r)
				return
			}
			w.WriteHeader(http.StatusMethodNotAllowed)
		})
	}
}
//copy past from chi middleware but chi exec this if content-length is more than 0, but I want exec it every request 
func allowedContentType(contentTypes ...string) func(h http.Handler) http.Handler {
	allowedContentTypes := make(map[string]struct{}, len(contentTypes))
	for _, ctype := range contentTypes {
		allowedContentTypes[strings.TrimSpace(strings.ToLower(ctype))] = struct{}{}
	}
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			s := strings.ToLower(strings.TrimSpace(r.Header.Get("Content-Type")))
			if i := strings.Index(s, ";"); i > -1 {
				s = s[0:i]
			}

			if _, ok := allowedContentTypes[s]; ok {
				next.ServeHTTP(w, r)
				return
			}

			w.WriteHeader(http.StatusUnsupportedMediaType)
		}
		return http.HandlerFunc(fn)
	}
}

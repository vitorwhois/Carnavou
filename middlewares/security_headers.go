package middlewares

import (
	"net/http"
)

func SetSecurityHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Security-Policy", `
		default-src 'self';
		script-src 'self' https://ajax.googleapis.com https://cdn.jsdelivr.net https://www.googletagmanager.com https://maps.googleapis.com;
		style-src 'self' https://fonts.googleapis.com https://cdn.jsdelivr.net;
		font-src 'self' https://fonts.gstatic.com;
		img-src 'self' data: https://www.google.com https://maps.gstatic.com https://www.googletagmanager.com;
		frame-src 'self' https://www.google.com https://maps.google.com;
	`)
		w.Header().Set("Strict-Transport-Security", "max-age=31536000; includeSubDomains")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("X-Frame-Options", "DENY")
		w.Header().Set("X-XSS-Protection", "1; mode=block")
		next.ServeHTTP(w, r)
	})
}

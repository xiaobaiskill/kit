package middleware

import "net/http"

type MiddleWareHeadle func(handler http.Handler) http.Handler

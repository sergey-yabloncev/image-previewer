package downloader

import "net/http"

func checkExtension(r *http.Response, contentType string) bool {
	return r.Header.Get("Content-Type") == contentType
}

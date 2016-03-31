package http

import "net/http"

// CloneRequest will clone the provided request and its derreferenced pointers
func CloneRequest(req *http.Request) *http.Request {
  clone := *req

  // if clone.URL != nil && clone.URL.User != nil {
  // user := *req.URL.User
  // clone.URL.User = &user
  // }

  return &clone
}

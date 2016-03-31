package http

import (
  "testing"
  "reflect"
  "net/http"
  "net/url"
)

func TestCloneRequestReturnsAnotherRequest(t *testing.T) {
  req := &http.Request{}
  clonedReq := CloneRequest(req)

  if req == clonedReq {
    t.Error("Original and cloned point to same memory address")
  }
}

func TestCloneRequestIsDeeplyEqual(t *testing.T) {
  url, _ := url.Parse("https://github.com/golang")

  req := &http.Request{
    URL: url,
    Proto: "HTTP/1.0",
    ProtoMajor: 1,
    ProtoMinor: 0,
    Header: map[string][]string{
      "Accept-Encoding": {"gzip, deflate"}, "Accept-Language": {"en-us"}, "Foo": {"Bar", "two"},
    },
    Host: "github.com",
  }
  clonedReq := CloneRequest(req)

  if false == reflect.DeepEqual(req, clonedReq) {
    t.Error("Request clone is not deeply equal")
  }
}

func TestCloneRequestURLIsDerreferenced(t *testing.T) {
  url, _ := url.Parse("https://user:password@github.com/golang")
  req := &http.Request{URL: url}

  clonedReq := CloneRequest(req)

  t.Error(&req.URL, &clonedReq.URL)

  switch {
  case &req.URL == &clonedReq.URL:
      t.Error("Original and cloned URL point to same memory address")

    // case &req.URL.User == &clonedReq.URL.User:
    //   t.Error("Original and cloned URL.User point to same memory address")
  }
}

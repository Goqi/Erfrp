// Copyright 2017 fatedier, fatedier@gmail.com
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package net

import (
	"compress/gzip"
	"io"
	"net/http"
	"strings"
)

type HTTPAuthWraper struct {
	h      http.Handler
	user   string
	passwd string
}

func NewHTTPBasicAuthWraper(h http.Handler, user, passwd string) http.Handler {
	return &HTTPAuthWraper{
		h:      h,
		user:   user,
		passwd: passwd,
	}
}

func (aw *HTTPAuthWraper) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user, passwd, hasAuth := r.BasicAuth()
	if (aw.user == "" && aw.passwd == "") || (hasAuth && user == aw.user && passwd == aw.passwd) {
		aw.h.ServeHTTP(w, r)
	} else {
		w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
	}
}

type HTTPAuthMiddleware struct {
	user   string
	passwd string
}

func NewHTTPAuthMiddleware(user, passwd string) *HTTPAuthMiddleware {
	return &HTTPAuthMiddleware{
		user:   user,
		passwd: passwd,
	}
}

func (authMid *HTTPAuthMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqUser, reqPasswd, hasAuth := r.BasicAuth()
		if (authMid.user == "" && authMid.passwd == "") ||
			(hasAuth && reqUser == authMid.user && reqPasswd == authMid.passwd) {
			next.ServeHTTP(w, r)
		} else {
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		}
	})
}

func HTTPBasicAuth(h http.HandlerFunc, user, passwd string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqUser, reqPasswd, hasAuth := r.BasicAuth()
		if (user == "" && passwd == "") ||
			(hasAuth && reqUser == user && reqPasswd == passwd) {
			h.ServeHTTP(w, r)
		} else {
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		}
	}
}

type HTTPGzipWraper struct {
	h http.Handler
}

func (gw *HTTPGzipWraper) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
		gw.h.ServeHTTP(w, r)
		return
	}
	w.Header().Set("Content-Encoding", "gzip")
	gz := gzip.NewWriter(w)
	defer gz.Close()
	gzr := gzipResponseWriter{Writer: gz, ResponseWriter: w}
	gw.h.ServeHTTP(gzr, r)
}

func MakeHTTPGzipHandler(h http.Handler) http.Handler {
	return &HTTPGzipWraper{
		h: h,
	}
}

type gzipResponseWriter struct {
	io.Writer
	http.ResponseWriter
}

func (w gzipResponseWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}

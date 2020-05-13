// Copyright {{Year}} {{Author}}. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package {{RepoName}}

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"
)

func testRequest(t *testing.T, req *http.Request, middleware func(h http.Handler) http.Handler) {
	w := httptest.NewRecorder()

	r := chi.NewRouter()
	r.Use(middleware)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		// TODO: Do something here

		w.Write([]byte("Hello World"))
	})
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "Response status code")

	// TODO: Return some test data
}

func Test{{MiddlewareFuncName}}(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)

	testRequest(t, req, {{MiddlewareFuncName}})

	// TODO: Check retuned data
}

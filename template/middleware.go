// Copyright {{Year}} {{Author}}. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package {{RepoName}}

import (
	"net/http"
)

func {{MiddlewareFuncName}}(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		// TODO: Implement

		h.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

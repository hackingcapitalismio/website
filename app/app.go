// Copyright © 2021 Kris Nóva <kris@nivenly.com>
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

package app

import (
	"net/http"
	"os"
	"sync"

	"github.com/kris-nova/logger"

	"github.com/kris-nova/anchovies"
	"github.com/kris-nova/bjorno"
)

// Values are the values we can access with text/template
// anywhere within the static website.
type Values struct {
	// Add exported fields here
}

type HackingCapitalism struct {
	sync.Mutex
}

func (h HackingCapitalism) Values(request *http.Request) interface{} {
	var v Values
	h.Lock()
	defer h.Unlock()
	return &v
}

func (h HackingCapitalism) Refresh() {
	//
}

func ListenAndServe() error {

	hc := &HackingCapitalism{}

	// Persistent storage
	anchovies.SetDir("/data/anchovies")

	// The bjorno config
	cfg := &bjorno.ServerConfig{

		// InterpolateExtensions are the file types that
		// bjorno will interpolate at runtime.
		InterpolateExtensions: []string{
			".html",
		},

		// BindAddress is the bind address and port following
		// the host:port string convention in Go.
		BindAddress: ":1315",

		// ServerDirectory looks weird but it's fine (I promise).
		//
		// We want to be able to "go run" the application for quick
		// development, and we also want to not have to change this
		// in the container image. The dockerfile takes care of
		// dropping off the binary where it needs to go.
		//
		// Always run from the "cmd" directory.
		ServeDirectory: "public",

		// DefaultIndexFiles index.html
		//
		// We started with Hugo for the website, so these are the
		// names of the files to look for in a directory.
		DefaultIndexFiles: []string{
			"index.html",
		},

		// Have fun Nóva you can now put shit online.
		Endpoints: []*bjorno.Endpoint{

			// [/empty] is a just a sample endpoint for us
			//          to copy because we are lazy.
			//{
			//	Pattern: "/empty",
			//	Handler: &EmptyHandler{},
			//},
		},
	}

	// 404 handling
	bytes, err := os.ReadFile("public/404.html")
	if err != nil {
		logger.Warning("Unable to load custom 404 path: %v", err)
		cfg.Content404 = []byte(bjorno.StatusDefault404)
	} else {
		cfg.Content404 = bytes
	}

	return bjorno.Runtime(cfg, hc)
}

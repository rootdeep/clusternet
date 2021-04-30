/*
Copyright 2021 The Clusternet Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package exchanger

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"sync"

	"github.com/gorilla/mux"
	"github.com/rancher/remotedialer"
	"github.com/sirupsen/logrus"
	"k8s.io/apimachinery/pkg/util/proxy"
	"k8s.io/apiserver/pkg/registry/rest"
	"k8s.io/klog/v2"

	proxies "github.com/clusternet/clusternet/pkg/apis/proxies/v1alpha1"
)

type Exchanger struct {
	// use cachedTransports to speed up networking connections
	cachedTransports map[string]*http.Transport

	lock sync.Mutex

	// dialerServer is used for serving websocket connection
	dialerServer *remotedialer.Server
}

var (
	urlPrefix = fmt.Sprintf("/apis/%s/sockets/", proxies.SchemeGroupVersion.String())
)

func NewExchanger(tunnelLogging bool) *Exchanger {
	if tunnelLogging {
		logrus.SetLevel(logrus.DebugLevel)
		remotedialer.PrintTunnelData = true
	}

	e := &Exchanger{
		cachedTransports: map[string]*http.Transport{},
	}
	e.dialerServer = remotedialer.New(e.authorizer, remotedialer.DefaultErrorWriter)
	return e
}

func (e *Exchanger) authorizer(req *http.Request) (string, bool, error) {
	vars := mux.Vars(req)
	clusterID := vars["cluster"]
	return clusterID, clusterID != "", nil
}

func (e *Exchanger) getTransport(server *remotedialer.Server, clusterID string) *http.Transport {
	e.lock.Lock()
	defer e.lock.Unlock()

	transport := e.cachedTransports[clusterID]
	if transport != nil {
		return transport
	}

	dialer := server.Dialer(clusterID)
	transport = &http.Transport{
		DialContext: dialer,
	}

	e.cachedTransports[clusterID] = transport
	return transport
}

func (e *Exchanger) Connect(ctx context.Context, id string, opts *proxies.Socket, responder rest.Responder) (http.Handler, error) {
	router := mux.NewRouter().PathPrefix(fmt.Sprintf("%s{cluster}", urlPrefix)).Subrouter()
	// serve websocket connection
	router.Handle("", e.dialerServer)
	router.Handle("/", e.dialerServer)

	// proxy and re-dial through websocket connection
	router.HandleFunc("/{scheme}/{host}{path:.*}",
		func(writer http.ResponseWriter, request *http.Request) {
			vars := mux.Vars(request)
			location := &url.URL{
				Scheme:   vars["scheme"],
				Host:     vars["host"],
				Path:     vars["path"],
				RawQuery: request.URL.RawQuery,
			}
			clusterID := vars["cluster"]

			klog.V(4).Infof("Request to %q will be redialed from cluster %q", location.String(), clusterID)

			// TODO: add metrics

			// if location.Path is empty, a status code 301 will be returned by below handler with a new location
			// ends with a '/'. This is essentially a hack for http://issue.k8s.io/4958.
			handler := proxy.NewUpgradeAwareHandler(location,
				e.getTransport(e.dialerServer, clusterID),
				false,
				false,
				proxy.NewErrorResponder(responder))
			handler.UseLocationHost = true
			handler.ServeHTTP(writer, request)
		})
	return router, nil
}
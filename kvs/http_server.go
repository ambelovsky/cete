// Copyright (c) 2019 Minoru Osuka
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 		http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package kvs

import (
	"log"
	"net"
	"net/http"

	"github.com/gorilla/mux"
	accesslog "github.com/mash/go-accesslog"
	cetehttp "github.com/mosuka/cete/http"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type HTTPServer struct {
	listener net.Listener
	router   *mux.Router

	grpcClient *GRPCClient

	logger     *log.Logger
	httpLogger *log.Logger
}

func NewHTTPServer(httpAddr string, grpcClient *GRPCClient, logger *log.Logger, httpLogger *log.Logger) (*HTTPServer, error) {
	listener, err := net.Listen("tcp", httpAddr)
	if err != nil {
		return nil, err
	}

	router := mux.NewRouter()
	router.StrictSlash(true)

	router.Handle("/", NewRootHandler(logger)).Methods("GET")
	router.Handle("/store/{path:.*}", NewPutHandler(grpcClient, logger)).Methods("PUT")
	router.Handle("/store/{path:.*}", NewGetHandler(grpcClient, logger)).Methods("GET")
	router.Handle("/store/{path:.*}", NewDeleteHandler(grpcClient, logger)).Methods("DELETE")
	router.Handle("/metrics", promhttp.Handler()).Methods("GET")

	return &HTTPServer{
		listener:   listener,
		router:     router,
		grpcClient: grpcClient,
		logger:     logger,
		httpLogger: httpLogger,
	}, nil
}

func (s *HTTPServer) Start() error {
	err := http.Serve(
		s.listener,
		accesslog.NewLoggingHandler(
			s.router,
			cetehttp.ApacheCombinedLogger{
				Logger: s.httpLogger,
			},
		),
	)
	if err != nil {
		return err
	}

	return nil
}

func (s *HTTPServer) Stop() error {
	err := s.listener.Close()
	if err != nil {
		return err
	}

	return nil
}

// Copyright 2020 Envoyproxy Authors
//
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.

// Package sotw provides an implementation of GRPC SoTW (State of The World) part of XDS server
package sotw

import (
	"context"
	"errors"
	"strconv"

	discovery "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v3"
	"github.com/envoyproxy/go-control-plane/pkg/cache/v3"
	"github.com/envoyproxy/go-control-plane/pkg/server/stream/v3"
)

type Server interface {
	StreamHandler(stream stream.Stream, typeURL string) error
}

type Callbacks interface {
	// OnStreamOpen is called once an xDS stream is open with a stream ID and the type URL (or "" for ADS).
	// Returning an error will end processing and close the stream. OnStreamClosed will still be called.
	OnStreamOpen(context.Context, int64, string) error
	// OnStreamClosed is called immediately prior to closing an xDS stream with a stream ID.
	OnStreamClosed(int64)
	// OnStreamRequest is called once a request is received on a stream.
	// Returning an error will end processing and close the stream. OnStreamClosed will still be called.
	OnStreamRequest(int64, *discovery.DiscoveryRequest) error
	// OnStreamResponse is called immediately prior to sending a response on a stream.
	OnStreamResponse(context.Context, int64, *discovery.DiscoveryRequest, *discovery.DiscoveryResponse)
}

// NewServer creates handlers from a config watcher and callbacks.
func NewServer(ctx context.Context, config cache.ConfigWatcher, callbacks Callbacks) Server {
	return &server{cache: config, callbacks: callbacks, ctx: ctx}
}

type server struct {
	cache     cache.ConfigWatcher
	callbacks Callbacks
	ctx       context.Context

	// streamCount for counting bi-di streams
	streamCount int64
}

type streamWrapper struct {
	stream      stream.Stream // parent stream object
	ID          int64         // stream ID in relation to total stream count
	nonce       int64         // nonce per stream
	watches     watches
	callbacks   Callbacks
	streamState stream.StreamState
}

func (s *streamWrapper) send(resp cache.Response) (string, error) {
	if resp == nil {
		return "", errors.New("missing response")
	}

	out, err := resp.GetDiscoveryResponse()
	if err != nil {
		return "", err
	}

	// increment nonce
	s.nonce = s.nonce + 1
	out.Nonce = strconv.FormatInt(s.nonce, 10)

	lastResponse := lastDiscoveryResponse{
		nonce:     out.Nonce,
		resources: make(map[string]struct{}),
	}
	for _, r := range resp.GetRequest().ResourceNames {
		lastResponse.resources[r] = struct{}{}
	}
	// lastDiscoveryResponses[resp.GetRequest().TypeUrl] = lastResponse

	if s.callbacks != nil {
		s.callbacks.OnStreamResponse(resp.GetContext(), s.ID, resp.GetRequest(), out)
	}
	return out.Nonce, s.stream.Send(out)
}

// Discovery response that is sent over GRPC stream
// We need to record what resource names are already sent to a client
// So if the client requests a new name we can respond back
// regardless current snapshot version (even if it is not changed yet)
type lastDiscoveryResponse struct {
	nonce     string
	resources map[string]struct{}
}

// StreamHandler converts a blocking read call to channels and initiates stream processing
func (s *server) StreamHandler(stream stream.Stream, typeURL string) error {
	// a channel for receiving incoming requests
	reqCh := make(chan *discovery.DiscoveryRequest)
	go func() {
		defer close(reqCh)
		for {
			req, err := stream.Recv()
			if err != nil {
				return
			}
			select {
			case reqCh <- req:
			case <-stream.Context().Done():
				return
			case <-s.ctx.Done():
				return
			}
		}
	}()

	return s.process(stream, reqCh, typeURL)
}

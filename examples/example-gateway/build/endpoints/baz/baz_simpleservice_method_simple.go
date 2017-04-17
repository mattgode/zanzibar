// Code generated by zanzibar
// @generated

package baz

import (
	"context"
	"net/http"

	"github.com/uber/zanzibar/examples/example-gateway/build/clients"
	zanzibar "github.com/uber/zanzibar/runtime"
	"go.uber.org/zap"
)

// HandleSimpleRequest handles "/baz/simple-path".
func HandleSimpleRequest(
	ctx context.Context,
	req *zanzibar.ServerHTTPRequest,
	res *zanzibar.ServerHTTPResponse,
	clients *clients.Clients,
) {

	workflow := SimpleEndpoint{
		Clients: clients,
		Logger:  req.Logger,
		Request: req,
	}

	respHeaders, err := workflow.Handle(ctx, req.Header)
	if err != nil {
		req.Logger.Warn("Workflow for endpoint returned error",
			zap.String("error", err.Error()),
		)
		res.SendErrorString(500, "Unexpected server error")
		return
	}

	res.WriteJSONBytes(204, respHeaders, nil)
}

// SimpleEndpoint calls thrift client Baz.Simple
type SimpleEndpoint struct {
	Clients *clients.Clients
	Logger  *zap.Logger
	Request *zanzibar.ServerHTTPRequest
}

// Handle calls thrift client.
func (w SimpleEndpoint) Handle(
	ctx context.Context,
	// TODO(sindelar): Switch to zanzibar.Headers when tchannel
	// generation is implemented.
	headers http.Header,
) (map[string]string, error) {

	clientHeaders := map[string]string{}
	for k, v := range map[string]string{} {
		clientHeaders[v] = headers.Get(k)
	}

	_, err := w.Clients.Baz.Simple(ctx, clientHeaders)

	if err != nil {
		w.Logger.Warn("Could not make client request",
			zap.String("error", err.Error()),
		)
		return nil, err
	}

	// Filter and map response headers from client to server response.
	endRespHead := map[string]string{}

	return endRespHead, nil
}

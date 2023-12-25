package ChainResponsability

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAuthenticator_Handle(t *testing.T) {
	authenticator := &Authenticator{}
	result := authenticator.Handle(nil)
	assert.Equal(t, 500, result.StatusCode)
}

func TestHTTPRequest_Handle(t *testing.T) {
	// Create a new HTTP request
	httpReq := &HTTPRequest{Url: "/api/data", Method: "GET", Params: "user=test&pass=1234"}
	authenticator := &Authenticator{}

	// Start the chain
	result := authenticator.Handle(httpReq)
	assert.Equal(t, 200, result.StatusCode)

	result = authenticator.Handle(&HTTPRequest{Url: "/api/data", Method: "GET", Params: "user=notvalid&pass=1234"})
	assert.Equal(t, 401, result.StatusCode)
}

func TestFullChain(t *testing.T) {
	authenticator := &Authenticator{}
	logger := &Logger{}
	sanitizer := &Sanitizer{}
	businessHandler := &BusinessHandler{}

	authenticator.SetNext(logger).SetNext(sanitizer).SetNext(businessHandler)

	httpReq := &HTTPRequest{Url: "/api/data", Method: "GET", Params: "user=test&pass=1234"}

	result := authenticator.Handle(httpReq)
	assert.NotNil(t, result)
	assert.Equal(t, 200, result.StatusCode)
}

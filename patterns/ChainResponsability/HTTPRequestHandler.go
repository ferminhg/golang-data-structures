package ChainResponsability

import (
	"fmt"
	"strings"
)

// Handler interface declares a method for processing requests.
type Handler interface {
	SetNext(handler Handler) Handler
	Handle(request *HTTPRequest) *HandlerResult
}

// BaseHandler provides a skeletal Handler implementation.
type BaseHandler struct {
	next Handler
}

// HandlerResult contains information about the result of a handler's operation.
type HandlerResult struct {
	StatusCode int
	Message    string
}

// SetNext sets the next handler in the chain and returns it.
func (b *BaseHandler) SetNext(handler Handler) Handler {
	b.next = handler
	return handler
}

func (b *BaseHandler) Handle(request *HTTPRequest) *HandlerResult {
	if b.next != nil {
		// Get the result from the next handler in the chain
		result := b.next.Handle(request)
		// Check if result is not nil before accessing its properties
		if result != nil && result.StatusCode != 200 {
			fmt.Println("An error occurred:", result.Message)
			return result
		}
	}
	// If no next handler, or next handler returned success (which should not happen in BaseHandler), return a default success result.
	return &HandlerResult{StatusCode: 200, Message: "Request processed successfully by BaseHandler"}
}

// HTTPRequest represents an incoming request.
type HTTPRequest struct {
	Url    string
	Method string
	Params string
}

// Authenticator checks if the request has valid credentials.
type Authenticator struct {
	BaseHandler
}

func (a *Authenticator) Handle(request *HTTPRequest) *HandlerResult {
	fmt.Println("Authenticator checking request")
	// Here should be the authentication logic.
	if request == nil {
		return &HandlerResult{StatusCode: 500, Message: "Request is nil"}
	}
	if !strings.Contains(request.Params, "user=test") ||
		!strings.Contains(request.Params, "pass=1234") {
		return &HandlerResult{StatusCode: 401, Message: "Unauthorized"}
	}
	// If authenticated, pass the request along the chain.
	return a.BaseHandler.Handle(request)
}

// Logger logs request details.
type Logger struct {
	BaseHandler
}

func (l *Logger) Handle(request *HTTPRequest) *HandlerResult {
	fmt.Println("Logger logging request:", request)
	// Perform logging here.
	return l.BaseHandler.Handle(request)
}

// Sanitizer cleans the request parameters to prevent attacks.
type Sanitizer struct {
	BaseHandler
}

func (s *Sanitizer) Handle(request *HTTPRequest) *HandlerResult {
	fmt.Println("Sanitizer sanitizing request")
	// Sanitize request.Params...
	return s.BaseHandler.Handle(request)
}

// BusinessHandler processes the business logic of the request.
type BusinessHandler struct {
	BaseHandler
}

func (b *BusinessHandler) Handle(_ *HTTPRequest) *HandlerResult {
	fmt.Println("BusinessHandler processing request")
	// Perform business logic.
	// If the business logic is successful, return a result.
	return &HandlerResult{
		StatusCode: 200,
		Message:    "Request processed successfully",
	}
}

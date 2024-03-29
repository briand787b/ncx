package cslogtest

import (
	"context"
	"io"
	"sync"

	"github.com/briand787b/ncx/src/svc/pkg/cslog"
)

var _ cslog.Logger = &MockLogger{}

// MockLogger is a mocked implementation of plog.Logger
type MockLogger struct {
	mu sync.Mutex

	CloseCallCount int
	CloseArgCloser []io.Closer

	ErrorCallCount int
	ErrorArgMsg    []string
	ErrorArgArgs   [][]interface{}

	InfoCallCount int
	InfoArgMsg    []string
	InfoArgArgs   [][]interface{}

	InvalidCallCount int
	InvalidArgSub    []interface{}
	InvalidArgReason []string

	QueryCallCount int
	QueryArgQry    []string
	QueryArgArgs   [][]interface{}
}

// NewMockLogger instantiates and returns a new MockLogger
func NewMockLogger() *MockLogger {
	return &MockLogger{
		mu: sync.Mutex{},
	}
}

// Close x
func (m *MockLogger) Close(ctx context.Context, c io.Closer) {
	m.mu.Lock()
	defer m.mu.Unlock()
	defer func() { m.CloseCallCount++ }()
	m.CloseArgCloser = append(m.CloseArgCloser, c)
}

// Error x
func (m *MockLogger) Error(ctx context.Context, msg string, args ...interface{}) {
	m.mu.Lock()
	defer m.mu.Unlock()
	defer func() { m.ErrorCallCount++ }()
	m.ErrorArgMsg = append(m.ErrorArgMsg, msg)
	m.ErrorArgArgs = append(m.ErrorArgArgs, args)
}

// Info x
func (m *MockLogger) Info(ctx context.Context, msg string, args ...interface{}) {
	m.mu.Lock()
	defer m.mu.Unlock()
	defer func() { m.InfoCallCount++ }()
	m.InfoArgMsg = append(m.InfoArgMsg, msg)
	m.InfoArgArgs = append(m.InfoArgArgs, args)
}

// Invalid x
func (m *MockLogger) Invalid(ctx context.Context, subj interface{}, reason string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	defer func() { m.InvalidCallCount++ }()
	m.InvalidArgSub = append(m.InvalidArgSub, subj)
	m.InvalidArgReason = append(m.InvalidArgReason, reason)
}

// Query x
func (m *MockLogger) Query(ctx context.Context, qry string, args ...interface{}) {
	m.mu.Lock()
	defer m.mu.Unlock()
	defer func() { m.QueryCallCount++ }()
	m.QueryArgQry = append(m.QueryArgQry, qry)
	m.QueryArgArgs = append(m.QueryArgArgs, args)
}

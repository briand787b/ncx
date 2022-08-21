package cserr

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/briand787b/ncx/src/svc/pkg/cslog"

	"github.com/pkg/errors"
)

// Error backs all custom error types of perr.
// It allows direct comparisons of error values
type Error string

// Error allows Error to satisfy the error interface
func (e Error) Error() string { return string(e) }

const (
	// ErrInvalid is when validation of the resource failed
	ErrInvalid = Error("request invalid")

	// ErrNotFound is when the request resource does not exist
	ErrNotFound = Error("resource could not be found")

	// ErrUnauthorized is when the requestor is unauthorized to perform
	// the requested action
	ErrUnauthorized = Error("authorization failed")

	// ErrInternal is when an error results from internal software failures
	ErrInternal = Error("internal server error")
)

// SameType determines whether the given error is the same type as thee
// known Error provided
func SameType(e error, known Error) bool {
	err, ok := errors.Cause(e).(Error)
	if !ok {
		return false
	}

	return err == known
}

// IsInternalServerError returns boolean indicating whether error is caused by the
// system itself.  True means that the error is either known to have been caused
// by a fault in the system or is of unknown type
func IsInternalServerError(ctx context.Context, l cslog.Logger, e error) bool {
	if e == nil {
		l.Error(ctx, "nil error passed to 'GetExternalMsg'")
		return false
	}

	switch c := errors.Cause(e); {
	case c == ErrInvalid,
		c == ErrNotFound,
		c == ErrUnauthorized:
		return false
	}

	return true
}

// GetExternalMsg extracts the message for the error that is suitable
// for displaying externally
func GetExternalMsg(ctx context.Context, l cslog.Logger, e error) string {
	if e == nil {
		l.Error(ctx, "nil error passed to 'GetExternalMsg'")
		return ""
	}

	switch c := errors.Cause(e); {
	case c == ErrNotFound:
		return "Resource Not Found"
	case c == ErrUnauthorized:
		return "Request Not Authorized to Perform Action"
	case c == ErrInvalid:
		if es := strings.Split(e.Error(), ":"); len(es) > 1 {
			return strings.TrimSpace(fmt.Sprintf("%s: %s", es[len(es)-1], es[len(es)-2]))
		}

		fallthrough
	default:
		return "Internal Server Error"
	}
}

// NewErrInvalid returns a wrapped ErrInvalid
func NewErrInvalid(reasonMsg string) error {
	err := errors.Wrap(ErrInvalid, reasonMsg)
	return err
}

// NewErrNotFound returns a wrapped ErrNotFound
func NewErrNotFound(e error) error {
	if e == nil {
		e = errors.New("WARNING: nil error provided to `NewErrNotFound()`")
	}

	return errors.Wrap(ErrNotFound, e.Error())
}

// NewErrInternal returns a wrapped ErrNewInternal
func NewErrInternal(e error) error {
	if e == nil {
		e = errors.New("WARNING: nil error provided to `NewErrInternal()`")
	}

	return errors.Wrap(ErrInternal, e.Error())
}

// NewErrInternalWithMsg returns a wrapped ErrNewInternal
func NewErrInternalWithMsg(e error, msg string) error {
	if e == nil {
		e = errors.New("WARNING: nil error provided to `NewErrInternal()`")
	}

	return errors.Wrap(ErrInternal, errors.Wrap(e, msg).Error())
}

// NewErrUnauthorized returns a wrapped ErrUnauthorized
func NewErrUnauthorized(e error) error {
	if e == nil {
		e = errors.New("WARNING: nil error provided to `NewErrInternal()`")
	}

	return errors.Wrap(ErrUnauthorized, e.Error())
}

// HandleHTTPResponse handles errors for HTTP clients.  Please note that it also
// reads and closes the body of the response if a non-nil error is returned.
func HandleHTTPResponse(ctx context.Context, l cslog.Logger, r *http.Response) error {
	baseErr := errors.Errorf("status code %d receieved", r.StatusCode)
	switch code := r.StatusCode; true {
	case code < 300:
		return nil
	case code == http.StatusBadRequest:
		baseErr = NewErrInvalid(baseErr.Error())
	case code == http.StatusNotFound:
		baseErr = NewErrNotFound(baseErr)
	case code == http.StatusForbidden, code == http.StatusUnauthorized:
		baseErr = NewErrUnauthorized(baseErr)
	default:
		baseErr = NewErrInternal(baseErr)
	}

	// failed requests (which are the only ones reaching this code) shoud be closed out here
	defer r.Body.Close()

	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return NewErrInternalWithMsg(err, baseErr.Error())
	}

	if len(bodyBytes) < 1 {
		return baseErr
	}

	// c.l.Error(ctx, "error getting user object", "error", string(body), "status_code", resp.StatusCode)
	return errors.Wrap(baseErr, string(bodyBytes))
}

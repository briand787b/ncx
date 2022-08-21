package cserr_test

import (
	"context"
	"database/sql"
	"testing"

	"github.com/briand787b/ncx/src/svc/pkg/cserr"
	"github.com/briand787b/ncx/src/svc/pkg/cslog/cslogtest"

	"github.com/pkg/errors"
)

func TestErrCause(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name         string
		err          error
		wracserrMsgs []string
		expCause     error
	}{
		{"1_layer_in_tact", sql.ErrNoRows, []string{"a"}, sql.ErrNoRows},
		{"2_layers_in_tact", sql.ErrNoRows, []string{"a", "b"}, sql.ErrNoRows},
		{"3_layers_in_tact", sql.ErrNoRows, []string{"a", "b", "c"}, sql.ErrNoRows},
		{"empty_val_in_tact", cserr.NewErrInvalid(""), []string{"a"}, cserr.ErrInvalid},
		{"full_val_in_tact", cserr.NewErrInvalid("z"), []string{"a"}, cserr.ErrInvalid},
		{"2_layer_full_val_in_tact", cserr.NewErrInvalid("z"), []string{"a", "b"}, cserr.ErrInvalid},
		{"empty_not_found_in_tact", cserr.NewErrNotFound(nil), []string{"a"}, cserr.ErrNotFound},
		{"full_not_found_in_tact", cserr.NewErrNotFound(sql.ErrNoRows), []string{"a"}, cserr.ErrNotFound},
		{"2_layer_not_found_in_tact", cserr.NewErrNotFound(sql.ErrNoRows), []string{"a", "b"}, cserr.ErrNotFound},
		{"empty_internal_in_tact", cserr.NewErrInternal(nil), []string{"a"}, cserr.ErrInternal},
		{"full_internal_in_tact", cserr.NewErrInternal(sql.ErrNoRows), []string{"a"}, cserr.ErrInternal},
		{"2_layer_internal_in_tact", cserr.NewErrInternal(sql.ErrNoRows), []string{"a", "b"}, cserr.ErrInternal},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var err error
			for _, msg := range tt.wracserrMsgs {
				err = errors.Wrap(tt.err, msg)
			}

			if retErr := errors.Cause(err); tt.expCause != retErr {
				t.Fatalf("expected error to be %v, was %v", tt.expCause, retErr)
			}
		})
	}
}

func TestGetExternalMgs(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		err    error
		expMsg string
	}{
		{"nil_err_returns_empty_msg", nil, ""},
		{"unknown_err_returns_internal_server_err", errors.New("a"), "Internal Server Error"},
		{"auth_err_returns_unauth_msg", cserr.ErrUnauthorized, "Request Not Authorized to Perform Action"},
		{"invalid_err_returns_internal_server_err", cserr.NewErrInvalid("b"), "request invalid: b"},
		{"uninit_invalid_err_returns_internal_server_err", cserr.ErrInvalid, "Internal Server Error"},
		{"not_found_err_returns_invalid_msg", cserr.ErrNotFound, "Resource Not Found"},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ctx := context.Background()

			if retMsg := cserr.GetExternalMsg(ctx, cslogtest.NewMockLogger(), tt.err); tt.expMsg != retMsg {
				t.Fatalf("expected msg to be %s, was %s", tt.expMsg, retMsg)
			}
		})
	}
}

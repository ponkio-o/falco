package function

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/ysugimoto/falco/interpreter"
	"github.com/ysugimoto/falco/interpreter/context"
	"github.com/ysugimoto/falco/interpreter/function/errors"
	"github.com/ysugimoto/falco/interpreter/value"
)

func Test_Assert_restart(t *testing.T) {

	tests := []struct {
		args   []value.Value
		ip     *interpreter.Interpreter
		err    error
		expect *value.Boolean
	}{
		{
			args: []value.Value{},
			ip: &interpreter.Interpreter{
				TestingState: interpreter.ERROR,
			},
			expect: &value.Boolean{Value: false},
			err:    &errors.AssertionError{},
		},
		{
			args: []value.Value{},
			ip: &interpreter.Interpreter{
				TestingState: interpreter.RESTART,
			},
			expect: &value.Boolean{Value: true},
		},
	}

	for i := range tests {
		_, err := Assert_restart(
			&context.Context{},
			tests[i].ip,
			tests[i].args...,
		)
		if diff := cmp.Diff(
			tests[i].err,
			err,
			cmpopts.IgnoreFields(errors.AssertionError{}, "Message", "Actual"),
			cmpopts.IgnoreFields(errors.TestingError{}, "Message"),
		); diff != "" {
			t.Errorf("Assert_restart()[%d] error: diff=%s", i, diff)
		}
	}
}

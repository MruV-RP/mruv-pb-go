// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: objects/movable.proto

package objects

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/golang/protobuf/ptypes"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = ptypes.DynamicAny{}
)

// define the regex for a UUID once up-front
var _movable_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on State with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *State) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Name

	// no validation rules for X

	// no validation rules for Y

	// no validation rules for Z

	// no validation rules for Rx

	// no validation rules for Ry

	// no validation rules for Rz

	// no validation rules for TransitionSpeed

	return nil
}

// StateValidationError is the validation error returned by State.Validate if
// the designated constraints aren't met.
type StateValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StateValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StateValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StateValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StateValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StateValidationError) ErrorName() string { return "StateValidationError" }

// Error satisfies the builtin error interface
func (e StateValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sState.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StateValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StateValidationError{}

// Validate checks the field values on MovableObject with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *MovableObject) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetObject()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return MovableObjectValidationError{
				field:  "Object",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetStates() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return MovableObjectValidationError{
					field:  fmt.Sprintf("States[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// MovableObjectValidationError is the validation error returned by
// MovableObject.Validate if the designated constraints aren't met.
type MovableObjectValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MovableObjectValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MovableObjectValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MovableObjectValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MovableObjectValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MovableObjectValidationError) ErrorName() string { return "MovableObjectValidationError" }

// Error satisfies the builtin error interface
func (e MovableObjectValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMovableObject.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MovableObjectValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MovableObjectValidationError{}

// Validate checks the field values on CreateMovableObjectRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateMovableObjectRequest) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetObject()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateMovableObjectRequestValidationError{
				field:  "Object",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetStates() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CreateMovableObjectRequestValidationError{
					field:  fmt.Sprintf("States[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// CreateMovableObjectRequestValidationError is the validation error returned
// by CreateMovableObjectRequest.Validate if the designated constraints aren't met.
type CreateMovableObjectRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateMovableObjectRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateMovableObjectRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateMovableObjectRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateMovableObjectRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateMovableObjectRequestValidationError) ErrorName() string {
	return "CreateMovableObjectRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateMovableObjectRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateMovableObjectRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateMovableObjectRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateMovableObjectRequestValidationError{}

// Validate checks the field values on CreateMovableObjectResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateMovableObjectResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	return nil
}

// CreateMovableObjectResponseValidationError is the validation error returned
// by CreateMovableObjectResponse.Validate if the designated constraints
// aren't met.
type CreateMovableObjectResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateMovableObjectResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateMovableObjectResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateMovableObjectResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateMovableObjectResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateMovableObjectResponseValidationError) ErrorName() string {
	return "CreateMovableObjectResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateMovableObjectResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateMovableObjectResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateMovableObjectResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateMovableObjectResponseValidationError{}

// Validate checks the field values on GetMovableObjectRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GetMovableObjectRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	return nil
}

// GetMovableObjectRequestValidationError is the validation error returned by
// GetMovableObjectRequest.Validate if the designated constraints aren't met.
type GetMovableObjectRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetMovableObjectRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetMovableObjectRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetMovableObjectRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetMovableObjectRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetMovableObjectRequestValidationError) ErrorName() string {
	return "GetMovableObjectRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetMovableObjectRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetMovableObjectRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetMovableObjectRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetMovableObjectRequestValidationError{}

// Validate checks the field values on GetMovableObjectResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GetMovableObjectResponse) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetObject()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetMovableObjectResponseValidationError{
				field:  "Object",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetStates() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetMovableObjectResponseValidationError{
					field:  fmt.Sprintf("States[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for CurrentStateId

	// no validation rules for CurrentStateName

	return nil
}

// GetMovableObjectResponseValidationError is the validation error returned by
// GetMovableObjectResponse.Validate if the designated constraints aren't met.
type GetMovableObjectResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetMovableObjectResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetMovableObjectResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetMovableObjectResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetMovableObjectResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetMovableObjectResponseValidationError) ErrorName() string {
	return "GetMovableObjectResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetMovableObjectResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetMovableObjectResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetMovableObjectResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetMovableObjectResponseValidationError{}

// Validate checks the field values on UpdateMovableObjectRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *UpdateMovableObjectRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for ObjectId

	for idx, item := range m.GetStates() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return UpdateMovableObjectRequestValidationError{
					field:  fmt.Sprintf("States[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// UpdateMovableObjectRequestValidationError is the validation error returned
// by UpdateMovableObjectRequest.Validate if the designated constraints aren't met.
type UpdateMovableObjectRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateMovableObjectRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateMovableObjectRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateMovableObjectRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateMovableObjectRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateMovableObjectRequestValidationError) ErrorName() string {
	return "UpdateMovableObjectRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateMovableObjectRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateMovableObjectRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateMovableObjectRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateMovableObjectRequestValidationError{}

// Validate checks the field values on UpdateMovableObjectResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *UpdateMovableObjectResponse) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// UpdateMovableObjectResponseValidationError is the validation error returned
// by UpdateMovableObjectResponse.Validate if the designated constraints
// aren't met.
type UpdateMovableObjectResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateMovableObjectResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateMovableObjectResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateMovableObjectResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateMovableObjectResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateMovableObjectResponseValidationError) ErrorName() string {
	return "UpdateMovableObjectResponseValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateMovableObjectResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateMovableObjectResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateMovableObjectResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateMovableObjectResponseValidationError{}

// Validate checks the field values on DeleteMovableObjectRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DeleteMovableObjectRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	return nil
}

// DeleteMovableObjectRequestValidationError is the validation error returned
// by DeleteMovableObjectRequest.Validate if the designated constraints aren't met.
type DeleteMovableObjectRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteMovableObjectRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteMovableObjectRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteMovableObjectRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteMovableObjectRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteMovableObjectRequestValidationError) ErrorName() string {
	return "DeleteMovableObjectRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteMovableObjectRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteMovableObjectRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteMovableObjectRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteMovableObjectRequestValidationError{}

// Validate checks the field values on DeleteMovableObjectResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DeleteMovableObjectResponse) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// DeleteMovableObjectResponseValidationError is the validation error returned
// by DeleteMovableObjectResponse.Validate if the designated constraints
// aren't met.
type DeleteMovableObjectResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteMovableObjectResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteMovableObjectResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteMovableObjectResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteMovableObjectResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteMovableObjectResponseValidationError) ErrorName() string {
	return "DeleteMovableObjectResponseValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteMovableObjectResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteMovableObjectResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteMovableObjectResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteMovableObjectResponseValidationError{}

// Validate checks the field values on MoveObjectRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *MoveObjectRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for State

	return nil
}

// MoveObjectRequestValidationError is the validation error returned by
// MoveObjectRequest.Validate if the designated constraints aren't met.
type MoveObjectRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MoveObjectRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MoveObjectRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MoveObjectRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MoveObjectRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MoveObjectRequestValidationError) ErrorName() string {
	return "MoveObjectRequestValidationError"
}

// Error satisfies the builtin error interface
func (e MoveObjectRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMoveObjectRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MoveObjectRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MoveObjectRequestValidationError{}

// Validate checks the field values on MoveObjectResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *MoveObjectResponse) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// MoveObjectResponseValidationError is the validation error returned by
// MoveObjectResponse.Validate if the designated constraints aren't met.
type MoveObjectResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MoveObjectResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MoveObjectResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MoveObjectResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MoveObjectResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MoveObjectResponseValidationError) ErrorName() string {
	return "MoveObjectResponseValidationError"
}

// Error satisfies the builtin error interface
func (e MoveObjectResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMoveObjectResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MoveObjectResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MoveObjectResponseValidationError{}

// Validate checks the field values on MoveObjectNextRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *MoveObjectNextRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	return nil
}

// MoveObjectNextRequestValidationError is the validation error returned by
// MoveObjectNextRequest.Validate if the designated constraints aren't met.
type MoveObjectNextRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MoveObjectNextRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MoveObjectNextRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MoveObjectNextRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MoveObjectNextRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MoveObjectNextRequestValidationError) ErrorName() string {
	return "MoveObjectNextRequestValidationError"
}

// Error satisfies the builtin error interface
func (e MoveObjectNextRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMoveObjectNextRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MoveObjectNextRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MoveObjectNextRequestValidationError{}

// Validate checks the field values on MoveObjectNextResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *MoveObjectNextResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for StateId

	// no validation rules for StateName

	return nil
}

// MoveObjectNextResponseValidationError is the validation error returned by
// MoveObjectNextResponse.Validate if the designated constraints aren't met.
type MoveObjectNextResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MoveObjectNextResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MoveObjectNextResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MoveObjectNextResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MoveObjectNextResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MoveObjectNextResponseValidationError) ErrorName() string {
	return "MoveObjectNextResponseValidationError"
}

// Error satisfies the builtin error interface
func (e MoveObjectNextResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMoveObjectNextResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MoveObjectNextResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MoveObjectNextResponseValidationError{}

// Validate checks the field values on MoveObjectPreviousRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *MoveObjectPreviousRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	return nil
}

// MoveObjectPreviousRequestValidationError is the validation error returned by
// MoveObjectPreviousRequest.Validate if the designated constraints aren't met.
type MoveObjectPreviousRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MoveObjectPreviousRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MoveObjectPreviousRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MoveObjectPreviousRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MoveObjectPreviousRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MoveObjectPreviousRequestValidationError) ErrorName() string {
	return "MoveObjectPreviousRequestValidationError"
}

// Error satisfies the builtin error interface
func (e MoveObjectPreviousRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMoveObjectPreviousRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MoveObjectPreviousRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MoveObjectPreviousRequestValidationError{}

// Validate checks the field values on MoveObjectPreviousResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *MoveObjectPreviousResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for StateId

	// no validation rules for StateName

	return nil
}

// MoveObjectPreviousResponseValidationError is the validation error returned
// by MoveObjectPreviousResponse.Validate if the designated constraints aren't met.
type MoveObjectPreviousResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MoveObjectPreviousResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MoveObjectPreviousResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MoveObjectPreviousResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MoveObjectPreviousResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MoveObjectPreviousResponseValidationError) ErrorName() string {
	return "MoveObjectPreviousResponseValidationError"
}

// Error satisfies the builtin error interface
func (e MoveObjectPreviousResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMoveObjectPreviousResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MoveObjectPreviousResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MoveObjectPreviousResponseValidationError{}

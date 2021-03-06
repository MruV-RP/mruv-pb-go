// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: elevators/elevators.proto

package elevators

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
var _elevators_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on CreateElevatorRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateElevatorRequest) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// CreateElevatorRequestValidationError is the validation error returned by
// CreateElevatorRequest.Validate if the designated constraints aren't met.
type CreateElevatorRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateElevatorRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateElevatorRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateElevatorRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateElevatorRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateElevatorRequestValidationError) ErrorName() string {
	return "CreateElevatorRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateElevatorRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateElevatorRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateElevatorRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateElevatorRequestValidationError{}

// Validate checks the field values on CreateElevatorResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateElevatorResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	return nil
}

// CreateElevatorResponseValidationError is the validation error returned by
// CreateElevatorResponse.Validate if the designated constraints aren't met.
type CreateElevatorResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateElevatorResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateElevatorResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateElevatorResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateElevatorResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateElevatorResponseValidationError) ErrorName() string {
	return "CreateElevatorResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateElevatorResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateElevatorResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateElevatorResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateElevatorResponseValidationError{}

// Validate checks the field values on GetElevatorRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GetElevatorRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	return nil
}

// GetElevatorRequestValidationError is the validation error returned by
// GetElevatorRequest.Validate if the designated constraints aren't met.
type GetElevatorRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetElevatorRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetElevatorRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetElevatorRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetElevatorRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetElevatorRequestValidationError) ErrorName() string {
	return "GetElevatorRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetElevatorRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetElevatorRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetElevatorRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetElevatorRequestValidationError{}

// Validate checks the field values on GetElevatorResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GetElevatorResponse) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// GetElevatorResponseValidationError is the validation error returned by
// GetElevatorResponse.Validate if the designated constraints aren't met.
type GetElevatorResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetElevatorResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetElevatorResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetElevatorResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetElevatorResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetElevatorResponseValidationError) ErrorName() string {
	return "GetElevatorResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetElevatorResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetElevatorResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetElevatorResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetElevatorResponseValidationError{}

// Validate checks the field values on UpdateElevatorRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *UpdateElevatorRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	return nil
}

// UpdateElevatorRequestValidationError is the validation error returned by
// UpdateElevatorRequest.Validate if the designated constraints aren't met.
type UpdateElevatorRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateElevatorRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateElevatorRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateElevatorRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateElevatorRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateElevatorRequestValidationError) ErrorName() string {
	return "UpdateElevatorRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateElevatorRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateElevatorRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateElevatorRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateElevatorRequestValidationError{}

// Validate checks the field values on UpdateElevatorResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *UpdateElevatorResponse) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// UpdateElevatorResponseValidationError is the validation error returned by
// UpdateElevatorResponse.Validate if the designated constraints aren't met.
type UpdateElevatorResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateElevatorResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateElevatorResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateElevatorResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateElevatorResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateElevatorResponseValidationError) ErrorName() string {
	return "UpdateElevatorResponseValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateElevatorResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateElevatorResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateElevatorResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateElevatorResponseValidationError{}

// Validate checks the field values on DeleteElevatorRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DeleteElevatorRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	return nil
}

// DeleteElevatorRequestValidationError is the validation error returned by
// DeleteElevatorRequest.Validate if the designated constraints aren't met.
type DeleteElevatorRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteElevatorRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteElevatorRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteElevatorRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteElevatorRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteElevatorRequestValidationError) ErrorName() string {
	return "DeleteElevatorRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteElevatorRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteElevatorRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteElevatorRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteElevatorRequestValidationError{}

// Validate checks the field values on DeleteElevatorResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DeleteElevatorResponse) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// DeleteElevatorResponseValidationError is the validation error returned by
// DeleteElevatorResponse.Validate if the designated constraints aren't met.
type DeleteElevatorResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteElevatorResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteElevatorResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteElevatorResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteElevatorResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteElevatorResponseValidationError) ErrorName() string {
	return "DeleteElevatorResponseValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteElevatorResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteElevatorResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteElevatorResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteElevatorResponseValidationError{}

// Validate checks the field values on GetElevatorFloorsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GetElevatorFloorsRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	return nil
}

// GetElevatorFloorsRequestValidationError is the validation error returned by
// GetElevatorFloorsRequest.Validate if the designated constraints aren't met.
type GetElevatorFloorsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetElevatorFloorsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetElevatorFloorsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetElevatorFloorsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetElevatorFloorsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetElevatorFloorsRequestValidationError) ErrorName() string {
	return "GetElevatorFloorsRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetElevatorFloorsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetElevatorFloorsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetElevatorFloorsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetElevatorFloorsRequestValidationError{}

// Validate checks the field values on GetElevatorFloorsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GetElevatorFloorsResponse) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// GetElevatorFloorsResponseValidationError is the validation error returned by
// GetElevatorFloorsResponse.Validate if the designated constraints aren't met.
type GetElevatorFloorsResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetElevatorFloorsResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetElevatorFloorsResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetElevatorFloorsResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetElevatorFloorsResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetElevatorFloorsResponseValidationError) ErrorName() string {
	return "GetElevatorFloorsResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetElevatorFloorsResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetElevatorFloorsResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetElevatorFloorsResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetElevatorFloorsResponseValidationError{}

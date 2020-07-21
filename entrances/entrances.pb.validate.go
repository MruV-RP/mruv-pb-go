// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: entrances/entrances.proto

package entrances

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
var _entrances_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on Entrance with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Entrance) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Name

	if v, ok := interface{}(m.GetOut()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return EntranceValidationError{
				field:  "Out",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetIn()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return EntranceValidationError{
				field:  "In",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// EntranceValidationError is the validation error returned by
// Entrance.Validate if the designated constraints aren't met.
type EntranceValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EntranceValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EntranceValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EntranceValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EntranceValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EntranceValidationError) ErrorName() string { return "EntranceValidationError" }

// Error satisfies the builtin error interface
func (e EntranceValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEntrance.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EntranceValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EntranceValidationError{}

// Validate checks the field values on CreateEntranceRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateEntranceRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Name

	if v, ok := interface{}(m.GetOut()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateEntranceRequestValidationError{
				field:  "Out",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetIn()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateEntranceRequestValidationError{
				field:  "In",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// CreateEntranceRequestValidationError is the validation error returned by
// CreateEntranceRequest.Validate if the designated constraints aren't met.
type CreateEntranceRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateEntranceRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateEntranceRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateEntranceRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateEntranceRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateEntranceRequestValidationError) ErrorName() string {
	return "CreateEntranceRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateEntranceRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateEntranceRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateEntranceRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateEntranceRequestValidationError{}

// Validate checks the field values on CreateEntranceResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateEntranceResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	return nil
}

// CreateEntranceResponseValidationError is the validation error returned by
// CreateEntranceResponse.Validate if the designated constraints aren't met.
type CreateEntranceResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateEntranceResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateEntranceResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateEntranceResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateEntranceResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateEntranceResponseValidationError) ErrorName() string {
	return "CreateEntranceResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateEntranceResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateEntranceResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateEntranceResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateEntranceResponseValidationError{}

// Validate checks the field values on GetEntranceRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GetEntranceRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	return nil
}

// GetEntranceRequestValidationError is the validation error returned by
// GetEntranceRequest.Validate if the designated constraints aren't met.
type GetEntranceRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetEntranceRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetEntranceRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetEntranceRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetEntranceRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetEntranceRequestValidationError) ErrorName() string {
	return "GetEntranceRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetEntranceRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetEntranceRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetEntranceRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetEntranceRequestValidationError{}

// Validate checks the field values on GetEntranceResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GetEntranceResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Name

	if v, ok := interface{}(m.GetInSpot()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetEntranceResponseValidationError{
				field:  "InSpot",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetOutSpot()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetEntranceResponseValidationError{
				field:  "OutSpot",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// GetEntranceResponseValidationError is the validation error returned by
// GetEntranceResponse.Validate if the designated constraints aren't met.
type GetEntranceResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetEntranceResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetEntranceResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetEntranceResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetEntranceResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetEntranceResponseValidationError) ErrorName() string {
	return "GetEntranceResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetEntranceResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetEntranceResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetEntranceResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetEntranceResponseValidationError{}

// Validate checks the field values on UpdateEntranceRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *UpdateEntranceRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for InSpotId

	// no validation rules for OutSpotId

	return nil
}

// UpdateEntranceRequestValidationError is the validation error returned by
// UpdateEntranceRequest.Validate if the designated constraints aren't met.
type UpdateEntranceRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateEntranceRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateEntranceRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateEntranceRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateEntranceRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateEntranceRequestValidationError) ErrorName() string {
	return "UpdateEntranceRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateEntranceRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateEntranceRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateEntranceRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateEntranceRequestValidationError{}

// Validate checks the field values on UpdateEntranceResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *UpdateEntranceResponse) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// UpdateEntranceResponseValidationError is the validation error returned by
// UpdateEntranceResponse.Validate if the designated constraints aren't met.
type UpdateEntranceResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateEntranceResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateEntranceResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateEntranceResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateEntranceResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateEntranceResponseValidationError) ErrorName() string {
	return "UpdateEntranceResponseValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateEntranceResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateEntranceResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateEntranceResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateEntranceResponseValidationError{}

// Validate checks the field values on DeleteEntranceRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DeleteEntranceRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	return nil
}

// DeleteEntranceRequestValidationError is the validation error returned by
// DeleteEntranceRequest.Validate if the designated constraints aren't met.
type DeleteEntranceRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteEntranceRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteEntranceRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteEntranceRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteEntranceRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteEntranceRequestValidationError) ErrorName() string {
	return "DeleteEntranceRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteEntranceRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteEntranceRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteEntranceRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteEntranceRequestValidationError{}

// Validate checks the field values on DeleteEntranceResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DeleteEntranceResponse) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// DeleteEntranceResponseValidationError is the validation error returned by
// DeleteEntranceResponse.Validate if the designated constraints aren't met.
type DeleteEntranceResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteEntranceResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteEntranceResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteEntranceResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteEntranceResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteEntranceResponseValidationError) ErrorName() string {
	return "DeleteEntranceResponseValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteEntranceResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteEntranceResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteEntranceResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteEntranceResponseValidationError{}

// Validate checks the field values on LockRequest with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *LockRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	return nil
}

// LockRequestValidationError is the validation error returned by
// LockRequest.Validate if the designated constraints aren't met.
type LockRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LockRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LockRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LockRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LockRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LockRequestValidationError) ErrorName() string { return "LockRequestValidationError" }

// Error satisfies the builtin error interface
func (e LockRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLockRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LockRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LockRequestValidationError{}

// Validate checks the field values on LockResponse with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *LockResponse) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// LockResponseValidationError is the validation error returned by
// LockResponse.Validate if the designated constraints aren't met.
type LockResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LockResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LockResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LockResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LockResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LockResponseValidationError) ErrorName() string { return "LockResponseValidationError" }

// Error satisfies the builtin error interface
func (e LockResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLockResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LockResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LockResponseValidationError{}

// Validate checks the field values on UnlockRequest with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *UnlockRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	return nil
}

// UnlockRequestValidationError is the validation error returned by
// UnlockRequest.Validate if the designated constraints aren't met.
type UnlockRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UnlockRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UnlockRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UnlockRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UnlockRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UnlockRequestValidationError) ErrorName() string { return "UnlockRequestValidationError" }

// Error satisfies the builtin error interface
func (e UnlockRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUnlockRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UnlockRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UnlockRequestValidationError{}

// Validate checks the field values on UnlockResponse with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *UnlockResponse) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// UnlockResponseValidationError is the validation error returned by
// UnlockResponse.Validate if the designated constraints aren't met.
type UnlockResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UnlockResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UnlockResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UnlockResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UnlockResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UnlockResponseValidationError) ErrorName() string { return "UnlockResponseValidationError" }

// Error satisfies the builtin error interface
func (e UnlockResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUnlockResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UnlockResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UnlockResponseValidationError{}

// Validate checks the field values on FindNearestEntranceRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *FindNearestEntranceRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for X

	// no validation rules for Y

	// no validation rules for Z

	// no validation rules for MaxDistance

	return nil
}

// FindNearestEntranceRequestValidationError is the validation error returned
// by FindNearestEntranceRequest.Validate if the designated constraints aren't met.
type FindNearestEntranceRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FindNearestEntranceRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FindNearestEntranceRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FindNearestEntranceRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FindNearestEntranceRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FindNearestEntranceRequestValidationError) ErrorName() string {
	return "FindNearestEntranceRequestValidationError"
}

// Error satisfies the builtin error interface
func (e FindNearestEntranceRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFindNearestEntranceRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FindNearestEntranceRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FindNearestEntranceRequestValidationError{}

// Validate checks the field values on FindNearestEntranceResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *FindNearestEntranceResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for Distance

	return nil
}

// FindNearestEntranceResponseValidationError is the validation error returned
// by FindNearestEntranceResponse.Validate if the designated constraints
// aren't met.
type FindNearestEntranceResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FindNearestEntranceResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FindNearestEntranceResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FindNearestEntranceResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FindNearestEntranceResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FindNearestEntranceResponseValidationError) ErrorName() string {
	return "FindNearestEntranceResponseValidationError"
}

// Error satisfies the builtin error interface
func (e FindNearestEntranceResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFindNearestEntranceResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FindNearestEntranceResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FindNearestEntranceResponseValidationError{}

// Validate checks the field values on EnterRequest with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *EnterRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	return nil
}

// EnterRequestValidationError is the validation error returned by
// EnterRequest.Validate if the designated constraints aren't met.
type EnterRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EnterRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EnterRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EnterRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EnterRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EnterRequestValidationError) ErrorName() string { return "EnterRequestValidationError" }

// Error satisfies the builtin error interface
func (e EnterRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEnterRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EnterRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EnterRequestValidationError{}

// Validate checks the field values on EnterResponse with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *EnterResponse) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// EnterResponseValidationError is the validation error returned by
// EnterResponse.Validate if the designated constraints aren't met.
type EnterResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EnterResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EnterResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EnterResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EnterResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EnterResponseValidationError) ErrorName() string { return "EnterResponseValidationError" }

// Error satisfies the builtin error interface
func (e EnterResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEnterResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EnterResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EnterResponseValidationError{}

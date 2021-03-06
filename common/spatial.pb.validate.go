// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: common/spatial.proto

package common

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
var _spatial_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on Position with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Position) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for X

	// no validation rules for Y

	// no validation rules for Z

	return nil
}

// PositionValidationError is the validation error returned by
// Position.Validate if the designated constraints aren't met.
type PositionValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PositionValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PositionValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PositionValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PositionValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PositionValidationError) ErrorName() string { return "PositionValidationError" }

// Error satisfies the builtin error interface
func (e PositionValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPosition.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PositionValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PositionValidationError{}

// Validate checks the field values on Rotation with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Rotation) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Rx

	// no validation rules for Ry

	// no validation rules for Rz

	return nil
}

// RotationValidationError is the validation error returned by
// Rotation.Validate if the designated constraints aren't met.
type RotationValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RotationValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RotationValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RotationValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RotationValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RotationValidationError) ErrorName() string { return "RotationValidationError" }

// Error satisfies the builtin error interface
func (e RotationValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRotation.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RotationValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RotationValidationError{}

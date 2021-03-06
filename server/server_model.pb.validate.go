// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: server/server_model.proto

package server

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
var _server_model_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on ServerID with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *ServerID) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	return nil
}

// ServerIDValidationError is the validation error returned by
// ServerID.Validate if the designated constraints aren't met.
type ServerIDValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ServerIDValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ServerIDValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ServerIDValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ServerIDValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ServerIDValidationError) ErrorName() string { return "ServerIDValidationError" }

// Error satisfies the builtin error interface
func (e ServerIDValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sServerID.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ServerIDValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ServerIDValidationError{}

// Validate checks the field values on ServerInfo with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *ServerInfo) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for Name

	// no validation rules for Host

	// no validation rules for Port

	// no validation rules for Platform

	// no validation rules for Status

	// no validation rules for Players

	return nil
}

// ServerInfoValidationError is the validation error returned by
// ServerInfo.Validate if the designated constraints aren't met.
type ServerInfoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ServerInfoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ServerInfoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ServerInfoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ServerInfoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ServerInfoValidationError) ErrorName() string { return "ServerInfoValidationError" }

// Error satisfies the builtin error interface
func (e ServerInfoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sServerInfo.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ServerInfoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ServerInfoValidationError{}

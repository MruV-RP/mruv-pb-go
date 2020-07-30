// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: texturestudio/texturestudio.proto

package texturestudio

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
var _texturestudio_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on CreateServerRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateServerRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for OwnerId

	return nil
}

// CreateServerRequestValidationError is the validation error returned by
// CreateServerRequest.Validate if the designated constraints aren't met.
type CreateServerRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateServerRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateServerRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateServerRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateServerRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateServerRequestValidationError) ErrorName() string {
	return "CreateServerRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateServerRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateServerRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateServerRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateServerRequestValidationError{}

// Validate checks the field values on CreateServerResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateServerResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for Port

	return nil
}

// CreateServerResponseValidationError is the validation error returned by
// CreateServerResponse.Validate if the designated constraints aren't met.
type CreateServerResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateServerResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateServerResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateServerResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateServerResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateServerResponseValidationError) ErrorName() string {
	return "CreateServerResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateServerResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateServerResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateServerResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateServerResponseValidationError{}

// Validate checks the field values on MyServerRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *MyServerRequest) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// MyServerRequestValidationError is the validation error returned by
// MyServerRequest.Validate if the designated constraints aren't met.
type MyServerRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MyServerRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MyServerRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MyServerRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MyServerRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MyServerRequestValidationError) ErrorName() string { return "MyServerRequestValidationError" }

// Error satisfies the builtin error interface
func (e MyServerRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMyServerRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MyServerRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MyServerRequestValidationError{}

// Validate checks the field values on MyServerResponse with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *MyServerResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	return nil
}

// MyServerResponseValidationError is the validation error returned by
// MyServerResponse.Validate if the designated constraints aren't met.
type MyServerResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MyServerResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MyServerResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MyServerResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MyServerResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MyServerResponseValidationError) ErrorName() string { return "MyServerResponseValidationError" }

// Error satisfies the builtin error interface
func (e MyServerResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMyServerResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MyServerResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MyServerResponseValidationError{}

// Validate checks the field values on DeleteServerRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DeleteServerRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	return nil
}

// DeleteServerRequestValidationError is the validation error returned by
// DeleteServerRequest.Validate if the designated constraints aren't met.
type DeleteServerRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteServerRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteServerRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteServerRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteServerRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteServerRequestValidationError) ErrorName() string {
	return "DeleteServerRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteServerRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteServerRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteServerRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteServerRequestValidationError{}

// Validate checks the field values on DeleteServerResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DeleteServerResponse) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// DeleteServerResponseValidationError is the validation error returned by
// DeleteServerResponse.Validate if the designated constraints aren't met.
type DeleteServerResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteServerResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteServerResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteServerResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteServerResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteServerResponseValidationError) ErrorName() string {
	return "DeleteServerResponseValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteServerResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteServerResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteServerResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteServerResponseValidationError{}

// Validate checks the field values on StartServerRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *StartServerRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	return nil
}

// StartServerRequestValidationError is the validation error returned by
// StartServerRequest.Validate if the designated constraints aren't met.
type StartServerRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StartServerRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StartServerRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StartServerRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StartServerRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StartServerRequestValidationError) ErrorName() string {
	return "StartServerRequestValidationError"
}

// Error satisfies the builtin error interface
func (e StartServerRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStartServerRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StartServerRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StartServerRequestValidationError{}

// Validate checks the field values on StartServerResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *StartServerResponse) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// StartServerResponseValidationError is the validation error returned by
// StartServerResponse.Validate if the designated constraints aren't met.
type StartServerResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StartServerResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StartServerResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StartServerResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StartServerResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StartServerResponseValidationError) ErrorName() string {
	return "StartServerResponseValidationError"
}

// Error satisfies the builtin error interface
func (e StartServerResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStartServerResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StartServerResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StartServerResponseValidationError{}

// Validate checks the field values on StopServerRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *StopServerRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	return nil
}

// StopServerRequestValidationError is the validation error returned by
// StopServerRequest.Validate if the designated constraints aren't met.
type StopServerRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StopServerRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StopServerRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StopServerRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StopServerRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StopServerRequestValidationError) ErrorName() string {
	return "StopServerRequestValidationError"
}

// Error satisfies the builtin error interface
func (e StopServerRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStopServerRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StopServerRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StopServerRequestValidationError{}

// Validate checks the field values on StopServerResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *StopServerResponse) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// StopServerResponseValidationError is the validation error returned by
// StopServerResponse.Validate if the designated constraints aren't met.
type StopServerResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StopServerResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StopServerResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StopServerResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StopServerResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StopServerResponseValidationError) ErrorName() string {
	return "StopServerResponseValidationError"
}

// Error satisfies the builtin error interface
func (e StopServerResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStopServerResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StopServerResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StopServerResponseValidationError{}

// Validate checks the field values on RestartServerRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *RestartServerRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	return nil
}

// RestartServerRequestValidationError is the validation error returned by
// RestartServerRequest.Validate if the designated constraints aren't met.
type RestartServerRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RestartServerRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RestartServerRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RestartServerRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RestartServerRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RestartServerRequestValidationError) ErrorName() string {
	return "RestartServerRequestValidationError"
}

// Error satisfies the builtin error interface
func (e RestartServerRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRestartServerRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RestartServerRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RestartServerRequestValidationError{}

// Validate checks the field values on RestartServerResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *RestartServerResponse) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// RestartServerResponseValidationError is the validation error returned by
// RestartServerResponse.Validate if the designated constraints aren't met.
type RestartServerResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RestartServerResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RestartServerResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RestartServerResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RestartServerResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RestartServerResponseValidationError) ErrorName() string {
	return "RestartServerResponseValidationError"
}

// Error satisfies the builtin error interface
func (e RestartServerResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRestartServerResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RestartServerResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RestartServerResponseValidationError{}

// Validate checks the field values on ServerStatusRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ServerStatusRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	return nil
}

// ServerStatusRequestValidationError is the validation error returned by
// ServerStatusRequest.Validate if the designated constraints aren't met.
type ServerStatusRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ServerStatusRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ServerStatusRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ServerStatusRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ServerStatusRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ServerStatusRequestValidationError) ErrorName() string {
	return "ServerStatusRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ServerStatusRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sServerStatusRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ServerStatusRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ServerStatusRequestValidationError{}

// Validate checks the field values on ServerStatusResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ServerStatusResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Port

	// no validation rules for Status

	return nil
}

// ServerStatusResponseValidationError is the validation error returned by
// ServerStatusResponse.Validate if the designated constraints aren't met.
type ServerStatusResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ServerStatusResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ServerStatusResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ServerStatusResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ServerStatusResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ServerStatusResponseValidationError) ErrorName() string {
	return "ServerStatusResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ServerStatusResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sServerStatusResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ServerStatusResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ServerStatusResponseValidationError{}

// Validate checks the field values on UploadProjectRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *UploadProjectRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for ServerId

	// no validation rules for ProjectName

	// no validation rules for ProjectFormat

	switch m.Project.(type) {

	case *UploadProjectRequest_Code:
		// no validation rules for Code

	case *UploadProjectRequest_File:
		// no validation rules for File

	}

	return nil
}

// UploadProjectRequestValidationError is the validation error returned by
// UploadProjectRequest.Validate if the designated constraints aren't met.
type UploadProjectRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UploadProjectRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UploadProjectRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UploadProjectRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UploadProjectRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UploadProjectRequestValidationError) ErrorName() string {
	return "UploadProjectRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UploadProjectRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUploadProjectRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UploadProjectRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UploadProjectRequestValidationError{}

// Validate checks the field values on UploadProjectResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *UploadProjectResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Success

	return nil
}

// UploadProjectResponseValidationError is the validation error returned by
// UploadProjectResponse.Validate if the designated constraints aren't met.
type UploadProjectResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UploadProjectResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UploadProjectResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UploadProjectResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UploadProjectResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UploadProjectResponseValidationError) ErrorName() string {
	return "UploadProjectResponseValidationError"
}

// Error satisfies the builtin error interface
func (e UploadProjectResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUploadProjectResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UploadProjectResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UploadProjectResponseValidationError{}

// Validate checks the field values on GetProjectRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *GetProjectRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for ServerId

	// no validation rules for Name

	return nil
}

// GetProjectRequestValidationError is the validation error returned by
// GetProjectRequest.Validate if the designated constraints aren't met.
type GetProjectRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetProjectRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetProjectRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetProjectRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetProjectRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetProjectRequestValidationError) ErrorName() string {
	return "GetProjectRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetProjectRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetProjectRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetProjectRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetProjectRequestValidationError{}

// Validate checks the field values on GetProjectResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GetProjectResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Code

	return nil
}

// GetProjectResponseValidationError is the validation error returned by
// GetProjectResponse.Validate if the designated constraints aren't met.
type GetProjectResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetProjectResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetProjectResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetProjectResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetProjectResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetProjectResponseValidationError) ErrorName() string {
	return "GetProjectResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetProjectResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetProjectResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetProjectResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetProjectResponseValidationError{}

// Validate checks the field values on GetProjectsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GetProjectsRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for ServerId

	return nil
}

// GetProjectsRequestValidationError is the validation error returned by
// GetProjectsRequest.Validate if the designated constraints aren't met.
type GetProjectsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetProjectsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetProjectsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetProjectsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetProjectsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetProjectsRequestValidationError) ErrorName() string {
	return "GetProjectsRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetProjectsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetProjectsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetProjectsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetProjectsRequestValidationError{}

// Validate checks the field values on GetProjectsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GetProjectsResponse) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// GetProjectsResponseValidationError is the validation error returned by
// GetProjectsResponse.Validate if the designated constraints aren't met.
type GetProjectsResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetProjectsResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetProjectsResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetProjectsResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetProjectsResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetProjectsResponseValidationError) ErrorName() string {
	return "GetProjectsResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetProjectsResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetProjectsResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetProjectsResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetProjectsResponseValidationError{}

// Validate checks the field values on SubscribeToProjectsChangesRequest with
// the rules defined in the proto definition for this message. If any rules
// are violated, an error is returned.
func (m *SubscribeToProjectsChangesRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for ServerId

	return nil
}

// SubscribeToProjectsChangesRequestValidationError is the validation error
// returned by SubscribeToProjectsChangesRequest.Validate if the designated
// constraints aren't met.
type SubscribeToProjectsChangesRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SubscribeToProjectsChangesRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SubscribeToProjectsChangesRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SubscribeToProjectsChangesRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SubscribeToProjectsChangesRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SubscribeToProjectsChangesRequestValidationError) ErrorName() string {
	return "SubscribeToProjectsChangesRequestValidationError"
}

// Error satisfies the builtin error interface
func (e SubscribeToProjectsChangesRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSubscribeToProjectsChangesRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SubscribeToProjectsChangesRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SubscribeToProjectsChangesRequestValidationError{}

// Validate checks the field values on SubscribeToProjectsChangesResponse with
// the rules defined in the proto definition for this message. If any rules
// are violated, an error is returned.
func (m *SubscribeToProjectsChangesResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Name

	return nil
}

// SubscribeToProjectsChangesResponseValidationError is the validation error
// returned by SubscribeToProjectsChangesResponse.Validate if the designated
// constraints aren't met.
type SubscribeToProjectsChangesResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SubscribeToProjectsChangesResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SubscribeToProjectsChangesResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SubscribeToProjectsChangesResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SubscribeToProjectsChangesResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SubscribeToProjectsChangesResponseValidationError) ErrorName() string {
	return "SubscribeToProjectsChangesResponseValidationError"
}

// Error satisfies the builtin error interface
func (e SubscribeToProjectsChangesResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSubscribeToProjectsChangesResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SubscribeToProjectsChangesResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SubscribeToProjectsChangesResponseValidationError{}
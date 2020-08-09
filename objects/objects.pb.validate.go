// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: objects/objects.proto

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
var _objects_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on ObjectModel with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *ObjectModel) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Model

	// no validation rules for ModelName

	// no validation rules for Name

	// no validation rules for Category

	// no validation rules for ProductId

	// no validation rules for Length

	// no validation rules for Width

	// no validation rules for Height

	// no validation rules for Size

	// no validation rules for HasCollision

	// no validation rules for BreaksOnHit

	// no validation rules for HasAnimation

	// no validation rules for VisibleByTime

	// no validation rules for VisibleFrom

	// no validation rules for VisibleTo

	return nil
}

// ObjectModelValidationError is the validation error returned by
// ObjectModel.Validate if the designated constraints aren't met.
type ObjectModelValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ObjectModelValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ObjectModelValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ObjectModelValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ObjectModelValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ObjectModelValidationError) ErrorName() string { return "ObjectModelValidationError" }

// Error satisfies the builtin error interface
func (e ObjectModelValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sObjectModel.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ObjectModelValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ObjectModelValidationError{}

// Validate checks the field values on Object with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Object) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Model

	// no validation rules for EstateId

	// no validation rules for X

	// no validation rules for Y

	// no validation rules for Z

	// no validation rules for Rx

	// no validation rules for Ry

	// no validation rules for Rz

	// no validation rules for WorldId

	// no validation rules for InteriorId

	// no validation rules for PlayerId

	// no validation rules for AreaId

	// no validation rules for StreamDistance

	// no validation rules for DrawDistance

	// no validation rules for Priority

	return nil
}

// ObjectValidationError is the validation error returned by Object.Validate if
// the designated constraints aren't met.
type ObjectValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ObjectValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ObjectValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ObjectValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ObjectValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ObjectValidationError) ErrorName() string { return "ObjectValidationError" }

// Error satisfies the builtin error interface
func (e ObjectValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sObject.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ObjectValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ObjectValidationError{}

// Validate checks the field values on CreateObjectModelRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateObjectModelRequest) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetObjectType()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateObjectModelRequestValidationError{
				field:  "ObjectType",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// CreateObjectModelRequestValidationError is the validation error returned by
// CreateObjectModelRequest.Validate if the designated constraints aren't met.
type CreateObjectModelRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateObjectModelRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateObjectModelRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateObjectModelRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateObjectModelRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateObjectModelRequestValidationError) ErrorName() string {
	return "CreateObjectModelRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateObjectModelRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateObjectModelRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateObjectModelRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateObjectModelRequestValidationError{}

// Validate checks the field values on CreateObjectModelResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateObjectModelResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	return nil
}

// CreateObjectModelResponseValidationError is the validation error returned by
// CreateObjectModelResponse.Validate if the designated constraints aren't met.
type CreateObjectModelResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateObjectModelResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateObjectModelResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateObjectModelResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateObjectModelResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateObjectModelResponseValidationError) ErrorName() string {
	return "CreateObjectModelResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateObjectModelResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateObjectModelResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateObjectModelResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateObjectModelResponseValidationError{}

// Validate checks the field values on GetObjectModelRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GetObjectModelRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Model

	return nil
}

// GetObjectModelRequestValidationError is the validation error returned by
// GetObjectModelRequest.Validate if the designated constraints aren't met.
type GetObjectModelRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetObjectModelRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetObjectModelRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetObjectModelRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetObjectModelRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetObjectModelRequestValidationError) ErrorName() string {
	return "GetObjectModelRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetObjectModelRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetObjectModelRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetObjectModelRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetObjectModelRequestValidationError{}

// Validate checks the field values on GetObjectModelResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GetObjectModelResponse) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetObjectType()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetObjectModelResponseValidationError{
				field:  "ObjectType",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// GetObjectModelResponseValidationError is the validation error returned by
// GetObjectModelResponse.Validate if the designated constraints aren't met.
type GetObjectModelResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetObjectModelResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetObjectModelResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetObjectModelResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetObjectModelResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetObjectModelResponseValidationError) ErrorName() string {
	return "GetObjectModelResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetObjectModelResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetObjectModelResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetObjectModelResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetObjectModelResponseValidationError{}

// Validate checks the field values on UpdateObjectModelRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *UpdateObjectModelRequest) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetObjectType()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateObjectModelRequestValidationError{
				field:  "ObjectType",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// UpdateObjectModelRequestValidationError is the validation error returned by
// UpdateObjectModelRequest.Validate if the designated constraints aren't met.
type UpdateObjectModelRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateObjectModelRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateObjectModelRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateObjectModelRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateObjectModelRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateObjectModelRequestValidationError) ErrorName() string {
	return "UpdateObjectModelRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateObjectModelRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateObjectModelRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateObjectModelRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateObjectModelRequestValidationError{}

// Validate checks the field values on UpdateObjectModelResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *UpdateObjectModelResponse) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// UpdateObjectModelResponseValidationError is the validation error returned by
// UpdateObjectModelResponse.Validate if the designated constraints aren't met.
type UpdateObjectModelResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateObjectModelResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateObjectModelResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateObjectModelResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateObjectModelResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateObjectModelResponseValidationError) ErrorName() string {
	return "UpdateObjectModelResponseValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateObjectModelResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateObjectModelResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateObjectModelResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateObjectModelResponseValidationError{}

// Validate checks the field values on DeleteObjectModelRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DeleteObjectModelRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Model

	return nil
}

// DeleteObjectModelRequestValidationError is the validation error returned by
// DeleteObjectModelRequest.Validate if the designated constraints aren't met.
type DeleteObjectModelRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteObjectModelRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteObjectModelRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteObjectModelRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteObjectModelRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteObjectModelRequestValidationError) ErrorName() string {
	return "DeleteObjectModelRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteObjectModelRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteObjectModelRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteObjectModelRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteObjectModelRequestValidationError{}

// Validate checks the field values on DeleteObjectModelResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DeleteObjectModelResponse) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// DeleteObjectModelResponseValidationError is the validation error returned by
// DeleteObjectModelResponse.Validate if the designated constraints aren't met.
type DeleteObjectModelResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteObjectModelResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteObjectModelResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteObjectModelResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteObjectModelResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteObjectModelResponseValidationError) ErrorName() string {
	return "DeleteObjectModelResponseValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteObjectModelResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteObjectModelResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteObjectModelResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteObjectModelResponseValidationError{}

// Validate checks the field values on CreateObjectRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateObjectRequest) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetObject()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateObjectRequestValidationError{
				field:  "Object",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// CreateObjectRequestValidationError is the validation error returned by
// CreateObjectRequest.Validate if the designated constraints aren't met.
type CreateObjectRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateObjectRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateObjectRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateObjectRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateObjectRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateObjectRequestValidationError) ErrorName() string {
	return "CreateObjectRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateObjectRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateObjectRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateObjectRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateObjectRequestValidationError{}

// Validate checks the field values on CreateObjectResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateObjectResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	return nil
}

// CreateObjectResponseValidationError is the validation error returned by
// CreateObjectResponse.Validate if the designated constraints aren't met.
type CreateObjectResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateObjectResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateObjectResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateObjectResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateObjectResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateObjectResponseValidationError) ErrorName() string {
	return "CreateObjectResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateObjectResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateObjectResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateObjectResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateObjectResponseValidationError{}

// Validate checks the field values on GetObjectRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *GetObjectRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	return nil
}

// GetObjectRequestValidationError is the validation error returned by
// GetObjectRequest.Validate if the designated constraints aren't met.
type GetObjectRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetObjectRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetObjectRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetObjectRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetObjectRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetObjectRequestValidationError) ErrorName() string { return "GetObjectRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetObjectRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetObjectRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetObjectRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetObjectRequestValidationError{}

// Validate checks the field values on GetObjectResponse with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *GetObjectResponse) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetObject()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetObjectResponseValidationError{
				field:  "Object",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// GetObjectResponseValidationError is the validation error returned by
// GetObjectResponse.Validate if the designated constraints aren't met.
type GetObjectResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetObjectResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetObjectResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetObjectResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetObjectResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetObjectResponseValidationError) ErrorName() string {
	return "GetObjectResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetObjectResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetObjectResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetObjectResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetObjectResponseValidationError{}

// Validate checks the field values on UpdateObjectRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *UpdateObjectRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	if v, ok := interface{}(m.GetObject()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateObjectRequestValidationError{
				field:  "Object",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// UpdateObjectRequestValidationError is the validation error returned by
// UpdateObjectRequest.Validate if the designated constraints aren't met.
type UpdateObjectRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateObjectRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateObjectRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateObjectRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateObjectRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateObjectRequestValidationError) ErrorName() string {
	return "UpdateObjectRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateObjectRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateObjectRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateObjectRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateObjectRequestValidationError{}

// Validate checks the field values on UpdateObjectResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *UpdateObjectResponse) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// UpdateObjectResponseValidationError is the validation error returned by
// UpdateObjectResponse.Validate if the designated constraints aren't met.
type UpdateObjectResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateObjectResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateObjectResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateObjectResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateObjectResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateObjectResponseValidationError) ErrorName() string {
	return "UpdateObjectResponseValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateObjectResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateObjectResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateObjectResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateObjectResponseValidationError{}

// Validate checks the field values on DeleteObjectRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DeleteObjectRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	return nil
}

// DeleteObjectRequestValidationError is the validation error returned by
// DeleteObjectRequest.Validate if the designated constraints aren't met.
type DeleteObjectRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteObjectRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteObjectRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteObjectRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteObjectRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteObjectRequestValidationError) ErrorName() string {
	return "DeleteObjectRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteObjectRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteObjectRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteObjectRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteObjectRequestValidationError{}

// Validate checks the field values on DeleteObjectResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DeleteObjectResponse) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// DeleteObjectResponseValidationError is the validation error returned by
// DeleteObjectResponse.Validate if the designated constraints aren't met.
type DeleteObjectResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteObjectResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteObjectResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteObjectResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteObjectResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteObjectResponseValidationError) ErrorName() string {
	return "DeleteObjectResponseValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteObjectResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteObjectResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteObjectResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteObjectResponseValidationError{}

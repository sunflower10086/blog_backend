// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: blog/v1/common.proto

package v1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
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
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on IdList with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *IdList) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on IdList with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in IdListMultiError, or nil if none found.
func (m *IdList) ValidateAll() error {
	return m.validate(true)
}

func (m *IdList) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return IdListMultiError(errors)
	}

	return nil
}

// IdListMultiError is an error wrapping multiple validation errors returned by
// IdList.ValidateAll() if the designated constraints aren't met.
type IdListMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m IdListMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m IdListMultiError) AllErrors() []error { return m }

// IdListValidationError is the validation error returned by IdList.Validate if
// the designated constraints aren't met.
type IdListValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e IdListValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e IdListValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e IdListValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e IdListValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e IdListValidationError) ErrorName() string { return "IdListValidationError" }

// Error satisfies the builtin error interface
func (e IdListValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sIdList.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = IdListValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = IdListValidationError{}

// Validate checks the field values on Id with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Id) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Id with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in IdMultiError, or nil if none found.
func (m *Id) ValidateAll() error {
	return m.validate(true)
}

func (m *Id) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetId() <= 0 {
		err := IdValidationError{
			field:  "Id",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return IdMultiError(errors)
	}

	return nil
}

// IdMultiError is an error wrapping multiple validation errors returned by
// Id.ValidateAll() if the designated constraints aren't met.
type IdMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m IdMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m IdMultiError) AllErrors() []error { return m }

// IdValidationError is the validation error returned by Id.Validate if the
// designated constraints aren't met.
type IdValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e IdValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e IdValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e IdValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e IdValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e IdValidationError) ErrorName() string { return "IdValidationError" }

// Error satisfies the builtin error interface
func (e IdValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sId.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = IdValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = IdValidationError{}

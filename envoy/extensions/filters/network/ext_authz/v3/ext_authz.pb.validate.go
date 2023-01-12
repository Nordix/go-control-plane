// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/extensions/filters/network/ext_authz/v3/ext_authz.proto

package ext_authzv3

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

	v3 "github.com/Nordix/go-control-plane/envoy/config/core/v3"
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

	_ = v3.ApiVersion(0)
)

// Validate checks the field values on ExtAuthz with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ExtAuthz) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ExtAuthz with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ExtAuthzMultiError, or nil
// if none found.
func (m *ExtAuthz) ValidateAll() error {
	return m.validate(true)
}

func (m *ExtAuthz) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetStatPrefix()) < 1 {
		err := ExtAuthzValidationError{
			field:  "StatPrefix",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetGrpcService()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ExtAuthzValidationError{
					field:  "GrpcService",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ExtAuthzValidationError{
					field:  "GrpcService",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetGrpcService()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ExtAuthzValidationError{
				field:  "GrpcService",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for FailureModeAllow

	// no validation rules for IncludePeerCertificate

	if _, ok := v3.ApiVersion_name[int32(m.GetTransportApiVersion())]; !ok {
		err := ExtAuthzValidationError{
			field:  "TransportApiVersion",
			reason: "value must be one of the defined enum values",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetFilterEnabledMetadata()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ExtAuthzValidationError{
					field:  "FilterEnabledMetadata",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ExtAuthzValidationError{
					field:  "FilterEnabledMetadata",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetFilterEnabledMetadata()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ExtAuthzValidationError{
				field:  "FilterEnabledMetadata",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for BootstrapMetadataLabelsKey

	if len(errors) > 0 {
		return ExtAuthzMultiError(errors)
	}

	return nil
}

// ExtAuthzMultiError is an error wrapping multiple validation errors returned
// by ExtAuthz.ValidateAll() if the designated constraints aren't met.
type ExtAuthzMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ExtAuthzMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ExtAuthzMultiError) AllErrors() []error { return m }

// ExtAuthzValidationError is the validation error returned by
// ExtAuthz.Validate if the designated constraints aren't met.
type ExtAuthzValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ExtAuthzValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ExtAuthzValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ExtAuthzValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ExtAuthzValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ExtAuthzValidationError) ErrorName() string { return "ExtAuthzValidationError" }

// Error satisfies the builtin error interface
func (e ExtAuthzValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sExtAuthz.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ExtAuthzValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ExtAuthzValidationError{}

// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/service/rate_limit_quota/v3/rlqs.proto

package rate_limit_quotav3

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

// Validate checks the field values on RateLimitQuotaUsageReports with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *RateLimitQuotaUsageReports) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RateLimitQuotaUsageReports with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// RateLimitQuotaUsageReportsMultiError, or nil if none found.
func (m *RateLimitQuotaUsageReports) ValidateAll() error {
	return m.validate(true)
}

func (m *RateLimitQuotaUsageReports) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetDomain()) < 1 {
		err := RateLimitQuotaUsageReportsValidationError{
			field:  "Domain",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(m.GetBucketQuotaUsages()) < 1 {
		err := RateLimitQuotaUsageReportsValidationError{
			field:  "BucketQuotaUsages",
			reason: "value must contain at least 1 item(s)",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	for idx, item := range m.GetBucketQuotaUsages() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, RateLimitQuotaUsageReportsValidationError{
						field:  fmt.Sprintf("BucketQuotaUsages[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, RateLimitQuotaUsageReportsValidationError{
						field:  fmt.Sprintf("BucketQuotaUsages[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return RateLimitQuotaUsageReportsValidationError{
					field:  fmt.Sprintf("BucketQuotaUsages[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return RateLimitQuotaUsageReportsMultiError(errors)
	}

	return nil
}

// RateLimitQuotaUsageReportsMultiError is an error wrapping multiple
// validation errors returned by RateLimitQuotaUsageReports.ValidateAll() if
// the designated constraints aren't met.
type RateLimitQuotaUsageReportsMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RateLimitQuotaUsageReportsMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RateLimitQuotaUsageReportsMultiError) AllErrors() []error { return m }

// RateLimitQuotaUsageReportsValidationError is the validation error returned
// by RateLimitQuotaUsageReports.Validate if the designated constraints aren't met.
type RateLimitQuotaUsageReportsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RateLimitQuotaUsageReportsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RateLimitQuotaUsageReportsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RateLimitQuotaUsageReportsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RateLimitQuotaUsageReportsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RateLimitQuotaUsageReportsValidationError) ErrorName() string {
	return "RateLimitQuotaUsageReportsValidationError"
}

// Error satisfies the builtin error interface
func (e RateLimitQuotaUsageReportsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRateLimitQuotaUsageReports.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RateLimitQuotaUsageReportsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RateLimitQuotaUsageReportsValidationError{}

// Validate checks the field values on RateLimitQuotaResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *RateLimitQuotaResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RateLimitQuotaResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// RateLimitQuotaResponseMultiError, or nil if none found.
func (m *RateLimitQuotaResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *RateLimitQuotaResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(m.GetBucketAction()) < 1 {
		err := RateLimitQuotaResponseValidationError{
			field:  "BucketAction",
			reason: "value must contain at least 1 item(s)",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	for idx, item := range m.GetBucketAction() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, RateLimitQuotaResponseValidationError{
						field:  fmt.Sprintf("BucketAction[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, RateLimitQuotaResponseValidationError{
						field:  fmt.Sprintf("BucketAction[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return RateLimitQuotaResponseValidationError{
					field:  fmt.Sprintf("BucketAction[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return RateLimitQuotaResponseMultiError(errors)
	}

	return nil
}

// RateLimitQuotaResponseMultiError is an error wrapping multiple validation
// errors returned by RateLimitQuotaResponse.ValidateAll() if the designated
// constraints aren't met.
type RateLimitQuotaResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RateLimitQuotaResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RateLimitQuotaResponseMultiError) AllErrors() []error { return m }

// RateLimitQuotaResponseValidationError is the validation error returned by
// RateLimitQuotaResponse.Validate if the designated constraints aren't met.
type RateLimitQuotaResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RateLimitQuotaResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RateLimitQuotaResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RateLimitQuotaResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RateLimitQuotaResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RateLimitQuotaResponseValidationError) ErrorName() string {
	return "RateLimitQuotaResponseValidationError"
}

// Error satisfies the builtin error interface
func (e RateLimitQuotaResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRateLimitQuotaResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RateLimitQuotaResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RateLimitQuotaResponseValidationError{}

// Validate checks the field values on BucketId with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *BucketId) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on BucketId with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in BucketIdMultiError, or nil
// if none found.
func (m *BucketId) ValidateAll() error {
	return m.validate(true)
}

func (m *BucketId) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(m.GetBucket()) < 1 {
		err := BucketIdValidationError{
			field:  "Bucket",
			reason: "value must contain at least 1 pair(s)",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	{
		sorted_keys := make([]string, len(m.GetBucket()))
		i := 0
		for key := range m.GetBucket() {
			sorted_keys[i] = key
			i++
		}
		sort.Slice(sorted_keys, func(i, j int) bool { return sorted_keys[i] < sorted_keys[j] })
		for _, key := range sorted_keys {
			val := m.GetBucket()[key]
			_ = val

			if utf8.RuneCountInString(key) < 1 {
				err := BucketIdValidationError{
					field:  fmt.Sprintf("Bucket[%v]", key),
					reason: "value length must be at least 1 runes",
				}
				if !all {
					return err
				}
				errors = append(errors, err)
			}

			if utf8.RuneCountInString(val) < 1 {
				err := BucketIdValidationError{
					field:  fmt.Sprintf("Bucket[%v]", key),
					reason: "value length must be at least 1 runes",
				}
				if !all {
					return err
				}
				errors = append(errors, err)
			}

		}
	}

	if len(errors) > 0 {
		return BucketIdMultiError(errors)
	}

	return nil
}

// BucketIdMultiError is an error wrapping multiple validation errors returned
// by BucketId.ValidateAll() if the designated constraints aren't met.
type BucketIdMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m BucketIdMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m BucketIdMultiError) AllErrors() []error { return m }

// BucketIdValidationError is the validation error returned by
// BucketId.Validate if the designated constraints aren't met.
type BucketIdValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BucketIdValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BucketIdValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BucketIdValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BucketIdValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BucketIdValidationError) ErrorName() string { return "BucketIdValidationError" }

// Error satisfies the builtin error interface
func (e BucketIdValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBucketId.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BucketIdValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BucketIdValidationError{}

// Validate checks the field values on
// RateLimitQuotaUsageReports_BucketQuotaUsage with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *RateLimitQuotaUsageReports_BucketQuotaUsage) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on
// RateLimitQuotaUsageReports_BucketQuotaUsage with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in
// RateLimitQuotaUsageReports_BucketQuotaUsageMultiError, or nil if none found.
func (m *RateLimitQuotaUsageReports_BucketQuotaUsage) ValidateAll() error {
	return m.validate(true)
}

func (m *RateLimitQuotaUsageReports_BucketQuotaUsage) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetBucketId() == nil {
		err := RateLimitQuotaUsageReports_BucketQuotaUsageValidationError{
			field:  "BucketId",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetBucketId()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, RateLimitQuotaUsageReports_BucketQuotaUsageValidationError{
					field:  "BucketId",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, RateLimitQuotaUsageReports_BucketQuotaUsageValidationError{
					field:  "BucketId",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetBucketId()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RateLimitQuotaUsageReports_BucketQuotaUsageValidationError{
				field:  "BucketId",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.GetTimeElapsed() == nil {
		err := RateLimitQuotaUsageReports_BucketQuotaUsageValidationError{
			field:  "TimeElapsed",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if d := m.GetTimeElapsed(); d != nil {
		dur, err := d.AsDuration(), d.CheckValid()
		if err != nil {
			err = RateLimitQuotaUsageReports_BucketQuotaUsageValidationError{
				field:  "TimeElapsed",
				reason: "value is not a valid duration",
				cause:  err,
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		} else {

			gt := time.Duration(0*time.Second + 0*time.Nanosecond)

			if dur <= gt {
				err := RateLimitQuotaUsageReports_BucketQuotaUsageValidationError{
					field:  "TimeElapsed",
					reason: "value must be greater than 0s",
				}
				if !all {
					return err
				}
				errors = append(errors, err)
			}

		}
	}

	// no validation rules for NumRequestsAllowed

	// no validation rules for NumRequestsDenied

	if len(errors) > 0 {
		return RateLimitQuotaUsageReports_BucketQuotaUsageMultiError(errors)
	}

	return nil
}

// RateLimitQuotaUsageReports_BucketQuotaUsageMultiError is an error wrapping
// multiple validation errors returned by
// RateLimitQuotaUsageReports_BucketQuotaUsage.ValidateAll() if the designated
// constraints aren't met.
type RateLimitQuotaUsageReports_BucketQuotaUsageMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RateLimitQuotaUsageReports_BucketQuotaUsageMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RateLimitQuotaUsageReports_BucketQuotaUsageMultiError) AllErrors() []error { return m }

// RateLimitQuotaUsageReports_BucketQuotaUsageValidationError is the validation
// error returned by RateLimitQuotaUsageReports_BucketQuotaUsage.Validate if
// the designated constraints aren't met.
type RateLimitQuotaUsageReports_BucketQuotaUsageValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RateLimitQuotaUsageReports_BucketQuotaUsageValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RateLimitQuotaUsageReports_BucketQuotaUsageValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RateLimitQuotaUsageReports_BucketQuotaUsageValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RateLimitQuotaUsageReports_BucketQuotaUsageValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RateLimitQuotaUsageReports_BucketQuotaUsageValidationError) ErrorName() string {
	return "RateLimitQuotaUsageReports_BucketQuotaUsageValidationError"
}

// Error satisfies the builtin error interface
func (e RateLimitQuotaUsageReports_BucketQuotaUsageValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRateLimitQuotaUsageReports_BucketQuotaUsage.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RateLimitQuotaUsageReports_BucketQuotaUsageValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RateLimitQuotaUsageReports_BucketQuotaUsageValidationError{}

// Validate checks the field values on RateLimitQuotaResponse_BucketAction with
// the rules defined in the proto definition for this message. If any rules
// are violated, the first error encountered is returned, or nil if there are
// no violations.
func (m *RateLimitQuotaResponse_BucketAction) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RateLimitQuotaResponse_BucketAction
// with the rules defined in the proto definition for this message. If any
// rules are violated, the result is a list of violation errors wrapped in
// RateLimitQuotaResponse_BucketActionMultiError, or nil if none found.
func (m *RateLimitQuotaResponse_BucketAction) ValidateAll() error {
	return m.validate(true)
}

func (m *RateLimitQuotaResponse_BucketAction) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetBucketId() == nil {
		err := RateLimitQuotaResponse_BucketActionValidationError{
			field:  "BucketId",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetBucketId()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, RateLimitQuotaResponse_BucketActionValidationError{
					field:  "BucketId",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, RateLimitQuotaResponse_BucketActionValidationError{
					field:  "BucketId",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetBucketId()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RateLimitQuotaResponse_BucketActionValidationError{
				field:  "BucketId",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	switch m.BucketAction.(type) {

	case *RateLimitQuotaResponse_BucketAction_QuotaAssignmentAction_:

		if all {
			switch v := interface{}(m.GetQuotaAssignmentAction()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, RateLimitQuotaResponse_BucketActionValidationError{
						field:  "QuotaAssignmentAction",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, RateLimitQuotaResponse_BucketActionValidationError{
						field:  "QuotaAssignmentAction",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetQuotaAssignmentAction()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return RateLimitQuotaResponse_BucketActionValidationError{
					field:  "QuotaAssignmentAction",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *RateLimitQuotaResponse_BucketAction_AbandonAction_:

		if all {
			switch v := interface{}(m.GetAbandonAction()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, RateLimitQuotaResponse_BucketActionValidationError{
						field:  "AbandonAction",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, RateLimitQuotaResponse_BucketActionValidationError{
						field:  "AbandonAction",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetAbandonAction()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return RateLimitQuotaResponse_BucketActionValidationError{
					field:  "AbandonAction",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		err := RateLimitQuotaResponse_BucketActionValidationError{
			field:  "BucketAction",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)

	}

	if len(errors) > 0 {
		return RateLimitQuotaResponse_BucketActionMultiError(errors)
	}

	return nil
}

// RateLimitQuotaResponse_BucketActionMultiError is an error wrapping multiple
// validation errors returned by
// RateLimitQuotaResponse_BucketAction.ValidateAll() if the designated
// constraints aren't met.
type RateLimitQuotaResponse_BucketActionMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RateLimitQuotaResponse_BucketActionMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RateLimitQuotaResponse_BucketActionMultiError) AllErrors() []error { return m }

// RateLimitQuotaResponse_BucketActionValidationError is the validation error
// returned by RateLimitQuotaResponse_BucketAction.Validate if the designated
// constraints aren't met.
type RateLimitQuotaResponse_BucketActionValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RateLimitQuotaResponse_BucketActionValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RateLimitQuotaResponse_BucketActionValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RateLimitQuotaResponse_BucketActionValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RateLimitQuotaResponse_BucketActionValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RateLimitQuotaResponse_BucketActionValidationError) ErrorName() string {
	return "RateLimitQuotaResponse_BucketActionValidationError"
}

// Error satisfies the builtin error interface
func (e RateLimitQuotaResponse_BucketActionValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRateLimitQuotaResponse_BucketAction.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RateLimitQuotaResponse_BucketActionValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RateLimitQuotaResponse_BucketActionValidationError{}

// Validate checks the field values on
// RateLimitQuotaResponse_BucketAction_QuotaAssignmentAction with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *RateLimitQuotaResponse_BucketAction_QuotaAssignmentAction) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on
// RateLimitQuotaResponse_BucketAction_QuotaAssignmentAction with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// RateLimitQuotaResponse_BucketAction_QuotaAssignmentActionMultiError, or nil
// if none found.
func (m *RateLimitQuotaResponse_BucketAction_QuotaAssignmentAction) ValidateAll() error {
	return m.validate(true)
}

func (m *RateLimitQuotaResponse_BucketAction_QuotaAssignmentAction) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if d := m.GetAssignmentTimeToLive(); d != nil {
		dur, err := d.AsDuration(), d.CheckValid()
		if err != nil {
			err = RateLimitQuotaResponse_BucketAction_QuotaAssignmentActionValidationError{
				field:  "AssignmentTimeToLive",
				reason: "value is not a valid duration",
				cause:  err,
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		} else {

			gte := time.Duration(0*time.Second + 0*time.Nanosecond)

			if dur < gte {
				err := RateLimitQuotaResponse_BucketAction_QuotaAssignmentActionValidationError{
					field:  "AssignmentTimeToLive",
					reason: "value must be greater than or equal to 0s",
				}
				if !all {
					return err
				}
				errors = append(errors, err)
			}

		}
	}

	if all {
		switch v := interface{}(m.GetRateLimitStrategy()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, RateLimitQuotaResponse_BucketAction_QuotaAssignmentActionValidationError{
					field:  "RateLimitStrategy",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, RateLimitQuotaResponse_BucketAction_QuotaAssignmentActionValidationError{
					field:  "RateLimitStrategy",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetRateLimitStrategy()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RateLimitQuotaResponse_BucketAction_QuotaAssignmentActionValidationError{
				field:  "RateLimitStrategy",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return RateLimitQuotaResponse_BucketAction_QuotaAssignmentActionMultiError(errors)
	}

	return nil
}

// RateLimitQuotaResponse_BucketAction_QuotaAssignmentActionMultiError is an
// error wrapping multiple validation errors returned by
// RateLimitQuotaResponse_BucketAction_QuotaAssignmentAction.ValidateAll() if
// the designated constraints aren't met.
type RateLimitQuotaResponse_BucketAction_QuotaAssignmentActionMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RateLimitQuotaResponse_BucketAction_QuotaAssignmentActionMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RateLimitQuotaResponse_BucketAction_QuotaAssignmentActionMultiError) AllErrors() []error {
	return m
}

// RateLimitQuotaResponse_BucketAction_QuotaAssignmentActionValidationError is
// the validation error returned by
// RateLimitQuotaResponse_BucketAction_QuotaAssignmentAction.Validate if the
// designated constraints aren't met.
type RateLimitQuotaResponse_BucketAction_QuotaAssignmentActionValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RateLimitQuotaResponse_BucketAction_QuotaAssignmentActionValidationError) Field() string {
	return e.field
}

// Reason function returns reason value.
func (e RateLimitQuotaResponse_BucketAction_QuotaAssignmentActionValidationError) Reason() string {
	return e.reason
}

// Cause function returns cause value.
func (e RateLimitQuotaResponse_BucketAction_QuotaAssignmentActionValidationError) Cause() error {
	return e.cause
}

// Key function returns key value.
func (e RateLimitQuotaResponse_BucketAction_QuotaAssignmentActionValidationError) Key() bool {
	return e.key
}

// ErrorName returns error name.
func (e RateLimitQuotaResponse_BucketAction_QuotaAssignmentActionValidationError) ErrorName() string {
	return "RateLimitQuotaResponse_BucketAction_QuotaAssignmentActionValidationError"
}

// Error satisfies the builtin error interface
func (e RateLimitQuotaResponse_BucketAction_QuotaAssignmentActionValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRateLimitQuotaResponse_BucketAction_QuotaAssignmentAction.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RateLimitQuotaResponse_BucketAction_QuotaAssignmentActionValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RateLimitQuotaResponse_BucketAction_QuotaAssignmentActionValidationError{}

// Validate checks the field values on
// RateLimitQuotaResponse_BucketAction_AbandonAction with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *RateLimitQuotaResponse_BucketAction_AbandonAction) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on
// RateLimitQuotaResponse_BucketAction_AbandonAction with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in
// RateLimitQuotaResponse_BucketAction_AbandonActionMultiError, or nil if none found.
func (m *RateLimitQuotaResponse_BucketAction_AbandonAction) ValidateAll() error {
	return m.validate(true)
}

func (m *RateLimitQuotaResponse_BucketAction_AbandonAction) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return RateLimitQuotaResponse_BucketAction_AbandonActionMultiError(errors)
	}

	return nil
}

// RateLimitQuotaResponse_BucketAction_AbandonActionMultiError is an error
// wrapping multiple validation errors returned by
// RateLimitQuotaResponse_BucketAction_AbandonAction.ValidateAll() if the
// designated constraints aren't met.
type RateLimitQuotaResponse_BucketAction_AbandonActionMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RateLimitQuotaResponse_BucketAction_AbandonActionMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RateLimitQuotaResponse_BucketAction_AbandonActionMultiError) AllErrors() []error { return m }

// RateLimitQuotaResponse_BucketAction_AbandonActionValidationError is the
// validation error returned by
// RateLimitQuotaResponse_BucketAction_AbandonAction.Validate if the
// designated constraints aren't met.
type RateLimitQuotaResponse_BucketAction_AbandonActionValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RateLimitQuotaResponse_BucketAction_AbandonActionValidationError) Field() string {
	return e.field
}

// Reason function returns reason value.
func (e RateLimitQuotaResponse_BucketAction_AbandonActionValidationError) Reason() string {
	return e.reason
}

// Cause function returns cause value.
func (e RateLimitQuotaResponse_BucketAction_AbandonActionValidationError) Cause() error {
	return e.cause
}

// Key function returns key value.
func (e RateLimitQuotaResponse_BucketAction_AbandonActionValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RateLimitQuotaResponse_BucketAction_AbandonActionValidationError) ErrorName() string {
	return "RateLimitQuotaResponse_BucketAction_AbandonActionValidationError"
}

// Error satisfies the builtin error interface
func (e RateLimitQuotaResponse_BucketAction_AbandonActionValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRateLimitQuotaResponse_BucketAction_AbandonAction.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RateLimitQuotaResponse_BucketAction_AbandonActionValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RateLimitQuotaResponse_BucketAction_AbandonActionValidationError{}
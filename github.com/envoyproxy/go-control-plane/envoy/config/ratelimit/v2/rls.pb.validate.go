// Code generated by protoc-gen-validate
// source: envoy/config/ratelimit/v2/rls.proto
// DO NOT EDIT!!!

package v2

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

	"github.com/gogo/protobuf/types"
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
	_ = types.DynamicAny{}
)

// Validate checks the field values on RateLimitServiceConfig with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *RateLimitServiceConfig) Validate() error {
	if m == nil {
		return nil
	}

	switch m.ServiceSpecifier.(type) {

	case *RateLimitServiceConfig_ClusterName:

		if len(m.GetClusterName()) < 1 {
			return RateLimitServiceConfigValidationError{
				Field:  "ClusterName",
				Reason: "value length must be at least 1 bytes",
			}
		}

	case *RateLimitServiceConfig_GrpcService:

		if v, ok := interface{}(m.GetGrpcService()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return RateLimitServiceConfigValidationError{
					Field:  "GrpcService",
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	default:
		return RateLimitServiceConfigValidationError{
			Field:  "ServiceSpecifier",
			Reason: "value is required",
		}

	}

	return nil
}

// RateLimitServiceConfigValidationError is the validation error returned by
// RateLimitServiceConfig.Validate if the designated constraints aren't met.
type RateLimitServiceConfigValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e RateLimitServiceConfigValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRateLimitServiceConfig.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = RateLimitServiceConfigValidationError{}

// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: harmonytypes/v1/types.proto

package v1

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
var _types_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on HarmonyMethodMetadata with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *HarmonyMethodMetadata) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for RequiresAuthentication

	// no validation rules for RequiresLocal

	// no validation rules for RequiresPermissionNode

	return nil
}

// HarmonyMethodMetadataValidationError is the validation error returned by
// HarmonyMethodMetadata.Validate if the designated constraints aren't met.
type HarmonyMethodMetadataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HarmonyMethodMetadataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HarmonyMethodMetadataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HarmonyMethodMetadataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HarmonyMethodMetadataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HarmonyMethodMetadataValidationError) ErrorName() string {
	return "HarmonyMethodMetadataValidationError"
}

// Error satisfies the builtin error interface
func (e HarmonyMethodMetadataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHarmonyMethodMetadata.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HarmonyMethodMetadataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HarmonyMethodMetadataValidationError{}

// Validate checks the field values on Override with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Override) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Name

	// no validation rules for Avatar

	switch m.Reason.(type) {

	case *Override_UserDefined:
		// no validation rules for UserDefined

	case *Override_Webhook:

		if v, ok := interface{}(m.GetWebhook()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return OverrideValidationError{
					field:  "Webhook",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *Override_SystemPlurality:

		if v, ok := interface{}(m.GetSystemPlurality()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return OverrideValidationError{
					field:  "SystemPlurality",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *Override_SystemMessage:

		if v, ok := interface{}(m.GetSystemMessage()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return OverrideValidationError{
					field:  "SystemMessage",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *Override_Bridge:

		if v, ok := interface{}(m.GetBridge()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return OverrideValidationError{
					field:  "Bridge",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// OverrideValidationError is the validation error returned by
// Override.Validate if the designated constraints aren't met.
type OverrideValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OverrideValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OverrideValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OverrideValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OverrideValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OverrideValidationError) ErrorName() string { return "OverrideValidationError" }

// Error satisfies the builtin error interface
func (e OverrideValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOverride.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OverrideValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OverrideValidationError{}

// Validate checks the field values on Action with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Action) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Text

	// no validation rules for Url

	// no validation rules for Id

	// no validation rules for Type

	// no validation rules for Presentation

	for idx, item := range m.GetChildren() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ActionValidationError{
					field:  fmt.Sprintf("Children[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ActionValidationError is the validation error returned by Action.Validate if
// the designated constraints aren't met.
type ActionValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ActionValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ActionValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ActionValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ActionValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ActionValidationError) ErrorName() string { return "ActionValidationError" }

// Error satisfies the builtin error interface
func (e ActionValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAction.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ActionValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ActionValidationError{}

// Validate checks the field values on EmbedHeading with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *EmbedHeading) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Text

	// no validation rules for Subtext

	// no validation rules for Url

	// no validation rules for Icon

	return nil
}

// EmbedHeadingValidationError is the validation error returned by
// EmbedHeading.Validate if the designated constraints aren't met.
type EmbedHeadingValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EmbedHeadingValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EmbedHeadingValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EmbedHeadingValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EmbedHeadingValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EmbedHeadingValidationError) ErrorName() string { return "EmbedHeadingValidationError" }

// Error satisfies the builtin error interface
func (e EmbedHeadingValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEmbedHeading.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EmbedHeadingValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EmbedHeadingValidationError{}

// Validate checks the field values on EmbedField with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *EmbedField) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Title

	// no validation rules for Subtitle

	// no validation rules for Body

	// no validation rules for ImageUrl

	// no validation rules for Presentation

	for idx, item := range m.GetActions() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return EmbedFieldValidationError{
					field:  fmt.Sprintf("Actions[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// EmbedFieldValidationError is the validation error returned by
// EmbedField.Validate if the designated constraints aren't met.
type EmbedFieldValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EmbedFieldValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EmbedFieldValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EmbedFieldValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EmbedFieldValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EmbedFieldValidationError) ErrorName() string { return "EmbedFieldValidationError" }

// Error satisfies the builtin error interface
func (e EmbedFieldValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEmbedField.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EmbedFieldValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EmbedFieldValidationError{}

// Validate checks the field values on Embed with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Embed) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Title

	// no validation rules for Body

	// no validation rules for Color

	if v, ok := interface{}(m.GetHeader()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return EmbedValidationError{
				field:  "Header",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetFooter()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return EmbedValidationError{
				field:  "Footer",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetFields() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return EmbedValidationError{
					field:  fmt.Sprintf("Fields[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	for idx, item := range m.GetActions() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return EmbedValidationError{
					field:  fmt.Sprintf("Actions[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// EmbedValidationError is the validation error returned by Embed.Validate if
// the designated constraints aren't met.
type EmbedValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EmbedValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EmbedValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EmbedValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EmbedValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EmbedValidationError) ErrorName() string { return "EmbedValidationError" }

// Error satisfies the builtin error interface
func (e EmbedValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEmbed.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EmbedValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EmbedValidationError{}

// Validate checks the field values on Attachment with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Attachment) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for Name

	// no validation rules for Type

	// no validation rules for Size

	return nil
}

// AttachmentValidationError is the validation error returned by
// Attachment.Validate if the designated constraints aren't met.
type AttachmentValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AttachmentValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AttachmentValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AttachmentValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AttachmentValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AttachmentValidationError) ErrorName() string { return "AttachmentValidationError" }

// Error satisfies the builtin error interface
func (e AttachmentValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAttachment.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AttachmentValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AttachmentValidationError{}

// Validate checks the field values on Metadata with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Metadata) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Kind

	for key, val := range m.GetExtension() {
		_ = val

		// no validation rules for Extension[key]

		if v, ok := interface{}(val).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return MetadataValidationError{
					field:  fmt.Sprintf("Extension[%v]", key),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// MetadataValidationError is the validation error returned by
// Metadata.Validate if the designated constraints aren't met.
type MetadataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MetadataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MetadataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MetadataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MetadataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MetadataValidationError) ErrorName() string { return "MetadataValidationError" }

// Error satisfies the builtin error interface
func (e MetadataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMetadata.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MetadataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MetadataValidationError{}

// Validate checks the field values on ContentText with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *ContentText) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Content

	return nil
}

// ContentTextValidationError is the validation error returned by
// ContentText.Validate if the designated constraints aren't met.
type ContentTextValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ContentTextValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ContentTextValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ContentTextValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ContentTextValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ContentTextValidationError) ErrorName() string { return "ContentTextValidationError" }

// Error satisfies the builtin error interface
func (e ContentTextValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sContentText.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ContentTextValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ContentTextValidationError{}

// Validate checks the field values on Photo with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Photo) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetPhoto()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PhotoValidationError{
				field:  "Photo",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Caption

	return nil
}

// PhotoValidationError is the validation error returned by Photo.Validate if
// the designated constraints aren't met.
type PhotoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PhotoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PhotoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PhotoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PhotoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PhotoValidationError) ErrorName() string { return "PhotoValidationError" }

// Error satisfies the builtin error interface
func (e PhotoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPhoto.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PhotoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PhotoValidationError{}

// Validate checks the field values on ContentPhoto with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *ContentPhoto) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetPhotos() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ContentPhotoValidationError{
					field:  fmt.Sprintf("Photos[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ContentPhotoValidationError is the validation error returned by
// ContentPhoto.Validate if the designated constraints aren't met.
type ContentPhotoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ContentPhotoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ContentPhotoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ContentPhotoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ContentPhotoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ContentPhotoValidationError) ErrorName() string { return "ContentPhotoValidationError" }

// Error satisfies the builtin error interface
func (e ContentPhotoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sContentPhoto.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ContentPhotoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ContentPhotoValidationError{}

// Validate checks the field values on ContentEmbed with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *ContentEmbed) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetEmbeds() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ContentEmbedValidationError{
					field:  fmt.Sprintf("Embeds[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ContentEmbedValidationError is the validation error returned by
// ContentEmbed.Validate if the designated constraints aren't met.
type ContentEmbedValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ContentEmbedValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ContentEmbedValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ContentEmbedValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ContentEmbedValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ContentEmbedValidationError) ErrorName() string { return "ContentEmbedValidationError" }

// Error satisfies the builtin error interface
func (e ContentEmbedValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sContentEmbed.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ContentEmbedValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ContentEmbedValidationError{}

// Validate checks the field values on ContentFiles with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *ContentFiles) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetAttachments() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ContentFilesValidationError{
					field:  fmt.Sprintf("Attachments[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ContentFilesValidationError is the validation error returned by
// ContentFiles.Validate if the designated constraints aren't met.
type ContentFilesValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ContentFilesValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ContentFilesValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ContentFilesValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ContentFilesValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ContentFilesValidationError) ErrorName() string { return "ContentFilesValidationError" }

// Error satisfies the builtin error interface
func (e ContentFilesValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sContentFiles.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ContentFilesValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ContentFilesValidationError{}

// Validate checks the field values on Content with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Content) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetActions() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ContentValidationError{
					field:  fmt.Sprintf("Actions[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	switch m.Content.(type) {

	case *Content_TextMessage:

		if v, ok := interface{}(m.GetTextMessage()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ContentValidationError{
					field:  "TextMessage",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *Content_PhotoMessage:

		if v, ok := interface{}(m.GetPhotoMessage()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ContentValidationError{
					field:  "PhotoMessage",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *Content_EmbedMessage:

		if v, ok := interface{}(m.GetEmbedMessage()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ContentValidationError{
					field:  "EmbedMessage",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *Content_FilesMessage:

		if v, ok := interface{}(m.GetFilesMessage()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ContentValidationError{
					field:  "FilesMessage",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ContentValidationError is the validation error returned by Content.Validate
// if the designated constraints aren't met.
type ContentValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ContentValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ContentValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ContentValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ContentValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ContentValidationError) ErrorName() string { return "ContentValidationError" }

// Error satisfies the builtin error interface
func (e ContentValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sContent.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ContentValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ContentValidationError{}

// Validate checks the field values on Message with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Message) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetMetadata()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return MessageValidationError{
				field:  "Metadata",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetOverrides()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return MessageValidationError{
				field:  "Overrides",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for GuildId

	// no validation rules for ChannelId

	// no validation rules for MessageId

	// no validation rules for AuthorId

	if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return MessageValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetEditedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return MessageValidationError{
				field:  "EditedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for InReplyTo

	if v, ok := interface{}(m.GetContent()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return MessageValidationError{
				field:  "Content",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// MessageValidationError is the validation error returned by Message.Validate
// if the designated constraints aren't met.
type MessageValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MessageValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MessageValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MessageValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MessageValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MessageValidationError) ErrorName() string { return "MessageValidationError" }

// Error satisfies the builtin error interface
func (e MessageValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMessage.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MessageValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MessageValidationError{}

// Validate checks the field values on Error with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Error) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Identifier

	// no validation rules for HumanMessage

	// no validation rules for MoreDetails

	return nil
}

// ErrorValidationError is the validation error returned by Error.Validate if
// the designated constraints aren't met.
type ErrorValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ErrorValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ErrorValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ErrorValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ErrorValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ErrorValidationError) ErrorName() string { return "ErrorValidationError" }

// Error satisfies the builtin error interface
func (e ErrorValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sError.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ErrorValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ErrorValidationError{}

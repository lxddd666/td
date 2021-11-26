// Code generated by gotdgen, DO NOT EDIT.

package tdapi

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"go.uber.org/multierr"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/tdjson"
	"github.com/gotd/td/tdp"
	"github.com/gotd/td/tgerr"
)

// No-op definition for keeping imports.
var (
	_ = bin.Buffer{}
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = multierr.AppendInto
	_ = sort.Ints
	_ = tdp.Format
	_ = tgerr.Error{}
	_ = tdjson.Encoder{}
)

// EditInlineMessageTextRequest represents TL type `editInlineMessageText#cd02c1e5`.
type EditInlineMessageTextRequest struct {
	// Inline message identifier
	InlineMessageID string
	// The new message reply markup; pass null if none
	ReplyMarkup ReplyMarkupClass
	// New text content of the message. Must be of type inputMessageText
	InputMessageContent InputMessageContentClass
}

// EditInlineMessageTextRequestTypeID is TL type id of EditInlineMessageTextRequest.
const EditInlineMessageTextRequestTypeID = 0xcd02c1e5

// Ensuring interfaces in compile-time for EditInlineMessageTextRequest.
var (
	_ bin.Encoder     = &EditInlineMessageTextRequest{}
	_ bin.Decoder     = &EditInlineMessageTextRequest{}
	_ bin.BareEncoder = &EditInlineMessageTextRequest{}
	_ bin.BareDecoder = &EditInlineMessageTextRequest{}
)

func (e *EditInlineMessageTextRequest) Zero() bool {
	if e == nil {
		return true
	}
	if !(e.InlineMessageID == "") {
		return false
	}
	if !(e.ReplyMarkup == nil) {
		return false
	}
	if !(e.InputMessageContent == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (e *EditInlineMessageTextRequest) String() string {
	if e == nil {
		return "EditInlineMessageTextRequest(nil)"
	}
	type Alias EditInlineMessageTextRequest
	return fmt.Sprintf("EditInlineMessageTextRequest%+v", Alias(*e))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*EditInlineMessageTextRequest) TypeID() uint32 {
	return EditInlineMessageTextRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*EditInlineMessageTextRequest) TypeName() string {
	return "editInlineMessageText"
}

// TypeInfo returns info about TL type.
func (e *EditInlineMessageTextRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "editInlineMessageText",
		ID:   EditInlineMessageTextRequestTypeID,
	}
	if e == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "InlineMessageID",
			SchemaName: "inline_message_id",
		},
		{
			Name:       "ReplyMarkup",
			SchemaName: "reply_markup",
		},
		{
			Name:       "InputMessageContent",
			SchemaName: "input_message_content",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (e *EditInlineMessageTextRequest) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode editInlineMessageText#cd02c1e5 as nil")
	}
	b.PutID(EditInlineMessageTextRequestTypeID)
	return e.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (e *EditInlineMessageTextRequest) EncodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode editInlineMessageText#cd02c1e5 as nil")
	}
	b.PutString(e.InlineMessageID)
	if e.ReplyMarkup == nil {
		return fmt.Errorf("unable to encode editInlineMessageText#cd02c1e5: field reply_markup is nil")
	}
	if err := e.ReplyMarkup.Encode(b); err != nil {
		return fmt.Errorf("unable to encode editInlineMessageText#cd02c1e5: field reply_markup: %w", err)
	}
	if e.InputMessageContent == nil {
		return fmt.Errorf("unable to encode editInlineMessageText#cd02c1e5: field input_message_content is nil")
	}
	if err := e.InputMessageContent.Encode(b); err != nil {
		return fmt.Errorf("unable to encode editInlineMessageText#cd02c1e5: field input_message_content: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (e *EditInlineMessageTextRequest) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode editInlineMessageText#cd02c1e5 to nil")
	}
	if err := b.ConsumeID(EditInlineMessageTextRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode editInlineMessageText#cd02c1e5: %w", err)
	}
	return e.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (e *EditInlineMessageTextRequest) DecodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode editInlineMessageText#cd02c1e5 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode editInlineMessageText#cd02c1e5: field inline_message_id: %w", err)
		}
		e.InlineMessageID = value
	}
	{
		value, err := DecodeReplyMarkup(b)
		if err != nil {
			return fmt.Errorf("unable to decode editInlineMessageText#cd02c1e5: field reply_markup: %w", err)
		}
		e.ReplyMarkup = value
	}
	{
		value, err := DecodeInputMessageContent(b)
		if err != nil {
			return fmt.Errorf("unable to decode editInlineMessageText#cd02c1e5: field input_message_content: %w", err)
		}
		e.InputMessageContent = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (e *EditInlineMessageTextRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if e == nil {
		return fmt.Errorf("can't encode editInlineMessageText#cd02c1e5 as nil")
	}
	b.ObjStart()
	b.PutID("editInlineMessageText")
	b.FieldStart("inline_message_id")
	b.PutString(e.InlineMessageID)
	b.FieldStart("reply_markup")
	if e.ReplyMarkup == nil {
		return fmt.Errorf("unable to encode editInlineMessageText#cd02c1e5: field reply_markup is nil")
	}
	if err := e.ReplyMarkup.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode editInlineMessageText#cd02c1e5: field reply_markup: %w", err)
	}
	b.FieldStart("input_message_content")
	if e.InputMessageContent == nil {
		return fmt.Errorf("unable to encode editInlineMessageText#cd02c1e5: field input_message_content is nil")
	}
	if err := e.InputMessageContent.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode editInlineMessageText#cd02c1e5: field input_message_content: %w", err)
	}
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (e *EditInlineMessageTextRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if e == nil {
		return fmt.Errorf("can't decode editInlineMessageText#cd02c1e5 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("editInlineMessageText"); err != nil {
				return fmt.Errorf("unable to decode editInlineMessageText#cd02c1e5: %w", err)
			}
		case "inline_message_id":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode editInlineMessageText#cd02c1e5: field inline_message_id: %w", err)
			}
			e.InlineMessageID = value
		case "reply_markup":
			value, err := DecodeTDLibJSONReplyMarkup(b)
			if err != nil {
				return fmt.Errorf("unable to decode editInlineMessageText#cd02c1e5: field reply_markup: %w", err)
			}
			e.ReplyMarkup = value
		case "input_message_content":
			value, err := DecodeTDLibJSONInputMessageContent(b)
			if err != nil {
				return fmt.Errorf("unable to decode editInlineMessageText#cd02c1e5: field input_message_content: %w", err)
			}
			e.InputMessageContent = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetInlineMessageID returns value of InlineMessageID field.
func (e *EditInlineMessageTextRequest) GetInlineMessageID() (value string) {
	return e.InlineMessageID
}

// GetReplyMarkup returns value of ReplyMarkup field.
func (e *EditInlineMessageTextRequest) GetReplyMarkup() (value ReplyMarkupClass) {
	return e.ReplyMarkup
}

// GetInputMessageContent returns value of InputMessageContent field.
func (e *EditInlineMessageTextRequest) GetInputMessageContent() (value InputMessageContentClass) {
	return e.InputMessageContent
}

// EditInlineMessageText invokes method editInlineMessageText#cd02c1e5 returning error if any.
func (c *Client) EditInlineMessageText(ctx context.Context, request *EditInlineMessageTextRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}

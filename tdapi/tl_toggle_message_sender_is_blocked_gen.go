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

// ToggleMessageSenderIsBlockedRequest represents TL type `toggleMessageSenderIsBlocked#884f0ed5`.
type ToggleMessageSenderIsBlockedRequest struct {
	// Identifier of a message sender to block/unblock
	SenderID MessageSenderClass
	// New value of is_blocked
	IsBlocked bool
}

// ToggleMessageSenderIsBlockedRequestTypeID is TL type id of ToggleMessageSenderIsBlockedRequest.
const ToggleMessageSenderIsBlockedRequestTypeID = 0x884f0ed5

// Ensuring interfaces in compile-time for ToggleMessageSenderIsBlockedRequest.
var (
	_ bin.Encoder     = &ToggleMessageSenderIsBlockedRequest{}
	_ bin.Decoder     = &ToggleMessageSenderIsBlockedRequest{}
	_ bin.BareEncoder = &ToggleMessageSenderIsBlockedRequest{}
	_ bin.BareDecoder = &ToggleMessageSenderIsBlockedRequest{}
)

func (t *ToggleMessageSenderIsBlockedRequest) Zero() bool {
	if t == nil {
		return true
	}
	if !(t.SenderID == nil) {
		return false
	}
	if !(t.IsBlocked == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (t *ToggleMessageSenderIsBlockedRequest) String() string {
	if t == nil {
		return "ToggleMessageSenderIsBlockedRequest(nil)"
	}
	type Alias ToggleMessageSenderIsBlockedRequest
	return fmt.Sprintf("ToggleMessageSenderIsBlockedRequest%+v", Alias(*t))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ToggleMessageSenderIsBlockedRequest) TypeID() uint32 {
	return ToggleMessageSenderIsBlockedRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ToggleMessageSenderIsBlockedRequest) TypeName() string {
	return "toggleMessageSenderIsBlocked"
}

// TypeInfo returns info about TL type.
func (t *ToggleMessageSenderIsBlockedRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "toggleMessageSenderIsBlocked",
		ID:   ToggleMessageSenderIsBlockedRequestTypeID,
	}
	if t == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "SenderID",
			SchemaName: "sender_id",
		},
		{
			Name:       "IsBlocked",
			SchemaName: "is_blocked",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (t *ToggleMessageSenderIsBlockedRequest) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode toggleMessageSenderIsBlocked#884f0ed5 as nil")
	}
	b.PutID(ToggleMessageSenderIsBlockedRequestTypeID)
	return t.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (t *ToggleMessageSenderIsBlockedRequest) EncodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode toggleMessageSenderIsBlocked#884f0ed5 as nil")
	}
	if t.SenderID == nil {
		return fmt.Errorf("unable to encode toggleMessageSenderIsBlocked#884f0ed5: field sender_id is nil")
	}
	if err := t.SenderID.Encode(b); err != nil {
		return fmt.Errorf("unable to encode toggleMessageSenderIsBlocked#884f0ed5: field sender_id: %w", err)
	}
	b.PutBool(t.IsBlocked)
	return nil
}

// Decode implements bin.Decoder.
func (t *ToggleMessageSenderIsBlockedRequest) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode toggleMessageSenderIsBlocked#884f0ed5 to nil")
	}
	if err := b.ConsumeID(ToggleMessageSenderIsBlockedRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode toggleMessageSenderIsBlocked#884f0ed5: %w", err)
	}
	return t.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (t *ToggleMessageSenderIsBlockedRequest) DecodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode toggleMessageSenderIsBlocked#884f0ed5 to nil")
	}
	{
		value, err := DecodeMessageSender(b)
		if err != nil {
			return fmt.Errorf("unable to decode toggleMessageSenderIsBlocked#884f0ed5: field sender_id: %w", err)
		}
		t.SenderID = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode toggleMessageSenderIsBlocked#884f0ed5: field is_blocked: %w", err)
		}
		t.IsBlocked = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (t *ToggleMessageSenderIsBlockedRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if t == nil {
		return fmt.Errorf("can't encode toggleMessageSenderIsBlocked#884f0ed5 as nil")
	}
	b.ObjStart()
	b.PutID("toggleMessageSenderIsBlocked")
	b.FieldStart("sender_id")
	if t.SenderID == nil {
		return fmt.Errorf("unable to encode toggleMessageSenderIsBlocked#884f0ed5: field sender_id is nil")
	}
	if err := t.SenderID.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode toggleMessageSenderIsBlocked#884f0ed5: field sender_id: %w", err)
	}
	b.FieldStart("is_blocked")
	b.PutBool(t.IsBlocked)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (t *ToggleMessageSenderIsBlockedRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if t == nil {
		return fmt.Errorf("can't decode toggleMessageSenderIsBlocked#884f0ed5 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("toggleMessageSenderIsBlocked"); err != nil {
				return fmt.Errorf("unable to decode toggleMessageSenderIsBlocked#884f0ed5: %w", err)
			}
		case "sender_id":
			value, err := DecodeTDLibJSONMessageSender(b)
			if err != nil {
				return fmt.Errorf("unable to decode toggleMessageSenderIsBlocked#884f0ed5: field sender_id: %w", err)
			}
			t.SenderID = value
		case "is_blocked":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode toggleMessageSenderIsBlocked#884f0ed5: field is_blocked: %w", err)
			}
			t.IsBlocked = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetSenderID returns value of SenderID field.
func (t *ToggleMessageSenderIsBlockedRequest) GetSenderID() (value MessageSenderClass) {
	return t.SenderID
}

// GetIsBlocked returns value of IsBlocked field.
func (t *ToggleMessageSenderIsBlockedRequest) GetIsBlocked() (value bool) {
	return t.IsBlocked
}

// ToggleMessageSenderIsBlocked invokes method toggleMessageSenderIsBlocked#884f0ed5 returning error if any.
func (c *Client) ToggleMessageSenderIsBlocked(ctx context.Context, request *ToggleMessageSenderIsBlockedRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}

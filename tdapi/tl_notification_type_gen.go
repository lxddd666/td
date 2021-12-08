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

// NotificationTypeNewMessage represents TL type `notificationTypeNewMessage#70691637`.
type NotificationTypeNewMessage struct {
	// The message
	Message Message
}

// NotificationTypeNewMessageTypeID is TL type id of NotificationTypeNewMessage.
const NotificationTypeNewMessageTypeID = 0x70691637

// construct implements constructor of NotificationTypeClass.
func (n NotificationTypeNewMessage) construct() NotificationTypeClass { return &n }

// Ensuring interfaces in compile-time for NotificationTypeNewMessage.
var (
	_ bin.Encoder     = &NotificationTypeNewMessage{}
	_ bin.Decoder     = &NotificationTypeNewMessage{}
	_ bin.BareEncoder = &NotificationTypeNewMessage{}
	_ bin.BareDecoder = &NotificationTypeNewMessage{}

	_ NotificationTypeClass = &NotificationTypeNewMessage{}
)

func (n *NotificationTypeNewMessage) Zero() bool {
	if n == nil {
		return true
	}
	if !(n.Message.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (n *NotificationTypeNewMessage) String() string {
	if n == nil {
		return "NotificationTypeNewMessage(nil)"
	}
	type Alias NotificationTypeNewMessage
	return fmt.Sprintf("NotificationTypeNewMessage%+v", Alias(*n))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*NotificationTypeNewMessage) TypeID() uint32 {
	return NotificationTypeNewMessageTypeID
}

// TypeName returns name of type in TL schema.
func (*NotificationTypeNewMessage) TypeName() string {
	return "notificationTypeNewMessage"
}

// TypeInfo returns info about TL type.
func (n *NotificationTypeNewMessage) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "notificationTypeNewMessage",
		ID:   NotificationTypeNewMessageTypeID,
	}
	if n == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Message",
			SchemaName: "message",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (n *NotificationTypeNewMessage) Encode(b *bin.Buffer) error {
	if n == nil {
		return fmt.Errorf("can't encode notificationTypeNewMessage#70691637 as nil")
	}
	b.PutID(NotificationTypeNewMessageTypeID)
	return n.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (n *NotificationTypeNewMessage) EncodeBare(b *bin.Buffer) error {
	if n == nil {
		return fmt.Errorf("can't encode notificationTypeNewMessage#70691637 as nil")
	}
	if err := n.Message.Encode(b); err != nil {
		return fmt.Errorf("unable to encode notificationTypeNewMessage#70691637: field message: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (n *NotificationTypeNewMessage) Decode(b *bin.Buffer) error {
	if n == nil {
		return fmt.Errorf("can't decode notificationTypeNewMessage#70691637 to nil")
	}
	if err := b.ConsumeID(NotificationTypeNewMessageTypeID); err != nil {
		return fmt.Errorf("unable to decode notificationTypeNewMessage#70691637: %w", err)
	}
	return n.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (n *NotificationTypeNewMessage) DecodeBare(b *bin.Buffer) error {
	if n == nil {
		return fmt.Errorf("can't decode notificationTypeNewMessage#70691637 to nil")
	}
	{
		if err := n.Message.Decode(b); err != nil {
			return fmt.Errorf("unable to decode notificationTypeNewMessage#70691637: field message: %w", err)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (n *NotificationTypeNewMessage) EncodeTDLibJSON(b tdjson.Encoder) error {
	if n == nil {
		return fmt.Errorf("can't encode notificationTypeNewMessage#70691637 as nil")
	}
	b.ObjStart()
	b.PutID("notificationTypeNewMessage")
	b.FieldStart("message")
	if err := n.Message.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode notificationTypeNewMessage#70691637: field message: %w", err)
	}
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (n *NotificationTypeNewMessage) DecodeTDLibJSON(b tdjson.Decoder) error {
	if n == nil {
		return fmt.Errorf("can't decode notificationTypeNewMessage#70691637 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("notificationTypeNewMessage"); err != nil {
				return fmt.Errorf("unable to decode notificationTypeNewMessage#70691637: %w", err)
			}
		case "message":
			if err := n.Message.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode notificationTypeNewMessage#70691637: field message: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetMessage returns value of Message field.
func (n *NotificationTypeNewMessage) GetMessage() (value Message) {
	return n.Message
}

// NotificationTypeNewSecretChat represents TL type `notificationTypeNewSecretChat#4771c6b0`.
type NotificationTypeNewSecretChat struct {
}

// NotificationTypeNewSecretChatTypeID is TL type id of NotificationTypeNewSecretChat.
const NotificationTypeNewSecretChatTypeID = 0x4771c6b0

// construct implements constructor of NotificationTypeClass.
func (n NotificationTypeNewSecretChat) construct() NotificationTypeClass { return &n }

// Ensuring interfaces in compile-time for NotificationTypeNewSecretChat.
var (
	_ bin.Encoder     = &NotificationTypeNewSecretChat{}
	_ bin.Decoder     = &NotificationTypeNewSecretChat{}
	_ bin.BareEncoder = &NotificationTypeNewSecretChat{}
	_ bin.BareDecoder = &NotificationTypeNewSecretChat{}

	_ NotificationTypeClass = &NotificationTypeNewSecretChat{}
)

func (n *NotificationTypeNewSecretChat) Zero() bool {
	if n == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (n *NotificationTypeNewSecretChat) String() string {
	if n == nil {
		return "NotificationTypeNewSecretChat(nil)"
	}
	type Alias NotificationTypeNewSecretChat
	return fmt.Sprintf("NotificationTypeNewSecretChat%+v", Alias(*n))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*NotificationTypeNewSecretChat) TypeID() uint32 {
	return NotificationTypeNewSecretChatTypeID
}

// TypeName returns name of type in TL schema.
func (*NotificationTypeNewSecretChat) TypeName() string {
	return "notificationTypeNewSecretChat"
}

// TypeInfo returns info about TL type.
func (n *NotificationTypeNewSecretChat) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "notificationTypeNewSecretChat",
		ID:   NotificationTypeNewSecretChatTypeID,
	}
	if n == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (n *NotificationTypeNewSecretChat) Encode(b *bin.Buffer) error {
	if n == nil {
		return fmt.Errorf("can't encode notificationTypeNewSecretChat#4771c6b0 as nil")
	}
	b.PutID(NotificationTypeNewSecretChatTypeID)
	return n.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (n *NotificationTypeNewSecretChat) EncodeBare(b *bin.Buffer) error {
	if n == nil {
		return fmt.Errorf("can't encode notificationTypeNewSecretChat#4771c6b0 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (n *NotificationTypeNewSecretChat) Decode(b *bin.Buffer) error {
	if n == nil {
		return fmt.Errorf("can't decode notificationTypeNewSecretChat#4771c6b0 to nil")
	}
	if err := b.ConsumeID(NotificationTypeNewSecretChatTypeID); err != nil {
		return fmt.Errorf("unable to decode notificationTypeNewSecretChat#4771c6b0: %w", err)
	}
	return n.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (n *NotificationTypeNewSecretChat) DecodeBare(b *bin.Buffer) error {
	if n == nil {
		return fmt.Errorf("can't decode notificationTypeNewSecretChat#4771c6b0 to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (n *NotificationTypeNewSecretChat) EncodeTDLibJSON(b tdjson.Encoder) error {
	if n == nil {
		return fmt.Errorf("can't encode notificationTypeNewSecretChat#4771c6b0 as nil")
	}
	b.ObjStart()
	b.PutID("notificationTypeNewSecretChat")
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (n *NotificationTypeNewSecretChat) DecodeTDLibJSON(b tdjson.Decoder) error {
	if n == nil {
		return fmt.Errorf("can't decode notificationTypeNewSecretChat#4771c6b0 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("notificationTypeNewSecretChat"); err != nil {
				return fmt.Errorf("unable to decode notificationTypeNewSecretChat#4771c6b0: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// NotificationTypeNewCall represents TL type `notificationTypeNewCall#66164179`.
type NotificationTypeNewCall struct {
	// Call identifier
	CallID int32
}

// NotificationTypeNewCallTypeID is TL type id of NotificationTypeNewCall.
const NotificationTypeNewCallTypeID = 0x66164179

// construct implements constructor of NotificationTypeClass.
func (n NotificationTypeNewCall) construct() NotificationTypeClass { return &n }

// Ensuring interfaces in compile-time for NotificationTypeNewCall.
var (
	_ bin.Encoder     = &NotificationTypeNewCall{}
	_ bin.Decoder     = &NotificationTypeNewCall{}
	_ bin.BareEncoder = &NotificationTypeNewCall{}
	_ bin.BareDecoder = &NotificationTypeNewCall{}

	_ NotificationTypeClass = &NotificationTypeNewCall{}
)

func (n *NotificationTypeNewCall) Zero() bool {
	if n == nil {
		return true
	}
	if !(n.CallID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (n *NotificationTypeNewCall) String() string {
	if n == nil {
		return "NotificationTypeNewCall(nil)"
	}
	type Alias NotificationTypeNewCall
	return fmt.Sprintf("NotificationTypeNewCall%+v", Alias(*n))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*NotificationTypeNewCall) TypeID() uint32 {
	return NotificationTypeNewCallTypeID
}

// TypeName returns name of type in TL schema.
func (*NotificationTypeNewCall) TypeName() string {
	return "notificationTypeNewCall"
}

// TypeInfo returns info about TL type.
func (n *NotificationTypeNewCall) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "notificationTypeNewCall",
		ID:   NotificationTypeNewCallTypeID,
	}
	if n == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "CallID",
			SchemaName: "call_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (n *NotificationTypeNewCall) Encode(b *bin.Buffer) error {
	if n == nil {
		return fmt.Errorf("can't encode notificationTypeNewCall#66164179 as nil")
	}
	b.PutID(NotificationTypeNewCallTypeID)
	return n.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (n *NotificationTypeNewCall) EncodeBare(b *bin.Buffer) error {
	if n == nil {
		return fmt.Errorf("can't encode notificationTypeNewCall#66164179 as nil")
	}
	b.PutInt32(n.CallID)
	return nil
}

// Decode implements bin.Decoder.
func (n *NotificationTypeNewCall) Decode(b *bin.Buffer) error {
	if n == nil {
		return fmt.Errorf("can't decode notificationTypeNewCall#66164179 to nil")
	}
	if err := b.ConsumeID(NotificationTypeNewCallTypeID); err != nil {
		return fmt.Errorf("unable to decode notificationTypeNewCall#66164179: %w", err)
	}
	return n.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (n *NotificationTypeNewCall) DecodeBare(b *bin.Buffer) error {
	if n == nil {
		return fmt.Errorf("can't decode notificationTypeNewCall#66164179 to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode notificationTypeNewCall#66164179: field call_id: %w", err)
		}
		n.CallID = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (n *NotificationTypeNewCall) EncodeTDLibJSON(b tdjson.Encoder) error {
	if n == nil {
		return fmt.Errorf("can't encode notificationTypeNewCall#66164179 as nil")
	}
	b.ObjStart()
	b.PutID("notificationTypeNewCall")
	b.FieldStart("call_id")
	b.PutInt32(n.CallID)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (n *NotificationTypeNewCall) DecodeTDLibJSON(b tdjson.Decoder) error {
	if n == nil {
		return fmt.Errorf("can't decode notificationTypeNewCall#66164179 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("notificationTypeNewCall"); err != nil {
				return fmt.Errorf("unable to decode notificationTypeNewCall#66164179: %w", err)
			}
		case "call_id":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode notificationTypeNewCall#66164179: field call_id: %w", err)
			}
			n.CallID = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetCallID returns value of CallID field.
func (n *NotificationTypeNewCall) GetCallID() (value int32) {
	return n.CallID
}

// NotificationTypeNewPushMessage represents TL type `notificationTypeNewPushMessage#d5949e32`.
type NotificationTypeNewPushMessage struct {
	// The message identifier. The message will not be available in the chat history, but the
	// ID can be used in viewMessages, or as reply_to_message_id
	MessageID int64
	// Identifier of the sender of the message. Corresponding user or chat may be
	// inaccessible
	SenderID MessageSenderClass
	// Name of the sender
	SenderName string
	// True, if the message is outgoing
	IsOutgoing bool
	// Push message content
	Content PushMessageContentClass
}

// NotificationTypeNewPushMessageTypeID is TL type id of NotificationTypeNewPushMessage.
const NotificationTypeNewPushMessageTypeID = 0xd5949e32

// construct implements constructor of NotificationTypeClass.
func (n NotificationTypeNewPushMessage) construct() NotificationTypeClass { return &n }

// Ensuring interfaces in compile-time for NotificationTypeNewPushMessage.
var (
	_ bin.Encoder     = &NotificationTypeNewPushMessage{}
	_ bin.Decoder     = &NotificationTypeNewPushMessage{}
	_ bin.BareEncoder = &NotificationTypeNewPushMessage{}
	_ bin.BareDecoder = &NotificationTypeNewPushMessage{}

	_ NotificationTypeClass = &NotificationTypeNewPushMessage{}
)

func (n *NotificationTypeNewPushMessage) Zero() bool {
	if n == nil {
		return true
	}
	if !(n.MessageID == 0) {
		return false
	}
	if !(n.SenderID == nil) {
		return false
	}
	if !(n.SenderName == "") {
		return false
	}
	if !(n.IsOutgoing == false) {
		return false
	}
	if !(n.Content == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (n *NotificationTypeNewPushMessage) String() string {
	if n == nil {
		return "NotificationTypeNewPushMessage(nil)"
	}
	type Alias NotificationTypeNewPushMessage
	return fmt.Sprintf("NotificationTypeNewPushMessage%+v", Alias(*n))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*NotificationTypeNewPushMessage) TypeID() uint32 {
	return NotificationTypeNewPushMessageTypeID
}

// TypeName returns name of type in TL schema.
func (*NotificationTypeNewPushMessage) TypeName() string {
	return "notificationTypeNewPushMessage"
}

// TypeInfo returns info about TL type.
func (n *NotificationTypeNewPushMessage) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "notificationTypeNewPushMessage",
		ID:   NotificationTypeNewPushMessageTypeID,
	}
	if n == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "MessageID",
			SchemaName: "message_id",
		},
		{
			Name:       "SenderID",
			SchemaName: "sender_id",
		},
		{
			Name:       "SenderName",
			SchemaName: "sender_name",
		},
		{
			Name:       "IsOutgoing",
			SchemaName: "is_outgoing",
		},
		{
			Name:       "Content",
			SchemaName: "content",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (n *NotificationTypeNewPushMessage) Encode(b *bin.Buffer) error {
	if n == nil {
		return fmt.Errorf("can't encode notificationTypeNewPushMessage#d5949e32 as nil")
	}
	b.PutID(NotificationTypeNewPushMessageTypeID)
	return n.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (n *NotificationTypeNewPushMessage) EncodeBare(b *bin.Buffer) error {
	if n == nil {
		return fmt.Errorf("can't encode notificationTypeNewPushMessage#d5949e32 as nil")
	}
	b.PutInt53(n.MessageID)
	if n.SenderID == nil {
		return fmt.Errorf("unable to encode notificationTypeNewPushMessage#d5949e32: field sender_id is nil")
	}
	if err := n.SenderID.Encode(b); err != nil {
		return fmt.Errorf("unable to encode notificationTypeNewPushMessage#d5949e32: field sender_id: %w", err)
	}
	b.PutString(n.SenderName)
	b.PutBool(n.IsOutgoing)
	if n.Content == nil {
		return fmt.Errorf("unable to encode notificationTypeNewPushMessage#d5949e32: field content is nil")
	}
	if err := n.Content.Encode(b); err != nil {
		return fmt.Errorf("unable to encode notificationTypeNewPushMessage#d5949e32: field content: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (n *NotificationTypeNewPushMessage) Decode(b *bin.Buffer) error {
	if n == nil {
		return fmt.Errorf("can't decode notificationTypeNewPushMessage#d5949e32 to nil")
	}
	if err := b.ConsumeID(NotificationTypeNewPushMessageTypeID); err != nil {
		return fmt.Errorf("unable to decode notificationTypeNewPushMessage#d5949e32: %w", err)
	}
	return n.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (n *NotificationTypeNewPushMessage) DecodeBare(b *bin.Buffer) error {
	if n == nil {
		return fmt.Errorf("can't decode notificationTypeNewPushMessage#d5949e32 to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode notificationTypeNewPushMessage#d5949e32: field message_id: %w", err)
		}
		n.MessageID = value
	}
	{
		value, err := DecodeMessageSender(b)
		if err != nil {
			return fmt.Errorf("unable to decode notificationTypeNewPushMessage#d5949e32: field sender_id: %w", err)
		}
		n.SenderID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode notificationTypeNewPushMessage#d5949e32: field sender_name: %w", err)
		}
		n.SenderName = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode notificationTypeNewPushMessage#d5949e32: field is_outgoing: %w", err)
		}
		n.IsOutgoing = value
	}
	{
		value, err := DecodePushMessageContent(b)
		if err != nil {
			return fmt.Errorf("unable to decode notificationTypeNewPushMessage#d5949e32: field content: %w", err)
		}
		n.Content = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (n *NotificationTypeNewPushMessage) EncodeTDLibJSON(b tdjson.Encoder) error {
	if n == nil {
		return fmt.Errorf("can't encode notificationTypeNewPushMessage#d5949e32 as nil")
	}
	b.ObjStart()
	b.PutID("notificationTypeNewPushMessage")
	b.FieldStart("message_id")
	b.PutInt53(n.MessageID)
	b.FieldStart("sender_id")
	if n.SenderID == nil {
		return fmt.Errorf("unable to encode notificationTypeNewPushMessage#d5949e32: field sender_id is nil")
	}
	if err := n.SenderID.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode notificationTypeNewPushMessage#d5949e32: field sender_id: %w", err)
	}
	b.FieldStart("sender_name")
	b.PutString(n.SenderName)
	b.FieldStart("is_outgoing")
	b.PutBool(n.IsOutgoing)
	b.FieldStart("content")
	if n.Content == nil {
		return fmt.Errorf("unable to encode notificationTypeNewPushMessage#d5949e32: field content is nil")
	}
	if err := n.Content.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode notificationTypeNewPushMessage#d5949e32: field content: %w", err)
	}
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (n *NotificationTypeNewPushMessage) DecodeTDLibJSON(b tdjson.Decoder) error {
	if n == nil {
		return fmt.Errorf("can't decode notificationTypeNewPushMessage#d5949e32 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("notificationTypeNewPushMessage"); err != nil {
				return fmt.Errorf("unable to decode notificationTypeNewPushMessage#d5949e32: %w", err)
			}
		case "message_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode notificationTypeNewPushMessage#d5949e32: field message_id: %w", err)
			}
			n.MessageID = value
		case "sender_id":
			value, err := DecodeTDLibJSONMessageSender(b)
			if err != nil {
				return fmt.Errorf("unable to decode notificationTypeNewPushMessage#d5949e32: field sender_id: %w", err)
			}
			n.SenderID = value
		case "sender_name":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode notificationTypeNewPushMessage#d5949e32: field sender_name: %w", err)
			}
			n.SenderName = value
		case "is_outgoing":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode notificationTypeNewPushMessage#d5949e32: field is_outgoing: %w", err)
			}
			n.IsOutgoing = value
		case "content":
			value, err := DecodeTDLibJSONPushMessageContent(b)
			if err != nil {
				return fmt.Errorf("unable to decode notificationTypeNewPushMessage#d5949e32: field content: %w", err)
			}
			n.Content = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetMessageID returns value of MessageID field.
func (n *NotificationTypeNewPushMessage) GetMessageID() (value int64) {
	return n.MessageID
}

// GetSenderID returns value of SenderID field.
func (n *NotificationTypeNewPushMessage) GetSenderID() (value MessageSenderClass) {
	return n.SenderID
}

// GetSenderName returns value of SenderName field.
func (n *NotificationTypeNewPushMessage) GetSenderName() (value string) {
	return n.SenderName
}

// GetIsOutgoing returns value of IsOutgoing field.
func (n *NotificationTypeNewPushMessage) GetIsOutgoing() (value bool) {
	return n.IsOutgoing
}

// GetContent returns value of Content field.
func (n *NotificationTypeNewPushMessage) GetContent() (value PushMessageContentClass) {
	return n.Content
}

// NotificationTypeClassName is schema name of NotificationTypeClass.
const NotificationTypeClassName = "NotificationType"

// NotificationTypeClass represents NotificationType generic type.
//
// Example:
//  g, err := tdapi.DecodeNotificationType(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tdapi.NotificationTypeNewMessage: // notificationTypeNewMessage#70691637
//  case *tdapi.NotificationTypeNewSecretChat: // notificationTypeNewSecretChat#4771c6b0
//  case *tdapi.NotificationTypeNewCall: // notificationTypeNewCall#66164179
//  case *tdapi.NotificationTypeNewPushMessage: // notificationTypeNewPushMessage#d5949e32
//  default: panic(v)
//  }
type NotificationTypeClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() NotificationTypeClass

	// TypeID returns type id in TL schema.
	//
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// TypeName returns name of type in TL schema.
	TypeName() string
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool

	EncodeTDLibJSON(b tdjson.Encoder) error
	DecodeTDLibJSON(b tdjson.Decoder) error
}

// DecodeNotificationType implements binary de-serialization for NotificationTypeClass.
func DecodeNotificationType(buf *bin.Buffer) (NotificationTypeClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case NotificationTypeNewMessageTypeID:
		// Decoding notificationTypeNewMessage#70691637.
		v := NotificationTypeNewMessage{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode NotificationTypeClass: %w", err)
		}
		return &v, nil
	case NotificationTypeNewSecretChatTypeID:
		// Decoding notificationTypeNewSecretChat#4771c6b0.
		v := NotificationTypeNewSecretChat{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode NotificationTypeClass: %w", err)
		}
		return &v, nil
	case NotificationTypeNewCallTypeID:
		// Decoding notificationTypeNewCall#66164179.
		v := NotificationTypeNewCall{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode NotificationTypeClass: %w", err)
		}
		return &v, nil
	case NotificationTypeNewPushMessageTypeID:
		// Decoding notificationTypeNewPushMessage#d5949e32.
		v := NotificationTypeNewPushMessage{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode NotificationTypeClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode NotificationTypeClass: %w", bin.NewUnexpectedID(id))
	}
}

// DecodeTDLibJSONNotificationType implements binary de-serialization for NotificationTypeClass.
func DecodeTDLibJSONNotificationType(buf tdjson.Decoder) (NotificationTypeClass, error) {
	id, err := buf.FindTypeID()
	if err != nil {
		return nil, err
	}
	switch id {
	case "notificationTypeNewMessage":
		// Decoding notificationTypeNewMessage#70691637.
		v := NotificationTypeNewMessage{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode NotificationTypeClass: %w", err)
		}
		return &v, nil
	case "notificationTypeNewSecretChat":
		// Decoding notificationTypeNewSecretChat#4771c6b0.
		v := NotificationTypeNewSecretChat{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode NotificationTypeClass: %w", err)
		}
		return &v, nil
	case "notificationTypeNewCall":
		// Decoding notificationTypeNewCall#66164179.
		v := NotificationTypeNewCall{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode NotificationTypeClass: %w", err)
		}
		return &v, nil
	case "notificationTypeNewPushMessage":
		// Decoding notificationTypeNewPushMessage#d5949e32.
		v := NotificationTypeNewPushMessage{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode NotificationTypeClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode NotificationTypeClass: %w", tdjson.NewUnexpectedID(id))
	}
}

// NotificationType boxes the NotificationTypeClass providing a helper.
type NotificationTypeBox struct {
	NotificationType NotificationTypeClass
}

// Decode implements bin.Decoder for NotificationTypeBox.
func (b *NotificationTypeBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode NotificationTypeBox to nil")
	}
	v, err := DecodeNotificationType(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.NotificationType = v
	return nil
}

// Encode implements bin.Encode for NotificationTypeBox.
func (b *NotificationTypeBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.NotificationType == nil {
		return fmt.Errorf("unable to encode NotificationTypeClass as nil")
	}
	return b.NotificationType.Encode(buf)
}

// DecodeTDLibJSON implements bin.Decoder for NotificationTypeBox.
func (b *NotificationTypeBox) DecodeTDLibJSON(buf tdjson.Decoder) error {
	if b == nil {
		return fmt.Errorf("unable to decode NotificationTypeBox to nil")
	}
	v, err := DecodeTDLibJSONNotificationType(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.NotificationType = v
	return nil
}

// EncodeTDLibJSON implements bin.Encode for NotificationTypeBox.
func (b *NotificationTypeBox) EncodeTDLibJSON(buf tdjson.Encoder) error {
	if b == nil || b.NotificationType == nil {
		return fmt.Errorf("unable to encode NotificationTypeClass as nil")
	}
	return b.NotificationType.EncodeTDLibJSON(buf)
}

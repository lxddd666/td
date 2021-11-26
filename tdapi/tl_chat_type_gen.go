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

// ChatTypePrivate represents TL type `chatTypePrivate#5e1e6374`.
type ChatTypePrivate struct {
	// User identifier
	UserID int64
}

// ChatTypePrivateTypeID is TL type id of ChatTypePrivate.
const ChatTypePrivateTypeID = 0x5e1e6374

// construct implements constructor of ChatTypeClass.
func (c ChatTypePrivate) construct() ChatTypeClass { return &c }

// Ensuring interfaces in compile-time for ChatTypePrivate.
var (
	_ bin.Encoder     = &ChatTypePrivate{}
	_ bin.Decoder     = &ChatTypePrivate{}
	_ bin.BareEncoder = &ChatTypePrivate{}
	_ bin.BareDecoder = &ChatTypePrivate{}

	_ ChatTypeClass = &ChatTypePrivate{}
)

func (c *ChatTypePrivate) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.UserID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChatTypePrivate) String() string {
	if c == nil {
		return "ChatTypePrivate(nil)"
	}
	type Alias ChatTypePrivate
	return fmt.Sprintf("ChatTypePrivate%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChatTypePrivate) TypeID() uint32 {
	return ChatTypePrivateTypeID
}

// TypeName returns name of type in TL schema.
func (*ChatTypePrivate) TypeName() string {
	return "chatTypePrivate"
}

// TypeInfo returns info about TL type.
func (c *ChatTypePrivate) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "chatTypePrivate",
		ID:   ChatTypePrivateTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "UserID",
			SchemaName: "user_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *ChatTypePrivate) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatTypePrivate#5e1e6374 as nil")
	}
	b.PutID(ChatTypePrivateTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ChatTypePrivate) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatTypePrivate#5e1e6374 as nil")
	}
	b.PutLong(c.UserID)
	return nil
}

// Decode implements bin.Decoder.
func (c *ChatTypePrivate) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatTypePrivate#5e1e6374 to nil")
	}
	if err := b.ConsumeID(ChatTypePrivateTypeID); err != nil {
		return fmt.Errorf("unable to decode chatTypePrivate#5e1e6374: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ChatTypePrivate) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatTypePrivate#5e1e6374 to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode chatTypePrivate#5e1e6374: field user_id: %w", err)
		}
		c.UserID = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (c *ChatTypePrivate) EncodeTDLibJSON(b tdjson.Encoder) error {
	if c == nil {
		return fmt.Errorf("can't encode chatTypePrivate#5e1e6374 as nil")
	}
	b.ObjStart()
	b.PutID("chatTypePrivate")
	b.FieldStart("user_id")
	b.PutLong(c.UserID)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (c *ChatTypePrivate) DecodeTDLibJSON(b tdjson.Decoder) error {
	if c == nil {
		return fmt.Errorf("can't decode chatTypePrivate#5e1e6374 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("chatTypePrivate"); err != nil {
				return fmt.Errorf("unable to decode chatTypePrivate#5e1e6374: %w", err)
			}
		case "user_id":
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode chatTypePrivate#5e1e6374: field user_id: %w", err)
			}
			c.UserID = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetUserID returns value of UserID field.
func (c *ChatTypePrivate) GetUserID() (value int64) {
	return c.UserID
}

// ChatTypeBasicGroup represents TL type `chatTypeBasicGroup#3a0c4c5c`.
type ChatTypeBasicGroup struct {
	// Basic group identifier
	BasicGroupID int64
}

// ChatTypeBasicGroupTypeID is TL type id of ChatTypeBasicGroup.
const ChatTypeBasicGroupTypeID = 0x3a0c4c5c

// construct implements constructor of ChatTypeClass.
func (c ChatTypeBasicGroup) construct() ChatTypeClass { return &c }

// Ensuring interfaces in compile-time for ChatTypeBasicGroup.
var (
	_ bin.Encoder     = &ChatTypeBasicGroup{}
	_ bin.Decoder     = &ChatTypeBasicGroup{}
	_ bin.BareEncoder = &ChatTypeBasicGroup{}
	_ bin.BareDecoder = &ChatTypeBasicGroup{}

	_ ChatTypeClass = &ChatTypeBasicGroup{}
)

func (c *ChatTypeBasicGroup) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.BasicGroupID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChatTypeBasicGroup) String() string {
	if c == nil {
		return "ChatTypeBasicGroup(nil)"
	}
	type Alias ChatTypeBasicGroup
	return fmt.Sprintf("ChatTypeBasicGroup%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChatTypeBasicGroup) TypeID() uint32 {
	return ChatTypeBasicGroupTypeID
}

// TypeName returns name of type in TL schema.
func (*ChatTypeBasicGroup) TypeName() string {
	return "chatTypeBasicGroup"
}

// TypeInfo returns info about TL type.
func (c *ChatTypeBasicGroup) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "chatTypeBasicGroup",
		ID:   ChatTypeBasicGroupTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "BasicGroupID",
			SchemaName: "basic_group_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *ChatTypeBasicGroup) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatTypeBasicGroup#3a0c4c5c as nil")
	}
	b.PutID(ChatTypeBasicGroupTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ChatTypeBasicGroup) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatTypeBasicGroup#3a0c4c5c as nil")
	}
	b.PutLong(c.BasicGroupID)
	return nil
}

// Decode implements bin.Decoder.
func (c *ChatTypeBasicGroup) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatTypeBasicGroup#3a0c4c5c to nil")
	}
	if err := b.ConsumeID(ChatTypeBasicGroupTypeID); err != nil {
		return fmt.Errorf("unable to decode chatTypeBasicGroup#3a0c4c5c: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ChatTypeBasicGroup) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatTypeBasicGroup#3a0c4c5c to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode chatTypeBasicGroup#3a0c4c5c: field basic_group_id: %w", err)
		}
		c.BasicGroupID = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (c *ChatTypeBasicGroup) EncodeTDLibJSON(b tdjson.Encoder) error {
	if c == nil {
		return fmt.Errorf("can't encode chatTypeBasicGroup#3a0c4c5c as nil")
	}
	b.ObjStart()
	b.PutID("chatTypeBasicGroup")
	b.FieldStart("basic_group_id")
	b.PutLong(c.BasicGroupID)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (c *ChatTypeBasicGroup) DecodeTDLibJSON(b tdjson.Decoder) error {
	if c == nil {
		return fmt.Errorf("can't decode chatTypeBasicGroup#3a0c4c5c to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("chatTypeBasicGroup"); err != nil {
				return fmt.Errorf("unable to decode chatTypeBasicGroup#3a0c4c5c: %w", err)
			}
		case "basic_group_id":
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode chatTypeBasicGroup#3a0c4c5c: field basic_group_id: %w", err)
			}
			c.BasicGroupID = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetBasicGroupID returns value of BasicGroupID field.
func (c *ChatTypeBasicGroup) GetBasicGroupID() (value int64) {
	return c.BasicGroupID
}

// ChatTypeSupergroup represents TL type `chatTypeSupergroup#a83a5a6a`.
type ChatTypeSupergroup struct {
	// Supergroup or channel identifier
	SupergroupID int64
	// True, if the supergroup is a channel
	IsChannel bool
}

// ChatTypeSupergroupTypeID is TL type id of ChatTypeSupergroup.
const ChatTypeSupergroupTypeID = 0xa83a5a6a

// construct implements constructor of ChatTypeClass.
func (c ChatTypeSupergroup) construct() ChatTypeClass { return &c }

// Ensuring interfaces in compile-time for ChatTypeSupergroup.
var (
	_ bin.Encoder     = &ChatTypeSupergroup{}
	_ bin.Decoder     = &ChatTypeSupergroup{}
	_ bin.BareEncoder = &ChatTypeSupergroup{}
	_ bin.BareDecoder = &ChatTypeSupergroup{}

	_ ChatTypeClass = &ChatTypeSupergroup{}
)

func (c *ChatTypeSupergroup) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.SupergroupID == 0) {
		return false
	}
	if !(c.IsChannel == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChatTypeSupergroup) String() string {
	if c == nil {
		return "ChatTypeSupergroup(nil)"
	}
	type Alias ChatTypeSupergroup
	return fmt.Sprintf("ChatTypeSupergroup%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChatTypeSupergroup) TypeID() uint32 {
	return ChatTypeSupergroupTypeID
}

// TypeName returns name of type in TL schema.
func (*ChatTypeSupergroup) TypeName() string {
	return "chatTypeSupergroup"
}

// TypeInfo returns info about TL type.
func (c *ChatTypeSupergroup) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "chatTypeSupergroup",
		ID:   ChatTypeSupergroupTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "SupergroupID",
			SchemaName: "supergroup_id",
		},
		{
			Name:       "IsChannel",
			SchemaName: "is_channel",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *ChatTypeSupergroup) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatTypeSupergroup#a83a5a6a as nil")
	}
	b.PutID(ChatTypeSupergroupTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ChatTypeSupergroup) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatTypeSupergroup#a83a5a6a as nil")
	}
	b.PutLong(c.SupergroupID)
	b.PutBool(c.IsChannel)
	return nil
}

// Decode implements bin.Decoder.
func (c *ChatTypeSupergroup) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatTypeSupergroup#a83a5a6a to nil")
	}
	if err := b.ConsumeID(ChatTypeSupergroupTypeID); err != nil {
		return fmt.Errorf("unable to decode chatTypeSupergroup#a83a5a6a: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ChatTypeSupergroup) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatTypeSupergroup#a83a5a6a to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode chatTypeSupergroup#a83a5a6a: field supergroup_id: %w", err)
		}
		c.SupergroupID = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode chatTypeSupergroup#a83a5a6a: field is_channel: %w", err)
		}
		c.IsChannel = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (c *ChatTypeSupergroup) EncodeTDLibJSON(b tdjson.Encoder) error {
	if c == nil {
		return fmt.Errorf("can't encode chatTypeSupergroup#a83a5a6a as nil")
	}
	b.ObjStart()
	b.PutID("chatTypeSupergroup")
	b.FieldStart("supergroup_id")
	b.PutLong(c.SupergroupID)
	b.FieldStart("is_channel")
	b.PutBool(c.IsChannel)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (c *ChatTypeSupergroup) DecodeTDLibJSON(b tdjson.Decoder) error {
	if c == nil {
		return fmt.Errorf("can't decode chatTypeSupergroup#a83a5a6a to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("chatTypeSupergroup"); err != nil {
				return fmt.Errorf("unable to decode chatTypeSupergroup#a83a5a6a: %w", err)
			}
		case "supergroup_id":
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode chatTypeSupergroup#a83a5a6a: field supergroup_id: %w", err)
			}
			c.SupergroupID = value
		case "is_channel":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode chatTypeSupergroup#a83a5a6a: field is_channel: %w", err)
			}
			c.IsChannel = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetSupergroupID returns value of SupergroupID field.
func (c *ChatTypeSupergroup) GetSupergroupID() (value int64) {
	return c.SupergroupID
}

// GetIsChannel returns value of IsChannel field.
func (c *ChatTypeSupergroup) GetIsChannel() (value bool) {
	return c.IsChannel
}

// ChatTypeSecret represents TL type `chatTypeSecret#3366ab31`.
type ChatTypeSecret struct {
	// Secret chat identifier
	SecretChatID int32
	// User identifier of the secret chat peer
	UserID int64
}

// ChatTypeSecretTypeID is TL type id of ChatTypeSecret.
const ChatTypeSecretTypeID = 0x3366ab31

// construct implements constructor of ChatTypeClass.
func (c ChatTypeSecret) construct() ChatTypeClass { return &c }

// Ensuring interfaces in compile-time for ChatTypeSecret.
var (
	_ bin.Encoder     = &ChatTypeSecret{}
	_ bin.Decoder     = &ChatTypeSecret{}
	_ bin.BareEncoder = &ChatTypeSecret{}
	_ bin.BareDecoder = &ChatTypeSecret{}

	_ ChatTypeClass = &ChatTypeSecret{}
)

func (c *ChatTypeSecret) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.SecretChatID == 0) {
		return false
	}
	if !(c.UserID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChatTypeSecret) String() string {
	if c == nil {
		return "ChatTypeSecret(nil)"
	}
	type Alias ChatTypeSecret
	return fmt.Sprintf("ChatTypeSecret%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChatTypeSecret) TypeID() uint32 {
	return ChatTypeSecretTypeID
}

// TypeName returns name of type in TL schema.
func (*ChatTypeSecret) TypeName() string {
	return "chatTypeSecret"
}

// TypeInfo returns info about TL type.
func (c *ChatTypeSecret) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "chatTypeSecret",
		ID:   ChatTypeSecretTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "SecretChatID",
			SchemaName: "secret_chat_id",
		},
		{
			Name:       "UserID",
			SchemaName: "user_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *ChatTypeSecret) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatTypeSecret#3366ab31 as nil")
	}
	b.PutID(ChatTypeSecretTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ChatTypeSecret) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatTypeSecret#3366ab31 as nil")
	}
	b.PutInt32(c.SecretChatID)
	b.PutLong(c.UserID)
	return nil
}

// Decode implements bin.Decoder.
func (c *ChatTypeSecret) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatTypeSecret#3366ab31 to nil")
	}
	if err := b.ConsumeID(ChatTypeSecretTypeID); err != nil {
		return fmt.Errorf("unable to decode chatTypeSecret#3366ab31: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ChatTypeSecret) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatTypeSecret#3366ab31 to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode chatTypeSecret#3366ab31: field secret_chat_id: %w", err)
		}
		c.SecretChatID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode chatTypeSecret#3366ab31: field user_id: %w", err)
		}
		c.UserID = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (c *ChatTypeSecret) EncodeTDLibJSON(b tdjson.Encoder) error {
	if c == nil {
		return fmt.Errorf("can't encode chatTypeSecret#3366ab31 as nil")
	}
	b.ObjStart()
	b.PutID("chatTypeSecret")
	b.FieldStart("secret_chat_id")
	b.PutInt32(c.SecretChatID)
	b.FieldStart("user_id")
	b.PutLong(c.UserID)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (c *ChatTypeSecret) DecodeTDLibJSON(b tdjson.Decoder) error {
	if c == nil {
		return fmt.Errorf("can't decode chatTypeSecret#3366ab31 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("chatTypeSecret"); err != nil {
				return fmt.Errorf("unable to decode chatTypeSecret#3366ab31: %w", err)
			}
		case "secret_chat_id":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode chatTypeSecret#3366ab31: field secret_chat_id: %w", err)
			}
			c.SecretChatID = value
		case "user_id":
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode chatTypeSecret#3366ab31: field user_id: %w", err)
			}
			c.UserID = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetSecretChatID returns value of SecretChatID field.
func (c *ChatTypeSecret) GetSecretChatID() (value int32) {
	return c.SecretChatID
}

// GetUserID returns value of UserID field.
func (c *ChatTypeSecret) GetUserID() (value int64) {
	return c.UserID
}

// ChatTypeClass represents ChatType generic type.
//
// Example:
//  g, err := tdapi.DecodeChatType(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tdapi.ChatTypePrivate: // chatTypePrivate#5e1e6374
//  case *tdapi.ChatTypeBasicGroup: // chatTypeBasicGroup#3a0c4c5c
//  case *tdapi.ChatTypeSupergroup: // chatTypeSupergroup#a83a5a6a
//  case *tdapi.ChatTypeSecret: // chatTypeSecret#3366ab31
//  default: panic(v)
//  }
type ChatTypeClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() ChatTypeClass

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

// DecodeChatType implements binary de-serialization for ChatTypeClass.
func DecodeChatType(buf *bin.Buffer) (ChatTypeClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case ChatTypePrivateTypeID:
		// Decoding chatTypePrivate#5e1e6374.
		v := ChatTypePrivate{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ChatTypeClass: %w", err)
		}
		return &v, nil
	case ChatTypeBasicGroupTypeID:
		// Decoding chatTypeBasicGroup#3a0c4c5c.
		v := ChatTypeBasicGroup{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ChatTypeClass: %w", err)
		}
		return &v, nil
	case ChatTypeSupergroupTypeID:
		// Decoding chatTypeSupergroup#a83a5a6a.
		v := ChatTypeSupergroup{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ChatTypeClass: %w", err)
		}
		return &v, nil
	case ChatTypeSecretTypeID:
		// Decoding chatTypeSecret#3366ab31.
		v := ChatTypeSecret{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ChatTypeClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode ChatTypeClass: %w", bin.NewUnexpectedID(id))
	}
}

// DecodeTDLibJSONChatType implements binary de-serialization for ChatTypeClass.
func DecodeTDLibJSONChatType(buf tdjson.Decoder) (ChatTypeClass, error) {
	id, err := buf.FindTypeID()
	if err != nil {
		return nil, err
	}
	switch id {
	case "chatTypePrivate":
		// Decoding chatTypePrivate#5e1e6374.
		v := ChatTypePrivate{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ChatTypeClass: %w", err)
		}
		return &v, nil
	case "chatTypeBasicGroup":
		// Decoding chatTypeBasicGroup#3a0c4c5c.
		v := ChatTypeBasicGroup{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ChatTypeClass: %w", err)
		}
		return &v, nil
	case "chatTypeSupergroup":
		// Decoding chatTypeSupergroup#a83a5a6a.
		v := ChatTypeSupergroup{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ChatTypeClass: %w", err)
		}
		return &v, nil
	case "chatTypeSecret":
		// Decoding chatTypeSecret#3366ab31.
		v := ChatTypeSecret{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ChatTypeClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode ChatTypeClass: %w", tdjson.NewUnexpectedID(id))
	}
}

// ChatType boxes the ChatTypeClass providing a helper.
type ChatTypeBox struct {
	ChatType ChatTypeClass
}

// Decode implements bin.Decoder for ChatTypeBox.
func (b *ChatTypeBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode ChatTypeBox to nil")
	}
	v, err := DecodeChatType(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.ChatType = v
	return nil
}

// Encode implements bin.Encode for ChatTypeBox.
func (b *ChatTypeBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.ChatType == nil {
		return fmt.Errorf("unable to encode ChatTypeClass as nil")
	}
	return b.ChatType.Encode(buf)
}

// DecodeTDLibJSON implements bin.Decoder for ChatTypeBox.
func (b *ChatTypeBox) DecodeTDLibJSON(buf tdjson.Decoder) error {
	if b == nil {
		return fmt.Errorf("unable to decode ChatTypeBox to nil")
	}
	v, err := DecodeTDLibJSONChatType(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.ChatType = v
	return nil
}

// EncodeTDLibJSON implements bin.Encode for ChatTypeBox.
func (b *ChatTypeBox) EncodeTDLibJSON(buf tdjson.Encoder) error {
	if b == nil || b.ChatType == nil {
		return fmt.Errorf("unable to encode ChatTypeClass as nil")
	}
	return b.ChatType.EncodeTDLibJSON(buf)
}

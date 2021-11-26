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

// ChatStatisticsAdministratorActionsInfo represents TL type `chatStatisticsAdministratorActionsInfo#e7c5cd7e`.
type ChatStatisticsAdministratorActionsInfo struct {
	// Administrator user identifier
	UserID int64
	// Number of messages deleted by the administrator
	DeletedMessageCount int32
	// Number of users banned by the administrator
	BannedUserCount int32
	// Number of users restricted by the administrator
	RestrictedUserCount int32
}

// ChatStatisticsAdministratorActionsInfoTypeID is TL type id of ChatStatisticsAdministratorActionsInfo.
const ChatStatisticsAdministratorActionsInfoTypeID = 0xe7c5cd7e

// Ensuring interfaces in compile-time for ChatStatisticsAdministratorActionsInfo.
var (
	_ bin.Encoder     = &ChatStatisticsAdministratorActionsInfo{}
	_ bin.Decoder     = &ChatStatisticsAdministratorActionsInfo{}
	_ bin.BareEncoder = &ChatStatisticsAdministratorActionsInfo{}
	_ bin.BareDecoder = &ChatStatisticsAdministratorActionsInfo{}
)

func (c *ChatStatisticsAdministratorActionsInfo) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.UserID == 0) {
		return false
	}
	if !(c.DeletedMessageCount == 0) {
		return false
	}
	if !(c.BannedUserCount == 0) {
		return false
	}
	if !(c.RestrictedUserCount == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChatStatisticsAdministratorActionsInfo) String() string {
	if c == nil {
		return "ChatStatisticsAdministratorActionsInfo(nil)"
	}
	type Alias ChatStatisticsAdministratorActionsInfo
	return fmt.Sprintf("ChatStatisticsAdministratorActionsInfo%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChatStatisticsAdministratorActionsInfo) TypeID() uint32 {
	return ChatStatisticsAdministratorActionsInfoTypeID
}

// TypeName returns name of type in TL schema.
func (*ChatStatisticsAdministratorActionsInfo) TypeName() string {
	return "chatStatisticsAdministratorActionsInfo"
}

// TypeInfo returns info about TL type.
func (c *ChatStatisticsAdministratorActionsInfo) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "chatStatisticsAdministratorActionsInfo",
		ID:   ChatStatisticsAdministratorActionsInfoTypeID,
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
		{
			Name:       "DeletedMessageCount",
			SchemaName: "deleted_message_count",
		},
		{
			Name:       "BannedUserCount",
			SchemaName: "banned_user_count",
		},
		{
			Name:       "RestrictedUserCount",
			SchemaName: "restricted_user_count",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *ChatStatisticsAdministratorActionsInfo) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatStatisticsAdministratorActionsInfo#e7c5cd7e as nil")
	}
	b.PutID(ChatStatisticsAdministratorActionsInfoTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ChatStatisticsAdministratorActionsInfo) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatStatisticsAdministratorActionsInfo#e7c5cd7e as nil")
	}
	b.PutLong(c.UserID)
	b.PutInt32(c.DeletedMessageCount)
	b.PutInt32(c.BannedUserCount)
	b.PutInt32(c.RestrictedUserCount)
	return nil
}

// Decode implements bin.Decoder.
func (c *ChatStatisticsAdministratorActionsInfo) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatStatisticsAdministratorActionsInfo#e7c5cd7e to nil")
	}
	if err := b.ConsumeID(ChatStatisticsAdministratorActionsInfoTypeID); err != nil {
		return fmt.Errorf("unable to decode chatStatisticsAdministratorActionsInfo#e7c5cd7e: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ChatStatisticsAdministratorActionsInfo) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatStatisticsAdministratorActionsInfo#e7c5cd7e to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode chatStatisticsAdministratorActionsInfo#e7c5cd7e: field user_id: %w", err)
		}
		c.UserID = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode chatStatisticsAdministratorActionsInfo#e7c5cd7e: field deleted_message_count: %w", err)
		}
		c.DeletedMessageCount = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode chatStatisticsAdministratorActionsInfo#e7c5cd7e: field banned_user_count: %w", err)
		}
		c.BannedUserCount = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode chatStatisticsAdministratorActionsInfo#e7c5cd7e: field restricted_user_count: %w", err)
		}
		c.RestrictedUserCount = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (c *ChatStatisticsAdministratorActionsInfo) EncodeTDLibJSON(b tdjson.Encoder) error {
	if c == nil {
		return fmt.Errorf("can't encode chatStatisticsAdministratorActionsInfo#e7c5cd7e as nil")
	}
	b.ObjStart()
	b.PutID("chatStatisticsAdministratorActionsInfo")
	b.FieldStart("user_id")
	b.PutLong(c.UserID)
	b.FieldStart("deleted_message_count")
	b.PutInt32(c.DeletedMessageCount)
	b.FieldStart("banned_user_count")
	b.PutInt32(c.BannedUserCount)
	b.FieldStart("restricted_user_count")
	b.PutInt32(c.RestrictedUserCount)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (c *ChatStatisticsAdministratorActionsInfo) DecodeTDLibJSON(b tdjson.Decoder) error {
	if c == nil {
		return fmt.Errorf("can't decode chatStatisticsAdministratorActionsInfo#e7c5cd7e to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("chatStatisticsAdministratorActionsInfo"); err != nil {
				return fmt.Errorf("unable to decode chatStatisticsAdministratorActionsInfo#e7c5cd7e: %w", err)
			}
		case "user_id":
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode chatStatisticsAdministratorActionsInfo#e7c5cd7e: field user_id: %w", err)
			}
			c.UserID = value
		case "deleted_message_count":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode chatStatisticsAdministratorActionsInfo#e7c5cd7e: field deleted_message_count: %w", err)
			}
			c.DeletedMessageCount = value
		case "banned_user_count":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode chatStatisticsAdministratorActionsInfo#e7c5cd7e: field banned_user_count: %w", err)
			}
			c.BannedUserCount = value
		case "restricted_user_count":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode chatStatisticsAdministratorActionsInfo#e7c5cd7e: field restricted_user_count: %w", err)
			}
			c.RestrictedUserCount = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetUserID returns value of UserID field.
func (c *ChatStatisticsAdministratorActionsInfo) GetUserID() (value int64) {
	return c.UserID
}

// GetDeletedMessageCount returns value of DeletedMessageCount field.
func (c *ChatStatisticsAdministratorActionsInfo) GetDeletedMessageCount() (value int32) {
	return c.DeletedMessageCount
}

// GetBannedUserCount returns value of BannedUserCount field.
func (c *ChatStatisticsAdministratorActionsInfo) GetBannedUserCount() (value int32) {
	return c.BannedUserCount
}

// GetRestrictedUserCount returns value of RestrictedUserCount field.
func (c *ChatStatisticsAdministratorActionsInfo) GetRestrictedUserCount() (value int32) {
	return c.RestrictedUserCount
}

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

// ReportSupergroupSpamRequest represents TL type `reportSupergroupSpam#d0c20173`.
type ReportSupergroupSpamRequest struct {
	// Supergroup identifier
	SupergroupID int64
	// User identifier
	UserID int64
	// Identifiers of messages sent in the supergroup by the user. This list must be
	// non-empty
	MessageIDs []int64
}

// ReportSupergroupSpamRequestTypeID is TL type id of ReportSupergroupSpamRequest.
const ReportSupergroupSpamRequestTypeID = 0xd0c20173

// Ensuring interfaces in compile-time for ReportSupergroupSpamRequest.
var (
	_ bin.Encoder     = &ReportSupergroupSpamRequest{}
	_ bin.Decoder     = &ReportSupergroupSpamRequest{}
	_ bin.BareEncoder = &ReportSupergroupSpamRequest{}
	_ bin.BareDecoder = &ReportSupergroupSpamRequest{}
)

func (r *ReportSupergroupSpamRequest) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.SupergroupID == 0) {
		return false
	}
	if !(r.UserID == 0) {
		return false
	}
	if !(r.MessageIDs == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *ReportSupergroupSpamRequest) String() string {
	if r == nil {
		return "ReportSupergroupSpamRequest(nil)"
	}
	type Alias ReportSupergroupSpamRequest
	return fmt.Sprintf("ReportSupergroupSpamRequest%+v", Alias(*r))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ReportSupergroupSpamRequest) TypeID() uint32 {
	return ReportSupergroupSpamRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ReportSupergroupSpamRequest) TypeName() string {
	return "reportSupergroupSpam"
}

// TypeInfo returns info about TL type.
func (r *ReportSupergroupSpamRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "reportSupergroupSpam",
		ID:   ReportSupergroupSpamRequestTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "SupergroupID",
			SchemaName: "supergroup_id",
		},
		{
			Name:       "UserID",
			SchemaName: "user_id",
		},
		{
			Name:       "MessageIDs",
			SchemaName: "message_ids",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (r *ReportSupergroupSpamRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode reportSupergroupSpam#d0c20173 as nil")
	}
	b.PutID(ReportSupergroupSpamRequestTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *ReportSupergroupSpamRequest) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode reportSupergroupSpam#d0c20173 as nil")
	}
	b.PutLong(r.SupergroupID)
	b.PutLong(r.UserID)
	b.PutInt(len(r.MessageIDs))
	for _, v := range r.MessageIDs {
		b.PutLong(v)
	}
	return nil
}

// Decode implements bin.Decoder.
func (r *ReportSupergroupSpamRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode reportSupergroupSpam#d0c20173 to nil")
	}
	if err := b.ConsumeID(ReportSupergroupSpamRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode reportSupergroupSpam#d0c20173: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *ReportSupergroupSpamRequest) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode reportSupergroupSpam#d0c20173 to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode reportSupergroupSpam#d0c20173: field supergroup_id: %w", err)
		}
		r.SupergroupID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode reportSupergroupSpam#d0c20173: field user_id: %w", err)
		}
		r.UserID = value
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode reportSupergroupSpam#d0c20173: field message_ids: %w", err)
		}

		if headerLen > 0 {
			r.MessageIDs = make([]int64, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode reportSupergroupSpam#d0c20173: field message_ids: %w", err)
			}
			r.MessageIDs = append(r.MessageIDs, value)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (r *ReportSupergroupSpamRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if r == nil {
		return fmt.Errorf("can't encode reportSupergroupSpam#d0c20173 as nil")
	}
	b.ObjStart()
	b.PutID("reportSupergroupSpam")
	b.FieldStart("supergroup_id")
	b.PutLong(r.SupergroupID)
	b.FieldStart("user_id")
	b.PutLong(r.UserID)
	b.FieldStart("message_ids")
	b.ArrStart()
	for _, v := range r.MessageIDs {
		b.PutLong(v)
	}
	b.ArrEnd()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (r *ReportSupergroupSpamRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if r == nil {
		return fmt.Errorf("can't decode reportSupergroupSpam#d0c20173 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("reportSupergroupSpam"); err != nil {
				return fmt.Errorf("unable to decode reportSupergroupSpam#d0c20173: %w", err)
			}
		case "supergroup_id":
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode reportSupergroupSpam#d0c20173: field supergroup_id: %w", err)
			}
			r.SupergroupID = value
		case "user_id":
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode reportSupergroupSpam#d0c20173: field user_id: %w", err)
			}
			r.UserID = value
		case "message_ids":
			if err := b.Arr(func(b tdjson.Decoder) error {
				value, err := b.Long()
				if err != nil {
					return fmt.Errorf("unable to decode reportSupergroupSpam#d0c20173: field message_ids: %w", err)
				}
				r.MessageIDs = append(r.MessageIDs, value)
				return nil
			}); err != nil {
				return fmt.Errorf("unable to decode reportSupergroupSpam#d0c20173: field message_ids: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetSupergroupID returns value of SupergroupID field.
func (r *ReportSupergroupSpamRequest) GetSupergroupID() (value int64) {
	return r.SupergroupID
}

// GetUserID returns value of UserID field.
func (r *ReportSupergroupSpamRequest) GetUserID() (value int64) {
	return r.UserID
}

// GetMessageIDs returns value of MessageIDs field.
func (r *ReportSupergroupSpamRequest) GetMessageIDs() (value []int64) {
	return r.MessageIDs
}

// ReportSupergroupSpam invokes method reportSupergroupSpam#d0c20173 returning error if any.
func (c *Client) ReportSupergroupSpam(ctx context.Context, request *ReportSupergroupSpamRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}

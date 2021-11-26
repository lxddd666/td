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

// DeepLinkInfo represents TL type `deepLinkInfo#6f1ba0fe`.
type DeepLinkInfo struct {
	// Text to be shown to the user
	Text FormattedText
	// True, if the user must be asked to update the application
	NeedUpdateApplication bool
}

// DeepLinkInfoTypeID is TL type id of DeepLinkInfo.
const DeepLinkInfoTypeID = 0x6f1ba0fe

// Ensuring interfaces in compile-time for DeepLinkInfo.
var (
	_ bin.Encoder     = &DeepLinkInfo{}
	_ bin.Decoder     = &DeepLinkInfo{}
	_ bin.BareEncoder = &DeepLinkInfo{}
	_ bin.BareDecoder = &DeepLinkInfo{}
)

func (d *DeepLinkInfo) Zero() bool {
	if d == nil {
		return true
	}
	if !(d.Text.Zero()) {
		return false
	}
	if !(d.NeedUpdateApplication == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (d *DeepLinkInfo) String() string {
	if d == nil {
		return "DeepLinkInfo(nil)"
	}
	type Alias DeepLinkInfo
	return fmt.Sprintf("DeepLinkInfo%+v", Alias(*d))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*DeepLinkInfo) TypeID() uint32 {
	return DeepLinkInfoTypeID
}

// TypeName returns name of type in TL schema.
func (*DeepLinkInfo) TypeName() string {
	return "deepLinkInfo"
}

// TypeInfo returns info about TL type.
func (d *DeepLinkInfo) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "deepLinkInfo",
		ID:   DeepLinkInfoTypeID,
	}
	if d == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Text",
			SchemaName: "text",
		},
		{
			Name:       "NeedUpdateApplication",
			SchemaName: "need_update_application",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (d *DeepLinkInfo) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode deepLinkInfo#6f1ba0fe as nil")
	}
	b.PutID(DeepLinkInfoTypeID)
	return d.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (d *DeepLinkInfo) EncodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode deepLinkInfo#6f1ba0fe as nil")
	}
	if err := d.Text.Encode(b); err != nil {
		return fmt.Errorf("unable to encode deepLinkInfo#6f1ba0fe: field text: %w", err)
	}
	b.PutBool(d.NeedUpdateApplication)
	return nil
}

// Decode implements bin.Decoder.
func (d *DeepLinkInfo) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode deepLinkInfo#6f1ba0fe to nil")
	}
	if err := b.ConsumeID(DeepLinkInfoTypeID); err != nil {
		return fmt.Errorf("unable to decode deepLinkInfo#6f1ba0fe: %w", err)
	}
	return d.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (d *DeepLinkInfo) DecodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode deepLinkInfo#6f1ba0fe to nil")
	}
	{
		if err := d.Text.Decode(b); err != nil {
			return fmt.Errorf("unable to decode deepLinkInfo#6f1ba0fe: field text: %w", err)
		}
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode deepLinkInfo#6f1ba0fe: field need_update_application: %w", err)
		}
		d.NeedUpdateApplication = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (d *DeepLinkInfo) EncodeTDLibJSON(b tdjson.Encoder) error {
	if d == nil {
		return fmt.Errorf("can't encode deepLinkInfo#6f1ba0fe as nil")
	}
	b.ObjStart()
	b.PutID("deepLinkInfo")
	b.FieldStart("text")
	if err := d.Text.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode deepLinkInfo#6f1ba0fe: field text: %w", err)
	}
	b.FieldStart("need_update_application")
	b.PutBool(d.NeedUpdateApplication)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (d *DeepLinkInfo) DecodeTDLibJSON(b tdjson.Decoder) error {
	if d == nil {
		return fmt.Errorf("can't decode deepLinkInfo#6f1ba0fe to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("deepLinkInfo"); err != nil {
				return fmt.Errorf("unable to decode deepLinkInfo#6f1ba0fe: %w", err)
			}
		case "text":
			if err := d.Text.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode deepLinkInfo#6f1ba0fe: field text: %w", err)
			}
		case "need_update_application":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode deepLinkInfo#6f1ba0fe: field need_update_application: %w", err)
			}
			d.NeedUpdateApplication = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetText returns value of Text field.
func (d *DeepLinkInfo) GetText() (value FormattedText) {
	return d.Text
}

// GetNeedUpdateApplication returns value of NeedUpdateApplication field.
func (d *DeepLinkInfo) GetNeedUpdateApplication() (value bool) {
	return d.NeedUpdateApplication
}

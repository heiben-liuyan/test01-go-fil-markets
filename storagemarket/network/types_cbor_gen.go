// Code generated by github.com/whyrusleeping/cbor-gen. DO NOT EDIT.

package network

import (
	"fmt"
	"io"

	"github.com/LIUYAN-0626/test01-go-fil-markets/storagemarket"
	"github.com/filecoin-project/specs-actors/actors/builtin/market"
	"github.com/filecoin-project/specs-actors/actors/crypto"
	cbg "github.com/whyrusleeping/cbor-gen"
	xerrors "golang.org/x/xerrors"
)

var _ = xerrors.Errorf

func (t *AskRequest) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{129}); err != nil {
		return err
	}

	// t.Miner (address.Address) (struct)
	if err := t.Miner.MarshalCBOR(w); err != nil {
		return err
	}
	return nil
}

func (t *AskRequest) UnmarshalCBOR(r io.Reader) error {
	br := cbg.GetPeeker(r)

	maj, extra, err := cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 1 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Miner (address.Address) (struct)

	{

		if err := t.Miner.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.Miner: %w", err)
		}

	}
	return nil
}

func (t *AskResponse) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{129}); err != nil {
		return err
	}

	// t.Ask (storagemarket.SignedStorageAsk) (struct)
	if err := t.Ask.MarshalCBOR(w); err != nil {
		return err
	}
	return nil
}

func (t *AskResponse) UnmarshalCBOR(r io.Reader) error {
	br := cbg.GetPeeker(r)

	maj, extra, err := cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 1 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Ask (storagemarket.SignedStorageAsk) (struct)

	{

		pb, err := br.PeekByte()
		if err != nil {
			return err
		}
		if pb == cbg.CborNull[0] {
			var nbuf [1]byte
			if _, err := br.Read(nbuf[:]); err != nil {
				return err
			}
		} else {
			t.Ask = new(storagemarket.SignedStorageAsk)
			if err := t.Ask.UnmarshalCBOR(br); err != nil {
				return xerrors.Errorf("unmarshaling t.Ask pointer: %w", err)
			}
		}

	}
	return nil
}

func (t *Proposal) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{130}); err != nil {
		return err
	}

	// t.DealProposal (market.ClientDealProposal) (struct)
	if err := t.DealProposal.MarshalCBOR(w); err != nil {
		return err
	}

	// t.Piece (storagemarket.DataRef) (struct)
	if err := t.Piece.MarshalCBOR(w); err != nil {
		return err
	}
	return nil
}

func (t *Proposal) UnmarshalCBOR(r io.Reader) error {
	br := cbg.GetPeeker(r)

	maj, extra, err := cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 2 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.DealProposal (market.ClientDealProposal) (struct)

	{

		pb, err := br.PeekByte()
		if err != nil {
			return err
		}
		if pb == cbg.CborNull[0] {
			var nbuf [1]byte
			if _, err := br.Read(nbuf[:]); err != nil {
				return err
			}
		} else {
			t.DealProposal = new(market.ClientDealProposal)
			if err := t.DealProposal.UnmarshalCBOR(br); err != nil {
				return xerrors.Errorf("unmarshaling t.DealProposal pointer: %w", err)
			}
		}

	}
	// t.Piece (storagemarket.DataRef) (struct)

	{

		pb, err := br.PeekByte()
		if err != nil {
			return err
		}
		if pb == cbg.CborNull[0] {
			var nbuf [1]byte
			if _, err := br.Read(nbuf[:]); err != nil {
				return err
			}
		} else {
			t.Piece = new(storagemarket.DataRef)
			if err := t.Piece.UnmarshalCBOR(br); err != nil {
				return xerrors.Errorf("unmarshaling t.Piece pointer: %w", err)
			}
		}

	}
	return nil
}

func (t *Response) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{132}); err != nil {
		return err
	}

	// t.State (uint64) (uint64)

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, uint64(t.State))); err != nil {
		return err
	}

	// t.Message (string) (string)
	if len(t.Message) > cbg.MaxLength {
		return xerrors.Errorf("Value in field t.Message was too long")
	}

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajTextString, uint64(len(t.Message)))); err != nil {
		return err
	}
	if _, err := w.Write([]byte(t.Message)); err != nil {
		return err
	}

	// t.Proposal (cid.Cid) (struct)

	if err := cbg.WriteCid(w, t.Proposal); err != nil {
		return xerrors.Errorf("failed to write cid field t.Proposal: %w", err)
	}

	// t.PublishMessage (cid.Cid) (struct)

	if t.PublishMessage == nil {
		if _, err := w.Write(cbg.CborNull); err != nil {
			return err
		}
	} else {
		if err := cbg.WriteCid(w, *t.PublishMessage); err != nil {
			return xerrors.Errorf("failed to write cid field t.PublishMessage: %w", err)
		}
	}

	return nil
}

func (t *Response) UnmarshalCBOR(r io.Reader) error {
	br := cbg.GetPeeker(r)

	maj, extra, err := cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 4 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.State (uint64) (uint64)

	{

		maj, extra, err = cbg.CborReadHeader(br)
		if err != nil {
			return err
		}
		if maj != cbg.MajUnsignedInt {
			return fmt.Errorf("wrong type for uint64 field")
		}
		t.State = uint64(extra)

	}
	// t.Message (string) (string)

	{
		sval, err := cbg.ReadString(br)
		if err != nil {
			return err
		}

		t.Message = string(sval)
	}
	// t.Proposal (cid.Cid) (struct)

	{

		c, err := cbg.ReadCid(br)
		if err != nil {
			return xerrors.Errorf("failed to read cid field t.Proposal: %w", err)
		}

		t.Proposal = c

	}
	// t.PublishMessage (cid.Cid) (struct)

	{

		pb, err := br.PeekByte()
		if err != nil {
			return err
		}
		if pb == cbg.CborNull[0] {
			var nbuf [1]byte
			if _, err := br.Read(nbuf[:]); err != nil {
				return err
			}
		} else {

			c, err := cbg.ReadCid(br)
			if err != nil {
				return xerrors.Errorf("failed to read cid field t.PublishMessage: %w", err)
			}

			t.PublishMessage = &c
		}

	}
	return nil
}

func (t *SignedResponse) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{130}); err != nil {
		return err
	}

	// t.Response (network.Response) (struct)
	if err := t.Response.MarshalCBOR(w); err != nil {
		return err
	}

	// t.Signature (crypto.Signature) (struct)
	if err := t.Signature.MarshalCBOR(w); err != nil {
		return err
	}
	return nil
}

func (t *SignedResponse) UnmarshalCBOR(r io.Reader) error {
	br := cbg.GetPeeker(r)

	maj, extra, err := cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 2 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Response (network.Response) (struct)

	{

		if err := t.Response.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.Response: %w", err)
		}

	}
	// t.Signature (crypto.Signature) (struct)

	{

		pb, err := br.PeekByte()
		if err != nil {
			return err
		}
		if pb == cbg.CborNull[0] {
			var nbuf [1]byte
			if _, err := br.Read(nbuf[:]); err != nil {
				return err
			}
		} else {
			t.Signature = new(crypto.Signature)
			if err := t.Signature.UnmarshalCBOR(br); err != nil {
				return xerrors.Errorf("unmarshaling t.Signature pointer: %w", err)
			}
		}

	}
	return nil
}

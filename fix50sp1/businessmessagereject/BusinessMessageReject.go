//Package businessmessagereject msg type = j.
package businessmessagereject

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fixt11"
)

//Message is a BusinessMessageReject FIX Message
type Message struct {
	FIXMsgType string `fix:"j"`
	Header     fixt11.Header
	//RefSeqNum is a non-required field for BusinessMessageReject.
	RefSeqNum *int `fix:"45"`
	//RefMsgType is a required field for BusinessMessageReject.
	RefMsgType string `fix:"372"`
	//BusinessRejectRefID is a non-required field for BusinessMessageReject.
	BusinessRejectRefID *string `fix:"379"`
	//BusinessRejectReason is a required field for BusinessMessageReject.
	BusinessRejectReason int `fix:"380"`
	//Text is a non-required field for BusinessMessageReject.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for BusinessMessageReject.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for BusinessMessageReject.
	EncodedText *string `fix:"355"`
	//RefApplVerID is a non-required field for BusinessMessageReject.
	RefApplVerID *string `fix:"1130"`
	//RefApplExtID is a non-required field for BusinessMessageReject.
	RefApplExtID *int `fix:"1406"`
	//RefCstmApplVerID is a non-required field for BusinessMessageReject.
	RefCstmApplVerID *string `fix:"1131"`
	Trailer          fixt11.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Mesage type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		m := new(Message)
		if err := quickfix.Unmarshal(msg, m); err != nil {
			return err
		}
		return router(*m, sessionID)
	}
	return enum.ApplVerID_FIX50SP1, "j", r
}
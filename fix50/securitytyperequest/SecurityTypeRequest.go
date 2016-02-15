//Package securitytyperequest msg type = v.
package securitytyperequest

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fixt11"
)

//Message is a SecurityTypeRequest FIX Message
type Message struct {
	FIXMsgType string `fix:"v"`
	Header     fixt11.Header
	//SecurityReqID is a required field for SecurityTypeRequest.
	SecurityReqID string `fix:"320"`
	//Text is a non-required field for SecurityTypeRequest.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for SecurityTypeRequest.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for SecurityTypeRequest.
	EncodedText *string `fix:"355"`
	//TradingSessionID is a non-required field for SecurityTypeRequest.
	TradingSessionID *string `fix:"336"`
	//TradingSessionSubID is a non-required field for SecurityTypeRequest.
	TradingSessionSubID *string `fix:"625"`
	//Product is a non-required field for SecurityTypeRequest.
	Product *int `fix:"460"`
	//SecurityType is a non-required field for SecurityTypeRequest.
	SecurityType *string `fix:"167"`
	//SecuritySubType is a non-required field for SecurityTypeRequest.
	SecuritySubType *string `fix:"762"`
	Trailer         fixt11.Trailer
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
	return enum.BeginStringFIX50, "v", r
}
//Package securitylistrequest msg type = x.
package securitylistrequest

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix43"
	"github.com/quickfixgo/quickfix/fix43/instrument"
)

//Message is a SecurityListRequest FIX Message
type Message struct {
	FIXMsgType string `fix:"x"`
	Header     fix43.Header
	//SecurityReqID is a required field for SecurityListRequest.
	SecurityReqID string `fix:"320"`
	//SecurityListRequestType is a required field for SecurityListRequest.
	SecurityListRequestType int `fix:"559"`
	//Instrument Component
	Instrument instrument.Component
	//Currency is a non-required field for SecurityListRequest.
	Currency *string `fix:"15"`
	//Text is a non-required field for SecurityListRequest.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for SecurityListRequest.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for SecurityListRequest.
	EncodedText *string `fix:"355"`
	//TradingSessionID is a non-required field for SecurityListRequest.
	TradingSessionID *string `fix:"336"`
	//TradingSessionSubID is a non-required field for SecurityListRequest.
	TradingSessionSubID *string `fix:"625"`
	//SubscriptionRequestType is a non-required field for SecurityListRequest.
	SubscriptionRequestType *string `fix:"263"`
	Trailer                 fix43.Trailer
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
	return enum.BeginStringFIX43, "x", r
}
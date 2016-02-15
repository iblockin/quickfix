//Package quotecancel msg type = Z.
package quotecancel

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix50sp2/parties"
	"github.com/quickfixgo/quickfix/fix50sp2/quotcxlentriesgrp"
	"github.com/quickfixgo/quickfix/fix50sp2/targetparties"
	"github.com/quickfixgo/quickfix/fixt11"
)

//Message is a QuoteCancel FIX Message
type Message struct {
	FIXMsgType string `fix:"Z"`
	Header     fixt11.Header
	//QuoteReqID is a non-required field for QuoteCancel.
	QuoteReqID *string `fix:"131"`
	//QuoteID is a non-required field for QuoteCancel.
	QuoteID *string `fix:"117"`
	//QuoteCancelType is a required field for QuoteCancel.
	QuoteCancelType int `fix:"298"`
	//QuoteResponseLevel is a non-required field for QuoteCancel.
	QuoteResponseLevel *int `fix:"301"`
	//Parties Component
	Parties parties.Component
	//Account is a non-required field for QuoteCancel.
	Account *string `fix:"1"`
	//AcctIDSource is a non-required field for QuoteCancel.
	AcctIDSource *int `fix:"660"`
	//AccountType is a non-required field for QuoteCancel.
	AccountType *int `fix:"581"`
	//TradingSessionID is a non-required field for QuoteCancel.
	TradingSessionID *string `fix:"336"`
	//TradingSessionSubID is a non-required field for QuoteCancel.
	TradingSessionSubID *string `fix:"625"`
	//QuotCxlEntriesGrp Component
	QuotCxlEntriesGrp quotcxlentriesgrp.Component
	//QuoteMsgID is a non-required field for QuoteCancel.
	QuoteMsgID *string `fix:"1166"`
	//QuoteType is a non-required field for QuoteCancel.
	QuoteType *int `fix:"537"`
	//TargetParties Component
	TargetParties targetparties.Component
	Trailer       fixt11.Trailer
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
	return enum.ApplVerID_FIX50SP2, "Z", r
}
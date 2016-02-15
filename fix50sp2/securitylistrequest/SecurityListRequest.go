//Package securitylistrequest msg type = x.
package securitylistrequest

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix50sp2/financingdetails"
	"github.com/quickfixgo/quickfix/fix50sp2/instrmtleggrp"
	"github.com/quickfixgo/quickfix/fix50sp2/instrument"
	"github.com/quickfixgo/quickfix/fix50sp2/instrumentextension"
	"github.com/quickfixgo/quickfix/fix50sp2/undinstrmtgrp"
	"github.com/quickfixgo/quickfix/fixt11"
)

//Message is a SecurityListRequest FIX Message
type Message struct {
	FIXMsgType string `fix:"x"`
	Header     fixt11.Header
	//SecurityReqID is a required field for SecurityListRequest.
	SecurityReqID string `fix:"320"`
	//SecurityListRequestType is a required field for SecurityListRequest.
	SecurityListRequestType int `fix:"559"`
	//Instrument Component
	Instrument instrument.Component
	//InstrumentExtension Component
	InstrumentExtension instrumentextension.Component
	//FinancingDetails Component
	FinancingDetails financingdetails.Component
	//UndInstrmtGrp Component
	UndInstrmtGrp undinstrmtgrp.Component
	//InstrmtLegGrp Component
	InstrmtLegGrp instrmtleggrp.Component
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
	//MarketID is a non-required field for SecurityListRequest.
	MarketID *string `fix:"1301"`
	//MarketSegmentID is a non-required field for SecurityListRequest.
	MarketSegmentID *string `fix:"1300"`
	//SecurityListID is a non-required field for SecurityListRequest.
	SecurityListID *string `fix:"1465"`
	//SecurityListType is a non-required field for SecurityListRequest.
	SecurityListType *int `fix:"1470"`
	//SecurityListTypeSource is a non-required field for SecurityListRequest.
	SecurityListTypeSource *int `fix:"1471"`
	Trailer                fixt11.Trailer
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
	return enum.ApplVerID_FIX50SP2, "x", r
}
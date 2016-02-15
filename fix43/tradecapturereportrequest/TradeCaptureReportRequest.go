//Package tradecapturereportrequest msg type = AD.
package tradecapturereportrequest

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix43"
	"github.com/quickfixgo/quickfix/fix43/instrument"
	"github.com/quickfixgo/quickfix/fix43/parties"
	"time"
)

//NoDates is a repeating group in TradeCaptureReportRequest
type NoDates struct {
	//TradeDate is a non-required field for NoDates.
	TradeDate *string `fix:"75"`
	//TransactTime is a non-required field for NoDates.
	TransactTime *time.Time `fix:"60"`
}

//Message is a TradeCaptureReportRequest FIX Message
type Message struct {
	FIXMsgType string `fix:"AD"`
	Header     fix43.Header
	//TradeRequestID is a required field for TradeCaptureReportRequest.
	TradeRequestID string `fix:"568"`
	//TradeRequestType is a required field for TradeCaptureReportRequest.
	TradeRequestType int `fix:"569"`
	//SubscriptionRequestType is a non-required field for TradeCaptureReportRequest.
	SubscriptionRequestType *string `fix:"263"`
	//ExecID is a non-required field for TradeCaptureReportRequest.
	ExecID *string `fix:"17"`
	//OrderID is a non-required field for TradeCaptureReportRequest.
	OrderID *string `fix:"37"`
	//ClOrdID is a non-required field for TradeCaptureReportRequest.
	ClOrdID *string `fix:"11"`
	//MatchStatus is a non-required field for TradeCaptureReportRequest.
	MatchStatus *string `fix:"573"`
	//Parties Component
	Parties parties.Component
	//Instrument Component
	Instrument instrument.Component
	//NoDates is a non-required field for TradeCaptureReportRequest.
	NoDates []NoDates `fix:"580,omitempty"`
	//Side is a non-required field for TradeCaptureReportRequest.
	Side *string `fix:"54"`
	//Text is a non-required field for TradeCaptureReportRequest.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for TradeCaptureReportRequest.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for TradeCaptureReportRequest.
	EncodedText *string `fix:"355"`
	//TradeInputSource is a non-required field for TradeCaptureReportRequest.
	TradeInputSource *string `fix:"578"`
	//TradeInputDevice is a non-required field for TradeCaptureReportRequest.
	TradeInputDevice *string `fix:"579"`
	Trailer          fix43.Trailer
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
	return enum.BeginStringFIX43, "AD", r
}
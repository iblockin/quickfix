//Package neworderlist msg type = E.
package neworderlist

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix40"
	"time"
)

//Message is a NewOrderList FIX Message
type Message struct {
	FIXMsgType string `fix:"E"`
	Header     fix40.Header
	//ListID is a required field for NewOrderList.
	ListID string `fix:"66"`
	//WaveNo is a non-required field for NewOrderList.
	WaveNo *string `fix:"105"`
	//ListSeqNo is a required field for NewOrderList.
	ListSeqNo int `fix:"67"`
	//ListNoOrds is a required field for NewOrderList.
	ListNoOrds int `fix:"68"`
	//ListExecInst is a non-required field for NewOrderList.
	ListExecInst *string `fix:"69"`
	//ClOrdID is a required field for NewOrderList.
	ClOrdID string `fix:"11"`
	//ClientID is a non-required field for NewOrderList.
	ClientID *string `fix:"109"`
	//ExecBroker is a non-required field for NewOrderList.
	ExecBroker *string `fix:"76"`
	//Account is a non-required field for NewOrderList.
	Account *string `fix:"1"`
	//SettlmntTyp is a non-required field for NewOrderList.
	SettlmntTyp *string `fix:"63"`
	//FutSettDate is a non-required field for NewOrderList.
	FutSettDate *string `fix:"64"`
	//HandlInst is a required field for NewOrderList.
	HandlInst string `fix:"21"`
	//ExecInst is a non-required field for NewOrderList.
	ExecInst *string `fix:"18"`
	//MinQty is a non-required field for NewOrderList.
	MinQty *int `fix:"110"`
	//MaxFloor is a non-required field for NewOrderList.
	MaxFloor *int `fix:"111"`
	//ExDestination is a non-required field for NewOrderList.
	ExDestination *string `fix:"100"`
	//ProcessCode is a non-required field for NewOrderList.
	ProcessCode *string `fix:"81"`
	//Symbol is a required field for NewOrderList.
	Symbol string `fix:"55"`
	//SymbolSfx is a non-required field for NewOrderList.
	SymbolSfx *string `fix:"65"`
	//SecurityID is a non-required field for NewOrderList.
	SecurityID *string `fix:"48"`
	//IDSource is a non-required field for NewOrderList.
	IDSource *string `fix:"22"`
	//Issuer is a non-required field for NewOrderList.
	Issuer *string `fix:"106"`
	//SecurityDesc is a non-required field for NewOrderList.
	SecurityDesc *string `fix:"107"`
	//PrevClosePx is a non-required field for NewOrderList.
	PrevClosePx *float64 `fix:"140"`
	//Side is a required field for NewOrderList.
	Side string `fix:"54"`
	//LocateReqd is a non-required field for NewOrderList.
	LocateReqd *string `fix:"114"`
	//OrderQty is a required field for NewOrderList.
	OrderQty int `fix:"38"`
	//OrdType is a required field for NewOrderList.
	OrdType string `fix:"40"`
	//Price is a non-required field for NewOrderList.
	Price *float64 `fix:"44"`
	//StopPx is a non-required field for NewOrderList.
	StopPx *float64 `fix:"99"`
	//Currency is a non-required field for NewOrderList.
	Currency *string `fix:"15"`
	//TimeInForce is a non-required field for NewOrderList.
	TimeInForce *string `fix:"59"`
	//ExpireTime is a non-required field for NewOrderList.
	ExpireTime *time.Time `fix:"126"`
	//Commission is a non-required field for NewOrderList.
	Commission *float64 `fix:"12"`
	//CommType is a non-required field for NewOrderList.
	CommType *string `fix:"13"`
	//Rule80A is a non-required field for NewOrderList.
	Rule80A *string `fix:"47"`
	//ForexReq is a non-required field for NewOrderList.
	ForexReq *string `fix:"121"`
	//SettlCurrency is a non-required field for NewOrderList.
	SettlCurrency *string `fix:"120"`
	//Text is a non-required field for NewOrderList.
	Text    *string `fix:"58"`
	Trailer fix40.Trailer
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
	return enum.BeginStringFIX40, "E", r
}
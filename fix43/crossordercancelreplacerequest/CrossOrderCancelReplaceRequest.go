//Package crossordercancelreplacerequest msg type = t.
package crossordercancelreplacerequest

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix43"
	"github.com/quickfixgo/quickfix/fix43/commissiondata"
	"github.com/quickfixgo/quickfix/fix43/instrument"
	"github.com/quickfixgo/quickfix/fix43/nestedparties"
	"github.com/quickfixgo/quickfix/fix43/orderqtydata"
	"github.com/quickfixgo/quickfix/fix43/parties"
	"github.com/quickfixgo/quickfix/fix43/spreadorbenchmarkcurvedata"
	"github.com/quickfixgo/quickfix/fix43/stipulations"
	"github.com/quickfixgo/quickfix/fix43/yielddata"
	"time"
)

//NoSides is a repeating group in CrossOrderCancelReplaceRequest
type NoSides struct {
	//Side is a required field for NoSides.
	Side string `fix:"54"`
	//OrigClOrdID is a required field for NoSides.
	OrigClOrdID string `fix:"41"`
	//ClOrdID is a required field for NoSides.
	ClOrdID string `fix:"11"`
	//SecondaryClOrdID is a non-required field for NoSides.
	SecondaryClOrdID *string `fix:"526"`
	//ClOrdLinkID is a non-required field for NoSides.
	ClOrdLinkID *string `fix:"583"`
	//OrigOrdModTime is a non-required field for NoSides.
	OrigOrdModTime *time.Time `fix:"586"`
	//Parties Component
	Parties parties.Component
	//TradeOriginationDate is a non-required field for NoSides.
	TradeOriginationDate *string `fix:"229"`
	//Account is a non-required field for NoSides.
	Account *string `fix:"1"`
	//AccountType is a non-required field for NoSides.
	AccountType *int `fix:"581"`
	//DayBookingInst is a non-required field for NoSides.
	DayBookingInst *string `fix:"589"`
	//BookingUnit is a non-required field for NoSides.
	BookingUnit *string `fix:"590"`
	//PreallocMethod is a non-required field for NoSides.
	PreallocMethod *string `fix:"591"`
	//NoAllocs is a non-required field for NoSides.
	NoAllocs []NoAllocs `fix:"78,omitempty"`
	//QuantityType is a non-required field for NoSides.
	QuantityType *int `fix:"465"`
	//OrderQtyData Component
	OrderQtyData orderqtydata.Component
	//CommissionData Component
	CommissionData commissiondata.Component
	//OrderCapacity is a non-required field for NoSides.
	OrderCapacity *string `fix:"528"`
	//OrderRestrictions is a non-required field for NoSides.
	OrderRestrictions *string `fix:"529"`
	//CustOrderCapacity is a non-required field for NoSides.
	CustOrderCapacity *int `fix:"582"`
	//ForexReq is a non-required field for NoSides.
	ForexReq *bool `fix:"121"`
	//SettlCurrency is a non-required field for NoSides.
	SettlCurrency *string `fix:"120"`
	//Text is a non-required field for NoSides.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for NoSides.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for NoSides.
	EncodedText *string `fix:"355"`
	//PositionEffect is a non-required field for NoSides.
	PositionEffect *string `fix:"77"`
	//CoveredOrUncovered is a non-required field for NoSides.
	CoveredOrUncovered *int `fix:"203"`
	//CashMargin is a non-required field for NoSides.
	CashMargin *string `fix:"544"`
	//ClearingFeeIndicator is a non-required field for NoSides.
	ClearingFeeIndicator *string `fix:"635"`
	//SolicitedFlag is a non-required field for NoSides.
	SolicitedFlag *bool `fix:"377"`
	//SideComplianceID is a non-required field for NoSides.
	SideComplianceID *string `fix:"659"`
}

//NoAllocs is a repeating group in NoSides
type NoAllocs struct {
	//AllocAccount is a non-required field for NoAllocs.
	AllocAccount *string `fix:"79"`
	//IndividualAllocID is a non-required field for NoAllocs.
	IndividualAllocID *string `fix:"467"`
	//NestedParties Component
	NestedParties nestedparties.Component
	//AllocQty is a non-required field for NoAllocs.
	AllocQty *float64 `fix:"80"`
}

//NoTradingSessions is a repeating group in CrossOrderCancelReplaceRequest
type NoTradingSessions struct {
	//TradingSessionID is a non-required field for NoTradingSessions.
	TradingSessionID *string `fix:"336"`
	//TradingSessionSubID is a non-required field for NoTradingSessions.
	TradingSessionSubID *string `fix:"625"`
}

//Message is a CrossOrderCancelReplaceRequest FIX Message
type Message struct {
	FIXMsgType string `fix:"t"`
	Header     fix43.Header
	//OrderID is a non-required field for CrossOrderCancelReplaceRequest.
	OrderID *string `fix:"37"`
	//CrossID is a required field for CrossOrderCancelReplaceRequest.
	CrossID string `fix:"548"`
	//OrigCrossID is a required field for CrossOrderCancelReplaceRequest.
	OrigCrossID string `fix:"551"`
	//CrossType is a required field for CrossOrderCancelReplaceRequest.
	CrossType int `fix:"549"`
	//CrossPrioritization is a required field for CrossOrderCancelReplaceRequest.
	CrossPrioritization int `fix:"550"`
	//NoSides is a required field for CrossOrderCancelReplaceRequest.
	NoSides []NoSides `fix:"552"`
	//Instrument Component
	Instrument instrument.Component
	//SettlmntTyp is a non-required field for CrossOrderCancelReplaceRequest.
	SettlmntTyp *string `fix:"63"`
	//FutSettDate is a non-required field for CrossOrderCancelReplaceRequest.
	FutSettDate *string `fix:"64"`
	//HandlInst is a required field for CrossOrderCancelReplaceRequest.
	HandlInst string `fix:"21"`
	//ExecInst is a non-required field for CrossOrderCancelReplaceRequest.
	ExecInst *string `fix:"18"`
	//MinQty is a non-required field for CrossOrderCancelReplaceRequest.
	MinQty *float64 `fix:"110"`
	//MaxFloor is a non-required field for CrossOrderCancelReplaceRequest.
	MaxFloor *float64 `fix:"111"`
	//ExDestination is a non-required field for CrossOrderCancelReplaceRequest.
	ExDestination *string `fix:"100"`
	//NoTradingSessions is a non-required field for CrossOrderCancelReplaceRequest.
	NoTradingSessions []NoTradingSessions `fix:"386,omitempty"`
	//ProcessCode is a non-required field for CrossOrderCancelReplaceRequest.
	ProcessCode *string `fix:"81"`
	//PrevClosePx is a non-required field for CrossOrderCancelReplaceRequest.
	PrevClosePx *float64 `fix:"140"`
	//LocateReqd is a non-required field for CrossOrderCancelReplaceRequest.
	LocateReqd *bool `fix:"114"`
	//TransactTime is a required field for CrossOrderCancelReplaceRequest.
	TransactTime time.Time `fix:"60"`
	//Stipulations Component
	Stipulations stipulations.Component
	//OrdType is a required field for CrossOrderCancelReplaceRequest.
	OrdType string `fix:"40"`
	//PriceType is a non-required field for CrossOrderCancelReplaceRequest.
	PriceType *int `fix:"423"`
	//Price is a non-required field for CrossOrderCancelReplaceRequest.
	Price *float64 `fix:"44"`
	//StopPx is a non-required field for CrossOrderCancelReplaceRequest.
	StopPx *float64 `fix:"99"`
	//SpreadOrBenchmarkCurveData Component
	SpreadOrBenchmarkCurveData spreadorbenchmarkcurvedata.Component
	//YieldData Component
	YieldData yielddata.Component
	//Currency is a non-required field for CrossOrderCancelReplaceRequest.
	Currency *string `fix:"15"`
	//ComplianceID is a non-required field for CrossOrderCancelReplaceRequest.
	ComplianceID *string `fix:"376"`
	//IOIid is a non-required field for CrossOrderCancelReplaceRequest.
	IOIid *string `fix:"23"`
	//QuoteID is a non-required field for CrossOrderCancelReplaceRequest.
	QuoteID *string `fix:"117"`
	//TimeInForce is a non-required field for CrossOrderCancelReplaceRequest.
	TimeInForce *string `fix:"59"`
	//EffectiveTime is a non-required field for CrossOrderCancelReplaceRequest.
	EffectiveTime *time.Time `fix:"168"`
	//ExpireDate is a non-required field for CrossOrderCancelReplaceRequest.
	ExpireDate *string `fix:"432"`
	//ExpireTime is a non-required field for CrossOrderCancelReplaceRequest.
	ExpireTime *time.Time `fix:"126"`
	//GTBookingInst is a non-required field for CrossOrderCancelReplaceRequest.
	GTBookingInst *int `fix:"427"`
	//MaxShow is a non-required field for CrossOrderCancelReplaceRequest.
	MaxShow *float64 `fix:"210"`
	//PegDifference is a non-required field for CrossOrderCancelReplaceRequest.
	PegDifference *float64 `fix:"211"`
	//DiscretionInst is a non-required field for CrossOrderCancelReplaceRequest.
	DiscretionInst *string `fix:"388"`
	//DiscretionOffset is a non-required field for CrossOrderCancelReplaceRequest.
	DiscretionOffset *float64 `fix:"389"`
	//CancellationRights is a non-required field for CrossOrderCancelReplaceRequest.
	CancellationRights *string `fix:"480"`
	//MoneyLaunderingStatus is a non-required field for CrossOrderCancelReplaceRequest.
	MoneyLaunderingStatus *string `fix:"481"`
	//RegistID is a non-required field for CrossOrderCancelReplaceRequest.
	RegistID *string `fix:"513"`
	//Designation is a non-required field for CrossOrderCancelReplaceRequest.
	Designation *string `fix:"494"`
	//AccruedInterestRate is a non-required field for CrossOrderCancelReplaceRequest.
	AccruedInterestRate *float64 `fix:"158"`
	//AccruedInterestAmt is a non-required field for CrossOrderCancelReplaceRequest.
	AccruedInterestAmt *float64 `fix:"159"`
	//NetMoney is a non-required field for CrossOrderCancelReplaceRequest.
	NetMoney *float64 `fix:"118"`
	Trailer  fix43.Trailer
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
	return enum.BeginStringFIX43, "t", r
}
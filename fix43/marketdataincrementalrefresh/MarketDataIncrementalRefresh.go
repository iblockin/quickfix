//Package marketdataincrementalrefresh msg type = X.
package marketdataincrementalrefresh

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix43"
	"github.com/quickfixgo/quickfix/fix43/instrument"
	"time"
)

//NoMDEntries is a repeating group in MarketDataIncrementalRefresh
type NoMDEntries struct {
	//MDUpdateAction is a required field for NoMDEntries.
	MDUpdateAction string `fix:"279"`
	//DeleteReason is a non-required field for NoMDEntries.
	DeleteReason *string `fix:"285"`
	//MDEntryType is a non-required field for NoMDEntries.
	MDEntryType *string `fix:"269"`
	//MDEntryID is a non-required field for NoMDEntries.
	MDEntryID *string `fix:"278"`
	//MDEntryRefID is a non-required field for NoMDEntries.
	MDEntryRefID *string `fix:"280"`
	//Instrument Component
	Instrument instrument.Component
	//FinancialStatus is a non-required field for NoMDEntries.
	FinancialStatus *string `fix:"291"`
	//CorporateAction is a non-required field for NoMDEntries.
	CorporateAction *string `fix:"292"`
	//MDEntryPx is a non-required field for NoMDEntries.
	MDEntryPx *float64 `fix:"270"`
	//Currency is a non-required field for NoMDEntries.
	Currency *string `fix:"15"`
	//MDEntrySize is a non-required field for NoMDEntries.
	MDEntrySize *float64 `fix:"271"`
	//MDEntryDate is a non-required field for NoMDEntries.
	MDEntryDate *string `fix:"272"`
	//MDEntryTime is a non-required field for NoMDEntries.
	MDEntryTime *string `fix:"273"`
	//TickDirection is a non-required field for NoMDEntries.
	TickDirection *string `fix:"274"`
	//MDMkt is a non-required field for NoMDEntries.
	MDMkt *string `fix:"275"`
	//TradingSessionID is a non-required field for NoMDEntries.
	TradingSessionID *string `fix:"336"`
	//TradingSessionSubID is a non-required field for NoMDEntries.
	TradingSessionSubID *string `fix:"625"`
	//QuoteCondition is a non-required field for NoMDEntries.
	QuoteCondition *string `fix:"276"`
	//TradeCondition is a non-required field for NoMDEntries.
	TradeCondition *string `fix:"277"`
	//MDEntryOriginator is a non-required field for NoMDEntries.
	MDEntryOriginator *string `fix:"282"`
	//LocationID is a non-required field for NoMDEntries.
	LocationID *string `fix:"283"`
	//DeskID is a non-required field for NoMDEntries.
	DeskID *string `fix:"284"`
	//OpenCloseSettleFlag is a non-required field for NoMDEntries.
	OpenCloseSettleFlag *string `fix:"286"`
	//TimeInForce is a non-required field for NoMDEntries.
	TimeInForce *string `fix:"59"`
	//ExpireDate is a non-required field for NoMDEntries.
	ExpireDate *string `fix:"432"`
	//ExpireTime is a non-required field for NoMDEntries.
	ExpireTime *time.Time `fix:"126"`
	//MinQty is a non-required field for NoMDEntries.
	MinQty *float64 `fix:"110"`
	//ExecInst is a non-required field for NoMDEntries.
	ExecInst *string `fix:"18"`
	//SellerDays is a non-required field for NoMDEntries.
	SellerDays *int `fix:"287"`
	//OrderID is a non-required field for NoMDEntries.
	OrderID *string `fix:"37"`
	//QuoteEntryID is a non-required field for NoMDEntries.
	QuoteEntryID *string `fix:"299"`
	//MDEntryBuyer is a non-required field for NoMDEntries.
	MDEntryBuyer *string `fix:"288"`
	//MDEntrySeller is a non-required field for NoMDEntries.
	MDEntrySeller *string `fix:"289"`
	//NumberOfOrders is a non-required field for NoMDEntries.
	NumberOfOrders *int `fix:"346"`
	//MDEntryPositionNo is a non-required field for NoMDEntries.
	MDEntryPositionNo *int `fix:"290"`
	//Scope is a non-required field for NoMDEntries.
	Scope *string `fix:"546"`
	//TotalVolumeTraded is a non-required field for NoMDEntries.
	TotalVolumeTraded *float64 `fix:"387"`
	//TotalVolumeTradedDate is a non-required field for NoMDEntries.
	TotalVolumeTradedDate *string `fix:"449"`
	//TotalVolumeTradedTime is a non-required field for NoMDEntries.
	TotalVolumeTradedTime *string `fix:"450"`
	//NetChgPrevDay is a non-required field for NoMDEntries.
	NetChgPrevDay *float64 `fix:"451"`
	//Text is a non-required field for NoMDEntries.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for NoMDEntries.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for NoMDEntries.
	EncodedText *string `fix:"355"`
}

//Message is a MarketDataIncrementalRefresh FIX Message
type Message struct {
	FIXMsgType string `fix:"X"`
	Header     fix43.Header
	//MDReqID is a non-required field for MarketDataIncrementalRefresh.
	MDReqID *string `fix:"262"`
	//NoMDEntries is a required field for MarketDataIncrementalRefresh.
	NoMDEntries []NoMDEntries `fix:"268"`
	Trailer     fix43.Trailer
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
	return enum.BeginStringFIX43, "X", r
}
//Package marketdatasnapshotfullrefresh msg type = W.
package marketdatasnapshotfullrefresh

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix50sp2/applicationsequencecontrol"
	"github.com/quickfixgo/quickfix/fix50sp2/instrmtleggrp"
	"github.com/quickfixgo/quickfix/fix50sp2/instrument"
	"github.com/quickfixgo/quickfix/fix50sp2/mdfullgrp"
	"github.com/quickfixgo/quickfix/fix50sp2/routinggrp"
	"github.com/quickfixgo/quickfix/fix50sp2/undinstrmtgrp"
	"github.com/quickfixgo/quickfix/fixt11"
)

//Message is a MarketDataSnapshotFullRefresh FIX Message
type Message struct {
	FIXMsgType string `fix:"W"`
	Header     fixt11.Header
	//MDReqID is a non-required field for MarketDataSnapshotFullRefresh.
	MDReqID *string `fix:"262"`
	//Instrument Component
	Instrument instrument.Component
	//UndInstrmtGrp Component
	UndInstrmtGrp undinstrmtgrp.Component
	//InstrmtLegGrp Component
	InstrmtLegGrp instrmtleggrp.Component
	//FinancialStatus is a non-required field for MarketDataSnapshotFullRefresh.
	FinancialStatus *string `fix:"291"`
	//CorporateAction is a non-required field for MarketDataSnapshotFullRefresh.
	CorporateAction *string `fix:"292"`
	//NetChgPrevDay is a non-required field for MarketDataSnapshotFullRefresh.
	NetChgPrevDay *float64 `fix:"451"`
	//MDFullGrp Component
	MDFullGrp mdfullgrp.Component
	//ApplQueueDepth is a non-required field for MarketDataSnapshotFullRefresh.
	ApplQueueDepth *int `fix:"813"`
	//ApplQueueResolution is a non-required field for MarketDataSnapshotFullRefresh.
	ApplQueueResolution *int `fix:"814"`
	//MDReportID is a non-required field for MarketDataSnapshotFullRefresh.
	MDReportID *int `fix:"963"`
	//ClearingBusinessDate is a non-required field for MarketDataSnapshotFullRefresh.
	ClearingBusinessDate *string `fix:"715"`
	//MDBookType is a non-required field for MarketDataSnapshotFullRefresh.
	MDBookType *int `fix:"1021"`
	//MDFeedType is a non-required field for MarketDataSnapshotFullRefresh.
	MDFeedType *string `fix:"1022"`
	//TradeDate is a non-required field for MarketDataSnapshotFullRefresh.
	TradeDate *string `fix:"75"`
	//RoutingGrp Component
	RoutingGrp routinggrp.Component
	//MDSubBookType is a non-required field for MarketDataSnapshotFullRefresh.
	MDSubBookType *int `fix:"1173"`
	//MarketDepth is a non-required field for MarketDataSnapshotFullRefresh.
	MarketDepth *int `fix:"264"`
	//TotNumReports is a non-required field for MarketDataSnapshotFullRefresh.
	TotNumReports *int `fix:"911"`
	//RefreshIndicator is a non-required field for MarketDataSnapshotFullRefresh.
	RefreshIndicator *bool `fix:"1187"`
	//ApplicationSequenceControl Component
	ApplicationSequenceControl applicationsequencecontrol.Component
	//MDStreamID is a non-required field for MarketDataSnapshotFullRefresh.
	MDStreamID *string `fix:"1500"`
	Trailer    fixt11.Trailer
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
	return enum.ApplVerID_FIX50SP2, "W", r
}
//Package settlementobligationreport msg type = BQ.
package settlementobligationreport

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix50sp2/applicationsequencecontrol"
	"github.com/quickfixgo/quickfix/fix50sp2/settlobligationinstructions"
	"github.com/quickfixgo/quickfix/fixt11"
	"time"
)

//Message is a SettlementObligationReport FIX Message
type Message struct {
	FIXMsgType string `fix:"BQ"`
	Header     fixt11.Header
	//ClearingBusinessDate is a non-required field for SettlementObligationReport.
	ClearingBusinessDate *string `fix:"715"`
	//SettlementCycleNo is a non-required field for SettlementObligationReport.
	SettlementCycleNo *int `fix:"1153"`
	//SettlObligMsgID is a required field for SettlementObligationReport.
	SettlObligMsgID string `fix:"1160"`
	//SettlObligMode is a required field for SettlementObligationReport.
	SettlObligMode int `fix:"1159"`
	//Text is a non-required field for SettlementObligationReport.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for SettlementObligationReport.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for SettlementObligationReport.
	EncodedText *string `fix:"355"`
	//TransactTime is a non-required field for SettlementObligationReport.
	TransactTime *time.Time `fix:"60"`
	//SettlObligationInstructions Component
	SettlObligationInstructions settlobligationinstructions.Component
	//ApplicationSequenceControl Component
	ApplicationSequenceControl applicationsequencecontrol.Component
	Trailer                    fixt11.Trailer
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
	return enum.ApplVerID_FIX50SP2, "BQ", r
}
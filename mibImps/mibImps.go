package mibImps

import (
	"github.com/gophertribe/snmp"
	"github.com/gophertribe/snmp/mibImps/dismanEventMib"
	"github.com/gophertribe/snmp/mibImps/ifMib"
	"github.com/gophertribe/snmp/mibImps/ucdMib"
)

func init() {
	g_Logger = snmp.NewDiscardLogger()
}

var g_Logger snmp.ILogger

//SetupLogger Setups Logger for All sub mibs.
func SetupLogger(i snmp.ILogger) {
	g_Logger = i
	dismanEventMib.SetupLogger(i)
	ifMib.SetupLogger(i)
	ucdMib.SetupLogger(i)
}

// All function provides a list of common used OID
//    includes part of ucdMib, ifMib, and dismanEventMib
func All() []*snmp.PDUValueControlItem {
	toRet := []*snmp.PDUValueControlItem{}
	toRet = append(toRet, dismanEventMib.All()...)
	toRet = append(toRet, ifMib.All()...)
	toRet = append(toRet, ucdMib.All()...)
	return toRet
}

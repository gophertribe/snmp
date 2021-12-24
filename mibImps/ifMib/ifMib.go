package ifMib

import "github.com/gophertribe/snmp"

func init() {
	g_Logger = snmp.NewDiscardLogger()
}

var g_Logger snmp.ILogger

//SetupLogger Setups Logger for this mib
func SetupLogger(i snmp.ILogger) {
	g_Logger = i
}

// All function provides a list of common used OID in IF-MIB
func All() []*snmp.PDUValueControlItem {
	return NetworkOIDs()
}

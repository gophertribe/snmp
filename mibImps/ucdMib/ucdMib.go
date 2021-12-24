package ucdMib

import "github.com/gophertribe/snmp"

func init() {
	g_Logger = snmp.NewDiscardLogger()
}

var g_Logger snmp.ILogger

//SetupLogger Setups Logger for this mib
func SetupLogger(i snmp.ILogger) {
	g_Logger = i
}

// All function provides a list of common used OID in UCD-MIB
func All() []*snmp.PDUValueControlItem {
	var result []*snmp.PDUValueControlItem
	result = append(result, MemoryOIDs()...)
	result = append(result, SystemStatsOIDs()...)
	result = append(result, SystemLoadOIDs()...)
	result = append(result, DiskUsageOIDs()...)
	return result

}

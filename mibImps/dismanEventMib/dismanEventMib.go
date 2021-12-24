package dismanEventMib

import (
	"github.com/gophertribe/snmp"
	"github.com/gosnmp/gosnmp"
	"github.com/shirou/gopsutil/host"
)

func init() {
	g_Logger = snmp.NewDiscardLogger()
}

var g_Logger snmp.ILogger

//SetupLogger Setups Logger for this mib
func SetupLogger(i snmp.ILogger) {
	g_Logger = i
}

// DismanEventOids function provides sysUptime
//   see http://www.oid-info.com/get/1.3.6.1.2.1.1.3.0
//       http://www.net-snmp.org/docs/mibs/dismanEventMIB.html
func DismanEventOids() []*snmp.PDUValueControlItem {
	return []*snmp.PDUValueControlItem{
		{
			OID:  "1.3.6.1.2.1.1.3.0",
			Type: gosnmp.TimeTicks,
			OnGet: func() (value interface{}, err error) {
				if val, err := host.Uptime(); err != nil {
					return nil, err
				} else {
					return snmp.Asn1TimeTicksWrap(uint32(val)), nil
				}
			},
			Document: "Uptime",
		},
	}
}

// All function provides a list of common used OID in DISMAN-EVENT-MIB
func All() []*snmp.PDUValueControlItem {
	return DismanEventOids()
}

package vivado

import (
	"html/template"
	"sync"
)

var (
	tUpload = template.Must(template.New("upload").Parse(`open_hw
connect_hw_server
open_hw_target {{ .TARGET }}

current_hw_device [get_hw_devices]
refresh_hw_device -update_hw_probes false [lindex [get_hw_devices] 0]
# INFO: [Labtools 27-1434] Device xcvu9p (JTAG device index = 0) is programmed with a design that has no supported debu$
# WARNING: [Labtools 27-3361] The debug hub core was not detected.
# Resolution:
# 1. Make sure the clock connected to the debug hub (dbg_hub) core is a free running clock and is active.
# 2. Make sure the BSCAN_SWITCH_USER_MASK device property in Vivado Hardware Manager reflects the user scan chain setti$
# For more details on setting the scan chain property, consult the Vivado Debug and Programming User Guide (UG908).
set_property PROBES.FILE {} [get_hw_devices]
set_property FULL_PROBES.FILE {} [get_hw_devices]
set_property PROGRAM.FILE {{ .FILENAME }} [get_hw_devices]
program_hw_devices [get_hw_devices]
# INFO: [Labtools 27-3164] End of startup status: HIGH
# program_hw_devices: Time (s): cpu = 00:00:45 ; elapsed = 00:00:46 . Memory (MB): peak = 6272.652 ; gain = 0.000 ; fre$
refresh_hw_device [lindex [get_hw_devices] 0]
# INFO: [Labtools 27-1434] Device xcvu9p (JTAG device index = 0) is programmed with a design that has no supported debu$
# WARNING: [Labtools 27-3361] The debug hub core was not detected.
exit
`))
)

func (v *Vivado) uploadBitstream(bitstream, target string) error {
	return execTclTemplate(v.location, tUpload, map[string]string{
		"TARGET":   target,
		"FILENAME": bitstream,
	})
}

func (v *Vivado) UploadBitstream(bitstream string, targets ...string) error {
	var err error
	var wg sync.WaitGroup
	wg.Add(len(targets))
	for _, target := range targets {
		v.syncUpload(&wg, bitstream, target, &err)
	}
	wg.Wait()
	return err
}

func (v *Vivado) syncUpload(wg *sync.WaitGroup, bitstream, target string, ret *error) {
	if err := v.uploadBitstream(bitstream, target); err != nil {
		*ret = err
	}
	wg.Done()
}

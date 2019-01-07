package vivado

import (
	"os"
	"testing"
)

func TestTemplate(t *testing.T) {
	err := tUpload.Execute(os.Stdout, map[string]string{
		"TARGET":   "{localhost:3121/xilinx_tcf/Xilinx/SG1809000028A}",
		"FILENAME": "/root/Miner/DC-VCU1525-ZP-21-600/bitstream/DC-VCU1525-ZP-21-600.bit",
	})
	if err != nil {
		t.Log(err)
		t.Fail()
	}
}

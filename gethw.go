package vivado

import (
	"html/template"
	"regexp"

	utils "github.com/alivanz/go-utils"
)

var (
	tGetHw = template.Must(template.New("get_hw").Parse(`open_hw
connect_hw_server
set fp [open {{ .FOUTPUT }} w]
set all_luts   [get_hw_targets]
foreach lut $all_luts {puts $fp $lut}
close $fp
`))
	tFpgaId = regexp.MustCompile("/([A-Za-z0-9]+)$")
)

const (
	fGetHw = "target.txt"
)

func (v *Vivado) GetHwTargets() ([]string, error) {
	err := execTclTemplate(v.location, tGetHw, map[string]string{
		"FOUTPUT": fGetHw,
	})
	if err != nil {
		return nil, err
	}
	return utils.ReadFileLines(fGetHw)
}
func (v *Vivado) FpgaId() ([]string, error) {
	targets, err := v.GetHwTargets()
	if err != nil {
		return nil, err
	}
	return FpgaId(targets...)
}

func fpgaId(target string) (string, error) {
	if !tFpgaId.MatchString(target) {
		return "", WrongFormat
	}
	return tFpgaId.FindStringSubmatch(target)[1], nil
}
func FpgaId(targets ...string) ([]string, error) {
	var err error
	var id string
	out := make([]string, 0, len(targets))
	for _, target := range targets {
		id, err = fpgaId(target)
		if err != nil {
			return nil, err
		}
		out = append(out, id)
	}
	return out, nil
}

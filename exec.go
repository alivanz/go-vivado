package vivado

import (
	"html/template"
	"io"
	"os"
	"os/exec"
	"runtime"
)

func baseParam() []string {
	out := []string{
		"-nolog",
		"-nojournal",
	}
	if runtime.GOOS == "windows" {
		out = append(out, "-exec", "vivado")
	}
	return out
}

func execVivado(location string, addParams ...string) *exec.Cmd {
	if runtime.GOOS == "windows" {

	}
	cmd := exec.Command(
		location,
		append(baseParam(), addParams...)...,
	)
	cmd.Env = append(
		os.Environ(),
		"LC_ALL=C",
	)
	return cmd
}

func execBatch(location, source string) error {
	return execVivado(
		location,
		"-mode", "batch",
		"-source", source,
	).Run()
}

func execTcl(location string) (*exec.Cmd, io.WriteCloser, error) {
	cmd := execVivado(
		location,
		"-mode", "tcl",
	)
	stdin, err := cmd.StdinPipe()
	if err != nil {
		panic(err)
	}
	cmd.Stdout = stdout
	err = cmd.Start()
	if err != nil {
		return nil, nil, err
	}
	return cmd, stdin, nil
}

func execTclTemplate(location string, t *template.Template, data interface{}) error {
	cmd, stdin, err := execTcl(location)
	if err != nil {
		return err
	}
	err = templateWriteTo(stdin, t, data)
	if err != nil {
		return err
	}
	return cmd.Wait()
}

func templateWriteTo(w io.WriteCloser, t *template.Template, data interface{}) error {
	defer w.Close()
	return t.Execute(w, data)
}

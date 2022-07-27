package ui

import (
	"github.com/elek/abced/abcfile"
	"github.com/elek/abced/midi"
	"github.com/elek/abced/parser"
	"github.com/zeebo/errs/v2"
	"io/ioutil"
	"log"
	"os/exec"
	"strings"
)

func Show(e *Editor) error {
	c := e.GetABC()
	firstLine := strings.Split(c, "\n")[0]
	id := strings.TrimSpace(strings.Split(firstLine, ":")[1])

	err := ioutil.WriteFile("/tmp/a.abc", []byte(c), 0644)
	if err != nil {
		return err
	}

	command := exec.Command("abcm2ps", "-e", id, "-E", "-O", "/tmp/out", "-i", "/tmp/a.abc")
	x, err := command.CombinedOutput()
	if err != nil {
		return err
	}
	log.Println(string(x))
	command = exec.Command("evince", "/tmp/out001.eps")
	_, err = command.CombinedOutput()
	if err != nil {
		return err
	}
	return nil
}

func PlayExternal(e *Editor) error {
	c := e.GetABC()
	firstLine := strings.Split(c, "\n")[0]
	id := strings.TrimSpace(strings.Split(firstLine, ":")[1])
	err := ioutil.WriteFile("/tmp/a.abc", []byte(c), 0644)
	if err != nil {
		return err
	}

	command := exec.Command("abc2midi", "/tmp/a.abc", id, "-o", "/tmp/out.midi")
	out, err := command.CombinedOutput()
	if err != nil {
		return err
	}
	if strings.Contains(string(out), "Error") {
		lines := strings.Split(strings.Trim(string(out), "\n"), "\n")
		return errs.Errorf("%s", lines[len(lines)-1])
	}

	command = exec.Command("timidity", "/tmp/out.midi")
	_, err = command.CombinedOutput()
	if err != nil {
		return err
	}
	return nil
}

func PlayInternal(e *Editor) error {
	c := e.GetABC()
	t := abcfile.Tune{
		Raw: c,
	}
	parsed, err := parser.ParseTune(t)
	if err != nil {
		return err
	}
	err = midi.PlayTune(parsed)
	return err
}

package main

import (
	"encoding/binary"
	"bytes"
	"log"
	"os/exec"

	"github.com/xlab/pocketsphinx-go/sphinx"
	pulse "github.com/mesilliac/pulse-simple" // pulse-simple
)

var buffSize int

func readInt16(buf []byte) (val int16) {
	binary.Read(bytes.NewBuffer(buf), binary.LittleEndian, &val)
	return
}

func createStream() *pulse.Stream {
	ss := pulse.SampleSpec{pulse.SAMPLE_S16LE, 16000, 1}
	buffSize = int(ss.UsecToBytes(1 * 1000000))
	stream, err := pulse.Capture("pulse-simple test", "capture test", &ss)
	if err != nil {
		log.Panicln(err)
	}
	return stream
}

func listen(decoder *sphinx.Decoder) {
	stream := createStream()
	defer stream.Free()
	defer decoder.Destroy()
	buf := make([]byte, buffSize)
	var bits []int16

	log.Println("Listening...")

	for {
		_, err := stream.Read(buf)
		if err != nil {
			log.Panicln(err)
		}

		for i := 0; i < buffSize; i += 2 {
			bits = append(bits, readInt16(buf[i:i+2]))
		}

		process(decoder, bits)
		bits = nil
	}
}

func process(dec *sphinx.Decoder, bits []int16) {
	if !dec.StartUtt() {
		panic("Decoder failed to start Utt")
	}
	
	dec.ProcessRaw(bits, false, false)
	dec.EndUtt()
	hyp, score := dec.Hypothesis()
	
	if score > -2500 {
		log.Println("Predicted:", hyp, score)
		handleAction(hyp)
	}
}

func executeCommand(commands ...string) {
	cmd := exec.Command(commands[0], commands[1:]...)
	cmd.Run()
}

func handleAction(hyp string) {
	switch hyp {
		case "SLEEP":
		executeCommand("loginctl", "lock-session")
		
		case "WAKE UP":
		executeCommand("loginctl", "unlock-session")

		case "POWEROFF":
		executeCommand("poweroff")
	}
}

func main() {
	cfg := sphinx.NewConfig(
		sphinx.HMMDirOption("/usr/local/share/pocketsphinx/model/en-us/en-us"),
		sphinx.DictFileOption("6129.dic"),
		sphinx.LMFileOption("6129.lm"),
		sphinx.LogFileOption("commander.log"),
	)
	
	dec, err := sphinx.NewDecoder(cfg)
	if err != nil {
		panic(err)
	}

	listen(dec)
}

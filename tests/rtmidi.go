package main

import (
	"fmt"
	"log"
	"testing"
	"time"

	"go.cld.moe/rtmidi"
)

var printf = fmt.Printf

func main() {
	midiout, err := rtmidi.NewMIDIOutDefault()
	if err != nil {
		panic(err)
	}
	name, err := midiout.PortName(0)
	if err != nil {
		panic(err)
	}
	printf("Device: %s\n", name)
	if err = midiout.OpenPort(0, name); err != nil {
		panic(err)
	}

	// Bits:
	// Note Off:   1000
	// Note On:    1001
	// Aftertouch: 1010

	// xxxx0000 x is msg type
	// 0000nnnn n is channel
	// 0kkkkkkk k is note
	// 0vvvvvvv v is velocity

	send := func(msgtype rtmidi.Control, channel rtmidi.Channel, args ...byte) {
		if  msgtype & rtmidi.Midi_ControlMask != msgtype ||
			channel & rtmidi.Midi_ChannelMask != channel ||
			len(args) < 1 {
			printf("Err: Invalid Input!?\n")
		}

		buffer := []byte{(msgtype & rtmidi.Midi_ControlMask) | 
						 (channel & rtmidi.Midi_ChannelMask)}
		buffer = append(buffer, args...)

		for _,v := range(buffer) {
			printf("%08b", v)	
		}
		printf("\n")
		

		err := midiout.SendMessage(buffer)
		if err != nil {
			printf("Err: %s\n", err)
		}
	}

	for i := 0;i < 128;i++ {
		send(rtmidi.Midi_NoteOff, 9, 41, 100)
		send(rtmidi.Midi_ProgramChange, 9, byte(i))
		send(rtmidi.Midi_NoteOn, 9, 41, 100)
		time.Sleep(300 * time.Millisecond)
	}

	time.Sleep(3 * time.Second)
	midiout.Destroy()
}

func TestExampleMIDIIn_SetCallback(tst *testing.T) {
	in, err := rtmidi.NewMIDIInDefault()
	if err != nil {
		log.Fatal(err)
	}
	defer in.Destroy()
	if err := in.OpenPort(0, "RtMidi"); err != nil {
		log.Fatal(err)
	}
	defer in.Close()
	in.SetCallback(func(m rtmidi.MIDIIn, msg []byte, t float64) {
		log.Println(msg, t)
	})
	<-make(chan struct{})
}

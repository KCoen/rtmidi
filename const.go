package rtmidi

type Control = byte
type Instr   = int8
type DrumKey = int8
type DrumKit = int8
type Channel = uint8

const (
	Midi_ControlMask Control = 0b1111_0000
	Midi_ChannelMask Control = 0b0000_1111

	Midi_NoteOff            Control = 0b1000_0000
	Midi_NoteOn             Control = 0b1001_0000
	Midi_Polyphonic         Control = 0b1010_0000
	Midi_ControlChange      Control = 0b1011_0000
	Midi_ProgramChange      Control = 0b1100_0000
	Midi_ChannelPressure    Control = 0b1101_0000
	Midi_PitchBendChange    Control = 0b1110_0000
	Midi_ChannelModeMessage Control = 0b1011_0000
	Midi_SystemExclusive    Control = 0b1111_0000

	Control_Off uint8 = 0
	Control_On  uint8 = 127

	Control_BankSelectMSB             Control = 0
	Control_ModulationWheelMSB        Control = 1
	Control_BreathControllerMSB       Control = 2
	Control_FootPedalMSB              Control = 4
	Control_PortamentoTimeMSB         Control = 5
	Control_DataEntryMSB              Control = 6
	Control_VolumeMSB                 Control = 7
	Control_BalanceMSB                Control = 8
	Control_PanPositionMSB            Control = 10
	Control_ExpressionMSB             Control = 11
	Control_EffectControl1MSB         Control = 12
	Control_EffectControl2MSB         Control = 13
	Control_GeneralPurposeSlider1     Control = 16
	Control_GeneralPurposeSlider2     Control = 17
	Control_GeneralPurposeSlider3     Control = 18
	Control_GeneralPurposeSlider4     Control = 19
	Control_BankSelectLSB             Control = 32
	Control_ModulationWheelLSB        Control = 33
	Control_BreathControllerLSB       Control = 34
	Control_FootPedalLSB              Control = 36
	Control_PortamentoTimeLSB         Control = 37
	Control_DataEntryLSB              Control = 38
	Control_VolumeLSB                 Control = 39
	Control_BalanceLSB                Control = 40
	Control_PanPositionLSB            Control = 42
	Control_ExpressionLSB             Control = 43
	Control_EffectControl1LSB         Control = 44
	Control_EffectControl2LSB         Control = 45
	Control_SoundVariation            Control = 70
	Control_SoundTimbre               Control = 71
	Control_SoundReleaseTime          Control = 72
	Control_SoundAttackTime           Control = 73
	Control_SoundBrightness           Control = 74
	Control_SoundControl6             Control = 75
	Control_SoundControl7             Control = 76
	Control_SoundControl8             Control = 77
	Control_SoundControl9             Control = 78
	Control_SoundControl10            Control = 79
	Control_EffectsLevel              Control = 91
	Control_TremuloLevel              Control = 92
	Control_ChorusLevel               Control = 93
	Control_CelesteLevel              Control = 94
	Control_PhaserLevel               Control = 95
	Control_DataButtonIncrement       Control = 96
	Control_DataButtonDecrement       Control = 97
	Control_NonRegisteredParameterLSB Control = 98
	Control_NonRegisteredParameterMSB Control = 99
	Control_RegisteredParameterLSB    Control = 100
	Control_RegisteredParameterMSB    Control = 101

	Control_AllSoundOff       Control = 120 // send it with value of 0/Off
	Control_AllControllersOff Control = 121 // send it with value of 0/Off
	Control_AllNotesOff       Control = 123 // send it with value of 0/Off

	Control_OmniModeOff Control = 124 // send it with value of 0/Off
	Control_OmniModeOn  Control = 125 // send it with value of 0

	Control_MonoOperation Control = 126

	Control_PolyOperation Control = 127 // send it with value of 0

	Control_LocalKeyboardSwitch  Control = 122 // send it with value of 127/On or 0/Off
	Control_HoldPedalSwitch      Control = 64  // send it with value of 127/On or 0/Off
	Control_PortamentoSwitch     Control = 65  // send it with value of 127/On or 0/Off
	Control_SustenutoPedalSwitch Control = 66  // send it with value of 127/On or 0/Off

	Control_SoftPedalSwitch             Control = 67 // send it with value of 127/On or 0/Off
	Control_LegatoPedalSwitch           Control = 68 // send it with value of 127/On or 0/Off
	Control_Hold2PedalSwitch            Control = 69 // send it with value of 127/On or 0/Off
	Control_GeneralPurposeButton1Switch Control = 80 // send it with value of 127/On or 0/Off
	Control_GeneralPurposeButton2Switch Control = 81 // send it with value of 127/On or 0/Off
	Control_GeneralPurposeButton3Switch Control = 82 // send it with value of 127/On or 0/Off
	Control_GeneralPurposeButton4Switch Control = 83 // send it with value of 127/On or 0/Off

	DrumKit_Standard    DrumKit = 0
	DrumKit_Standard1   DrumKit = 1
	DrumKit_Standard2   DrumKit = 2
	DrumKit_Standard3   DrumKit = 3
	DrumKit_Standard4   DrumKit = 4
	DrumKit_Standard5   DrumKit = 5
	DrumKit_Standard6   DrumKit = 6
	DrumKit_Standard7   DrumKit = 7
	DrumKit_Room        DrumKit = 8
	DrumKit_Room1       DrumKit = 9
	DrumKit_Room2       DrumKit = 10
	DrumKit_Room3       DrumKit = 11
	DrumKit_Room4       DrumKit = 12
	DrumKit_Room5       DrumKit = 13
	DrumKit_Room6       DrumKit = 14
	DrumKit_Room7       DrumKit = 15
	DrumKit_Power       DrumKit = 16
	DrumKit_Power1      DrumKit = 17
	DrumKit_Power2      DrumKit = 18
	DrumKit_Power3      DrumKit = 19
	DrumKit_Power4      DrumKit = 20
	DrumKit_Power5      DrumKit = 21
	DrumKit_Power6      DrumKit = 22
	DrumKit_Power7      DrumKit = 23
	DrumKit_Electronic  DrumKit = 24
	DrumKit_Electronic1 DrumKit = 25
	DrumKit_Electronic2 DrumKit = 26
	DrumKit_Electronic3 DrumKit = 27
	DrumKit_Electronic4 DrumKit = 28
	DrumKit_Electronic5 DrumKit = 29
	DrumKit_Electronic6 DrumKit = 30
	DrumKit_Electronic7 DrumKit = 31
	DrumKit_Tr808       DrumKit = 25
	DrumKit_Jazz        DrumKit = 32
	DrumKit_Jazz1       DrumKit = 33
	DrumKit_Jazz2       DrumKit = 34
	DrumKit_Jazz3       DrumKit = 35
	DrumKit_Jazz4       DrumKit = 36
	DrumKit_Jazz5       DrumKit = 37
	DrumKit_Jazz6       DrumKit = 38
	DrumKit_Jazz7       DrumKit = 39
	DrumKit_Brush       DrumKit = 40
	DrumKit_Brush1      DrumKit = 41
	DrumKit_Brush2      DrumKit = 42
	DrumKit_Brush3      DrumKit = 43
	DrumKit_Brush4      DrumKit = 44
	DrumKit_Brush5      DrumKit = 45
	DrumKit_Brush6      DrumKit = 46
	DrumKit_Brush7      DrumKit = 47
	DrumKit_Orchestra   DrumKit = 48
	DrumKit_Orchestra1  DrumKit = 49
	DrumKit_Orchestra2  DrumKit = 50
	DrumKit_Orchestra3  DrumKit = 51
	DrumKit_Orchestra4  DrumKit = 52
	DrumKit_Orchestra5  DrumKit = 53
	DrumKit_Orchestra6  DrumKit = 54
	DrumKit_Orchestra7  DrumKit = 55
	DrumKit_SoundFX     DrumKit = 56
	DrumKit_SoundFX1    DrumKit = 57
	DrumKit_SoundFX2    DrumKit = 58
	DrumKit_SoundFX3    DrumKit = 59
	DrumKit_SoundFX4    DrumKit = 60
	DrumKit_SoundFX5    DrumKit = 61
	DrumKit_SoundFX6    DrumKit = 62
	DrumKit_SoundFX7    DrumKit = 63

	Instr_AcousticGrandPiano  Instr = 0
	Instr_BrightAcousticPiano Instr = 1
	Instr_ElectricGrandPiano  Instr = 2
	Instr_HonkytonkPiano      Instr = 3
	Instr_ElectricPiano1      Instr = 4
	Instr_ElectricPiano2      Instr = 5
	Instr_Harpsichord         Instr = 6
	Instr_Clavi               Instr = 7
	Instr_Celesta             Instr = 8
	Instr_Glockenspiel        Instr = 9
	Instr_MusicBox            Instr = 10
	Instr_Vibraphone          Instr = 11
	Instr_Marimba             Instr = 12
	Instr_Xylophone           Instr = 13
	Instr_TubularBells        Instr = 14
	Instr_Dulcimer            Instr = 15
	Instr_DrawbarOrgan        Instr = 16
	Instr_PercussiveOrgan     Instr = 17
	Instr_RockOrgan           Instr = 18
	Instr_ChurchOrgan         Instr = 19
	Instr_ReedOrgan           Instr = 20
	Instr_Accordion           Instr = 21
	Instr_Harmonica           Instr = 22
	Instr_TangoAccordion      Instr = 23
	Instr_AcousticGuitarNylon Instr = 24
	Instr_AcousticGuitarSteel Instr = 25
	Instr_ElectricGuitarJazz  Instr = 26
	Instr_ElectricGuitarClean Instr = 27
	Instr_ElectricGuitarMuted Instr = 28
	Instr_OverdrivenGuitar    Instr = 29
	Instr_DistortionGuitar    Instr = 30
	Instr_Guitarharmonics     Instr = 31
	Instr_AcousticBass        Instr = 32
	Instr_ElectricBassFinger  Instr = 33
	Instr_ElectricBassPick    Instr = 34
	Instr_FretlessBass        Instr = 35
	Instr_SlapBass1           Instr = 36
	Instr_SlapBass2           Instr = 37
	Instr_SynthBass1          Instr = 38
	Instr_SynthBass2          Instr = 39
	Instr_Violin              Instr = 40
	Instr_Viola               Instr = 41
	Instr_Cello               Instr = 42
	Instr_Contrabass          Instr = 43
	Instr_TremoloStrings      Instr = 44
	Instr_PizzicatoStrings    Instr = 45
	Instr_OrchestralHarp      Instr = 46
	Instr_Timpani             Instr = 47
	Instr_StringEnsemble1     Instr = 48
	Instr_StringEnsemble2     Instr = 49
	Instr_SynthStrings1       Instr = 50
	Instr_SynthStrings2       Instr = 51
	Instr_ChoirAahs           Instr = 52
	Instr_VoiceOohs           Instr = 53
	Instr_SynthVoice          Instr = 54
	Instr_OrchestraHit        Instr = 55
	Instr_Trumpet             Instr = 56
	Instr_Trombone            Instr = 57
	Instr_Tuba                Instr = 58
	Instr_MutedTrumpet        Instr = 59
	Instr_FrenchHorn          Instr = 60
	Instr_BrassSection        Instr = 61
	Instr_SynthBrass1         Instr = 62
	Instr_SynthBrass2         Instr = 63
	Instr_SopranoSax          Instr = 64
	Instr_AltoSax             Instr = 65
	Instr_TenorSax            Instr = 66
	Instr_BaritoneSax         Instr = 67
	Instr_Oboe                Instr = 68
	Instr_EnglishHorn         Instr = 69
	Instr_Bassoon             Instr = 70
	Instr_Clarinet            Instr = 71
	Instr_Piccolo             Instr = 72
	Instr_Flute               Instr = 73
	Instr_Recorder            Instr = 74
	Instr_PanFlute            Instr = 75
	Instr_BlownBottle         Instr = 76
	Instr_Shakuhachi          Instr = 77
	Instr_Whistle             Instr = 78
	Instr_Ocarina             Instr = 79
	Instr_Lead1Square         Instr = 80
	Instr_Lead2Sawtooth       Instr = 81
	Instr_Lead3Calliope       Instr = 82
	Instr_Lead4Chiff          Instr = 83
	Instr_Lead5Charang        Instr = 84
	Instr_Lead6Voice          Instr = 85
	Instr_Lead7Fifths         Instr = 86
	Instr_Lead8Basslead       Instr = 87
	Instr_Pad1Newage          Instr = 88
	Instr_Pad2Warm            Instr = 89
	Instr_Pad3Polysynth       Instr = 90
	Instr_Pad4Choir           Instr = 91
	Instr_Pad5Bowed           Instr = 92
	Instr_Pad6Metallic        Instr = 93
	Instr_Pad7Halo            Instr = 94
	Instr_Pad8Sweep           Instr = 95
	Instr_FX1Rain             Instr = 96
	Instr_FX2Soundtrack       Instr = 97
	Instr_FX3Crystal          Instr = 98
	Instr_FX4Atmosphere       Instr = 99
	Instr_FX5Brightness       Instr = 100
	Instr_FX6Goblins          Instr = 101
	Instr_FX7Echoes           Instr = 102
	Instr_FX8Scifi            Instr = 103
	Instr_Sitar               Instr = 104
	Instr_Banjo               Instr = 105
	Instr_Shamisen            Instr = 106
	Instr_Koto                Instr = 107
	Instr_Kalimba             Instr = 108
	Instr_Bagpipe             Instr = 109
	Instr_Fiddle              Instr = 110
	Instr_Shanai              Instr = 111
	Instr_TinkleBell          Instr = 112
	Instr_Agogo               Instr = 113
	Instr_SteelDrums          Instr = 114
	Instr_Woodblock           Instr = 115
	Instr_TaikoDrum           Instr = 116
	Instr_MelodicTom          Instr = 117
	Instr_SynthDrum           Instr = 118
	Instr_ReverseCymbal       Instr = 119
	Instr_GuitarFretNoise     Instr = 120
	Instr_BreathNoise         Instr = 121
	Instr_Seashore            Instr = 122
	Instr_BirdTweet           Instr = 123
	Instr_TelephoneRing       Instr = 124
	Instr_Helicopter          Instr = 125
	Instr_Applause            Instr = 126
	Instr_Gunshot             Instr = 127

	DrumKey_AcousticBassDrum DrumKey = 34
	DrumKey_BassDrum1        DrumKey = 35
	DrumKey_SideStick        DrumKey = 36
	DrumKey_AcousticSnare    DrumKey = 37
	DrumKey_HandClap         DrumKey = 38
	DrumKey_ElectricSnare    DrumKey = 39
	DrumKey_LowFloorTom      DrumKey = 40
	DrumKey_ClosedHiHat      DrumKey = 41
	DrumKey_HighFloorTom     DrumKey = 42
	DrumKey_PedalHiHat       DrumKey = 43
	DrumKey_LowTom           DrumKey = 44
	DrumKey_OpenHiHat        DrumKey = 45
	DrumKey_LowMidTom        DrumKey = 46
	DrumKey_HiMidTom         DrumKey = 47
	DrumKey_CrashCymbal1     DrumKey = 48
	DrumKey_HighTom          DrumKey = 49
	DrumKey_RideCymbal1      DrumKey = 50
	DrumKey_ChineseCymbal    DrumKey = 51
	DrumKey_RideBell         DrumKey = 52
	DrumKey_Tambourine       DrumKey = 53
	DrumKey_SplashCymbal     DrumKey = 54
	DrumKey_Cowbell          DrumKey = 55
	DrumKey_CrashCymbal2     DrumKey = 56
	DrumKey_Vibraslap        DrumKey = 57
	DrumKey_RideCymbal2      DrumKey = 58
	DrumKey_HiBongo          DrumKey = 59
	DrumKey_LowBongo         DrumKey = 60
	DrumKey_MuteHiConga      DrumKey = 61
	DrumKey_OpenHiConga      DrumKey = 62
	DrumKey_LowConga         DrumKey = 63
	DrumKey_HighTimbale      DrumKey = 64
	DrumKey_LowTimbale       DrumKey = 65
	DrumKey_HighAgogo        DrumKey = 66
	DrumKey_LowAgogo         DrumKey = 67
	DrumKey_Cabasa           DrumKey = 68
	DrumKey_Maracas          DrumKey = 69
	DrumKey_ShortWhistle     DrumKey = 70
	DrumKey_LongWhistle      DrumKey = 71
	DrumKey_ShortGuiro       DrumKey = 72
	DrumKey_LongGuiro        DrumKey = 73
	DrumKey_Claves           DrumKey = 74
	DrumKey_HiWoodBlock      DrumKey = 75
	DrumKey_LowWoodBlock     DrumKey = 76
	DrumKey_MuteCuica        DrumKey = 77
	DrumKey_OpenCuica        DrumKey = 78
	DrumKey_MuteTriangle     DrumKey = 79
	DrumKey_OpenTriangle     DrumKey = 80
)

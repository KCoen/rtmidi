Previous Go Midi libraries seemed to have various latency issues, because the latency imposed by the golang scheduler causes noticable latency.
This midi library is just a thin wrapper for https://github.com/thestk/rtmidi so users can manually deal with latency concerns.

Compatible with Linux and Windows, but requires CGO

// THE AUTOGENERATED LICENSE. ALL THE RIGHTS ARE RESERVED BY ROBOTS.

// WARNING: This file has automatically been generated on Thu, 08 Sep 2016 18:36:12 MSK.
// By http://git.io/cgogen. DO NOT EDIT.

package pm

/*
#cgo LDFLAGS: -lportmidi
#include "portmidi.h"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"

const (
	// DefaultSysexBufferSize as defined in pm/portmidi.h:123
	DefaultSysexBufferSize = 1024
	// HostErrorMsgLen as defined in pm/portmidi.h:196
	HostErrorMsgLen = uint32(256)
	// nodevice as defined in pm/portmidi.h:206
	nodevice = -1
	// FiltActive as defined in pm/portmidi.h:390
	FiltActive = (1 << 0x0E)
	// FiltSysex as defined in pm/portmidi.h:392
	FiltSysex = (1 << 0x00)
	// FiltClock as defined in pm/portmidi.h:394
	FiltClock = (1 << 0x08)
	// FiltPlay as defined in pm/portmidi.h:396
	FiltPlay = ((1 << 0x0A) | (1 << 0x0C) | (1 << 0x0B))
	// FiltTick as defined in pm/portmidi.h:398
	FiltTick = (1 << 0x09)
	// FiltFd as defined in pm/portmidi.h:400
	FiltFd = (1 << 0x0D)
	// FiltUndefined as defined in pm/portmidi.h:402
	FiltUndefined = FiltFd
	// FiltReset as defined in pm/portmidi.h:404
	FiltReset = (1 << 0x0F)
	// FiltRealtime as defined in pm/portmidi.h:406
	FiltRealtime = (FiltActive | FiltSysex | FiltClock | FiltPlay | FiltUndefined | FiltReset | FiltTick)
	// FiltNote as defined in pm/portmidi.h:409
	FiltNote = ((1 << 0x19) | (1 << 0x18))
	// FiltChannelAftertouch as defined in pm/portmidi.h:411
	FiltChannelAftertouch = (1 << 0x1D)
	// FiltPolyAftertouch as defined in pm/portmidi.h:413
	FiltPolyAftertouch = (1 << 0x1A)
	// FiltAftertouch as defined in pm/portmidi.h:415
	FiltAftertouch = (FiltChannelAftertouch | FiltPolyAftertouch)
	// FiltProgram as defined in pm/portmidi.h:417
	FiltProgram = (1 << 0x1C)
	// FiltControl as defined in pm/portmidi.h:419
	FiltControl = (1 << 0x1B)
	// FiltPitchbend as defined in pm/portmidi.h:421
	FiltPitchbend = (1 << 0x1E)
	// FiltMtc as defined in pm/portmidi.h:423
	FiltMtc = (1 << 0x01)
	// FiltSongPosition as defined in pm/portmidi.h:425
	FiltSongPosition = (1 << 0x02)
	// FiltSongSelect as defined in pm/portmidi.h:427
	FiltSongSelect = (1 << 0x03)
	// FiltTune as defined in pm/portmidi.h:429
	FiltTune = (1 << 0x06)
	// FiltSystemcommon as defined in pm/portmidi.h:431
	FiltSystemcommon = (FiltMtc | FiltSongPosition | FiltSongSelect | FiltTune)
)

// Error as declared in pm/portmidi.h:147
type Error int32

// Error enumeration from pm/portmidi.h:147
const (
	noerror            Error = iota
	nodata             Error = 0
	gotdata            Error = 1
	hosterror          Error = -10000
	invaliddeviceid    Error = -9999
	insufficientmemory Error = -9998
	buffertoosmall     Error = -9997
	bufferoverflow     Error = -9996
	badptr             Error = -9995
	baddata            Error = -9994
	internalerror      Error = -9993
	buffermaxsize      Error = -9992
)
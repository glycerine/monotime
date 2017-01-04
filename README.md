monotime
--------

Provides access to a fast monotonic clock source, to fill in the gap in the Go standard library, which lacks one: see https://github.com/golang/go/issues/12914. Don't use time.Now() in code that needs to time things or otherwise assume that time passes at a constant rate. Instead use monotime.Now().

On linux, this uses the same vdso as clock_gettime(CLOCK_MONOTONIC,...) to retreive a monotonic timestamp.

You cannot just eplace calls to time.Now() with time.Unix(0, int64(monotime.Now())), because monotime.Now() doesn't produce an actual timestamp value. Instead it is relative to the os boot time (on linux; or some arbitrary fixed point in time in the past on others).

Courtesy of https://github.com/aristanetworks/goarista and https://github.com/aristanetworks/goarista/blob/master/monotime/nanotime.go

Copyright (C) 2016  Arista Networks, Inc.
LICENSE: Apache2; see LICENSE file.

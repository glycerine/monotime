monotime
--------

Provides access to a fast monotonic clock source, to fill in the gap in the Go standard library, which lacks one: see https://github.com/golang/go/issues/12914. Don't use time.Now() in code that needs to time things or otherwise assume that time passes at a constant rate, instead use monotime.Now().

This uses clock_gettime(CLOCK_MONOTONIC,...) to retreive a monotonic timestamp.

Replace calls to time.Now() with time.Unix(0, int64(monotime.Now())).

Courtesy of https://github.com/glycerine/goarista and https://github.com/glycerine/goarista/blob/master/monotime/nanotime.go

Copyright (C) 2016  Arista Networks, Inc.
LICENSE: Apache2; see LICENSE file.

// Copyright 2013 Marc Weistroff. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

/*
Package log implements logging in a similar fashion than monolog/logbook

    Usage:

    import (
        "github.com/marcw/log"
    )

    // Will log to stdout every message where severity >= DEBUG
    h1 := log.NewStdoutLogger(DEBUG)
    // Will log to stderr every message where severity >= ERROR
    h2 := log.NewStderrLogger(ERR)

    logger := log.NewLogger("channel_name")

    // Will add to log lines some informations about the go runtime
    logger.PushProcessor(log.RuntimeProcessor)

    // Will output to stdout "This is debug"
    logger.Debug("This is debug")

    // Will output to both stdout and stderr "This is critical"
    logger.Critical("This is critical")
*/
package log

// Copyright 2013 Marc Weistroff. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

/*
gogol is a log package heavily inspired by monolog/logbook.

    Usage:

    import (
        log "github.com/marcw/gogol"
    )

    h1 := log.NewStdoutLogger(DEBUG)  // Will log to stdout every message where severity >= DEBUG
    h2 := log.NewStderrLogger(ERR)    // Will log to stderr every message where severity >= ERROR

    logger := log.NewLogger("channel_name", []log.Handler{h1, h2}, []log.Processor{})
    logger.PushProcessor(log.RuntimeProcessor) // Will add to log lines some informations about the go runtime
    logger.Debug("This is debug")              // Will output to stdout "This is debug"
    logger.Critical("This is critical")        // Will output to both stdout and stderr "This is critical"
*/
package gogol

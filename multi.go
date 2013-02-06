// Copyright 2013 Marc Weistroff. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package gogol

// MultiLogger creates a Logger that duplicates its writes to all the
// provided loggers
func MultiLogger(loggers ...Logger) Logger {
	return &multiLogger{loggers}
}

type multiLogger struct {
	loggers []Logger
}

func (logger *multiLogger) Close() (err error) {
	for _, l := range logger.loggers {
		if err := l.Close(); err != nil {
			return err
		}
	}

	return nil
}

func (logger *multiLogger) Write(b []byte) (int, error) {
	for _, l := range logger.loggers {
		if c, err := l.Write(b); err != nil {
			return c, err
		}
	}

	return len(b), nil
}

func (logger *multiLogger) Warning(m string) (err error) {
	for _, l := range logger.loggers {
		if err := l.Warning(m); err != nil {
			return err
		}
	}

	return nil
}

func (logger *multiLogger) Notice(m string) (err error) {
	for _, l := range logger.loggers {
		if err := l.Notice(m); err != nil {
			return err
		}
	}

	return nil
}

func (logger *multiLogger) Info(m string) (err error) {
	for _, l := range logger.loggers {
		if err := l.Info(m); err != nil {
			return err
		}
	}

	return nil
}

func (logger *multiLogger) Err(m string) (err error) {
	for _, l := range logger.loggers {
		if err := l.Err(m); err != nil {
			return err
		}
	}

	return nil
}

func (logger *multiLogger) Emerg(m string) (err error) {
	for _, l := range logger.loggers {
		if err := l.Emerg(m); err != nil {
			return err
		}
	}

	return nil
}

func (logger *multiLogger) Debug(m string) (err error) {
	for _, l := range logger.loggers {
		if err := l.Debug(m); err != nil {
			return err
		}
	}

	return nil
}

func (logger *multiLogger) Crit(m string) (err error) {
	for _, l := range logger.loggers {
		if err := l.Crit(m); err != nil {
			return err
		}
	}

	return nil
}

func (logger *multiLogger) Alert(m string) (err error) {
	for _, l := range logger.loggers {
		if err := l.Alert(m); err != nil {
			return err
		}
	}

	return nil
}

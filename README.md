# Gogol

A `log/syslog` interface to `log` for the [go language](http://golang.org).

## Why?

Current go's implementation of `log.Logger` and `syslog.Writer` do not share a
common interface. Sometimes, you might want to combine the syslog output to a
stderr/stdout/whatever output. Gogol is here to solve this particular problem.

## Documentation

Documentation is available at [godoc.org](http://godoc.org/github.com/marcw/gogol)

## License

The Gogol code is free to use and distribute, under the [MIT
license](https://github.com/marcw/gogol/blob/master/LICENSE).

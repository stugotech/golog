# Golog

A simple no-nonsense logging library.

## Documentation

See the [godoc page](https://godoc.org/github.com/stugotech/golog).

## Example

    logger = golog.NewPackageLogger()

    logger.Info("Hello world", 
      golog.String("package", "golog"),
      golog.Int32("answer", 42)
    )

Output:

    github.com/acme/package/main.go:12 [INFO] Hello world
            package = golang
            answer = 42

## Configuration

The configuration may be altered by setting the members of the Config struct.  The default is as follows:

    golog.CurrentConfig = golog.Config{
      IncludeCallerInfo: true,
      Level:             golog.LevelDebug,
      Writer:            golog.ConsoleWriter,
    }
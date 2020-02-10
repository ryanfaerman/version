# Version

This is a simple package where most of its values should be defined at compile time.

If you don't use LDFLAGs already, you should! They are pretty neat.
Essentially, the `-ldflags` flag can be used to change the value of variables
while your application compiles. This makes it possible to inject build
variables into your binary at compile time.

## Motivation

In most of my tooling, I found myself re-implementing or copying packages
similar to this. It's just a collection of version data that supports the
Stringer interface. This makes it pretty easy to share version data by using
LDFLAGs at compile time.

# Usage

The correct syntax for an LDFlag is the following:

```bash
go build -ldflags '-X "var=val with spaces"'
```

For this package, that means something resembling:


```bash
go build -ldflags="-X github.com/ryanfaerman/version.BuildTag=v0.0.9 -X github.com/ryanfaerman/version.CommitHash=fdcad803 -X github.com/ryanfaerman/version.ApplicationName=my-awesome-app -X github.com/ryanfaerman/version.BuildDate=`date -u +%Y-%m-%d-%H:%M:%S`"
```

See `_examples/*` for some usage examples.

### Available Variables

The following variables can be modified with an ldflag:

 - `ApplicationName`
 - `CommitHash`
 - `BuildDate`
 - `BuildTag`
 - `Template`

See `version.Info` for the variables available to your template.


JsonRouter
==========

JSON Router is a simple library to construct a routing table mapping
[JsonPath](http://jsonpath.com/) expressions to target strings.

Check [samples/config.json]() contains an example configuration of the
routing table.

The CLI reads a JSON from STDIN and takes a routing table from a
configuration file.  To invoke it:

    cat samples/data.json | go run cmd/main.go -routes=samples/config.json

## Use cases

Connect an event stream to multiple FAAS endpoints that handle a subset
of the events in the stream.

In this case, the targets would be each endpoint URI.  JsonRouter allows
defining a set of routes based on the JSON payload.  If the event's
payload matches a route, it can be forwarded to the target.

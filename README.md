# Logo
**Attention: very early stage**

A simple logging library for golang.

## Techs

### Logger
The Logger wraps multiple Receivers.

### Receivers
A Receiver holds a `io.Writer` interface which it writes to. Also attributes if the output should be colored, if the Receiver is active and which Level should be logged.

## Todo
- Support multiple arguments and format methods
- Custom log format
- Tests

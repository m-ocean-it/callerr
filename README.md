# `callerr`

*Automatically annotate errors with method names.*

*Just call `callerr.Wrap(err)`.*

## Example

`./pkg/barer.go`:
```go
package pkg

import "github.com/m-ocean-it/callerr"

type Barer struct {}

func (b *Barer) Bar() {
    return callerr.New("oops")
}
```

`./main.go`:
```go
package main

import "github.com/m-ocean-it/callerr"

func main() {
    err := foo()
    fmt.Println(err) // main.foo: pkg.Barer.Bar: oops
}

func foo() {
    barer := pkg.Barer{}
    err := barer.Bar()
    
    return callerr.Wrap(err)
}
```

## Available functions

`New` creates a new error with the supplied message and annotates it with the caller's name.
```go
func New(msg string) error
```

`Wrap` annotates an error with the caller's name.
```go
func Wrap(err error) error
```

`WrapWithMsg` annotates an error with the caller's name and the supplied message.
```go
func WrapWithMsg(err error, msg string) error
```

`WithMsg` annotates an error with the supplied message without the caller's name.
```go
func WithMsg(err error, msg string) error
```

## Edge cases

### Generics miss their type parameters in the error messages

The output of the following code
```go
func try() {
        tryer := pkg.Tryer[int]{}
        fmt.Println(callerr.Wrap(tryer.TryToConnect()))
}
```
might look like this:
```
main.try: pkg.(*Tryer[...]).TryToConnect: ...
```

### Function inlining does not affect the resulting error messages

The library uses the `runtime.CallersFrames` method that accounts for function inlining, so it's not a factor here.

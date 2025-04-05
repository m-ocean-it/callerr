# `werror`

*Automatically annotate errors with method names. Just call `Wrap(err)`.*

## Example

`./pkg/barer.go`:
```go
package pkg

type Barer struct {}

func (b *Barer) Bar() {
    return werror.New("oops")
}
```

`./main.go`:
```go
package main 

func main() {
    err := foo()
    fmt.Println(err) // main.foo: pkg.Barer.Bar: oops
}

func foo() {
    barer := pkg.Barer{}
    err := barer.Bar()
    
    return werror.Wrap(err)
}
```
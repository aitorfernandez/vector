# vector

Module to describe a three dimensional vector

## Installation

To install it, run:

```sh
go get github.com/aitorfernandez/vector
```

## Tests

To run it:

```sh
go test
```

## Example

```go
package main

import (
  "fmt"

  vector "github.com/aitorfernandez/vector"
)

func main() {
  v := vector.Vec3{X: 1, Y: 2, Z: 3}
  fmt.Println(v.ToString()) // Vec3 struct: 1.00, 2.00, 3.00

  // methods

  v.Add(3)
  fmt.Println(v.ToString()) // Vec3 struct: 4.00, 2.00, 3.00
  // or
  v.Add(vector.Vec3{X: 4, Y: 5, Z: 6})
  fmt.Println(v.ToString()) // Vec3 struct: 8.00, 7.00, 9.00

  v.Sub(2, 2)
  fmt.Println(v.ToString()) // Vec3 struct: 6.00, 5.00, 9.00
  // or
  v.Sub(vector.Vec3{X: 1, Y: 1, Z: 1})
  fmt.Println(v.ToString()) // Vec3 struct: 5.00, 4.00, 8.00

  v.Mult(3)
  fmt.Println(v.ToString()) // Vec3 struct: 15.00, 12.00, 24.00

  v.Div(2)
  fmt.Println(v.ToString()) // Vec3 struct: 7.50, 6.00, 12.00

  mag := v.Magnitude()
  fmt.Println(mag) // 15.370426148939398

  v.Normalize()
  fmt.Println(v.ToString()) // Vec3 struct: 0.49, 0.39, 0.78
}
```

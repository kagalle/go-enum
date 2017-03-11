# go-enum
Basic type-safe enum for golang in the spirit of a Java enum.

In Java, an enum is a type safe langage construct where the elements are an enum are related to one another by the enum's type.
```java
public enum Apple {
    RED, GREEN, YELLOW
}

```

In Go, there is no enum type.  Instead, constant ints are used and `iota` is used to initialize the constant values within a `const` block, which is similar to how a Java enum wraps a set of constant values.

```go
const (
	redApple    = iota  // redApple    == 0
	greenApple  = iota  // greenApple  == 1
	yellowApple = iota  // yellowApple == 2
)
```

The issue with this is that one int constant is no different than any other int constant, as seen below:
```
package main

import (
	"fmt"
	"reflect"
)

const (
	apple0 = iota  // apple0 == 0
	apple1 = iota  // apple1 == 1
	apple2 = iota  // apple2 == 2
)

const ( // iota is reset to 0
	orange0 = iota  // orange0 == 0
	orange1 = iota  // orange1 == 1
	orange2 = iota  // orange2 == 2
)

func main() {
	fmt.Printf("Type of apple0: %s; type of orange0: %s\n", reflect.TypeOf(apple0), reflect.TypeOf(orange0))
	AcceptApple(orange0)
}

func AcceptApple(a int) {
	fmt.Printf("a should be an apple, but it is actually an orange: %d\n", a)
}
```
Java's enum is a special type where all enums inherit from `java.lang.Enum`.  This provides type safety.

Go-enum does the same thing using `struct` as a basis and an interface to define the available enum operations.

See `enum_test.go` for an example of useage.

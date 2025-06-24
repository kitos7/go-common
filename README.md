# Go Common

A collection of useful Go utilities and helpers built with Go 1.18, leveraging Go's generics.

## Installation

```bash
go get github.com/kitos7/go-common
```

## Requirements

- Go 1.18 or higher (for generics support)

## Features

### Slices Package

The `slices` package provides generic slice manipulation utilities and a generic Set implementation:

```go
import "github.com/kitos7/go-common/slices"
```

#### Set

A generic set implementation using maps:

```go
// Create a new set
set := slices.NewSet[string]()

// Add values
set.Set("item1")
set.Set("item2")

// Check if value exists
if set.Has("item1") {
    // Do something
}

// Remove a value
set.Delete("item1")

// Convert set to slice
items := set.ToSlice()
```

#### Slice Utilities

Perform functional operations on slices:

```go
// Map: transform each element
numbers := []int{1, 2, 3, 4, 5}
doubled := slices.Map(numbers, func(n int) int {
    return n * 2
})
// doubledNumbers = [2, 4, 6, 8, 10]

// Filter: keep only elements that match a condition
even := slices.Filter(numbers, func(n int) bool {
    return n%2 == 0
})
// even = [2, 4]

// Contains: check if a value exists in a slice
hasThree := slices.Contains(numbers, 3) // true

// NotContains: check if a value doesn't exist in a slice
hasSix := slices.NotContains(numbers, 6) // true

// ToMap: convert a slice to a map
users := []struct{ ID int; Name string }{{1, "Alice"}, {2, "Bob"}}
userMap := slices.ToMap(users, func(u struct{ID int; Name string}) (int, string) {
    return u.ID, u.Name
})
// userMap = map[1:"Alice", 2:"Bob"]
```

### Logger Package

The `logger` package provides a context-aware logging interface built on top of Go's `slog` package:

```go
import "github.com/kitos7/go-common/logger"
```

#### Features

- Context-aware logging with `slog`
- Key-value logging with `InfoKV`, `ErrorKV`, etc.
- Printf-style logging with `Infof`, `Errorf`, etc.

```go
// Create a logger and attach it to context
logr := slog.New(slog.NewJSONHandler(os.Stdout, nil))
ctx := logger.WithLogger(context.Background(), logr)

// Log with key-value pairs
logger.InfoKV(ctx, "User logged in", "user_id", 123, "source", "api")

// Log with printf formatting
logger.Infof(ctx, "Processing item %d of %d", current, total)

// Get logger from context
logr = logger.FromContext(ctx)
```

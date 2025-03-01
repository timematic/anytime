# anytime

a user friendly `time.Time` parser which no need specify the time `Layout`.

## Example

```
package main

import (
	"fmt"

	"github.com/longqimin/anytime"
)

func main() {
	datetime, err := anytime.Parse("2006-01-02T15:04:05+08")
	if err != nil {
		panic(err)
	}
	fmt.Println(datetime) // 2006-01-02 15:04:05 +0800 CST
}
```

## APIs

### `anytime.Parse(value string) (time.Time, error)`

automatically figure out the time `Layout` and parse to `time.Time`.
**possible implementation**:

```go
Parse(value string) (time.Time, error){
    layout, err := ExtractLayout(str)
    if err != nil {
        return time.Time{}, err
    }
    return time.Parse(layout, value)
}
```

### `anytime.ParseInLocation(value string, loc *time.Location) (time.Time, error)`

**possible implementation**:

```go
ParseInLocation(value string, loc *time.Location) (time.Time, error){
    layout, err := ExtractLayout(str)
    if err != nil {
        return time.Time{}, err
    }
    return time.ParseInLocation(layout, value, loc)
}
```

## Benchmark

|                         | ns/op       |
| ----------------------- | ----------- |
| time.Parse date_only    | 38.98 ns/op |
| anytime.Parse date_only | 28.14 ns/op |
| time.Parse rfc3339      | 26.79 ns/op |
| anytime.Parse rfc3339   | 52.53 ns/op |

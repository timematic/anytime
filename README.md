# anytime

a high performance `time.Time` parser without specify the layout.

## Usage

```
package main

import (
    "github.com/longqimin/anytime"
)

func main() {
    datetime, err := anytime.Parse("2006-01-02T15:04:05+08")
    if err != nil {
        panic(err)
    }
    println(datetime) // 2006-01-02 15:04:05 +0800 CST
}

```

## Benchmark

```
BenchmarkStdParse_DateOnly-4       	 9190088	       123.9 ns/op
BenchmarkAnytimeParse_DateOnly-4   	14495055	        77.98 ns/op
BenchmarkStdParse_RFC3339-4        	16821751	        67.50 ns/op
BenchmarkAnytimeParse_RFC3339-4    	 7461024	       160.0 ns/op
```

## APIs

### `anytime.Parse(value string) (time.Time, error)`

automatically figure out the time `Layout` and parse to `time.Time`.
possible implementation:

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

automatically figure out the time `Layout` and parse to `time.Time`.
possible implementation:

```go
Parse(value string, loc *time.Location) (time.Time, error){
    layout, err := ExtractLayout(str)
    if err != nil {
        return time.Time{}, err
    }
    return time.ParseInLocation(layout, value, loc)
}
```

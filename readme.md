# Will write some thing
## example
``` go
func main() {
    // Use backoff retry by default
    err := retry4go.Do(
        doSomething(),
        retry4go.MaxRetryTimes(10),
        retry4go.Retryable(retryable),
        retry4go.Interval(100 * time.Microsecond),
        retry4go.MaxJitterInterval(100 * time.Microsecond),
        retry4go.Multiplier(2.0)
    )
    fmt.Println(err)
}

func doSomething() error {
    return retryableErr
}

func retryable(err error) bool {
    if err == retryableErr {
        return true
    }
    return false
}
```

## default parameters
``` go
// Required for all types.
DefaultInterval        = 100.0 * time.Millisecond
DefaultMaxInterval     = 1000.0 * time.Millisecond
DefaultMaxRetryTimes   = 3
DefaultRetryable = func(error) bool // return false by default

// You can use backoff or regular type.
DefaultRetryType       = BackOffRetry

// Only required for backoff.
DefaultMultiplier      = 2.0
DefaultRandomFactor    = 0.5

// Only required for regular.
DefaultRegularInterval = DefaultInterval
DefaultJitterInterval  = DefaultInterval
```


## will write about type and interface
# Retry
[![codecov](https://codecov.io/gh/takuoki/clmconv/branch/master/graph/badge.svg)](https://codecov.io/gh/rossy0213/retry)
Retry with context or not.
Default using exponential backoff.

## Example
``` go
func DoSomething() error {
	err := retry.Do( // use DoWithContext() if you need context
		doSomeThing(),
		retry.ChckeRetryable(checkRetryable),
		retry.MaxRetryTimes(uint(5)),
		retry.Interval(100.0*time.Millisecond),
		retry.MaxInterval(5000.0*time.Millisecond),
		retry.MaxJitterInterval(10.0*time.Millisecond),
		retry.MaxElapsedTime(10*time.Minute),
		retry.Multiplier(1.5),
	)
	return err
}

func doSomething() error {
    return retryableErr
}

func checkRetryable(err error) bool {
    if err == retryableErr {
        return true
    }
    return false
}
```

### Default parameters
``` go
DefaultMaxRetryTimes  = 10
DefaultInterval       = 100.0 * time.Millisecond
DefaultMaxInterval    = 1000.0 * time.Millisecond
DefaultJitterInterval = 30.0 * time.Millisecond
DefaultMultiplier     = 2.0
DefaultMaxElapsedTime = 5 * time.Minute
```

## Reference
[backoff](https://github.com/cenkalti/backoff/)

[retry-go](https://github.com/avast/retry-go)

[retry](k8s.io/client-go/util/retry)
[gregoryv/testnet](https://pkg.go.dev/github.com/gregoryv/testnet) - In memory network connection for testing

## Quick start

```
import "github.com/gregoryv/testnet"

...

conn, srvconn := testnet.Dial("tcp", "somehost:1234")

// closing one side closes the other
defer conn.Close()
```


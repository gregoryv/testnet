testnet - In memory network connection for testing

## Quick start

```
import "github.com/gregoryv/testnet"

...

clientIO, serverIO := testnet.Dial("tcp", "somehost:1234") // the values do not matter

defer clientIO.Close()
```


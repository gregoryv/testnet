[gregoryv/testnet](https://pkg.go.dev/github.com/gregoryv/testnet) - In memory network connection for testing

## Quick start

```
import "github.com/gregoryv/testnet"

...

client, server := testnet.Dial("tcp", "somehost:1234") 

defer client.Close()
defer server.Close()
```


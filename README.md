# Hamilton Commisioning Registry

This is a golang library for getting the encryption keys for a hamilton.

Typical usage:

```
import "github.com/immesys/hcr"
...

  key := os.Getenv("DEPLOYMENT_READ_KEY")
  info, err := hcr.GetMoteInfo(moteid, key)
  // handle err
  some_decrypt(info.AESK, ...)
```

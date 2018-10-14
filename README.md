# Intdian

Determine the **int**ernal en**dian**ness of a machine in go without a hassle

## Example

Replace:

```go
binary.BigEndian
```

with:

```go
import (
    "github.com/kg6zvp/go-intdian"
)

intdian.ByteOrder()
```

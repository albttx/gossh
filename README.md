# GOSSH

This package is simply running a ssh connection

## Feature

- [x] Arrow Keys
- [x] Ctrl- Keys
- [x] SSH connection with password
- [x] SSH connection with private keys

## Exemple

```go
// leave pass empty for connection with ssh keys
err := gossh.Prompt("user", "pass", "host", "port")
if err != nil {
    fmt.Println(err)
    os.Exit(1)
}
```
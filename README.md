# GOSSH

This package is simply running a ssh connection

## Feature

- [x] Arrow Keys
- [x] Ctrl- Keys
- [x] SSH connection with password
- [ ] SSH connection with private keys

## Exemple

```go
err := gossh.Prompt("user", "pass", "host", "port")
if err != nil {
    fmt.Println(err)
    os.Exit(1)
}
```
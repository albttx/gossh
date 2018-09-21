# GOSSH

This package simply prompt a ssh connection.

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

[![asciicast](https://asciinema.org/a/E1MswnMqQcVakjy3qU6RD4nuk.png)](https://asciinema.org/a/E1MswnMqQcVakjy3qU6RD4nuk)
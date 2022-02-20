kv
==

:old_key: A simple key-value store CLI.

Install:
```
go install github.com/allanbreyes/kv@latest
```

Usage:
```
$ kv set foo bar
bar

$ kv get foo
bar
```

## To-do

- [ ] Actually delete keys
- [ ] Add tests
- [ ] Encryption at rest?
- [ ] Switch to a faster database?

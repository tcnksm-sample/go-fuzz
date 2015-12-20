# Go-Fuzz

Play with [go-fuzz](https://github.com/dvyukov/go-fuzz).

```bash
$ go-fuzz-build github.com/tcnksm-sample/go-fuzz
$ go-fuzz -bin=nicepkg-fuzz.zip -workdir=workdir
```

### Example output

```bash
2015/12/20 21:23:54 slaves: 4, corpus: 3 (2s ago), crashers: 1, restarts: 1/0, execs: 0 (0/sec), cover: 0, uptime: 3s
```

## Author

[Taichi Nakashima](https://github.com/tcnksm)

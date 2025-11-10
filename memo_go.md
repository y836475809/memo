# sample11 debug mingw cgo
- https://doss.eidos.ic.i.u-tokyo.ac.jp/html/gdb_step_into_libraries.html
- https://stackoverflow.com/questions/9607155/make-gcc-put-relative-filenames-in-debug-information
- https://kazuhira-r.hatenablog.com/entry/2021/02/13/234024
- https://devlights.hatenablog.com/entry/2025/01/16/073000
- https://stackoverflow.com/questions/65567411/how-can-i-influence-the-the-source-directory-of-an-executable-or-library

```
$ gcc -g -shared -fPIC -o libcalc.dll calc.c
$ objcopy --only-keep-debug libcalc.dll libcalc.dll.debug
```

```
$ go build 
$ gdb ./debug_dll.exe 
(gdb) add-symbol-file ./libcalc.dll.debug
```

- set breakpoint -> strat -> stop -> c
  - Cannot insert breakpoint Cannot access memory at address
```
(gdb) info sources calc
(gdb) b ~calc.c:12
(gdb) start
(gdb) c
```
- strat -> stop -> set breakpoint -> c
```
(gdb) start
(gdb) info sources calc
(gdb) b ~calc.c:12
(gdb) c
```

- https://zenn.dev/saitoyutaka/articles/dbd8076aeca09a
- https://h-noson.hatenablog.jp/entry/2017/12/22/000000
- https://stackoverflow.com/questions/44368703/gdb-cannot-insert-breakpoint-cannot-access-memory-at-address-xxx
# Safe Buffer [![Build Status](https://travis-ci.org/dmathieu/safe_buffer.svg?branch=master)](https://travis-ci.org/dmathieu/safe_buffer)

A very simple buffer adding a mutex between each read and write operation so they can be used without causing race conditions.  
I am using this as a `os.Stdin` mock in tests.

## Usage

```go
stdin := safe_buffer.NewMock()
```

Then, use the mock as a normal ReadWriter.

```go
stdin.Write([]byte("hello"))
ioutil.ReadAll(stdin)
```

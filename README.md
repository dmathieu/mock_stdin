# Mock Stdin [![Build Status](https://travis-ci.org/dmathieu/mock_stdin.svg?branch=master)](https://travis-ci.org/dmathieu/mock_stdin)

This is a very simple library one can use to get a race-conditin free mock of `os.Stdin` in their tests.  
I have been using a plain `bytes.Buffer` in other tests, but that causes race condition issues, as a read and a write can happen at the same time.

## Usage

```go
stdin := mock_stdin.NewMock()
```

Then, use the mock as a normal ReadWriter.

```go
stdin.Write([]byte("hello"))
ioutil.ReadAll(stdin)
```

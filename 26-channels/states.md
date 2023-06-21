What happens if you try to receive, send or close on different kinds of channels?

```
    STATE       RECEIVE         SEND           CLOSE
    nil         block(*)        block(*)       panic
    empty       block           write          close
    p. full     read            write          read until empty
    full        read            block          read until empty
    closed      default(**)     panic          panic
    ---------------------------------------------------
    <- ch       OK                  compile error
    ch <-       comp. err                 OK

```

_Note_: do not buffer until analysing/needed (may hide a race condition)
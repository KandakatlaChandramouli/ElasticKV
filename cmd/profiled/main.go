package main

import (
    "log"
    "net/http"
    _ "net/http/pprof"
)

func main() {

    log.Println(
        "pprof listening on :6060",
    )

    err := http.ListenAndServe(
        ":6060",
        nil,
    )

    if err != nil {
        panic(err)
    }
}

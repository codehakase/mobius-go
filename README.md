
# Mobius Go API Client
[![GoDoc](https://godoc.org/gopkg.in/mailgun/mailgun-go.v1?status.svg)](https://godoc.org/gopkg.in/mailgun/mailgun-go.v1)

The Mobius Go (Golang) Client provides simple access to the Mobius API for applications written in Go

## Installation
Install with `go get`
```shell
$ go get github.com/codehakase/mobius-go
```

## QuickStart
To use the library, you'd need to create a client, with your Mobius `appuid` and `apikey`.

```go
package main

import (
  "github.com/codehakase/mobius-go"
  "fmt"
)

func main() {
  mc := mobius.New(yourApiKey, yourAppUID)

  // Retrive a struct to communicate with the DApp store
  appStore := mc.AppStore

  // Get a user's balance
  myBalance, err := appStore.Balance("myemail@example.com")
  if err != nil {
    log.Fatalf("can't fetch user's balance, err: %+v", err)
    return
  }
  fmt.Println("User got %d mobi credits available", usrBal.NumCredits)

  // Credit user with mobi credits
  if fundUser, err := appStore.Credit("user@example.com", 1000); err != nil {
    log.Fatalf("could not fund user, err: %+v", err)
    return
  }
  // Use 20 Mobi Credits from user@example.com
  if charge, err := appStore.Use("user@example.com", 20); err == nil {
    if charge.Success {
      fmt.Printf("User has been charged, and is left with %d mobi credits", charge.NumCredits)
    }
  } else {
    log.Fatalf("could not charge user, err: %+v", err)
  }
}
```

## Methods


- ##### `mobius.AppStore.Balance( email )`
  Get balance of credits for email.

- ##### `mobius.AppStore.Use( email, numCredits )`
  Use numCredits from user with email.

- ##### `mobius.Tokens.Register( tokenType, name, symbol, address )`
  Register a token.

- ##### `mobius.Tokens.Balance( tokenUid, address )`
  Query the number of tokens specified by the token.

- ##### `mobius.Tokens.CreateAddress( tokenUid, managed )`
  Create an address for the token.

- ##### `mobius.Tokens.RegisterAddress( tokenUid, address )`
  Register an address for the token.

- ##### `mobius.Tokens.TransferInfo( tokenAddressTransferUid )`
  Get the status and transaction hash of a Mobius managed token transfer.

Other methods can be found in the API reference


## API Reference
https://godoc.org/github.com/codehakase/mobius-go
Mobius API Docs - https://mobius.network/docs/

## TODOs
- Integration Testing
- CLI Tool

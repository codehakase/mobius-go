/*
Package mobius provides methods for interacting with the Mobius Network API.

It automates the HTTP request/response cycle, encodings, and other details needed by the API.
This SDK lets you do everything the API lets you, in a more Go-friendly way.

For further information please see the Official Mobius API documentation at
https://mobius.network/docs/

Original Author: Francis Sunday <twitter.com/codehakase>

QuickStart
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
      usrBal, err := appStore.Balance("myemail@example.com")
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
*/
package mobius

package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"grpc-client/proto"
	"io"
	"sync"
	"time"

	"github.com/mr-tron/base58"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
  url := "pixelglobal.shyft.to:8443"
  accounts := []string{
    "monacoUXKtUi6vKsQwaLyxmXKSievfNWEcYXTgkbCih",
    "M2mx93ekt1fmXSVkTrUL9xVFHkmME8HTUi5Cyc5aF7K",
  }
  conn, err := grpc.Dial(url, grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{})))

  if err != nil {
    fmt.Println("Failed to  dial")
    panic(err)
  }

  client := proto.NewGeyserClient(conn)

  ctx, cancel := context.WithCancel(context.Background())
  sub, err := client.Subscribe(ctx)

  if err != nil {
    fmt.Println("Failed to create subscribe")
    panic(err)
  }

  err = sub.Send(&proto.SubscribeRequest{
    Accounts: map[string]*proto.SubscribeRequestFilterAccounts{
      "account_sub": {
        Owner: accounts,
      },
    },
  })


  if err != nil {
    fmt.Println("Failed to send subscribe request")
    panic(err)
  }

  wg := sync.WaitGroup{}

  fmt.Println("Started listener on accounts", accounts)
  wg.Add(1)
  go receiveMessages(sub, &wg)

  t := time.NewTicker(100000000*time.Second)

  go func(cancel context.CancelFunc) {
    for {
      select {
      case <- t.C: {
        fmt.Println("\n\\n\n\n\n\n\n\n\n\nn\nSEnding clear signal\n\n\n\n\n\n\n\n")
        cancel()
      }
      }
    }
  }(cancel)
  wg.Wait()
}

func printAccountUpdate(acc *proto.SubscribeUpdateAccount) {
  innerAccount := acc.GetAccount()

  if innerAccount.GetTxnSignature() == nil {
    return
  }

  fmt.Printf(
    "Acocunt detials: slot: %v, pubkey: %v, owner: %v, txn: %v, lamports: %v, data-length: %v, \n",
    acc.Slot,
    base58.Encode(innerAccount.Pubkey),
    base58.Encode(innerAccount.Owner),
    base58.Encode(innerAccount.TxnSignature),
    innerAccount.Lamports,
    len(innerAccount.Data),
  )
}

func receiveMessages(sub proto.Geyser_SubscribeClient, wg *sync.WaitGroup) {
  for {
    message, err := sub.Recv()

    if err == io.EOF {
      fmt.Println("Stream ended")
      wg.Done()
    }

    printAccountUpdate(message.GetAccount())
  }
}

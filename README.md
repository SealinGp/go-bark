# go-bark

go client library for [Bark Server](https://github.com/Finb/bark-server)

# Documentation

## Install
```bash
go get github.com/SealinGp/go-bark@latest
```

## Basic Usage

```go
  import "github.com/SealinGp/go-bark"

  authKey := ""
  barkCli, err := gobark.NewClient(authKey)
  if err != nil {
    panic(err)
  }

  //text
  err = barkCli.Bark(context.Background(),&gobark.BarkRequest{
    Text: &gobark.Text{
      Title: "xxx",
      Content:"xxx"
    }
  })

  //sound
  err = barkCli.Bark(ctx, &gobark.BarkRequest{
    BarkRequestOptions: gobark.BarkRequestOptions{
      Sound: "minuet",
    },
    Text: &gobark.Text{
      Title: "推送铃声",
    },
  })
```

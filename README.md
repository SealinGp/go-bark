# go-bark

go client library for [Bark Server](https://github.com/Finb/bark-server)

# Documentation

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
  err = barkCli.Bark(context.Background(),&gobark.BarkRequest{
    Ring: &gobark.Ring{
      Sound: "minuet"
    }
  })
```

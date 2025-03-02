# Go-API-Template (Single...)

## swagger 등록하는 법

```sh
    ## in repository
    go install github.com/swaggo/swag/cmd/swag@latest

    ## swag error
    zsh: command not found: swag

    ## swag error solution
    export PATH=$(go env GOPATH)/bin:$PATH

    ## fmt (formatting)
    swag fmt

    ## init (init)
    swag init

    ## import 등록

    import (
        ...
        _ "{application}/docs
    )
```

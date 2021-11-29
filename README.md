# captcha

Verify hCaptcha token validity.


[![Go](https://github.com/gmvbr/captcha/actions/workflows/go.yml/badge.svg)](https://github.com/gmvbr/captcha/actions/workflows/go.yml)

## Instalation

```bash
go get -u github.com/gmvbr/captcha
```

## Usage

```go
package main

import (
    "github.com/gmvbr/captcha"  
)

func main() {

    service := captcha.NewHCaptcha("hcaptcha secret")
    result, err := service.Verify("enter site response")

    if err != nil {
        // handle errors
    }

    if result.Success == true {
        // ok
    } else {
        // error captcha
    }
}

```

## Testing

See tests in [captcha_test.go](captcha_test.go)

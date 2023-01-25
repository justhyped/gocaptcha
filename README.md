# gocaptcha
An API wrapper for popular captcha solvers such as AntiCaptcha and 2Captcha in Golang

## Installation
```sh
go get github.com/justhyped/gocaptcha
```

## Support
| Type          | 2Captcha | AntiCaptcha |
|:--------------|:---------|:------------|
 | RecaptchaV2   | ✅        | ✅           |
 | RecaptchaV3   | ✅        | ✅           |
| Image Captcha | ✅        | ✅           |
| HCaptcha      | ✅        | ✅           |
| Turnstile     | ✅        | ✅           |

Software like XEVil and CapMonster are also supported. You can also implement your own provider by 
using the `IProvider` interface.

## Usage
- [2Captcha](https://github.com/justhyped/gocaptcha/blob/main/examples/twocaptcha/main.go)
- [2Captcha (with custom domain)](https://github.com/justhyped/gocaptcha/blob/main/examples/twocaptcha_custom/main.go)
- [AntiCaptcha](https://github.com/justhyped/gocaptcha/blob/main/examples/anticaptcha_custom/main.go)
- [AntiCaptcha (with custom domain)](https://github.com/justhyped/gocaptcha/blob/main/examples/anticaptcha_custom/main.go)
- [Custom provider](https://github.com/justhyped/gocaptcha/blob/main/examples/custom_provider/main.go)

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)

# gocaptcha
An API wrapper for popular captcha solvers such as AntiCaptcha and 2Captcha in Golang

## Installation
Run the following command in your project folder:
``go get github.com/justhyped/gocaptcha``

## Support
| Type            | 2Captcha | AntiCaptcha | CapMonster Cloud |
|:----------------|:---------|:------------|:-----------------|
 | RecaptchaV2     | ✅        | ✅           | ✅                |
 | RecaptchaV3     | ✅        | ✅           | ✅                |
 | Image Recaptcha | ✅        | ✅           | ✅                |

## Usage
### RecaptchaV2
```go
    payload := gocaptcha.RecaptchaV2Payload{
	EndpointUrl:   "https://www.google.com/recaptcha/api2/demo",
        EndpointKey:   "6Le-wvkSAAAAAPBMRTvw0Q4Muexq9bi0DJwx_mJ-",
        ServiceApiKey: "key",
        ServiceName:   "2Captcha",
	}
	
    captcha, err := gocaptcha.SolveRecaptchaV2(&payload)
```

These are all supported variables to use in RecaptchaV2Payload:
```go
    	// This is the endpoint that has Recaptcha Protection
EndpointUrl string

// This is the Recaptcha Key
// Can be found on the Endpoint URL page
EndpointKey string

// The API key for your captcha service
ServiceApiKey string

// The name of the captcha service
// Can be AntiCaptcha, 2Captcha or CapMonster Cloud
ServiceName string

// Enable if endpoint has invisible Recaptcha V2
IsInvisibleCaptcha bool

// Set this in case you're using a custom solver like CapMonster (not cloud)
CustomServiceUrl string

// The time to wait before starting to poll result
InitialWaitTime int

// The time to wait between polling results
PollInterval int

// Max amount of poll attempts
MaxRetries int
```

### RecaptchaV3
```go
    payload := gocaptcha.RecaptchaV3Payload{
        EndpointUrl:   "https://recaptcha-demo.appspot.com/recaptcha-v3-request-scores.php",
        EndpointKey:   "6LdyC2cUAAAAACGuDKpXeDorzUDWXmdqeg-xy696",
        ServiceApiKey: "key",
        ServiceName:   "2Captcha",
        Action:        "examples/v3scores",
    }

    captcha, err := gocaptcha.SolveRecaptchaV3(&payload)
```

These are all supported variables to use in RecaptchaV3Payload:
```go
// This is the endpoint that has Recaptcha Protection
EndpointUrl string

// This is the Recaptcha Key
// Can be found on the Endpoint URL page
EndpointKey string

// The API key for your captcha service
ServiceApiKey string

// The name of the captcha service
// Can be AntiCaptcha, 2Captcha or CapMonster Cloud
ServiceName string

// The action name of the recaptcha, you can find it in source code of site
Action string

// Set this in case you're using a custom solver like CapMonster (not cloud)
CustomServiceUrl string

// Set to true if it's V3 Enterprise
IsEnterprise bool

// Defaults to 0.3, accepted values are 0.3, 0.6, 0.9
MinScore float32

// The time to wait before starting to poll result
InitialWaitTime int

// The time to wait between polling results
PollInterval int

// Max amount of poll attempts
MaxRetries int
```

### Image Captcha
```go
payload := gocaptcha.ImageCaptchaPayload{
    ServiceApiKey: "key", // your api key
    ServiceName:   "2Captcha", // the provider, can be 2Captcha, AntiCaptcha or Capmonster Cloud
    Base64String:  imageBase64, // the image converted to a base64 string
}

captcha, err := gocaptcha.SolveImageCaptcha(&payload)
```

These are all supported variables to use in ImageCaptchaPayload:
```go
    // This is the base64 that represents the image captcha
	Base64String string

	// The API key for your captcha service
	ServiceApiKey string

	// The name of the captcha service
	// Can be AntiCaptcha, 2Captcha or CapMonster Cloud
	ServiceName string

	// Set this in case you're using a custom solver like CapMonster (not cloud)
	CustomServiceUrl string

	// Set to true if captcha is case sensitive
	CaseSensitive bool

	// Set this if the human solver needs additional information
	// about how to solve the captcha
	InstructionsForSolver string

	// The time to wait before starting to poll result
	InitialWaitTime int

	// The time to wait between polling results
	PollInterval int

	// Max amount of poll attempts
	MaxRetries int
```


## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)
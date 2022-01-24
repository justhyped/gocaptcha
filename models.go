package gocaptcha

func (p *RecaptchaV2Payload) SetDefaultValues() {
	if p.InitialWaitTime == 0 {
		p.InitialWaitTime = 10
	}

	if p.PollInterval == 0 {
		p.PollInterval = 5
	}

	if p.MaxRetries == 0 {
		p.MaxRetries = 15
	}
}

func (p *RecaptchaV3Payload) SetDefaultValues() {
	if p.InitialWaitTime == 0 {
		p.InitialWaitTime = 10
	}

	if p.PollInterval == 0 {
		p.PollInterval = 5
	}

	if p.MaxRetries == 0 {
		p.MaxRetries = 15
	}

	if p.MinScore != 0.3 && p.MinScore != 0.6 && p.MinScore != 0.9 {
		p.MinScore = 0.3
	}
}

func (p *ImageCaptchaPayload) SetDefaultValues() {
	if p.InitialWaitTime == 0 {
		p.InitialWaitTime = 5
	}

	if p.PollInterval == 0 {
		p.PollInterval = 3
	}

	if p.MaxRetries == 0 {
		p.MaxRetries = 15
	}
}

func (p *ImageCaptchaPayload) CreateImageCaptchaResponse() *CaptchaResponse {
	return &CaptchaResponse{
		Service:       p.ServiceName,
		ServiceApiKey: p.ServiceApiKey,
		Endpoint:      p.CustomServiceUrl,
	}
}

func (p *RecaptchaV2Payload) CreateRecaptchaResponse() *CaptchaResponse {
	return &CaptchaResponse{
		Service:       p.ServiceName,
		ServiceApiKey: p.ServiceApiKey,
		Endpoint:      p.CustomServiceUrl,
	}
}

func (p *RecaptchaV3Payload) CreateRecaptchaResponse() *CaptchaResponse {
	return &CaptchaResponse{
		Service:       p.ServiceName,
		ServiceApiKey: p.ServiceApiKey,
		Endpoint:      p.CustomServiceUrl,
	}
}

type RecaptchaV2Payload struct {
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
}

type RecaptchaV3Payload struct {
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
}

type ImageCaptchaPayload struct {
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
}

type CaptchaResponse struct {
	// The task ID of the solved captcha
	TaskId string

	// The answer of the solved captcha
	Solution string

	// The service used for solving the captcha
	// this is used when reporting a good/bad captcha
	Service string

	// The service endpoint, it's copied from CustomServiceUrl
	Endpoint string

	// The API key for the service that's used for solving
	ServiceApiKey string
}

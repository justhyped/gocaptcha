package gocaptcha

import (
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

// Supported values are V2, V3, image, HCaptcha, AntiCaptcha, 2Captcha and CapMonster Cloud
var testGroup string

var twoCaptchaKey string
var antiCaptchaKey string
var capMonsterCloudKey string

func init() {
	err := godotenv.Load("secrets.env")
	if err != nil {
		log.Fatalf("Error loading secrets.env: %s", err)
	}

	twoCaptchaKey = os.Getenv("2CAPTCHA")
	antiCaptchaKey = os.Getenv("ANTICAPTCHA")
	capMonsterCloudKey = os.Getenv("CAPMONSTERCLOUD")

	testGroup = "HCaptcha"
}

func TestSolveRecaptchaV2AntiCaptcha(t *testing.T) {
	if testGroup != "V2" && testGroup != "AntiCaptcha" && testGroup != "" {
		t.Skip()
	}

	payload := RecaptchaV2Payload{
		EndpointUrl:   "https://www.google.com/recaptcha/api2/demo",
		EndpointKey:   "6Le-wvkSAAAAAPBMRTvw0Q4Muexq9bi0DJwx_mJ-",
		ServiceApiKey: antiCaptchaKey,
		ServiceName:   "AntiCaptcha",
	}

	captcha, err := SolveRecaptchaV2(&payload)
	checkError(t, err)

	t.Log(captcha.Solution)
}

func TestSolveRecaptchaV2TwoCaptcha(t *testing.T) {
	if testGroup != "V2" && testGroup != "2Captcha" && testGroup != "" {
		t.Skip()
	}

	payload := RecaptchaV2Payload{
		EndpointUrl:   "https://www.google.com/recaptcha/api2/demo",
		EndpointKey:   "6Le-wvkSAAAAAPBMRTvw0Q4Muexq9bi0DJwx_mJ-",
		ServiceApiKey: twoCaptchaKey,
		ServiceName:   "2Captcha",
	}

	captcha, err := SolveRecaptchaV2(&payload)
	checkError(t, err)

	t.Log(captcha.Solution)
}

func TestSolveRecaptchaV2CapMonster(t *testing.T) {
	if testGroup != "V2" && testGroup != "CapMonster Cloud" && testGroup != "" {
		t.Skip()
	}

	payload := RecaptchaV2Payload{
		EndpointUrl:   "https://www.google.com/recaptcha/api2/demo",
		EndpointKey:   "6Le-wvkSAAAAAPBMRTvw0Q4Muexq9bi0DJwx_mJ-",
		ServiceApiKey: capMonsterCloudKey,
		ServiceName:   "CapMonster Cloud",
	}

	captcha, err := SolveRecaptchaV2(&payload)
	checkError(t, err)

	t.Log(captcha.Solution)
}

func TestSolveRecaptchaV3AntiCaptcha(t *testing.T) {
	if testGroup != "V3" && testGroup != "AntiCaptcha" && testGroup != "" {
		t.Skip()
	}

	payload := RecaptchaV3Payload{
		EndpointUrl:   "https://recaptcha-demo.appspot.com/recaptcha-v3-request-scores.php",
		EndpointKey:   "6LdyC2cUAAAAACGuDKpXeDorzUDWXmdqeg-xy696",
		ServiceApiKey: antiCaptchaKey,
		ServiceName:   "AntiCaptcha",
		Action:        "examples/v3scores",
	}

	captcha, err := SolveRecaptchaV3(&payload)
	checkError(t, err)

	t.Log(captcha.Solution)
}

func TestSolveRecaptchaV3TwoCaptcha(t *testing.T) {
	if testGroup != "V3" && testGroup != "2Captcha" && testGroup != "" {
		t.Skip()
	}

	payload := RecaptchaV3Payload{
		EndpointUrl:   "https://recaptcha-demo.appspot.com/recaptcha-v3-request-scores.php",
		EndpointKey:   "6LdyC2cUAAAAACGuDKpXeDorzUDWXmdqeg-xy696",
		ServiceApiKey: twoCaptchaKey,
		ServiceName:   "2Captcha",
		Action:        "examples/v3scores",
	}

	captcha, err := SolveRecaptchaV3(&payload)
	checkError(t, err)

	t.Log(captcha.Solution)
}

func TestSolveRecaptchaV3CapMonster(t *testing.T) {
	if testGroup != "V3" && testGroup != "CapMonster Cloud" && testGroup != "" {
		t.Skip()
	}

	payload := RecaptchaV3Payload{
		EndpointUrl:   "https://recaptcha-demo.appspot.com/recaptcha-v3-request-scores.php",
		EndpointKey:   "6LdyC2cUAAAAACGuDKpXeDorzUDWXmdqeg-xy696",
		ServiceApiKey: capMonsterCloudKey,
		ServiceName:   "CapMonster Cloud",
		Action:        "examples/v3scores",
	}

	captcha, err := SolveRecaptchaV3(&payload)
	checkError(t, err)

	t.Log(captcha.Solution)
}

func TestSolveHCaptchaAntiCaptcha(t *testing.T) {
	if testGroup != "HCaptcha" && testGroup != "AntiCaptcha" && testGroup != "" {
		t.Skip()
	}

	payload := HCaptchaPayload{
		EndpointUrl:   "https://www.hcaptcha.com/",
		EndpointKey:   "00000000-0000-0000-0000-000000000000",
		ServiceApiKey: antiCaptchaKey,
		ServiceName:   "AntiCaptcha",
	}

	captcha, err := SolveHCaptcha(&payload)
	checkError(t, err)

	t.Log(captcha.Solution)
}

func TestSolveHCaptchaTwoCaptcha(t *testing.T) {
	if testGroup != "HCaptcha" && testGroup != "2Captcha" && testGroup != "" {
		t.Skip()
	}

	payload := HCaptchaPayload{
		EndpointUrl:   "http://democaptcha.com/demo-form-eng/hcaptcha.html",
		EndpointKey:   "51829642-2cda-4b09-896c-594f89d700cc",
		ServiceApiKey: twoCaptchaKey,
		ServiceName:   "2Captcha",
	}

	captcha, err := SolveHCaptcha(&payload)
	checkError(t, err)

	t.Log(captcha.Solution)
}

func TestSolveHCaptchaCapMonster(t *testing.T) {
	if testGroup != "HCaptcha" && testGroup != "CapMonster Cloud" && testGroup != "" {
		t.Skip()
	}

	payload := HCaptchaPayload{
		EndpointUrl:   "https://www.hcaptcha.com/",
		EndpointKey:   "00000000-0000-0000-0000-000000000000",
		ServiceApiKey: capMonsterCloudKey,
		ServiceName:   "CapMonster Cloud",
	}

	captcha, err := SolveHCaptcha(&payload)
	checkError(t, err)

	t.Log(captcha.Solution)
}

const imageBase64 string = `R0lGODlhcwAjAPcAAAAAAAAAMwAAZgAAmQAAzAAA/wArAAArMwArZgArmQAr
zAAr/wBVAABVMwBVZgBVmQBVzABV/wCAAACAMwCAZgCAmQCAzACA/wCqAACqMwCqZgCqmQCqzACq/wDVAADVMwDVZg
DVmQDVzADV/wD/AAD/MwD/ZgD/mQD/zAD//zMAADMAMzMAZjMAmTMAzDMA/zMrADMrMzMrZjMrmTMrzDMr/zNVADNVM
zNVZjNVmTNVzDNV/zOAADOAMzOAZjOAmTOAzDOA/zOqADOqMzOqZjOqmTOqzDOq/zPVADPVMzPVZjPVmTPVzDPV/zP/A
DP/MzP/ZjP/mTP/zDP//2YAAGYAM2YAZmYAmWYAzGYA/2YrAGYrM2YrZmYrmWYrzGYr/2ZVAGZVM2ZVZmZVmWZVzGZV/
2aAAGaAM2aAZmaAmWaAzGaA/2aqAGaqM2aqZmaqmWaqzGaq/2bVAGbVM2bVZmbVmWbVzGbV/2b/AGb/M2b/Zmb/mWb/z
Gb//5kAAJkAM5kAZpkAmZkAzJkA/5krAJkrM5krZpkrmZkrzJkr/5lVAJlVM5lVZplVmZlVzJlV/5mAAJmAM5mAZpmAm
ZmAzJmA/5mqAJmqM5mqZpmqmZmqzJmq/5nVAJnVM5nVZpnVmZnVzJnV/5n/AJn/M5n/Zpn/mZn/zJn//8wAAMwAM8wAZ
swAmcwAzMwA/8wrAMwrM8wrZswrmcwrzMwr/8xVAMxVM8xVZsxVmcxVzMxV/8yAAMyAM8yAZsyAmcyAzMyA/8yqAMyqM
8yqZsyqmcyqzMyq/8zVAMzVM8zVZszVmczVzMzV/8z/AMz/M8z/Zsz/mcz/zMz///8AAP8AM/8AZv8Amf8AzP8A//8rA
P8rM/8rZv8rmf8rzP8r//9VAP9VM/9VZv9Vmf9VzP9V//+AAP+AM/+AZv+Amf+AzP+A//+qAP+qM/+qZv+qmf+qzP+q/
//VAP/VM//VZv/Vmf/VzP/V////AP//M///Zv//mf//zP///wAAAAAAAAAAAAAAACH5BAEAAPwALAAAAABzACMAAAj/A
PcJHKhM2UCCBg/uK6hwYcKDDBVGhPgQYcOJFiVWFIiR40aHBCFyFLmQZMWEJ0eGLLkyJUuVLmOanNmypk2YNHHeZClzJ
8qcPIH+XAlS48WPHYtSPMpUITGkUJsuNSpxqtWMVz1KxcpVK9WsSruGbfiSZNWLZA3WO8vWrFuib1W2Xbk2rV20eOfKj
Vt239qPfffCHSxYYl2+gBPf1Ru4seK8iMl6zVovacF6mTLVtRxVIrGj9aAd5rx1MtaeOlPzzPQ5qM+9yohlqvi54FOOx
J7mZh17YT1i0FyrHspzNOzjvgU+NdjzL/KCrIPKvr0wt7LZvGdzzFSWuPCXpL9y/80cHixD2QivZwo+1iF3853dcq58u
3Js6iCvt7atbHRrgfa1Ntp1By23lXVI/TcQfmIxpt5u2eUG0XSxkZeZRplltlx0ymm3j2yaOSbQhRxh5tJcaym42ISzP
aTeSRluhpmC/CknmzLsfchdhSdC5CFB70UG2ZDuqVjhf/39uN1DNyqU2Wi8vRgXfQrZR6SIfl30n4zrCRQcdD/ahuFhm
m3U4kCstaifWOWxqaJSqFWn5H7U8ShSiCMydNJnrZFno3YJKYgejJUJ9dxrH+5z4X5o6onZjiPZ6VeQKHF4H4wMVtoSe
t955ymbzPGmFIWh7pfQk+4F6ZBsOXZUEFR4av8VZnxswtekR48ilGmFWRm4JINRHmQirbiV1p5e2J10XUVKerQlbylq2
JuOmy0Y3UOWCnvmlY9J1hdvFFLoI4M6bnSjdpwaxFqrAPoZqovWFoYlY3C9OJ2otL2qkWKVecicg4v5u+KCA09YYJnG
csYcSqe+16ZpXln65sMfwnvTmXF+WtKNgWbb3aEr7TaccoaChxxm1X6ssrLYgahxxjBFBPPKZlkGInwJ9zcRxTwT255l
OIsH8c8fkVsdgKDmrLS38wrpVMFNE+ZbjvIC9mZjTvN1dbdZS801XCnTWzXUXIfdNcmQ1bU12d7yPLHPRve8tMFCw1ny
zHjfXTKNIKsQxqhKfNOMaN6DkxQ44SMHBAA7`

func TestSolveImageCaptchaTwoCaptcha(t *testing.T) {
	if testGroup != "image" && testGroup != "2Captcha" && testGroup != "" {
		t.Skip()
	}

	payload := ImageCaptchaPayload{
		ServiceApiKey: twoCaptchaKey,
		ServiceName:   "2Captcha",
		Base64String:  imageBase64,
	}

	captcha, err := SolveImageCaptcha(&payload)
	checkError(t, err)

	t.Log(captcha.Solution)
}

func TestSolveImageCaptchaAntiCaptcha(t *testing.T) {
	if testGroup != "image" && testGroup != "AntiCaptcha" && testGroup != "" {
		t.Skip()
	}

	payload := ImageCaptchaPayload{
		ServiceApiKey: antiCaptchaKey,
		ServiceName:   "AntiCaptcha",
		Base64String:  imageBase64,
	}

	captcha, err := SolveImageCaptcha(&payload)
	checkError(t, err)

	t.Log(captcha.Solution)
}

func TestSolveImageCaptchaCapMonster(t *testing.T) {
	if testGroup != "image" && testGroup != "CapMonster Cloud" && testGroup != "" {
		t.Skip()
	}

	payload := ImageCaptchaPayload{
		ServiceApiKey: capMonsterCloudKey,
		ServiceName:   "CapMonster Cloud",
		Base64String:  imageBase64,
	}

	captcha, err := SolveImageCaptcha(&payload)
	checkError(t, err)

	t.Log(captcha.Solution)
}

func checkError(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}

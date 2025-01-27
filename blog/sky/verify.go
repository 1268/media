package main

import (
   "io"
   "net/http"
   "net/url"
   "os"
   "strings"
)

func main() {
   var req http.Request
   req.Header = http.Header{}
   req.Header["Connection"] = []string{"Keep-Alive"}
   req.Header["Content-Length"] = []string{"3694"}
   req.Header["Content-Type"] = []string{"application/json"}
   req.Header["Host"] = []string{"3df51d5ca14f.12c6c787.eu-central-2.token.awswaf.com"}
   req.Header["User-Agent"] = []string{"Dalvik/2.1.0 (Linux; U; Android 9; AOSP on IA Emulator Build/PSR1.180720.122)"}
   req.Method = "POST"
   req.ProtoMajor = 1
   req.ProtoMinor = 1
   req.URL = &url.URL{}
   req.URL.Host = "3df51d5ca14f.12c6c787.eu-central-2.token.awswaf.com"
   req.URL.Path = "/3df51d5ca14f/4104523408d1/verify"
   req.URL.Scheme = "https"
   req.Body = io.NopCloser(body)
   resp, err := http.DefaultClient.Do(&req)
   if err != nil {
      panic(err)
   }
   defer resp.Body.Close()
   resp.Write(os.Stdout)
}

var body = strings.NewReader(`{"challenge":{"input":"eyJ2ZXJzaW9uIjoxLCJ1YmlkIjoiNGM1YzM2OWQtOGQ3OC00MGMzLWI3ZGMtOWQ4NmIyYjBlYTk0IiwiYXR0ZW1wdF9pZCI6ImVlYTE4YWY3LTNhNDAtNGFkNS05OGI2LTJlN2NkZjVjNDZlNiIsImNyZWF0ZV90aW1lIjoiMjAyNS0wMS0yN1QwMjozNTo1NC42MDc2MjM4MDJaIiwiZGlmZmljdWx0eSI6NCwiY2hhbGxlbmdlX3R5cGUiOiJIYXNoY2FzaFBCS0RGMiJ9","hmac":"RnFUusDTZdU8QSizHeblwuBobjJOVJ6Q53CDnfkNIBw\u003d"},"solution":"21","checksum":"1f09a11382ab588aa61e70622d5c98c697447118d1bef3f24e614bc150518a85","client":"android","domain":"clientapi.prd.sky.ch","signals":[{"name":"AndroidID","value":{"Present":""}},{"name":"GsfID","value":{"Present":""}},{"name":"ApplicationID","value":{"Present":"homedia.sky.sport"}},{"name":"MediaDRMID","value":{"Present":""}},{"name":"MediaDRMVendor","value":{"Present":""}},{"name":"MediaDRMVersion","value":{"Present":""}},{"name":"MediaDRMAlgorithms","value":{"Present":""}},{"name":"IsNewDevice","value":{"Present":"true"}},{"name":"WAFDeviceUUID","value":{"Present":"41a32a12-eb3d-4786-98c2-c8b43dfe0c16"}},{"name":"DeviceID","value":{"Present":"ff1bbbfe-6048-47bf-9311-cf5b8087f825-1"}},{"name":"DeviceModel","value":{"Present":""}},{"name":"DeviceBrand","value":{"Present":""}},{"name":"ManufacturerName","value":{"Present":""}},{"name":"AndroidVersion","value":{"Present":""}},{"name":"SDKVersion","value":{"Present":""}},{"name":"KernelVersion","value":{"Present":""}},{"name":"OSFingerprint","value":{"Present":""}},{"name":"BatteryCapacity","value":{"Present":""}},{"name":"CameraCount","value":{"Present":""}},{"name":"CameraInfo","value":{"Present":""}},{"name":"Processor","value":{"Present":""}},{"name":"CPUHardware","value":{"Present":""}},{"name":"CPUCoreCount","value":{"Present":""}},{"name":"GLEVersion","value":{"Present":""}},{"name":"CodecsSupported","value":{"Present":""}},{"name":"TotalRAM","value":{"Present":""}},{"name":"TotalInternalStorage","value":{"Present":""}},{"name":"TotalExternalStorage","value":{"Present":""}},{"name":"Sensors","value":{"Present":""}},{"name":"NetworkInterfaces","value":{"Present":""}},{"name":"DefaultLanguage","value":{"Present":""}},{"name":"Timezone","value":{"Present":""}},{"name":"Locales","value":{"Present":""}},{"name":"CountryCode","value":{"Present":""}},{"name":"LastKnownLocationLatitude","value":{"Present":""}},{"name":"LastKnownLocationLongitude","value":{"Present":""}},{"name":"StorageEncryptionStatus","value":{"Present":""}},{"name":"IsPinSecurityEnabled","value":{"Present":""}},{"name":"SecurityProviders","value":{"Present":""}},{"name":"ADBEnabled","value":{"Present":""}},{"name":"DevelopmentSettingEnabled","value":{"Present":""}},{"name":"HttpProxy","value":{"Present":""}},{"name":"TransitionAnimationScale","value":{"Present":""}},{"name":"WindowAnimationScale","value":{"Present":""}},{"name":"DataRoamingEnabled","value":{"Present":""}},{"name":"AccessibilityEnabled","value":{"Present":""}},{"name":"TouchExplorationEnabled","value":{"Present":""}},{"name":"AlarmAlertPath","value":{"Present":""}},{"name":"DateFormat","value":{"Present":""}},{"name":"FontScale","value":{"Present":""}},{"name":"EndButtonBehaviour","value":{"Present":""}},{"name":"ScreenOffTimeout","value":{"Present":""}},{"name":"TextAutoReplaceEnabled","value":{"Present":""}},{"name":"TextAutoPunctuateEnabled","value":{"Present":""}},{"name":"Time12Or24","value":{"Present":""}},{"name":"SupportedInputMethods","value":{"Present":""}}],"metrics":[{"name":"SignalExecutionTime","value":48.0,"unit":"Milliseconds"},{"name":"GetChallengeExecutionTime","value":4149.0,"unit":"Milliseconds"},{"name":"TokenRefreshStartTimestamp","value":1.737945350956E12,"unit":"Milliseconds"},{"name":"ChallengeExecutionTime","value":64.0,"unit":"Milliseconds"}]}`)


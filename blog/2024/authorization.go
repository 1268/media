package amc

type a struct {
   AccessToken string `json:"access_token"`
   RefreshToken string `json:"refresh_token"`
}

type B struct {
   Value a
   Raw []byte
}

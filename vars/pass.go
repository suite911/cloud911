package vars

var Pass = PassType{
	HTTP: ":80",
	HTTPS: ":443",
}

type PassType struct {
	CaptchaSecret    []byte `json:"captcha"`     // Google reCAPTCHA secret data
	DataBase         string `json:"db"`          // Path to the main sqlite3 database
	HTTP             string `json:"http"`        // Address on which to listen to HTTP traffic
	HTTPS            string `json:"https"`       // Address on which to listen to HTTPS traffic
	TLSCertData      []byte `json:"cert"`        // TLS certificate and key file data
	TLSKeyData       []byte `json:"key"`         // TLS certificate and key file data
	URLRegistered    string `json:"url_regd"`    // Path to which to 302 after registering
	URLBotRegistered string `json:"url_botregd"` // Path to which to 302 after a bot fails a registration captcha

	FeatureUserProfiles bool `json:"ft_user_profiles"` // Feature toggle for User Profiles
}

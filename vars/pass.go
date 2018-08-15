package vars

var Pass = PassType{
	HTTP: ":10080",
	HTTPS: ":10443",
}

type PassType struct {
	Cert  []byte `json:"cert"`  // TLS certificate and key file data
	HTTP  string `json:"http"`  // Address on which to listen to HTTP traffic
	HTTPS string `json:"https"` // Address on which to listen to HTTPS traffic
}

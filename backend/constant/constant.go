package constant

type Credential struct {
	Birthdate string // if 2002/07/27 then 20020727
	Sin       string
}

type VerifyData struct {
	HashedCredential string
	Signature        string
	PublicKey        string
}

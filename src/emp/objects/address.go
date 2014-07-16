package objects

type AddressDetail struct {
	String       string `json:"address"`
	Address      []byte `json:"address_bytes"`
	IsRegistered bool   `json:"registered"`
	Pubkey       []byte `json:"public_key"`
	Privkey      []byte `json:"private_key"`
	Label        string `json:"address_label"`
}
package microServs

type Message struct {
	Text []byte `json:"text"`
	Signature []byte `json:"signature"`
	PubKey []byte	`json:"pubKey"`
}
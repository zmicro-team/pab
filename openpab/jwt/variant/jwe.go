package variant

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"strings"
	"sync"

	"github.com/zmicro-team/pab/gm/sm2"
	"github.com/zmicro-team/pab/gm/sm4"
	"github.com/zmicro-team/pab/openpab/util"
)

var mu sync.Mutex

func EncryptJWE(kid string, publicKey *sm2.PublicKey, plainText string) (result string, err error) {
	mu.Lock()
	defer mu.Unlock()
	header := make(map[string]any)
	header["alg"] = "SM2"
	header["typ"] = "JWT"
	header["kid"] = kid
	header["enc"] = "SM4"

	sm4Key := []byte("1234567890abcdef")
	encryptedKey, err := sm2.Encrypt(publicKey, []byte(util.ByteToHex(sm4Key)), nil, sm2.C1C2C3)
	if err != nil {
		return
	}

	iv := []byte("0000000000000000") // 16字节
	sm4.SetIV(iv)
	cbcMsg, err := sm4.Sm4Cbc(sm4Key, []byte(plainText), true)
	if err != nil {
		return
	}
	ciperText := util.ByteToHex(cbcMsg)

	jwe := Jwe{
		Header:       header,
		EncryptedKey: util.ByteToHex(encryptedKey),
		Iv:           util.ByteToHex(iv),
		CipherText:   ciperText,
		Tag:          []byte("0"),
	}

	result = jwe.String()
	return
}

func DecryptJWE(jwe string, privateKey *sm2.PrivateKey) (result string, err error) {
	mu.Lock()
	defer mu.Unlock()

	j, err := Parse(jwe)
	if err != nil {
		return
	}
	// TODO: fix me
	sm4KeyHex, _ := sm2.Decrypt(privateKey, util.HexToByte(j.EncryptedKey), sm2.C1C2C3)
	// if err != nil {
	//	return
	// }
	sm4.SetIV(util.HexToByte(j.Iv))
	b, err := sm4.Sm4Cbc(util.HexToByte(string(sm4KeyHex)), util.HexToByte(j.CipherText), false)
	if err != nil {
		return
	}

	result = string(b)
	return
}

type Jwe struct {
	Header       map[string]any
	EncryptedKey string
	Iv           string
	CipherText   string
	Tag          []byte
}

func (j *Jwe) String() string {
	b, _ := json.Marshal(j.Header)
	p1 := base64.RawURLEncoding.EncodeToString(b)
	p2 := base64.RawURLEncoding.EncodeToString([]byte(j.EncryptedKey))
	p3 := base64.RawURLEncoding.EncodeToString([]byte(j.Iv))
	p4 := base64.RawURLEncoding.EncodeToString([]byte(j.CipherText))
	p5 := base64.RawURLEncoding.EncodeToString(j.Tag)

	return p1 + "." + p2 + "." + p3 + "." + p4 + "." + p5
}

func Parse(jwe string) (*Jwe, error) {
	parts := strings.Split(jwe, ".")
	if len(parts) != 5 {
		return nil, errors.New("malformed jwe")
	} else if parts[0] != "" && parts[1] != "" && parts[2] != "" && parts[3] != "" && parts[4] != "" {
		b, _ := base64.RawURLEncoding.DecodeString(parts[0])
		encryptedKey, _ := base64.RawURLEncoding.DecodeString(parts[1])
		iv, _ := base64.RawURLEncoding.DecodeString(parts[2])
		cipherText, _ := base64.RawURLEncoding.DecodeString(parts[3])
		tag, _ := base64.RawURLEncoding.DecodeString(parts[4])

		header := make(map[string]any)
		json.Unmarshal(b, &header)
		j := &Jwe{}
		j.Header = header
		j.Iv = string(iv)
		j.EncryptedKey = string(encryptedKey)
		j.CipherText = string(cipherText)
		j.Tag = tag

		return j, nil

	} else {
		return nil, errors.New("malformed jwe")
	}
}

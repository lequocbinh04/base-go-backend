package paseto

import (
	"aidanwoods.dev/go-paseto"
	"cronbrowser/appCommon"
	"cronbrowser/plugin/tokenprovider"
	"encoding/json"
	"flag"
	"fmt"
	"strconv"
	"time"
)

type pasetoProvider struct {
	prefix    string
	secretKey string
	publicKey string
}

func NewPasetoProvider(prefix string) *pasetoProvider {
	return &pasetoProvider{
		prefix: prefix,
	}
}

type myClaims struct {
	Payload appCommon.TokenPayload `json:"payload"`
}

type tokenRes struct {
	Token   string    `json:"token"`
	Created time.Time `json:"created"`
	Expiry  int       `json:"expiry"`
}

func (t *tokenRes) GetToken() string {
	return t.Token
}

func (p *pasetoProvider) Generate(data tokenprovider.TokenPayload, expiry int) (tokenprovider.Token, error) {
	now := time.Now()
	token := paseto.NewToken()
	token.SetIssuedAt(now)
	token.SetNotBefore(time.Now())
	token.SetExpiration(time.Now().Add(time.Duration(expiry) * time.Second))
	token.SetString("user_id", fmt.Sprintf("%d", data.UserId()))
	token.SetString("role", data.Role())

	key, err := paseto.NewV4AsymmetricSecretKeyFromHex(p.secretKey)
	if err != nil {
		return nil, err
	}

	signed := token.V4Sign(key, nil)

	return &tokenRes{
		Token:   signed,
		Expiry:  expiry,
		Created: now,
	}, nil
}

func (p *pasetoProvider) Validate(signed string) (tokenprovider.TokenPayload, error) {
	publicKey, _ := paseto.NewV4AsymmetricPublicKeyFromHex(p.publicKey)
	parser := paseto.NewParser()
	token, err := parser.ParseV4Public(publicKey, signed, nil)
	if err != nil {
		return nil, err
	}
	data := string(token.ClaimsJSON())
	dataJson := make(map[string]interface{})
	err = json.Unmarshal([]byte(data), &dataJson)
	if err != nil {
		return nil, err
	}
	i, _ := strconv.Atoi(dataJson["user_id"].(string))
	return appCommon.TokenPayload{
		UId:   int64(i),
		URole: dataJson["role"].(string),
	}, nil
}

func (p *pasetoProvider) String() string {
	return "Paseto implement Provider"
}

func (p *pasetoProvider) GetPrefix() string {
	return p.prefix
}

func (p *pasetoProvider) Get() interface{} {
	return p
}

func (p *pasetoProvider) Name() string {
	return "paseto"
}

func (p *pasetoProvider) InitFlags() {
	prefix := p.prefix
	if p.prefix != "" {
		prefix += "-"
	}
	flag.StringVar(&p.secretKey, prefix+"secret", "ad414b6cdefd2475fbd0d8c222135fd7f172fd659263edc2a05f3411a8f8480086dc3751ef04b3a87d2d44899c08946d70518a663320dd0a25ba2ffe3d525133", "Secret key for Paseto.")
	flag.StringVar(&p.publicKey, prefix+"public", "86dc3751ef04b3a87d2d44899c08946d70518a663320dd0a25ba2ffe3d525133", "Public key for Paseto.")
}

func (pasetoProvider) Configure() error {
	return nil
}

func (pasetoProvider) Run() error {
	return nil
}

func (pasetoProvider) Stop() <-chan bool {
	c := make(chan bool)
	go func() { c <- true }()
	return c
}

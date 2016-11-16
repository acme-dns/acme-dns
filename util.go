package main

import (
	"crypto/rand"
	"errors"
	"github.com/BurntSushi/toml"
	"github.com/satori/go.uuid"
	"math/big"
	"regexp"
	"strings"
)

func readConfig(fname string) (DNSConfig, error) {
	var conf DNSConfig
	if _, err := toml.DecodeFile(fname, &conf); err != nil {
		return DNSConfig{}, errors.New("Malformed configuration file")
	}
	return conf, nil
}

func SanitizeString(s string) string {
	// URL safe base64 alphabet without padding as defined in ACME
	re, err := regexp.Compile("[^A-Za-z\\-\\_0-9]+")
	if err != nil {
		log.Errorf("%v", err)
		return ""
	}
	return re.ReplaceAllString(s, "")
}

func GeneratePassword(length int) (string, error) {
	ret := make([]byte, length)
	const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1234567890-_"
	alphalen := big.NewInt(int64(len(alphabet)))
	for i := 0; i < length; i++ {
		c, err := rand.Int(rand.Reader, alphalen)
		if err != nil {
			return "", err
		}
		r := int(c.Int64())
		ret[i] = alphabet[r]
	}
	return string(ret), nil
}

func SanitizeDomainQuestion(d string) string {
	var dom string
	suffix := DNSConf.General.Domain + "."
	if strings.HasSuffix(d, suffix) {
		dom = d[0 : len(d)-len(suffix)]
	} else {
		dom = d
	}
	return dom
}

func NewACMETxt() (ACMETxt, error) {
	var a = ACMETxt{}
	password, err := GeneratePassword(40)
	if err != nil {
		return a, err
	}
	a.Username = uuid.NewV4()
	a.Password = password
	a.Subdomain = uuid.NewV4().String()
	return a, nil
}

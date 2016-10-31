package gravago

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/url"
)

const (
	baseURL       string = "//www.gravatar.com/avatar/"
	defaultIcons  string = "retro"
	defaultRating string = "g"
	defaultSize   int    = 64
)

// Gravatar - gravatar object
type Gravatar struct {
	gravatarURL string
	icons       string
	rating      string
	size        int
}

// New - create new gravatar object instance
func New() *Gravatar {
	g := &Gravatar{
		gravatarURL: baseURL,
		icons:       defaultIcons,
		rating:      defaultRating,
		size:        defaultSize,
	}

	return g
}

// SetSize - set size of gravatar image
func (g *Gravatar) SetSize(size int) {
	g.size = size
}

// SetRating - set rating
func (g *Gravatar) SetRating(rating string) {
	g.rating = rating
}

// SetIcons - set the default icon to return
func (g *Gravatar) SetIcons(iconset string) {
	g.icons = iconset
}

// UseHTTPS - force using HTTPS protocol
func (g *Gravatar) UseHTTPS(protocol bool) {
	if protocol {
		g.gravatarURL = "https://s.gravatar.com/avatar/"
	} else {
		g.gravatarURL = "http:" + baseURL
	}
}

// URL - generate gravatar URL string
func (g *Gravatar) URL(email string) string {
	hashedEmail := emailToHash(email)
	queryString := g.encodeParameters()

	gravURL := g.gravatarURL + hashedEmail + queryString

	return gravURL
}

func (g *Gravatar) encodeParameters() string {
	parameters := url.Values{}
	parameters.Add("s", fmt.Sprintf("%v", g.size))
	parameters.Add("r", fmt.Sprintf("%v", g.rating))
	parameters.Add("d", fmt.Sprintf("%v", g.icons))

	return "?" + parameters.Encode()
}

func emailToHash(email string) string {
	hasher := md5.New()
	hasher.Write([]byte(email))
	return fmt.Sprintf("%v", hex.EncodeToString(hasher.Sum(nil)))
}

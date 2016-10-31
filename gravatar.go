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

// gravatar object
type gravatar struct {
	GravatarURL string
	Icons       string
	Rating      string
	Size        int
}

// New - create new gravatar object instance
func New() *gravatar {
	g := &gravatar{
		GravatarURL: baseURL,
		Icons:       defaultIcons,
		Rating:      defaultRating,
		Size:        defaultSize,
	}

	return g
}

// SetSize - set size of gravatar image
func (g *gravatar) SetSize(size int) {
	g.Size = size
}

// SetRating - set rating
func (g *gravatar) SetRating(rating string) {
	g.Rating = rating
}

// SetIcons - set the default icon to return
func (g *gravatar) SetIcons(iconset string) {
	g.Icons = iconset
}

// UseHTTPS - force using HTTPS protocol
func (g *gravatar) UseHTTPS(protocol bool) {
	if protocol {
		g.GravatarURL = "https://s.gravatar.com/avatar/"
	} else {
		g.GravatarURL = "http:" + baseURL
	}
}

// URL - generate gravatar URL string
func (g *gravatar) URL(email string) string {
	hashedEmail := emailToHash(email)
	queryString := g.encodeParameters()

	gravURL := g.GravatarURL + hashedEmail + queryString

	return gravURL
}

func (g *gravatar) encodeParameters() string {
	parameters := url.Values{}
	parameters.Add("s", fmt.Sprintf("%v", g.Size))
	parameters.Add("r", fmt.Sprintf("%v", g.Rating))
	parameters.Add("d", fmt.Sprintf("%v", g.Icons))

	return "?" + parameters.Encode()
}

func emailToHash(email string) string {
	hasher := md5.New()
	hasher.Write([]byte(email))
	return fmt.Sprintf("%v", hex.EncodeToString(hasher.Sum(nil)))
}

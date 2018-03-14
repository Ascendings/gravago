package gravago

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/url"
	"strings"
)

// Gravatar - gravatar object
type Gravatar struct {
	gravatarURL string
	icons       string
	rating      string
	size        int
}

// New - create new gravatar object instance
func New(opts GravatarOptions) *Gravatar {
	// declare variables for this function
	var baseURL    string = ""
	var iconSet    string = ""
	var iconRating string = ""
	var iconSize   int    = 0

	// set the base URL
	if opts.ForceHTTPS == true {
		baseURL = strings.Join([]string{"https://s.", defaultBaseURL}, "")
	} else {
		baseURL = strings.Join([]string{"//www.", defaultBaseURL}, "")
	}
	// set the icon set
	if opts.IconSet != "" {
		iconSet = opts.IconSet
	} else {
		iconSet = defaultIconSet
	}
	// set the icon rating
	if opts.IconRating != "" {
		iconRating = opts.IconRating
	} else {
		iconRating = defaultIconRating
	}
	// set the icon size
	if opts.IconSize != 0 {
		iconSize = opts.IconSize
	} else {
		iconSize = defaultIconSize
	}

	g := &Gravatar{
		gravatarURL: baseURL,
		icons:       iconSet,
		rating:      iconRating,
		size:        iconSize,
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
		g.gravatarURL = "https://s." + defaultBaseURL
	} else {
		g.gravatarURL = "http://www." + defaultBaseURL
	}
}

// GetURL - generate gravatar URL string
func (g *Gravatar) GetURL(email string) string {
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

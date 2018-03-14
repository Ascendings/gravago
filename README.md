# Gravago - Gravatar library for Go

Simple go library for creating gravatar.com avatar URLs

**Current version:** 0.4.0

### Credits

This library has some inspiration from the  [eefret/gravatar](https://github.com/eefret/gravatar) Go library and the [node-gravatar](https://www.npmjs.com/package/gravatar) Node.js library.

#### Table of Contents

1. [Overview](#overview)
2. [Installation - Getting this package ready to roll](#installation)
3. [Usage - How to use this thing](#usage)
4. [License - Licensing information](#license)
5. [Contact - How to contact me](#contact)

## Installation

Just like any other Go package, you can install this baby with `go get ...`

`go get github.com/Ascendings/gravago`

## Usage

Now for the fun stuff... the meat! Seriously, though, this package isn't designed to be too difficult to use in a web application - just create a new instance, fiddle some options, then generate your URL!

For example:
```go
// of course we need to import this package
import "github.com/Ascendings/gravago"

// create a new instance
grav := gravago.New(gravago.GravatarOptions{
  ForceHTTPS: true,
  IconSet: "mosterid",
  IconRating: "x",
  IconSize: 80,
})

// change some options later on...
grav.SetSize(80) // set the icon size to 80
grav.SetIcons("monsterid") // use monsterid default icons
grav.SetRating("x") // allow x-rated icons
grav.UseHTTPS(true) // force HTTPS

// generate URL
gravURL := grav.URL("myemail@example.com")
```

## License

This package/program is licensed under the 2-Clause BSD License

## Contact

Email me at: brotherballantine@gmail.com

Feel free to hit me up at
  * [Mastodon](https://mastodon.rocks/@brotherballan)
  * [Google+](https://plus.google.com/+GregoryBallantine1)
  * [Wire](https://wire.com) @brotherballan

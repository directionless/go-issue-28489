# Reproduce for go issue 28489

https://github.com/golang/go/issues/28489


```
# Attempt upgrader
$ go get -u github.com/kolide/launcher
go: finding github.com/kolide/launcher latest
go: finding github.com/hailocab/go-hostpool latest
go: finding github.com/bmizerany/assert latest
go: finding github.com/bitly/go-hostpool latest
go: finding github.com/Shopify/logrus-bugsnag latest
go: finding github.com/kolide/osquery-go latest
go: finding github.com/kolide/updater latest
go: finding github.com/jinzhu/inflection latest
go: finding github.com/mixer/clock latest
go: finding github.com/kolide/kit latest
go: finding golang.org/x/sync latest
go: finding golang.org/x/time latest
go: finding golang.org/x/image latest
go: finding golang.org/x/sys latest
go: finding github.com/alexkohler/nakedret latest
go: finding github.com/WatchBeam/clock latest
go: finding github.com/cloudflare/cfssl latest
go: finding github.com/kr/logfmt latest
go: finding github.com/agl/ed25519 latest
go: github.com/Sirupsen/logrus@v1.4.1: parsing go.mod: unexpected module path "github.com/sirupsen/logrus"
go: finding github.com/serenize/snaker latest
go: finding github.com/groob/plist latest
go: finding golang.org/x/net latest
go get: error loading module requirements

# where is this coming from?
$ go mod why github.com/Sirupsen/logrus
# github.com/Sirupsen/logrus
(main module does not need package github.com/Sirupsen/logrus)

# where is this coming from?
$ go mod why github.com/sirupsen/logrus
# github.com/sirupsen/logrus
(main module does not need package github.com/sirupsen/logrus)

# maybe in the graph?
# (no, those are indirect depends, so this doesn't help me)
$ go mod graph | grep logrus
github.com/kolide/launcher@v0.0.0-20190227153714-69bed6381900 github.com/Shopify/logrus-bugsnag@v0.0.0-20171204204709-577dee27f20d
github.com/kolide/launcher@v0.0.0-20190227153714-69bed6381900 github.com/Sirupsen/logrus@v1.0.3
github.com/kolide/launcher@v0.0.0-20190227153714-69bed6381900 github.com/sirupsen/logrus@v1.2.0
github.com/kolide/launcher@v0.0.0-20190227153714-69bed6381900 gopkg.in/gemnasium/logrus-airbrake-hook.v2@v2.1.2
github.com/sirupsen/logrus@v1.2.0 github.com/davecgh/go-spew@v1.1.1
github.com/sirupsen/logrus@v1.2.0 github.com/konsorten/go-windows-terminal-sequences@v1.0.1
github.com/sirupsen/logrus@v1.2.0 github.com/pmezard/go-difflib@v1.0.0
github.com/sirupsen/logrus@v1.2.0 github.com/stretchr/objx@v0.1.1
github.com/sirupsen/logrus@v1.2.0 github.com/stretchr/testify@v1.2.2
github.com/sirupsen/logrus@v1.2.0 golang.org/x/crypto@v0.0.0-20180904163835-0709b304e793
github.com/sirupsen/logrus@v1.2.0 golang.org/x/sys@v0.0.0-20180905080454-ebe1bf3edb33
```

## The fix

don't upgrade dependancies

```
$ go get github.com/kolide/launcher
go: finding github.com/kolide/launcher latest
```

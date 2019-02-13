//go:generate statik -f -src=./data/files
package gnfinder

import (
	"github.com/gnames/bayes"
	"github.com/gnames/gnfinder/dict"
	"github.com/gnames/gnfinder/lang"
	"github.com/gnames/gnfinder/verifier"
)

// GNfinder is responsible for name-finding operations
type GNfinder struct {
	// Language of the text
	Language lang.Language
	// Bayes flag forces to run Bayes name-finding on unknown languages
	Bayes bool
	// BayesOddsThreshold sets the limit of posterior odds. Everything bigger
	// that this limit will go to the names output.
	BayesOddsThreshold float64
	// TextOdds captures "concentration" of names as it is found for the whole
	// text by heuristic name-finding. It should be close enough for real
	// number of names in text. We use it when we do not have local conentration
	// of names in a region of text.
	TextOdds bayes.LabelFreq

	// NameDistribution keeps data about position of names candidates and
	// their value according to heuristic and Bayes name-finding algorithms.
	// NameDistribution

	// Verifier for scientific names
	Verifier *verifier.Verifier
	// Dict contains black, grey, and white list dictionaries
	Dict *dict.Dictionary
}

// Option type for changing GNfinder settings.
type Option func(*GNfinder)

// OptLanguage sets a language of a text.
func OptLanguage(l lang.Language) Option {
	return func(gnf *GNfinder) {
		gnf.Language = l
	}
}

// OptBayes is an option that forces running bayes name-finding even when
// the language is not supported by training sets.
func OptBayes(b bool) Option {
	return func(gnf *GNfinder) {
		gnf.Bayes = b
	}
}

// OptBayesThreshold is an option for name finding, that sets new threshold
// for results from the Bayes name-finding. All the name candidates that have a
// higher threshold will appear in the resulting names output.
func OptBayesThreshold(odds float64) Option {
	return func(gnf *GNfinder) {
		gnf.BayesOddsThreshold = odds
	}
}

// OptVerify is sets Verifier that will be used for validation of
// name-strings against https://index.globalnames.org service.
func OptVerify(opts ...verifier.Option) Option {
	return func(gnf *GNfinder) {
		gnf.Verifier = verifier.NewVerifier(opts...)
	}
}

// OptDict allows to set already created dictionary for GNfinder.
// It saves time, because then dictionary does not have to be loaded at
// the construction time.
func OptDict(d *dict.Dictionary) Option {
	return func(gnf *GNfinder) {
		gnf.Dict = d
	}
}

// NewGNfinder creates GNfinder object with default data, or with data coming
// from opts.
func NewGNfinder(opts ...Option) *GNfinder {
	gnf := &GNfinder{
		Language:           lang.NotSet,
		BayesOddsThreshold: 100.0,
	}
	for _, opt := range opts {
		opt(gnf)
	}
	if gnf.Dict == nil {
		gnf.Dict = dict.LoadDictionary()
	}
	return gnf
}

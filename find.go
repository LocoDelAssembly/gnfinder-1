package gnfinder

import (
	"github.com/gnames/gnfinder/heuristic"
	"github.com/gnames/gnfinder/lang"
	"github.com/gnames/gnfinder/nlp"
	"github.com/gnames/gnfinder/output"
	"github.com/gnames/gnfinder/token"
)

// FindNamesJSON takes a text as bytes and returns JSON representation of
// scientific names found in the text
func (gnf *GNfinder) FindNamesJSON(data []byte) []byte {
	output := gnf.FindNames(data)
	return output.ToJSON()
}

// FindNames traverses a text and finds scientific names in it.
func (gnf *GNfinder) FindNames(data []byte) *output.Output {
	text := []rune(string(data))
	tokens := token.Tokenize(text)

	if !gnf.LanguageForced {
		gnf.LanguageUsed, gnf.LanguageDetected = lang.DetectLanguage(text)
	}
	if !gnf.BayesForced && gnf.LanguageUsed != lang.UnknownLanguage {
		gnf.Bayes = true
	}

	heuristic.TagTokens(tokens, gnf.Dict)
	if gnf.Bayes {
		nb := gnf.BayesWeights[gnf.LanguageUsed]
		nlp.TagTokens(tokens, gnf.Dict, nb, gnf.BayesOddsThreshold)
	}
	return output.TokensToOutput(tokens, text, gnf.LanguageUsed,
		gnf.LanguageDetected)
}

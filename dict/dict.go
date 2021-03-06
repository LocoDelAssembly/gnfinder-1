// package dict contains dictionaries for finding scientific names
package dict

import (
	"encoding/csv"
	"io"
	"log"

	"github.com/gnames/gnfinder/fs"
)

// DictionaryType describes available dictionaries
type DictionaryType int

func (d DictionaryType) String() string {
	types := [...]string{"NotSet", "WhiteGenus", "GreyGenus",
		"WhiteUninomial", "GreyUninomial", "BlackUninomial", "WhiteSpecies",
		"GreySpecies", "BlackSpecies", "CommonWords", "Rank", "NotInDictionary"}
	return types[d]
}

// DictionaryType dictionaries
const (
	NotSet DictionaryType = iota
	WhiteGenus
	GreyGenus
	WhiteUninomial
	GreyUninomial
	BlackUninomial
	WhiteSpecies
	GreySpecies
	BlackSpecies
	CommonWords
	Rank
	NotInDictionary
)

// Dictionary contains dictionaries used for detecting scientific names
type Dictionary struct {
	BlackUninomials map[string]struct{}
	BlackSpecies    map[string]struct{}
	CommonWords     map[string]struct{}
	GreyGenera      map[string]struct{}
	GreyGeneraSp    map[string]struct{}
	GreySpecies     map[string]struct{}
	GreyUninomials  map[string]struct{}
	WhiteGenera     map[string]struct{}
	WhiteSpecies    map[string]struct{}
	WhiteUninomials map[string]struct{}
	Ranks           map[string]struct{}
}

// LoadDictionary contain most popular words in European languages.
func LoadDictionary() *Dictionary {
	d := &Dictionary{
		BlackUninomials: readData("/black/uninomials.csv"),
		BlackSpecies:    readData("/black/species.csv"),
		CommonWords:     readData("/common/eu.csv"),
		GreyGenera:      readData("/grey/genera.csv"),
		GreyGeneraSp:    readData("/grey/genera_species.csv"),
		GreySpecies:     readData("/grey/species.csv"),
		GreyUninomials:  readData("/grey/uninomials.csv"),
		WhiteGenera:     readData("/white/genera.csv"),
		WhiteSpecies:    readData("/white/species.csv"),
		WhiteUninomials: readData("/white/uninomials.csv"),
		Ranks:           setRanks(),
	}
	return d
}

func readData(path string) map[string]struct{} {
	res := make(map[string]struct{})
	f, err := fs.Files.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	var empty struct{}

	defer func() {
		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	reader := csv.NewReader(f)
	for {
		v, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		res[v[0]] = empty
	}
	return res
}

func setRanks() map[string]struct{} {
	var empty struct{}
	ranks := map[string]struct{}{
		"nat": empty, "f.sp": empty, "mut.": empty, "morph.": empty,
		"nothosubsp.": empty, "convar.": empty, "pseudovar": empty, "sect.": empty,
		"ser.": empty, "subvar.": empty, "subf.": empty, "race": empty,
		"α": empty, "ββ": empty, "β": empty, "γ": empty, "δ": empty, "ε": empty,
		"φ": empty, "θ": empty, "μ": empty, "a.": empty, "b.": empty,
		"c.": empty, "d.": empty, "e.": empty, "g.": empty, "k.": empty,
		"pv.": empty, "pathovar.": empty, "ab.": empty, "st.": empty, "fm.": empty,
		"variety": empty, "var": empty, "var.": empty, "forma": empty, "fm": empty,
		"forma.": empty, "fma": empty, "fma.": empty, "form": empty, "form.": empty,
		"fo": empty, "fo.": empty, "f": empty, "f.": empty, "ssp": empty,
		"ssp.": empty, "subsp": empty, "subsp.": empty,
	}
	return ranks
}

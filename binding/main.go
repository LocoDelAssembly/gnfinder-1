package main

/*
	#include "stdlib.h"
	#include "callback_bridge.h"
*/
import "C"

import (
	"unsafe"

	"log"
	"encoding/json"
	"github.com/gnames/gnfinder"
	"github.com/gnames/gnfinder/lang"
	"github.com/gnames/gnfinder/verifier"
)

//export FindNamesToJSON
// NOTE: Read callback type as "void (*callback)(char *output)"
//export FindNamesToJSON
func FindNamesToJSON(txt *C.char, json_opts *C.char, callback unsafe.Pointer) {
	gotxt := C.GoString(txt)
	gnf := gnfinder.NewGNfinder()
	output := gnf.FindNamesJSON([]byte(gotxt), setOpts(json_opts)...)

	p := C.CString(string(output))
	defer C.free(unsafe.Pointer(p))

	C.callback_bridge(callback, p)
}

type Params struct {
	Text                 string   `json:"text,omitempty"`
	NoBayes              bool     `json:"no_bayes,omitempty"`
	Language             string   `json:"language,omitempty"`
	DetectLanguage       bool     `json:"detect_language,omitempty"`
	Verification         bool     `json:"verification,omitempty"`
	Sources              []int32  `json:"sources,omitempty"`
}

func setOpts(json_opts *C.char) []gnfinder.Option {
	var params Params

	opts := []gnfinder.Option{}

	err := json.Unmarshal([]byte(C.GoString(json_opts)), &params)

	if err != nil {
		log.Println(err)
		return opts
	}

	opts = append(opts, gnfinder.OptBayes(!params.NoBayes))

	if params.Verification {
		var verOpts []verifier.Option
		var sources []int
		for _, v := range params.Sources {
			sources = append(sources, int(v))
		}
		verOpts = append(verOpts, verifier.OptSources(sources))

		opts = append(opts, gnfinder.OptVerify(verOpts...))
	}

	if params.DetectLanguage {
		opts = append(opts, gnfinder.OptDetectLanguage(true))
	} else if len(params.Language) > 0 {
		l, err := lang.NewLanguage(params.Language)
		if err == nil {
			opts = append(opts, gnfinder.OptLanguage(l))
		}
	}

	return opts
}

func main() {}

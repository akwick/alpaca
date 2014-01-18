package alpaca

import (
	"bitbucket.org/pkg/inflect"
	"path/filepath"
	"strings"
)

const (
	Version = "0.2.0"
)

var (
	LibraryRoot string
)

type Data struct {
	Pkg PkgStruct
	Api map[string]interface{}
	Doc map[string]interface{}
	Fnc map[string]interface{}
}

type LanguageOptions struct {
	Php    bool `long:"no-php" description:"Do not write php library"`
	Python bool `long:"no-python" description:"Do not write python library"`
	Ruby   bool `long:"no-ruby" description:"Do not write ruby library"`
	Node   bool `long:"no-node" description:"Do not write node library"`
}

func WriteLibraries(directory string, opts *LanguageOptions) {
	var err error

	LibraryRoot, err = filepath.Abs(directory)
	HandleError(err)

	data := ReadData()
	ModifyData(data)

	if !opts.Php {
		CheckPhp(data)
		FunctionsPhp(data.Fnc)
		WritePhp(data)
	}

	if !opts.Python {
		CheckPython(data)
		FunctionsPython(data.Fnc)
		WritePython(data)
	}

	if !opts.Ruby {
		CheckRuby(data)
		FunctionsRuby(data.Fnc)
		WriteRuby(data)
	}

	if !opts.Node {
		CheckNode(data)
		FunctionsNode(data.Fnc)
		WriteNode(data)
	}
}

func ReadData() *Data {
	var pkg PkgStruct
	var api, doc map[string]interface{}

	ReadJSON("pkg.json", &pkg)
	ReadJSON("api.json", &api)
	ReadJSON("doc.json", &doc)

	return &Data{pkg, api, doc, make(map[string]interface{})}
}

func ModifyData(data *Data) {
	data.Api["alpaca_version"] = Version

	data.Api["classes"] = MapKeysToStringArray(data.Api["class"], []string{})

	data.Fnc["join"] = strings.Join
	data.Fnc["upper"] = strings.ToUpper

	data.Fnc["camelize"] = inflect.Camelize
	data.Fnc["camelizeDownFirst"] = inflect.CamelizeDownFirst
	data.Fnc["underscore"] = inflect.Underscore

	data.Fnc["methods"] = MethodList

	data.Fnc["args"] = make(map[string]interface{})
	data.Fnc["path"] = make(map[string]interface{})
	data.Fnc["prnt"] = make(map[string]interface{})
}
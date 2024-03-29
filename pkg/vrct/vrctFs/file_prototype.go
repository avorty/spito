package vrctFs

import (
	"encoding/json"
	"errors"
	"github.com/pelletier/go-toml/v2"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
	"strings"
)

type FilePrototype struct {
	Layers               []PrototypeLayer
	RealFileExists       bool
	FileType             FileType
	OriginalFileIncluded bool
	Path                 string `bson:"-"`
	Name                 string `bson:"-"`
}

func (p *FilePrototype) getDestinationPath() string {
	newPath := strings.TrimPrefix(p.getVirtualPath(), VirtualFsPathPrefix)

	// Remove first slash
	newPath = newPath[1:]

	firstSlashIndex := strings.Index(newPath, "/")
	newPath = newPath[firstSlashIndex:]

	return strings.TrimSuffix(newPath, ".prototype.bson")
}

func (p *FilePrototype) getVirtualPath() string {
	return filepath.Join(p.Path, p.Name+VirtualFilePostfix)
}

func (p *FilePrototype) SimulateFile() ([]byte, error) {
	finalLayer, err := p.mergeLayers()

	if err != nil {
		return nil, err
	}

	var filePath string

	if finalLayer.ContentPath == "" {
		filePath = p.getDestinationPath()
	} else {
		filePath = finalLayer.ContentPath
	}

	file, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var tempContentInterface map[string]interface{}
	if p.FileType != TextFile {
		err = bson.Unmarshal(file, &tempContentInterface)
		if err != nil {
			return file, err
		}
	}

	var fileContent []byte
	switch p.FileType {
	case JsonConfig:
		fileContent, err = json.Marshal(tempContentInterface)
		break
	case YamlConfig:
		fileContent, err = yaml.Marshal(tempContentInterface)
		break
	case TomlConfig:
		fileContent, err = toml.Marshal(tempContentInterface)
		break
	default:
		return file, nil
	}
	return fileContent, err
}

func (p *FilePrototype) Read(vrctPrefix string, realPath string) error {
	prototypeFilePath := filepath.Join(vrctPrefix, realPath)

	path := filepath.Dir(prototypeFilePath)
	name := filepath.Base(prototypeFilePath)

	p.Path = path
	p.Name = name
	file, err := os.ReadFile(p.getVirtualPath())

	if os.IsNotExist(err) {
		_, err := os.Stat(realPath)
		p.RealFileExists = !os.IsNotExist(err)

		return p.Save()
	} else if err != nil {
		return err
	}

	err = bson.Unmarshal(file, p)

	p.Path = path
	p.Name = name
	return err
}

func (p *FilePrototype) Save() error {
	rawBson, err := bson.Marshal(p)
	if err != nil {
		return err
	}
	return os.WriteFile(p.getVirtualPath(), rawBson, os.ModePerm)
}

func (p *FilePrototype) CreateLayer(content []byte, options []byte, isOptional bool) (PrototypeLayer, error) {
	if p.Path == "" {
		return PrototypeLayer{}, errors.New("file prototype hasn't been loaded yet")
	}

	dir := filepath.Dir(p.Path)

	fileNameOk := false
	var contentPath string
	for !fileNameOk {
		randFileName := randomLetters(5)
		contentPath = filepath.Join(dir, randFileName)
		_, err := os.Stat(contentPath)
		if err != nil {
			fileNameOk = true
		}
	}

	optionNameOk := false
	var optionsPath string
	for !optionNameOk {
		randOptsName := randomLetters(5)
		optionsPath = filepath.Join(dir, randOptsName)
		_, err := os.Stat(optionsPath)
		if err != nil {
			optionNameOk = true
		}
	}

	tempConvertedContent, err := GetMapFromBytes(content, p.FileType)

	if p.FileType != TextFile {
		content, err = bson.Marshal(tempConvertedContent)
		if err != nil {
			return PrototypeLayer{}, err
		}
	}

	if err = os.WriteFile(contentPath, content, os.ModePerm); err != nil {
		return PrototypeLayer{}, err
	}

	if options == nil {
		options = []byte("{}")
	}

	var tempOptionalKeysMap map[string]interface{}
	err = json.Unmarshal(options, &tempOptionalKeysMap)
	if err != nil {
		err = yaml.Unmarshal(options, &tempOptionalKeysMap)
		if err != nil {
			return PrototypeLayer{}, err
		}
	}

	optionalKeysBson, err := bson.Marshal(tempOptionalKeysMap)
	if err != nil {
		return PrototypeLayer{}, err
	}

	if err = os.WriteFile(optionsPath, optionalKeysBson, os.ModePerm); err != nil {
		return PrototypeLayer{}, err
	}

	newLayer := PrototypeLayer{
		ContentPath: contentPath,
		OptionsPath: optionsPath,
		IsOptional:  isOptional,
	}

	return newLayer, nil
}

func (p *FilePrototype) AddNewLayer(layer PrototypeLayer, isOriginal bool) error {
	backup := p.Layers
	p.Layers = append(p.Layers, layer)
	_, err := p.mergeLayers()
	if err != nil {
		p.Layers = backup
		return err
	}

	p.OriginalFileIncluded = isOriginal

	if err = p.Save(); err != nil {
		return err
	}

	return nil
}

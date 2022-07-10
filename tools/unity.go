package tools

import (
	"github.com/google/uuid"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type UnityMetadata struct {
	FileFormatVersion int    `yaml:"fileFormatVersion"`
	Guid              string `yaml:"guid"`
}

func GetGUID(path string, create bool) (string, error) {
	metadataPath := path + ".meta"
	if FileExists(metadataPath) {
		data, err := os.ReadFile(metadataPath)
		if err != nil {
			return "", err
		}

		var metadata UnityMetadata
		if err := yaml.Unmarshal(data, &metadata); err != nil {
			return "", err
		}

		return metadata.Guid, nil
	}
	if create {
		return CreateGUID(), nil
	}
	return "", nil
}

func CreateMetadata(path string, guid string) error {
	data, err := yaml.Marshal(&UnityMetadata{
		FileFormatVersion: 2,
		Guid:              guid,
	})
	if err != nil {
		log.Fatalln("Unable to create metadata file:", err)
		return err
	}

	err = ioutil.WriteFile(path+".meta", data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func CreateGUID() string {
	return strings.Replace(uuid.NewString(), "-", "", -1)
}

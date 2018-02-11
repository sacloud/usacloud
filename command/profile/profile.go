package profile

import (
	"encoding/json"
	"fmt"
	"github.com/mitchellh/go-homedir"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

const ProfileDirEnv = "USACLOUD_PROFILE_DIR"
const DefaultProfileName = "default"

var (
	configDirName   = ".usacloud"
	configFileName  = "config.json"
	currentFileName = "current"
)

func ValidateProfileName(profileName string, invalidRunes ...rune) error {
	invalids := invalidRunes
	if len(invalids) == 0 {
		// validate profileName
		invalids = []rune{filepath.ListSeparator, filepath.Separator}
	}

	for _, r := range invalids {
		if strings.ContainsRune(profileName, r) {
			return fmt.Errorf("ProfileName[%q] is invalid", profileName)
		}
	}
	return nil
}

func getProfileBaseDir() (string, error) {

	// from profileDirEnv var
	if path, isset := os.LookupEnv(ProfileDirEnv); isset {
		if err := ValidateProfileName(path, filepath.ListSeparator); err != nil {
			return "", fmt.Errorf("getProficeBaseDir(from %s) is failed:%s", ProfileDirEnv, err)
		}
		return filepath.Clean(path), nil
	}

	// default, use homedir
	homeDir, err := homedir.Dir()
	if err != nil {
		return "", fmt.Errorf("getProficeBaseDir(use HomeDir) is failed:%s", err)
	}
	return homeDir, nil
}

func GetConfigFilePath(profileName string) (string, error) {

	if err := ValidateProfileName(profileName); err != nil {
		return "", err
	}

	if profileName == "" {
		profileName = DefaultProfileName
	}
	baseDir, err := getProfileBaseDir()
	if err != nil {
		return "", fmt.Errorf("getting HomeDir is failed:%s", err)
	}
	return filepath.Clean(filepath.Join(baseDir, configDirName, filepath.Clean(profileName), configFileName)), nil
}

type ConfigFileValue struct {
	AccessToken       string
	AccessTokenSecret string
	Zone              string
	AcceptLanguage    string   `json:",omitempty"`
	RetryMax          int      `json:",omitempty"`
	RetryIntervalSec  int64    `json:",omitempty"`
	NoColor           bool     `json:",omitempty"`
	DefaultOutputType string   `json:",omitempty"`
	Zones             []string `json:",omitempty"`
	APIRootURL        string   `json:",omitempty"`
}

func (p *ConfigFileValue) IsEmpty() bool {
	return p.AccessToken == "" &&
		p.AccessTokenSecret == "" &&
		p.Zone == "" && p.DefaultOutputType == ""
}

func SaveConfigFile(profileName string, val *ConfigFileValue) error {

	if val == nil || val.IsEmpty() {
		return fmt.Errorf("ConfigFileValue is empty")
	}

	path, err := GetConfigFilePath(profileName)
	if err != nil {
		return err
	}

	// create dir
	dir := filepath.Dir(path)
	if _, err := os.Stat(dir); err != nil {
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			return fmt.Errorf("Profile: Mkdir(%q) is failed: %s", dir, err)
		}
	}

	rawBody, err := json.MarshalIndent(val, "", "\t")
	if err != nil {
		return fmt.Errorf("Profile: Marshal configFileValue to JSON is failed: %s", err)
	}

	err = ioutil.WriteFile(path, rawBody, 0600)
	if err != nil {
		return fmt.Errorf("Profile: Writing configFileValue to %q is failed: %s", path, err)
	}

	return nil
}

func LoadConfigFile(profileName string) (*ConfigFileValue, error) {
	v := &ConfigFileValue{}
	filePath, err := GetConfigFilePath(profileName)
	if err != nil {
		return v, err
	}

	// file exists?
	if _, err := os.Stat(filePath); err == nil {
		// read file
		buf, err := ioutil.ReadFile(filePath)
		if err != nil {
			return v, fmt.Errorf("Loading config file[%q] is failed: %s", filePath, err)
		}
		if err := json.Unmarshal(buf, v); err != nil {
			return v, fmt.Errorf("Parsing config file[%q] is failed: %s", filePath, err)
		}
	} else if profileName != DefaultProfileName {
		return nil, fmt.Errorf("Profile[%q] is not exists", profileName)
	}

	return v, nil
}

func RemoveConfigFile(profileName string) error {

	path, err := GetConfigFilePath(profileName)
	if err != nil {
		return err
	}

	dir := filepath.Dir(path)

	if _, err := os.Stat(dir); err != nil {
		return fmt.Errorf("Removing config file is failed: directory(%q) is not exists", dir)
	}

	if _, err := os.Stat(path); err != nil {
		return fmt.Errorf("Removing config file is failed: file(%q) is not exists", path)
	}

	// remove file
	if err := os.Remove(path); err != nil {
		return fmt.Errorf("Removing config file(%q) is failed: %s", path, err)
	}

	// remove dir if dir is empty
	info, err := ioutil.ReadDir(dir)
	if err != nil {
		return fmt.Errorf("Removing config file is failed: ReadDir(%q) is failed: %s", dir, err)
	}
	if len(info) == 0 {
		// remove dir
		if err := os.RemoveAll(dir); err != nil {
			return fmt.Errorf("Removing config dir(%q) is failed: %s", dir, err)
		}
	}

	current, err := GetCurrentName()
	if err != nil {
		return fmt.Errorf("Removing config file is failed: GetCurrentName if failed: %s", err)
	}

	if current == profileName {
		if err := SetCurrentName(DefaultProfileName); err != nil {
			return fmt.Errorf("Removing config file is failed: SetCurrentName if failed: %s", err)
		}
	}
	return nil
}

func GetCurrentName() (string, error) {
	baseDir, err := getProfileBaseDir()
	if err != nil {
		return "", err
	}

	profNameFile := filepath.Join(baseDir, configDirName, currentFileName)
	if _, err := os.Stat(profNameFile); err == nil {
		data, err := ioutil.ReadFile(profNameFile)
		if err != nil {
			return "", fmt.Errorf("Reading current profile file is failed:%s", err)
		}
		profileName := string(data)
		if err := ValidateProfileName(profileName); err != nil {
			return "", err
		}

		profileName = cleanupProfileName(profileName)
		if profileName == "" {
			profileName = DefaultProfileName
		}
		return profileName, nil
	}

	return DefaultProfileName, nil
}

func cleanupProfileName(profileName string) string {
	targets := []string{"　", "\t", "\n"}
	res := profileName
	for _, s := range targets {
		res = strings.Replace(res, s, "", -1)
	}
	return strings.Trim(res, " ")
}

func SetCurrentName(profileName string) error {

	if err := ValidateProfileName(profileName); err != nil {
		return err
	}

	profileName = cleanupProfileName(profileName)

	baseDir, err := getProfileBaseDir()
	if err != nil {
		return err
	}

	configDir := filepath.Join(baseDir, configDirName)
	if _, err := os.Stat(configDir); err != nil {
		err := os.MkdirAll(configDir, 0755)
		if err != nil {
			return fmt.Errorf("Creating config dir[%q] is failed:%s", configDir, err)
		}
	}

	if profileName != DefaultProfileName {
		profileConfigPath := filepath.Join(configDir, profileName, configFileName)
		if _, err := os.Stat(profileConfigPath); err != nil {
			return fmt.Errorf("Profile[%s] is not exists", profileName)
		}
	}

	profNameFile := filepath.Join(baseDir, configDirName, currentFileName)
	if err := ioutil.WriteFile(profNameFile, []byte(profileName), 0600); err != nil {
		return fmt.Errorf("Writing profile name file[%q] is failed:%s", profNameFile, err)
	}

	return nil
}

func List() ([]string, error) {
	res := []string{"default"}

	// get profile dirs under base dir
	baseDir, err := getProfileBaseDir()
	if err != nil {
		return []string{}, fmt.Errorf("List profiles is failed: %s", err)
	}
	configDirPath := filepath.Join(baseDir, configDirName)
	if _, err := os.Stat(configDirPath); err != nil {
		return res, nil
	}
	entries, err := ioutil.ReadDir(filepath.Join(baseDir, configDirName))
	if err != nil {
		return []string{}, fmt.Errorf("List profiles is failed: %s", err)
	}

	// validate each profile dir
	for _, fi := range entries {
		if fi.IsDir() {
			profile := filepath.Base(fi.Name())
			if profile != DefaultProfileName {
				// has valid config?
				_, err := LoadConfigFile(profile)
				if err == nil {
					res = append(res, profile)
				}
			}
		}
	}

	return res, nil
}

package constant

import (
	"os"
	P "path"
	"path/filepath"
)

// Path is used to get the configuration path
var Path *path

type path struct {
	homeDir    string
	indexDir   string
	originDir  string
	configFile string
}

func init() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		homeDir, _ = os.Getwd()
	}

	homeDir = P.Join(homeDir, Name)
	Path = &path{homeDir: homeDir, configFile: "config.yaml"}
	_, err = os.Stat(homeDir)
	if err != nil {
		err := os.Mkdir(homeDir, os.ModePerm)
		if err != nil {
			Path.homeDir = ""
		}
	}
	Path.originDir = filepath.Join(homeDir, "origin")
	_, err = os.Stat(Path.originDir)
	if err != nil {
		_ = os.Mkdir(Path.originDir, os.ModePerm)
	}
	Path.indexDir = filepath.Join(homeDir, "index")
	_, err = os.Stat(Path.indexDir)
	if err != nil {
		_ = os.Mkdir(Path.indexDir, os.ModePerm)
	}
}

// SetHomeDir is used to set the configuration path
func SetHomeDir(root string) {
	Path.homeDir = root
}

// SetConfig is used to set the configuration file
func SetConfig(file string) {
	Path.configFile = file
}

func (p *path) HomeDir() string {
	return p.homeDir
}

func (p *path) IndexDir() string {
	return p.indexDir
}

func (p *path) OriginDir() string {
	return p.originDir
}

func (p *path) Config() string {
	return p.configFile
}

// Resolve return a absolute path or a relative path with homedir
func (p *path) Resolve(path string) string {
	if !filepath.IsAbs(path) {
		return filepath.Join(p.HomeDir(), path)
	}

	return path
}

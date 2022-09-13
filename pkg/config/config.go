package config

import (
	"github.com/alecthomas/kong"
	log "github.com/sirupsen/logrus"
	"gopkg.in/ini.v1"
)

type Application struct {
	ConfigFile *ini.File
	ConfigPath string
}

func (a *Application) Section(sectionName string) *ini.Section {
	section, err := a.ConfigFile.GetSection(sectionName)
	if err != nil {
		return nil
	}
	return section
}

func (a *Application) GetRegion() string {
	value, err := a.Section("default").GetKey("region")
	if err != nil {
		return ""
	}
	return value.Value()
}

func (a *Application) GetRoleARN() string {
	value, err := a.Section("default").GetKey("role_arn")
	if err != nil {
		return ""
	}
	return value.Value()
}

func (a *Application) GetAssumeARN() string {
	value, err := a.Section("default").GetKey("assume_arn")
	if err != nil {
		return ""
	}
	return value.Value()
}

func (a *Application) Load() {
	absolutePath := kong.ExpandPath(a.ConfigPath)
	iniCfg, err := ini.Load(absolutePath)
	if err != nil {
		log.Fatalf("Could not read default config: %+v\n", err)
	}
	a.ConfigFile = iniCfg
}

func New(cfgPath string) *Application {
	return &Application{ConfigPath: cfgPath}
}

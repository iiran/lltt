package setting

import (
	"encoding/json"
	"github.com/iiran/lltt/pkg/helper"
	"io/ioutil"
)

// ServerConfigDB is database field in global config
type ServerConfigDB struct {
	Name     string `json:"name"`
	Dialect  string `json:"dialect"`
	Username string `json:"username"`
	Password string `json:"password"`
	Address  string `json:"address"`
	Port     int64  `json:"port"`
	Database string `json:"database"`
}

type ServerConfigLog struct {
	PathName   string `json:"file_path"`
	FileName   string `json:"file_name"`
	FileExt    string `json:"file_ext"`
	LinePrefix string `json:"line_prefix"`
	TraceDepth int    `json:"trace_depth"`
}

type ServerConfigSession struct {
	Secret string `json:"secret"`
	Count  int64  `json:"count"`
}

// ServerConfig is global config
type ServerConfig struct {
	Address      string              `json:"address"`
	Port         int64               `json:"port"`
	Mode         string              `json:"mode"`
	PageSize     int64               `json:"page_size"`
	JWTSecret    string              `json:"jwt_secret"`
	WriteTimeout int64               `json:"write_timeout"`
	Session      ServerConfigSession `json:"session"`
	Database     []ServerConfigDB    `json:"database"`
	Log          ServerConfigLog     `json:"log"`
}

type RunMode string

const (
	ReleaseMode RunMode = "release"
	DebugMode   RunMode = "debug"
)

// Cfg share config for all module
var Cfg = &ServerConfig{}

// Setup init Cfg by config file
func Setup(filepath string) {
	var (
		cfgData []byte
		err     error
	)

	cfgData, err = ioutil.ReadFile(filepath)
	helper.PanicIfErr(err)

	err = json.Unmarshal(cfgData, Cfg)
	helper.PanicIfErr(err)
}

// GetDBConfig return DBConfig slice in Cfg
func GetDBConfig() (dbcfg *[]ServerConfigDB) {
	if Cfg == nil {
		return nil
	}
	return &Cfg.Database
}

func GetLogConfig() (logcfg *ServerConfigLog) {
	if Cfg == nil {
		return nil
	}
	return &Cfg.Log
}

func GetMode() (m string) {
	if Cfg == nil {
		return string(DebugMode)
	}
	return Cfg.Mode
}

func GetSessionConfig() *ServerConfigSession {
	if Cfg == nil {
		return nil
	}
	return &Cfg.Session
}

package test

import (
	"github.com/iiran/lltt/pkg/setting"
	"testing"
)

func TestSetup(t *testing.T) {
	var (
		testFilepath string = "config.json"
	)
	setting.Setup(testFilepath)
	if setting.Cfg.Address != "127.0.0.1" {
		t.Error("address invalid")
	}
	if setting.Cfg.Port != 44540 {
		t.Error("port invalid")
	}
	if setting.Cfg.Mode != "debug" {
		t.Error("mode invalid")
	}
	if setting.Cfg.PageSize != 10 {
		t.Error("page size invalid")
	}
	if setting.Cfg.JWTSecret != "123456" {
		t.Error("jwt secret invalid")
	}
	if setting.Cfg.WriteTimeout != 6000 {
		t.Error("write timeout invalid")
	}
	if len(setting.Cfg.Database) != 2 {
		t.Error("database num invalid")
	}
	if setting.Cfg.Database[0].Dialect != "postgres" {
		t.Error("db1 dialect error")
	}
	if setting.Cfg.Database[0].Username != "yiranfeng" {
		t.Error("db1 username error")
	}
	if setting.Cfg.Database[0].Password != "" {
		t.Error("db1 password error")
	}
	if setting.Cfg.Database[0].Address != "127.0.0.1" {
		t.Error("db1 address error")
	}
	if setting.Cfg.Database[0].Port != 5432 {
		t.Error("db1 port error")
	}
	if setting.Cfg.Database[0].Database != "leas" {
		t.Error("db1 database error")
	}
	if setting.Cfg.Database[0].Name != "leas" {
		t.Error("db1 name error")
	}
	if len(setting.Cfg.Database[0].TableLink) != 4 {
		t.Error("db1 table link num error")
	}
	if setting.Cfg.Database[0].TableLink[3].From != "public.post_tag" {
		t.Error("db1 link4 from error")
	}
	if setting.Cfg.Database[0].TableLink[3].To != "post_tag" {
		t.Error("db1 link4 to error")
	}
	if setting.Cfg.Log.PathName != "logs/" {
		t.Error("log file path error")
	}
	if setting.Cfg.Log.FileName != "lltt" {
		t.Error("log file name error")
	}
	if setting.Cfg.Log.FileExt != "log" {
		t.Error("log file ext error")
	}
	if setting.Cfg.Log.LinePrefix != "lltt: " {
		t.Error("log line prefix error")
	}
	if setting.Cfg.Log.TraceDepth != 2 {
		t.Error("log trace depth error")
	}
}

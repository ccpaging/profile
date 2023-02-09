package profile

import (
	"github.com/ccpaging/profile/config"
	"github.com/spf13/cast"
)

var DELIMITER = config.DELIMITER

type Profile struct {
	*config.Config
	section string
}

func New(dsn string, section string) (*Profile, error) {
	return &Profile{
		section: section,
		Config:  config.New(dsn),
	}, nil
}

func (pf *Profile) WithSection(section string) *Profile {
	return &Profile{
		section: section,
		Config:  pf.Config,
	}
}

/* profile set and get */

func (pf *Profile) WriteString(key string, str string) {
	pf.Config.WriteValue(pf.section, key, str)
}

func (pf *Profile) WriteBool(key string, b bool) {
	pf.Config.WriteValue(pf.section, key, b)
}

func (pf *Profile) WriteInt(key string, n int) {
	pf.Config.WriteValue(pf.section, key, n)
}

func (pf *Profile) WriteInt64(key string, n int64) {
	pf.Config.WriteValue(pf.section, key, n)
}

func (pf *Profile) WriteValue(key string, v any) {
	pf.Config.WriteValue(pf.section, key, v)
}

func (pf *Profile) HasKey(key string) (interface{}, bool) {
	return pf.Config.HasKey(pf.section, key)
}

func (pf *Profile) GetString(key string, sDefault string) string {
	v := pf.Config.GetValue(pf.section, key, sDefault)
	return cast.ToString(v)
}

func (pf *Profile) GetBool(key string, bDefault bool) bool {
	v := pf.Config.GetValue(pf.section, key, bDefault)
	return cast.ToBool(v)
}

func (pf *Profile) GetInt(key string, nDefault int) int {
	v := pf.Config.GetValue(pf.section, key, nDefault)
	return cast.ToInt(v)
}

func (pf *Profile) GetInt64(key string, nDefault int64) int64 {
	v := pf.Config.GetValue(pf.section, key, nDefault)
	return cast.ToInt64(v)
}

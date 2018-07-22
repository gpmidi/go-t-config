package go_t_config

import "github.com/spf13/viper"

type SubConfigViper struct {
	name     string
	parent   *SubConfigI
	children map[string]*SubConfigI
	cfg      *viper.Viper
	l        *TLogI
}

func NewSubConfigViper(name string, v *viper.Viper, parent *SubConfigI, log *TLogI) (*SubConfigViper, error) {
	sc := SubConfigViper{parent: parent, l: log, cfg: v}

	return &sc, nil
}

func (s *SubConfigViper) GetLogger() *TLogI {
	return s.l
}

func (s *SubConfigViper) SetLogger(log *TLogI) {
	s.l = log
}

func (s *SubConfigViper) GetParent() *SubConfigI {
	return s.parent
}

func (s *SubConfigViper) SetParent(parent *SubConfigI) {
	s.parent = parent
}

func (s *SubConfigViper) GetChild(key string) *SubConfigI {
	return s.children[key]
}

func (s *SubConfigViper) SetChild(key string, child *SubConfigI) {
	s.children[key] = child
}

func (s *SubConfigViper) GetName() string {
	return s.name
}

func (s *SubConfigViper) GetString(key string) (string, error) {
	return s.cfg.GetString(key), nil
}

func (s *SubConfigViper) GetBool(key string) (bool, error) {
	return s.cfg.GetBool(key), nil
}

func (s *SubConfigViper) GetInt64(key string) (int64, error) {
	return s.cfg.GetInt64(key), nil
}

func (s *SubConfigViper) GetFloat64(key string) (float64, error) {
	return s.cfg.GetFloat64(key), nil
}

func (s *SubConfigViper) SetString(key string, value string) (error) {
	s.cfg.Set(key, value)
	return nil
}

func (s *SubConfigViper) SetBool(key string, value bool) (error) {
	s.cfg.Set(key, value)
	return nil
}

func (s *SubConfigViper) SetInt64(key string, value int64) (error) {
	s.cfg.Set(key, value)
	return nil
}

func (s *SubConfigViper) SetFloat64(key string, value float64) (error) {
	s.cfg.Set(key, value)
	return nil
}

package go_t_config

import "strings"

type TConfig struct {
	l          *TLogI
	rootConfig SubConfigI
	defaultSub string
}

func NewTConfig(loggers *TLogI) (*TConfig, error) {
	t := TConfig{l: loggers}
	return &t, nil
}

func (t *TConfig) SetDefaultWriteSub(sub string) {
	t.defaultSub = sub
}

func (t *TConfig) GetLogger() *TLogI {
	return t.l
}

func (t *TConfig) SetLogger(log *TLogI) {
	t.l = log
}

func (t *TConfig) GetSubConfig(sub string) (SubConfigI, error) {
	return t.rootConfig.GetChild(sub)
}

func (t *TConfig) SetSubConfig(sub string, child *SubConfigI) (error) {
	return t.rootConfig.SetChild(sub, child)
}

func (t *TConfig) splitName(sub string) []string {
	return strings.Split(sub, ",")
}
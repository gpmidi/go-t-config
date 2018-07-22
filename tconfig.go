package go_t_config

type TConfig struct {
	l          *TLogI
	rootConfig SubConfigI
}

func New(loggers *TLogI) (*TConfig, error) {
	t := TConfig{l: loggers}

	return &t, nil
}

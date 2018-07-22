package go_t_config

import ()

// Main config
type TConfigI interface {
	// Getters
	GetString(sub string, key string) (string, error)
	GetBool(sub string, key string) (bool, error)
	GetInt64(sub string, key string) (int64, error)
	GetFloat64(sub string, key string) (float64, error)
	// Setters
	SetString(sub string, key string, value string) (error)
	SetBool(sub string, key string, value bool) (error)
	SetInt64(sub string, key string, value int64) (error)
	SetFloat64(sub string, key string, value float64) (error)
	// Default params to use for lookups
	SetDefaultWriteSub(sub string)
	// Logging
	SetLogger(log *TLogI)
	GetLogger() *TLogI
	// Kids
	GetSubConfig(sub string) (SubConfigI, error)
	SetSubConfig(sub string, child SubConfigI) (error)
}

// Sub configs
type SubConfigI interface {
	// Getters
	GetString(key string) (string, error)
	GetBool(key string) (bool, error)
	GetInt64(key string) (int64, error)
	GetFloat64(key string) (float64, error)
	// Setters
	SetString(key string, value string) (error)
	SetBool(key string, value bool) (error)
	SetInt64(key string, value int64) (error)
	SetFloat64(key string, value float64) (error)
	// Related
	GetParent() *SubConfigI
	SetParent(parent *SubConfigI)
	GetChild(key string) *SubConfigI
	SetChild(key string, child *SubConfigI)
	// Get our relative name
	GetName() string

	// Logging
	SetLogger(log *TLogI)
	GetLogger() *TLogI
}

// Logging
type TLogI interface {
	Debug(string) error
	Info(string) error
	Warning(string) error
	Error(string) error
	Log(int, string) error
}

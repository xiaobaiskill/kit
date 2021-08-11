package password

import "sync"

var (
	passwd *passwdConfig
	once   sync.Once
)

func Get() *passwdConfig {
	once.Do(func() {
		passwd = MustLoadConfig()
	})
	return passwd
}

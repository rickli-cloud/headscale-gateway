package config

func Init() error {
	if _, err := Env.init(); err != nil {
		return err
	}

	return nil
}

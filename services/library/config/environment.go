package config

type Environment string

const (
	Production  Environment = "prod"
	Development Environment = "dev"
)

func EnvFromFlag(flag string) Environment {
	switch flag {
	case string(Production):
		return Production
	case string(Development):
		return Development
	}

	return Development
}

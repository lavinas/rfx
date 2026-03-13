package config

type Config struct {
	Validator ValidatorConfig `yaml:"validator"`
}

type ValidatorConfig struct {
	SkipHierarchyValidation []string `yaml:"skipHierarchyValidation"`
}

func (v ValidatorConfig) SkipMap() map[string]bool {

	skip := make(map[string]bool)

	for _, d := range v.SkipHierarchyValidation {
		skip[d] = true
	}

	return skip
}

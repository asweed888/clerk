package domain

type ClerkYaml struct {
    Export bool `yaml:"export"`
	Lang   string `yaml:"lang"`
	Spec   []struct {
		Location  string `yaml:"location"`
        Comment string `yaml:"comment"`
		Upstream []struct {
			Name    string   `yaml:"name"`
			Comment string   `yaml:"comment"`
            Function []string `yaml:"function"`
		} `yaml:"upstream"`
	} `yaml:"spec"`
}

type ClerkYamlRepository interface {
    Load() (*ClerkYaml, error)
}


type ClerkYamlService interface {}

type clerkYamlService struct {
    ClerkYamlRepository
}


func NewClerkYamlService(r ClerkYamlRepository) ClerkYamlService {
    return &clerkYamlService{r}
}

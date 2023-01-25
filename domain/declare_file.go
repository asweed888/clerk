package domain

type Declare struct {
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

type DeclareRepository interface {
    Load() (*Declare, error)
}


type DeclareService interface {}

type declareService struct {
    DeclareRepository
}


func NewDeclareService(r DeclareRepository) DeclareService {
    return &declareService{r}
}

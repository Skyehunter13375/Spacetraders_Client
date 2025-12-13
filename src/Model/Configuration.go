package Model

type Configs struct {
	DB struct {
		DbPath  string `yaml:"DbPath"`
		DbBuild string `yaml:"DbBuild"`
		DbReset string `yaml:"DbReset"`
	} `yaml:"DB"`

	LOG struct {
		ErrPath string `yaml:"ErrPath"`
		ActPath string `yaml:"ActPath"`
	} `yaml:"LOG"`

	API struct {
		AgentName    string `yaml:"AgentName"`
		AgentFaction string `yaml:"AgentFaction"`
		AccntToken   string `yaml:"AccntToken"`
		AgentToken   string `yaml:"AgentToken"`
	} `yaml:"API"`
}

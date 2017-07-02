package agents

//ConfigurationAgent Responsible to receive and parse configuration from master
type ConfigurationAgent struct {
	Port          int
	Configuration *map[string]interface{}
}
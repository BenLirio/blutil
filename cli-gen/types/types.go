package types

type Command struct {
	Use      string
	Short    string
	Long     string
	Commands []Command
}

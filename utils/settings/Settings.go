package settings

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type Settings struct {
	Token string `json:"token"`
	DisabledCommands []string `json:"disabled_commands"`
}

var Data = Settings{}

func New() string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Welcome to the Bot setup!\nToken: ")
	scanner.Scan()
	Data.Token = scanner.Text()
	Save()
	return scanner.Text()
}

func Load() {
	settings, err := os.ReadFile("settings.json")
	if err == nil {
		if err := json.Unmarshal(settings, &Data); err != nil {
			panic(err)
		}
	}
	if Data.Token == "" {
		New()
	}
}

func Save(){
	d, err := json.MarshalIndent(Data, "", "    ")
	if err != nil {
		fmt.Println("Warning: Error encoding json (settings will not save)\n" + err.Error())
	}
	err = os.WriteFile("settings.json", d, 0644)
	if err != nil {
		fmt.Println("Warning: Error saving settings file\n" + err.Error())
	}
}
package settings

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/prim69/wolframgo"
	"os"
)

type Settings struct {
	Token            string            `json:"token"`
	WolframAppID     string            `json:"wolfram_app_id"`
	DisabledCommands []string          `json:"disabled_commands"`
	MainXBLToken     string            `json:"main_xbl_token"`
	XBLTokens        map[string]string `json:"xbl_tokens"`
}

var Data = Settings{}

func New() string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome to the Bot setup!")
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
	if Data.WolframAppID != "" {
		wolframgo.AppID = Data.WolframAppID
	}
}

func Save() {
	d, err := json.MarshalIndent(Data, "", "    ")
	if err != nil {
		fmt.Println("Warning: Error encoding json (settings will not save)\n" + err.Error())
	}
	err = os.WriteFile("settings.json", d, 0644)
	if err != nil {
		fmt.Println("Warning: Error saving settings file\n" + err.Error())
	}
}

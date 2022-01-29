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
	XblApiToken      string            `json:"xbl_api_token"`
	DisabledCommands []string          `json:"disabled_commands"`
	MainXBLToken     string            `json:"main_xbl_token"`
	XBLTokens        map[string]string `json:"xbl_tokens"`
	SnipeLimit       uint64            `json:"snipe_limit"`
}

var Data = Settings{}

func init() {
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
	if Data.SnipeLimit == 0 {
		Data.SnipeLimit = 50
	}
}

func New() string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome to the Bot setup!")
	fmt.Print("Enter your token: ")
	scanner.Scan()
	Data.Token = scanner.Text()
	Save()
	return scanner.Text()
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

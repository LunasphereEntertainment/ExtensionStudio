package main

import (
	"github.com/LunasphereEntertainment/ExtensionStudio/model"
	"encoding/json"
	"fmt"
)

func main() {
	extension, err := LoadXML[model.ExtensionInfo]("E:\\SteamLibrary\\steamapps\\common\\Hacknet\\Extensions\\IntroExtension\\ExtensionInfo.xml")
	if err != nil {
		panic(err)
	}

	mission, err := extension.StartingMission.Load("E:\\SteamLibrary\\steamapps\\common\\Hacknet\\Extensions\\IntroExtension\\")
	if err != nil {
		panic(err)
	}

	data, err := json.MarshalIndent(extension, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Print(string(data))
	fmt.Print("\nStarting Mission: \n")
	data, err = json.MarshalIndent(mission, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Print(string(data))
}

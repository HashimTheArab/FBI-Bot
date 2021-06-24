package main

import "github.com/bwmarrin/discordgo"

var token = "ODU3MDQyMTY4MDI3OTM4ODY4.YNJ0tw.mC3aWE07jj7_gV3354qxDCkksxQ"

func newUpdateStatusData(idle int, activityType discordgo.ActivityType, name, url string) *discordgo.UpdateStatusData {
	usd := &discordgo.UpdateStatusData{
		Status: "online",
	}

	if idle > 0 {
		usd.IdleSince = &idle
	}

	if name != "" {
		usd.Activities = []*discordgo.Activity{{
			Name: name,
			Type: activityType,
			URL:  url,
		}}
	}

	return usd
}
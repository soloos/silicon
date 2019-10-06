package sidecar

import "soloos/common/iron"

type SolofsOptions struct {
	LogFilePath string
}

type SolomqOptions struct {
	LogFilePath string
}

type BadgerOptions struct {
	LogFilePath string
}

type Options struct {
	SolofsInsList []SolofsOptions
	SolomqInsList []SolofsOptions
	BadgerInsList []SolofsOptions
	WebServer     iron.Options
}

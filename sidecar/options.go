package sidecar

import "soloos/common/iron"

type SDFSOptions struct {
	LogFilePath string
}

type SWALOptions struct {
	LogFilePath string
}

type BadgerOptions struct {
	LogFilePath string
}

type Options struct {
	SDFSInsList   []SDFSOptions
	SWALInsList   []SDFSOptions
	BadgerInsList []SDFSOptions
	WebServer     iron.Options
}

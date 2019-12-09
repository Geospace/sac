package sac

type ConfigType uint64

const (
	JSON ConfigType = iota
	YAML
	UNKNOWN
)

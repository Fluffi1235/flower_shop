package models

import "time"

type Config struct {
	ConnectDb      string        `yaml:"connectDb"`
	Port           string        `yaml:"port"`
	ClientTimeout  time.Duration `yaml:"clientTimeout"`
	ContextTimeout time.Duration `yaml:"contextTimeout"`
}

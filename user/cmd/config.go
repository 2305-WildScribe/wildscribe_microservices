package main

type ServiceConfig struct {
	APIConfig apiConfig `yaml:"api"`
}
type apiConfig struct {
	Port string `yaml:"port"`
}

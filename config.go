package main

import (
	"os"
	"flag"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"log"
)

type config struct {
	port string
	git struct {
		url				string
		username		string
		password		string
		clonepath		string
	}
}


func setup() {

}

func flags() config {
	var flags config
	flags.port = *flag.String("port", "8888", "server port")
	flags.git.url = *flag.String("git-url", "", "git repository url")
	flags.git.username = *flag.String("git-username", "", "git repository username")
	flags.git.password = *flag.String("git-password", "", "git repository password")
	flags.git.clonepath = *flag.String("git-clonepath", ".filesystem-local","git clone directory path")

	return flags
}

func envvars() config {
	var envvars config
	envvars.port = os.Getenv("MYAC_PORT")
	envvars.git.url = os.Getenv("MYAC_GIT_URL")
	envvars.git.username = os.Getenv("MYAC_GIT_USERNAME")
	envvars.git.password = os.Getenv("MYAC_GIT_PASSWORD")
	envvars.git.clonepath = os.Getenv("MYAC_GIT_CLONEPATH")

	return envvars
}

func (c *config) yamlconfig() config {
	configpath := flag.String("config", "config.yml", "path to service config")
	flag.Parse()

	yamlFile, err := ioutil.ReadFile(*configpath)
	if err != nil {
		log.Printf("yamlFile.Get err #%v ", err)
	}

	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return *c
}

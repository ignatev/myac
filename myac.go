package main

import(
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"fmt"
)

type ServerConf struct {
	Server struct {
		Git struct {
			URI      string `json:"uri"`
			Username string `json:"username"`
			Password string `json:"password"`
		} `json:"git"`
	} `json:"server"`
}

func (c *ServerConf) getConf() *ServerConf {
	yamlFile, err := ioutil.ReadFile("config.yml")

	if err != nil {
		log.Printf("yamlFile.Get err #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}

func main() {
	var c ServerConf
	fmt.Println(c)
	c.getConf()
	fmt.Println(c)
}

//todo read configfile from server using git settings
//todo display configfile via http
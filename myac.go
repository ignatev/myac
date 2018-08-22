package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/yaml.v2"
)

type serverConf struct {
	Server struct {
		Port int `json:"port"`
		Git  struct {
			URL                 string `json:"url"`
			Username            string `json:"username"`
			Password            string `json:"password"`
			LocalRepositoryPath string //todo investigate why yaml.v2 can not parse dashed props from yaml config: local-repository-path is nil
		} `json:"git"`
	} `json:"server"`
}

type configHandler struct {
	configs *map[string]string
}

func (ch *configHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p := strings.TrimLeft(r.URL.Path, "/")
	log.Println(p)
	log.Println(r.RemoteAddr)
	c := ch.configs
	if val, ok := (*c)[p]; ok {
		serveConfigFile(w, r, val)
	}
}

func serveConfigFile(w http.ResponseWriter, r *http.Request, p string) {
	http.ServeFile(w, r, p)
}

func runServer(port string, configs *map[string]string) {
	err := http.ListenAndServe(port, &configHandler{configs})
	log.Println("Listening...")
	if err != nil {
		log.Println(err)
	}
}

func (c *serverConf) getConf(configPath string) *serverConf {
	yamlFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Printf("yamlFile.Get err #%v ", err)
	}

	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}

func cloneConfigRepo(repopath, url string) {
	_, err := git.PlainClone(repopath, false, &git.CloneOptions{
		URL:      url,
		Progress: os.Stdout,
	})
	if err != nil {
		log.Println(err)
	}
}

func main() {
	configPath := flag.String("config", "config.yml", "path to service config")
	flag.Parse()
	var c serverConf
	config := c.getConf(*configPath)
	url := config.Server.Git.URL
	repopath := config.Server.Git.LocalRepositoryPath
	port := ":" + strconv.Itoa(config.Server.Port)

	cloneConfigRepo(repopath, url)
	printServerStatus(port)
	tree := treebuilder(repopath)
	runServer(port, tree.finalmapping)
}

func printServerStatus(port string) {
	fmt.Println(`
 ._ _        _.   _
 | | |  \/  (_|  (_
        /`)
	fmt.Println("Configuration server")
	fmt.Println("")
	fmt.Println("Service running on port", port[1:])
}

//todo add method for mapping service:[multiple config files, e.g. dev, prod, test]
//todo organize logs
//todo add Dockerfile
//todo add tests
//todo add func for parsing absolute and relative repo paths
//todo add logging for clients (client connected from <address>, client request: URL, file)

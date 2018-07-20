package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"fmt"
	"gopkg.in/src-d/go-git.v4"
	"os"
	"path/filepath"
	"strings"
	"net/http"
	"flag"
//	"time"
)

type serverConf struct {
	Server struct {
		Port int `json:"port"`
		Git struct {
			URL                 string `json:"url"`
			Username            string `json:"username"`
			Password            string `json:"password"`
			LocalRepositoryPath string  //todo investigate why yaml.v2 can not parse dashed props from yaml config: local-repository-path is nil
		} `json:"git"`
	} `json:"server"`
}

type configHandler struct {
	configs map[string][]string
}

func (ch *configHandler) ServeHTTP(w http.ResponseWriter, r * http.Request) {
	p := strings.TrimLeft(r.URL.Path, "/")
	log.Println(p)
	c := ch.configs
	log.Println(c)
	if val, ok := c[p]; ok {
		log.Println("serving file %s", val[0])
		serveConfigFile(w, r, val[0])
	}
}
	

func serveConfigFile(w http.ResponseWriter, r *http.Request, p string) {
	http.ServeFile(w, r, p)
}


func runServer(port int, configs map[string][]string) {
	err := http.ListenAndServe(":9091", &configHandler{configs})
	log.Println("Listening...")
	if err != nil {
		log.Fatal(err)
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

func main() {
	configPath := flag.String("config", "config.yml", "path to service config")
	flag.Parse()
	var c serverConf
	var config = c.getConf(*configPath)
	url := config.Server.Git.URL
	localRepositoryPath := config.Server.Git.LocalRepositoryPath
	log.Println(config)
	log.Println("repo path:", localRepositoryPath)
	port := config.Server.Port

	_, err := git.PlainClone(localRepositoryPath, false, &git.CloneOptions {
		URL:      url,
		Progress: os.Stdout,
	})
	if err != nil {
		log.Println(err)
	}
	repo, err := listRepo(localRepositoryPath)
	if err != nil {
		log.Println(err)
	}
	
	runServer(port, createSliceWithPaths(repo))
}

func listRepo(root string) ([]string, error) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && !strings.HasPrefix(path, root + "/.git") && path != root {
			fmt.Println(path)
			files = append(files, path)
		}
		return nil
	})
	return files, err
}

func createSliceWithPaths(paths []string) map[string][]string {
	m := make(map[string][]string)
	for _, p := range paths {
		segs := strings.Split(p, "/")
		log.Println(segs)
		k := segs[len(segs) - 2]
		m[k] = append(m[k], p)
	}
	log.Println(m)
	return m
}


//todo create function for derivation web-server paths from git-repo structure, generate tree and run service using runServer() example
//todo add Dockerfile
//todo add tests

package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"fmt"
	"gopkg.in/src-d/go-git.v4"
	"os"
	"github.com/kataras/iris"
	"path/filepath"
	"strings"
	"strconv"
)

type ServerConf struct {
	Server struct {
		Port int `json:"port"`
		Git struct {
			URI                 string `json:"uri"`
			Username            string `json:"username"`
			Password            string `json:"password"`
			LocalRepositoryPath string `json:"local-repository-path"`
		} `json:"git"`
	} `json:"server"`
}

func runServer(port int, repo []string) {
	app := iris.New()
	app.StaticWeb("/service-1", "./.filesystem-repo/service-1/generic-service.yml")
	app.Run(iris.Addr(":" + strconv.Itoa(port)))
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
	var config = c.getConf()
	var url string = config.Server.Git.URI
	var localRepositoryPath string = config.Server.Git.LocalRepositoryPath
	var port int = config.Server.Port

	_, err1 := git.PlainClone(localRepositoryPath, false, &git.CloneOptions{
		URL:      url,
		Progress: os.Stdout,
	})
	if err1 != nil {
		log.Println(err1)
	}
	repo, err := listRepo(".filesystem-repo")
	if err != nil {
		log.Println(err)
	}
	runServer(port, repo)
}

func listRepo(root string) ([]string, error) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() && !strings.HasPrefix(path, ".filesystem-repo/.git") {
			fmt.Println(path)
			serviceURI := strings.TrimPrefix(path, ".filesystem-repo/")
			files = append(files, serviceURI)
		}
		return nil
	})
	return files, err
}

//todo create function for derivation web-server paths from git-repo structure, generate tree and run service using runServer() example
//todo add Dockerfile
//todo add tests

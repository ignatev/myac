package main

import(
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"fmt"
	"gopkg.in/src-d/go-git.v4"
	"os"
	"github.com/kataras/iris"
	"path/filepath"
	"strings"
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

func runServer() {
	app := iris.New()
	app.StaticWeb("/service-1", "./.filesystem-repo/service-1/generic-service.yml")
	app.Run(iris.Addr(":8081"))
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
	var url string = c.getConf().Server.Git.URI
	fmt.Println(url)


	_, err1 := git.PlainClone(".filesystem-repo", false, &git.CloneOptions{
		URL: url,
		Progress: os.Stdout,
	})
	if err1 != nil {
		log.Println(err1)
	}
	repo, err := listRepo(".filesystem-repo")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(repo)
	runServer()
}

func listRepo(root string) ([]string, error) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() && !strings.HasPrefix(path, ".filesystem-repo/.git"){
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
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


func runServer(port int, repo []string) {



//	http.Handle("/.filesystem-repo/", http.StripPrefix("/.filesystem-repo/", http.FileServer(http.Dir(".filesystem-repo"))))
//	http.HandleFunc("/service-1", serveConfigFile)

	log.Println("Listening...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}


}

func myTestHandler(w http.ResponseWriter, r *http.Request) {

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
	createSliceWithPaths(repo)
	runServer(port, repo)
}

func listRepo(root string) ([]string, error) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && !strings.HasPrefix(path, root + "/.git") && path != root {
			fmt.Println(path)
//			serviceURI := strings.TrimPrefix(path, root)
			files = append(files, path)
		}
		return nil
	})
	return files, err
}

func serveConfigFile(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, ".filesystem-repo/service-1/generic-service.yml")
}

type configurationToServe struct {
	FilePath	string
	FileName	string
	URL			string
}

func createSliceWithPaths(paths []string) []configurationToServe {
	var configs []configurationToServe
	for _, p := range paths {
		segs := strings.Split(p, "/")
		log.Println(segs)
		c := configurationToServe{p, segs[len(segs) -1 ], segs[len(segs) - 2]}
		configs = append(configs, c)
	}
	log.Println(configs)
	return configs
}


//todo create function for derivation web-server paths from git-repo structure, generate tree and run service using runServer() example
//todo add Dockerfile
//todo add tests

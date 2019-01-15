package main

import (
	"flag"
	"github.com/plouc/go-gitlab-client/gitlab"
	"fmt"
)

var msg string

func init()  {
	flag.StringVar(&msg,"msg","hello world","message to display")
	flag.Parse()
}

func main() {
	var gbClient *gitlab.Gitlab
	gbClient = gitlab.NewGitlab("http://192.168.1.249", "/api/v3/", "HsnYW8E_kusztj2Kzh8d")


	pc, response, err := gbClient.Projects(nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	} else if response.StatusCode != 200 {
		fmt.Printf("response error %6d \n",response.StatusCode)
		return
	}

	for _, project := range pc.Items {
		fmt.Printf("> %6d | %s\n", project.Id, project.Name)
	}

}
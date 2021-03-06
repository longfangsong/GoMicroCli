package main

import (
	"flag"
	"github.com/iancoleman/strcase"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"text/template"
)

func getProjectName() string {
	s, _ := os.Getwd()
	i := strings.LastIndex(s, "/")
	path := string(s[i+1:])
	return path
}

func downloadPrepared(filePath string) {
	pwd, _ := os.Getwd()
	fileName := pwd + "/" + filePath
	_, err := os.Stat(fileName)
	if err != nil {
		url := "https://raw.githubusercontent.com/longfangsong/GoMicroCli/master/template/" + filePath + ".template"
		response, _ := http.Get(url)
		body, _ := ioutil.ReadAll(response.Body)
		file, _ := os.Create(fileName)
		_, _ = file.Write(body)
	}
}

func downloadAndRender(filePath string) {
	pwd, _ := os.Getwd()
	fileName := pwd + "/" + filePath
	_, err := os.Stat(fileName)
	if err != nil {
		url := "https://raw.githubusercontent.com/longfangsong/GoMicroCli/master/template/" + filePath + ".template"
		response, _ := http.Get(url)
		body, _ := ioutil.ReadAll(response.Body)
		file, _ := os.Create(fileName)
		tmpl, _ := template.New("template").Parse(string(body))
		_ = tmpl.Execute(file, struct {
			PROJECT_NAME            string
			PROJECT_NAME_LOWER_CASE string
		}{
			getProjectName(),
			strcase.ToKebab(getProjectName()),
		})
	}
}

func downloadAndRenderModel(filePath string, modelName string) {
	pwd, _ := os.Getwd()
	fileName := pwd + "/" + filePath
	_, err := os.Stat(fileName)
	if err != nil {
		url := "https://raw.githubusercontent.com/longfangsong/GoMicroCli/master/template/model/model.go.template"
		response, _ := http.Get(url)
		body, _ := ioutil.ReadAll(response.Body)
		file, _ := os.Create(fileName)
		tmpl, _ := template.New("template").Parse(string(body))
		_ = tmpl.Execute(file, struct {
			PROJECT_NAME            string
			PROJECT_NAME_LOWER_CASE string
			MODEL_NAME              string
		}{
			getProjectName(),
			strcase.ToKebab(getProjectName()),
			modelName,
		})
	}
}

func main() {
	createWhat := flag.String("create", "", "Creat what?")
	flag.Parse()
	detail := flag.Args()
	switch *createWhat {
	case "basic":
		pwd, _ := os.Getwd()
		_ = os.MkdirAll(pwd+"/script", 0777)
		_ = os.MkdirAll(pwd+"/.github", 0777)
		_ = os.MkdirAll(pwd+"/.github/workflows", 0777)
		downloadPrepared("script/dev-server.sh")
		downloadAndRender(".github/workflows/dockerimage.yml")
		downloadAndRender("Dockerfile")
		downloadAndRender("main.go")
	case "handler":
		pwd, _ := os.Getwd()
		_ = os.MkdirAll(pwd+"/handler", 0777)
		downloadPrepared("handler/ping.go")
		downloadPrepared("handler/handler_test.go")
	case "service":
		pwd, _ := os.Getwd()
		_ = os.MkdirAll(pwd+"/service", 0777)
		switch detail[0] {
		case "token":
			_ = os.MkdirAll(pwd+"/service/token", 0777)
			downloadPrepared("service/token/token.go")
		}
	case "infr":
		pwd, _ := os.Getwd()
		_ = os.MkdirAll(pwd+"/infrastructure", 0777)
		if detail[0] == "postgres" {
			downloadPrepared("infrastructure/db.go")
		} else if detail[0] == "redis" {
			downloadPrepared("infrastructure/redis.go")
		}
	case "model":
		pwd, _ := os.Getwd()
		_ = os.MkdirAll(pwd+"/model", 0777)
		modelName := detail[0]
		downloadAndRenderModel("model/"+strings.ToLower(modelName)+".go", modelName)
	case "tools":
		pwd, _ := os.Getwd()
		_ = os.MkdirAll(pwd+"/tools", 0777)
		if detail[0] == "checkErr" {
			downloadPrepared("tools/checkErr.go")
		} else if detail[0] == "optionalString" {
			downloadPrepared("tools/optionalString.go")
		}
	}
}

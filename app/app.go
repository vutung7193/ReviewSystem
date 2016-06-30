package app

import (
  "sync"
  "path/filepath"
  "io/ioutil"
  "gopkg.in/yaml.v2"
)
type Database struct {
  Adapter       string `yaml:"adapter"`
  Host          string `yaml:"host"`
  Username      string `yaml:"username"`
  Password      string `yaml:"password"`
  Dbname        string `yaml:"dbname"`
  Encoding      string `yaml:"encoding"`
  Port          string `yaml:"port"`
}
type Config struct {
	Database Database `yaml:"database"`
}
type App struct {
	Config Config
}

var app *App
var once sync.Once

func GetInstance() *App {
  once.Do(func() {
    app = &App{}
    
    filename, _ := filepath.Abs("./config/database.yml")
    yamlFile, err := ioutil.ReadFile(filename)

    if err != nil {
        panic(err)
    }

    err = yaml.Unmarshal(yamlFile, &app.Config)
    if err != nil {
        panic(err)
    }
  })
  return app
}
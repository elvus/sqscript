package main

import (
	"encoding/json"
    "os"
	"os/exec"
	"io/ioutil"
	"runtime"
	"strings"
	"fmt"
)

type Properties struct{
	Url string  `json:"url"`
	Key string `json:"key"`
}

func main(){
	sonarpath, err := os.Executable()
	sonarpath = sonarpath[0:strings.LastIndex(sonarpath, "/")]
	fmt.Println(sonarpath+"properties.json")
	file, err := ioutil.ReadFile(sonarpath+"/properties.json")
	if os.IsNotExist(err) {
		fmt.Println("No se encontraron las configuraciones correspondientes")
		makeProperties()
	} else {
		props := &Properties{}
		json.Unmarshal([]byte(file), &props)
		if(len(os.Args)>1){
			runSonar(props, os.Args[1], os.Args[2])
		}else{
			ex, err := os.Getwd()
			if(err != nil){
				fmt.Println(err)
			}
			fmt.Print(ex[strings.LastIndex(ex, "/")+1:])
			runSonar(props, ex, ex[strings.LastIndex(ex, "/")+1:])
		}
	}
}

func makeProperties(){
	//Se ingresa la url donde esta servicio del sonar
    fmt.Print("Ingrese la url del sonar: ")
    var url string
	//ingresamos nuestra ssh key generada desde sonar
    fmt.Scanln(&url)
    fmt.Print("Ingrese su clave: ")
    var key string
    fmt.Scanln(&key)
	//Preparamos la estructura del archivo properties
	props := Properties{
		Url: url,
		Key: key,		
	}
	//Se crea el archivo properties.json
	jsonB, err := json.MarshalIndent(props, "", " ")
	err = ioutil.WriteFile("properties.json", jsonB, 0644)

	if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}

func runSonar(props *Properties, path string, name string){
	osCommand := map[string]string{"darwin": "/sonar-scanner", "linux": "/sonar-scanner", "windows": "/sonar-scanner.bat"}
	sonarpath, err := os.Executable()
	sonarpath = sonarpath[0:strings.LastIndex(sonarpath, "/")]
	command := "sudo "+sonarpath + osCommand[runtime.GOOS]+" -X -D sonar.login=\""+props.Key+"\" -D sonar.host.url="+props.Url+" -D sonar.projectKey="+name+" -D sonar.sources=apiproxy -D sonar.projectBaseDir="+path
	cmd := exec.Command("bash","-c", command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
    if err != nil {
        fmt.Println("Error: "+err.Error())
        return
    }
}
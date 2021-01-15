package shared

import (
	"fmt"
	"github.com/DanielHauge/goSpace/space"
)

const Manager = "manager"
const Coordinator = "coordinator"
const Worker = "worker"
const TimeFormat = "2006-01-02 15:04:05"



var (
	LogSpace space.Space
	IndexSpace space.Space
	StatusSpace space.Space
	TaskSpace space.Space
	ClientName string
	ProgrammingLanguages map[string]string
)

func InitUtility(){
	ProgrammingLanguages = map[string]string{
		".go" : "Go",
		".cs" : "C#",
		".java": "Java",
		".c" : "C",
		".cpp" : "C++",
		".h" : "C/C++",
		".js" : "Javascript",
		".ts" : "Typescript",
		".scala" : "Scala",
		".rb" : "Ruby",
		".sh" : "Shell",
		".md" : "Markdown",
		".rs" : "Rust",
		".py" : "Python",
		".php" : "PHP",
		".kt" : "Kotlin",
		".dart" : "Dart",
	}
}

func withDefault(s string, def string) string{
	if len(s) == 0 { return def }
	return s
}

func Log(message string){
	log := fmt.Sprintf("%v - %v\n", ClientName, message)
	fmt.Println(log)
	LogSpace.Put(log)
}

func CheckError(err error) bool{
	if err != nil {
		Log(fmt.Sprintf("Failed with: %v", err))
		return true
	}
	return false
}

func CheckErrorF(err error){
	if err != nil{
		Log(fmt.Sprintf("Failed with: %v", err))
		panic(err)
	}
}

func LogPanicError(){
	if r := recover(); r!= nil{
		fmt.Println(r)
		Log(fmt.Sprintf("Panic with error: %v", r))
	}
}



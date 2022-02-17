package _1

// Use map[string]interface{} to pair functions to name
// Could maybe use anonymous functions instead. Might be clean
// in certain cases
//var SolutionMap = map[string]interface{}{
//	"spencerreeves": spencerreeves,
//	"name":          name,
//}

func main() {
	callDynamically("hello")
	callDynamically("name", "Joe")
}

func callDynamically(name string, args ...interface{}) {
	//switch name {
	//case "hello":
	//	SolutionMap["hello"].(func())()
	//case "name":
	//	SolutionMap["name"].(func(string))(args[0].(string))
	//}

}

func Run() {

}

package main

func main() {
    action := Action{
        Human: Human{Name: "Alex", Age: 25},
        Role:  "Developer",
    }

    action.Greet()
    action.ShowRole()   
    action.Birthday()
    action.Greet()
}
package main

import ("fmt";"net/http";"html/template")
// struct это класс
type User struct {
  Name string
  Age uint16 //целое число не может быть отрицательным
  Money int16 //может быть отрицательным
  Avg_grades,Happiness float64 //Может быть с точкой
  Hobbies []string//список
}

//метод
//func (u *User) getAllInfo() string{
//  return fmt.Sprintf("User name is: %s." +
//    "He is %d and he has money equal: %d", u.Name, u.Age, u.Money)
//}
// * это ссылка на объект
//func (u *User) setNewName(newName string){
//  u.Name = newName
//}

func home_page(page http.ResponseWriter, r *http.Request){
  jake := User{"Jake",20,-50,3.8,0.4, []string{"Football","Skate","Dance"}}
  //jake.name="Alex"
  //jake.setNewName("Alex")
  //fmt.Fprintf(page, jake.getAllInfo())
  //fmt.Fprintf(page, `<h1>Main Text</h1>
  //  <b>Main Text</b>`) //вывод небольшого кусочка html
  tmpl, _:=template.ParseFiles("templates/home_page.html") //подгружаем html фронтэнд файл
  tmpl.Execute(page, jake)
}

func contacts_page(page http.ResponseWriter, r *http.Request){
  fmt.Fprintf(page, "dodelat")
}

func twitter_page(page http.ResponseWriter, r *http.Request){
  fmt.Fprintf(page, "https://vk.com/syppix")
}

func handleRequest(){
  //fmt.Println("Hello world")
  http.HandleFunc("/", home_page)//home_page это название метода
  http.HandleFunc("/firstpage/", contacts_page) //contacts просто название, может быть любым
  http.HandleFunc("/secondpage/", twitter_page)
  http.ListenAndServe(":8080", nil)//Порт, nil это null или none(настройки для самого сервера)
}

func main() {
  //var jake User=... создания объекта jake
  //jake := User{name:"Jake", age: 20, money: -50, avg_grades: 3.8, happiness: 0.4}
  handleRequest()
}

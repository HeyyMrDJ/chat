package main

import (
	"html/template"
	"net/http"
	"slices"
)

var chats = []string{}
func main() {
    http.HandleFunc("/", serveHome)
    http.HandleFunc("/chat", createChat)
    http.HandleFunc("/getchats", getChats)

    http.ListenAndServe(":8000", nil)
}

func addChat(newMessage string) {
    chats = append(chats, newMessage)
}

func serveHome(w http.ResponseWriter, r *http.Request) {
    tmpl := template.Must(template.ParseFiles("base.html", "content.html"))
    tmpl.Execute(w, "")
}

func createChat(w http.ResponseWriter, r *http.Request) {
    // Parse form data
    err := r.ParseForm()
    if err != nil {
        http.Error(w, "Failed to parse form", http.StatusBadRequest)
        return
    }

    // Get chat message from form
    chatMessage := r.FormValue("chat")

    // Here you would add the chat message to your chat system
    addChat(chatMessage) // Placeholder for your addChat function
}

func getChats(w http.ResponseWriter, r *http.Request) {
    tmpl := template.Must(template.ParseFiles("chats.html"))
    chats2 := slices.Clone(chats)
    slices.Reverse(chats2)
    tmpl.Execute(w, chats2)
}

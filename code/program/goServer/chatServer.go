package main

import (
   "net/http"
   "fmt"
   "time"
   "encoding/json"
)

var serverData ServerData
var onlineTimer int = 5


type ServerData struct {
   AllUsers []User `json:"allUsers"`
   ChatRooms []Room `json:"chatRooms"`
}

type User struct {
   Name string `json:"name"`
   CurrRoom *Room `json:"currRoom"`
   OnlineTimer int `json:"onlineTimer"`
}

type Room struct {
   Name string `json:"name"`
   Chat []Comment `json:"chat"`
}

type Comment struct {
   Name string `json:"person"`
   Text string `json:"text"`
}


/*
This function is run concurrently with all other server
operations. Every second, the list of all the users held in
serverData is gone through and their OnlineTimer is decremented
by 1.  Once a user's OnlineTimer is at 0, they are taken off
from the list of users: they are no longer online. This OnlineTimer
is reset to 2 everytime a user makes a get request to the server.
*/
func startClock(){
   for true {
      time.Sleep(1000 * time.Millisecond)
      var newUsers []User
      for _, usr := range serverData.AllUsers{
         usr.OnlineTimer -= 1
         if usr.OnlineTimer > 0{
            newUsers = append(newUsers, usr)
         }
      }
      serverData.AllUsers = newUsers
      fmt.Println(serverData.AllUsers)
      fmt.Println("================================")
      fmt.Println(serverData.ChatRooms)
      fmt.Println("================================")
   }
}

/*
Parameter: A comment from the Comment struct that contains as
strings the name of the person who made the comment, and the comment
itself

This function searches through all the users online until it finds
the name of the person who made the comment, at which point the comment
is posted in that users chat room.
*/
func postComment(comment Comment){
   for _, usr := range serverData.AllUsers{
      if (usr.Name == comment.Name){
         (*(usr.CurrRoom)).Chat = append((*(usr.CurrRoom)).Chat, comment)
      }
   }
}

/*
This function is the entry point for the Go Chat Room server.

When this function is called, the following occurs:
   #
      A clock is started to keep mark anyone inactive as no longer online
   #
      http.HandleFunc is invoked to call requestHandler() anytime a request
      is made to the server
   #
      The server is initialized to run on port 3000

All together, this method initializes the server.
*/
func main() {
   serverData = ServerData{}
   serverData.ChatRooms = append(serverData.ChatRooms, Room{})
   serverData.ChatRooms[0].Name = "general"

   //starts a thread to decrement the users online timers every second
   go startClock()

   http.HandleFunc("/" , requestHandler)

   fmt.Println(http.ListenAndServe(":3000", nil));
}

/*
This method is used for three main purposes:
   #
      To send the html page when a user first connects (done by the else statement)
   #
      To retrieve data from the server about comments and whose online
   #
      To register a new user as being online and to post a comment

Whenever a user makes a request to the server this method is called
*/
func requestHandler(w http.ResponseWriter, r *http.Request){
   fmt.Println(len(serverData.AllUsers))
   if r.FormValue("mode") == "getdata"{
      //send back JSON containing all ServerData
      for i := 0; i < len(serverData.AllUsers); i++{
         if (serverData.AllUsers[i].Name == r.FormValue("name")){
            serverData.AllUsers[i].OnlineTimer = onlineTimer;
         }
      }
      serverDataJSON, err := json.Marshal(serverData)
      if (err != nil){
         fmt.Println("JSON conversion failed")
      }
      w.Header().Set("Content-Type", "application/json")
      w.Write(serverDataJSON)
   }else if r.FormValue("mode") == "postcomment"{
      //put a comment in whatever chat room the user who posted it is in
      comment := Comment{r.FormValue("name"), r.FormValue("comment")}
      postComment(comment)
   }else if r.FormValue("mode") == "adduser"{
      //add parameterized user to serverData
      serverData.AllUsers = append(serverData.AllUsers, User{r.FormValue("name"), &serverData.ChatRooms[0], onlineTimer})
   }else {
      //send the html page file
      http.ServeFile(w, r, "../webpageUtility/entryPage.html")
   }
}
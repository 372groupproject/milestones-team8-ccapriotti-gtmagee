/*
Authors: Christian Capriotti
         Gavin Magee

Running this file using 'go run chatServer.go' will
start the server that can be connected to by going
to 'localhost:3000' on your computer (in your web browser).

Dependencies: entryPage.html
*/

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

type Validator struct {
   IsValid bool `json:"isValid"`
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
This function is designed to add a new chat room
to the list of chat rooms in serverData.  The name of the added
chat room is the parameter roomName

If the chatroom name already exists, it is not created.

Parameter: roomName is a string representing the room to be deleted
*/
func createChatRoom(roomName string){
   alreadyExists := false
   for i := 0; i < len(serverData.ChatRooms) && !alreadyExists; i++ {
      //check to see if the chat room already exists
      if roomName == serverData.ChatRooms[i].Name {
         alreadyExists = true
      }
   }
   if !alreadyExists {
      //create a new chat room if one doesnt already exist
      serverData.ChatRooms = append(serverData.ChatRooms,
       Room{roomName, []Comment{Comment{"server", "Welcome to " + roomName + " Chat!"}}})

      for i := 0; i < len(serverData.AllUsers); i++ {
         for j := 0; j < len(serverData.ChatRooms); j++{
            if ((*(serverData.AllUsers[i].CurrRoom)).Name == serverData.ChatRooms[j].Name){
               //update the current chat rooms of the users
               serverData.AllUsers[i].CurrRoom = &(serverData.ChatRooms[j])
            }
         }
      }
   }
}

/*
This function deletes the chat room with the given name
from the list of chatRooms under serverData.

If the room to be deleted doesn't exist, nothing happens

Any users in the deleted room are moved to general, which
can never be deleted.

Parameter: roomName is a string representing the chat room to be deleted
*/
func removeChatRoom(roomName string){
   if (roomName != "general"){
      var newRooms []Room
      for i := 0; i < len(serverData.ChatRooms); i++ {
         if serverData.ChatRooms[i].Name != roomName {
            newRooms = append(newRooms, serverData.ChatRooms[i])
         }else{
            for j := 0; j < len(serverData.AllUsers); j++{
               if roomName == (*(serverData.AllUsers[j].CurrRoom)).Name {
                  //change chat room to general for every user
                  //in deleted channel
                  serverData.AllUsers[j].CurrRoom = &(newRooms[0])
               }
            }
         }
      }
      serverData.ChatRooms = newRooms
      for i := 0; i < len(serverData.AllUsers); i++ {
         for j := 0; j < len(serverData.ChatRooms); j++{
            if ((*(serverData.AllUsers[i].CurrRoom)).Name == serverData.ChatRooms[j].Name){
               //update the current rooms of the users
               serverData.AllUsers[i].CurrRoom = &(serverData.ChatRooms[j])
            }
         }
      }
   }
}

/*
This function returns whether or not the suggested username
is valid

A username is valid if it does not already belong to another user.

Parameter: possName is a string of the potential username someone connects with

Return: a boolean on whether or not the username is valid (not already existing)
*/
func isValid(possName string) bool {
   for i := 0; i < len(serverData.AllUsers); i++{
      if (possName == serverData.AllUsers[i].Name){
         return false
      }
   }
   return true
}

/*
This method resets the OnlineTimer of the user
with the same name as the passed in string.

Parameter: accessor is a string of the name of the user
who just made a request to the server. This means that
they are online and their online timer should be reset
to 5.
*/
func updateUser(accessor string){
   for i := 0; i < len(serverData.AllUsers); i++{
      if serverData.AllUsers[i].Name == accessor{
         serverData.AllUsers[i].OnlineTimer = onlineTimer
      }
   }
}

/*
This function moves the user from the room they are currently
in to the rroom specified by the 'room' parameter. This change
is seen in serverData

Parameters: 
   #
      room is a string representation of the chat room the
      user is moving into
   #
      name is a string of the user who is moving rooms


*/
func moveRooms(room string, name string){
   var newRoom *Room
   for i := 0; i < len(serverData.ChatRooms); i++{
      if (serverData.ChatRooms[i].Name == room){
         newRoom = &(serverData.ChatRooms[i])
      }
   }
   for i := 0; i < len(serverData.AllUsers); i++{
      if (serverData.AllUsers[i].Name == name){
         serverData.AllUsers[i].CurrRoom = newRoom
         return
      }
   }
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
   if r.FormValue("mode") == "getdata"{
      //send back JSON containing all ServerData
      updateUser(r.FormValue("name"))
      serverDataJSON, _ := json.Marshal(serverData)
      w.Header().Set("Content-Type", "application/json")
      w.Write(serverDataJSON)
   }else if r.FormValue("mode") == "validate"{
      possName := r.FormValue("name")
      valid := isValid(possName)
      validJSON, _ := json.Marshal(Validator{valid})
      w.Header().Set("Content-Type", "application/json")
      w.Write(validJSON)
   }else if r.FormValue("mode") == "postcomment"{
      //put a comment in whatever chat room the user who posted it is in
      comment := Comment{r.FormValue("name"), r.FormValue("comment")}
      postComment(comment)
   }else if r.FormValue("mode") == "adduser"{
      //add parameterized user to serverData
      serverData.AllUsers = append(serverData.AllUsers, User{r.FormValue("name"),
       &(serverData.ChatRooms[0]), onlineTimer})
   }else if r.FormValue("mode") == "create"{
      //create a new room
      roomName := r.FormValue("name")
      createChatRoom(roomName)
   }else if r.FormValue("mode") == "remove"{
      //remove an existing room that is not general
      roomName := r.FormValue("name")
      removeChatRoom(roomName)
   }else if r.FormValue("mode") == "changeroom"{
      moveRooms(r.FormValue("room"), r.FormValue("name"))
   }else {
      //send the html page file
      http.ServeFile(w, r, "../webpageUtility/entryPage.html")
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
   serverData.ChatRooms[0].Chat = []Comment{Comment{"server", "Welcome to general Chat!"}}

   //starts a thread to decrement the users online timers every second
   go startClock()
   http.HandleFunc("/" , requestHandler)
   fmt.Println(http.ListenAndServe(":3000", nil));
}
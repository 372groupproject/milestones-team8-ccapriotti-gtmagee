(function() {
	"use strict";

	let username;
	let currRoom;

	function enterChatRooms(){
		document.getElementById("signin").style.display = "none";
		document.getElementById("chatmain").style.display = "block";
	}

	window.onload = function(){
		document.getElementById("submitname").onclick = enterChatRooms;
		document.getElementById("nameinput").onkeypress = function(event){
			if (event.keyCode == 13){
				enterChatRooms();
			}
		}
	};
})();
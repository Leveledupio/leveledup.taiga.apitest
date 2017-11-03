package taiga

import "net/http"

var client = &http.Client{}

//public token
//var token = "eyJ1c2VyX2F1dGhlbnRpY2F0aW9uX2lkIjoyNjAxODd9:1e9fUV:y2RMjwvBSDn6eKu42sDony85Wyc"

//local token
var token = "eyJ1c2VyX2F1dGhlbnRpY2F0aW9uX2lkIjoxNn0:1eAPdy:z-tzgGUMoeLn48xLHj-rRzHZaaY"
var baseURL = "http://127.0.0.1:8000"
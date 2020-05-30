# recraft-server
This repository is one of the part of the recraft project, it contains the code of the server.

**Note: this project is still in development, it may not work properly**

# Building server
First, clone this repository in your GOPATH, then you can start an example server:


  
    package main

    import server "github.com/recraft/recraft-server"
    
    func  main() {
    //Start server with debug logging
    err := server.NewServer("localhost", 25565, 2)
    if err != nil {
    panic(err) 
    }
    }

# Project status
At the moment only handshake and status states are supported.


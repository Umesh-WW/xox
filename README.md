# xox

# TDD
    - Write the doc
    - Write test 
    - Develop fuctionality
    - build API

# Release

# data strucure
1. user
    -> player - x 
    -> o -> gpt,gemine, other
2. game
    -> board
        [
            [],
            [].
            []
        ]
    -> move
    -> state : win,lose,play,err 

## version 1 - one player against Random/AI
-  api 
    1. v1/start -ai type -> give game id, 
    
    2. v1/move - state,update board

## improvement node
It is a good practice to start reading and understanding the code base of large open source projects.

For example:

- https://github.com/pocketbase/pocketbase

- https://github.com/memphisdev/memphis

- https://github.com/argoproj/argo-cd

- https://github.com/mattermost/mattermost-server/tree/master/server

Try to answer the questions during the study:

How HTTP requests are processed (where is handlers ?)

What does request validation look like

Where is the business logic of the application

How is the interaction with the database looks like

How the application is monitored: tracking, logging, sensors and metrics

Thanks to this, we will understand how bad the reality of programming in Go is.You can spend a lot of time studying about clean architecture, DDD, hexagonal architecture, golang project structure best practices. But the "reality" and how the projects actually look is priceless.
# Golang CRUD Database

![Project Image](https://i.guim.co.uk/img/media/027f99e50af24acf75d5f1a881c9db8813b1c990/0_75_4928_2957/master/4928.jpg?width=1200&quality=85&auto=format&fit=max&s=851cf9952544cbbf8472dfc1b5305c5f)

> This is a simple Golang Database that performs crud operations on a postgres Databse.

---

### Table of Contents
You're sections headers will be used to reference location of destination.

- [Description](#description)
- [How To Use](#how-to-use)
- [References](#references)
- [Author Info](#author-info)

---

## Description

You can use this to create CRUD operations on your postgres Database


#### Technologies

- Golang
- GORM
- GIN
- CompilerDaemon
- GotdotEnv

[Back To The Top](#read-me-template)

---

## How To Use

When you clone the project, you need to change the credentioals to match your database. I have a .env file
that contains my credentials and it's not pushed into github so if you want to run this locally make sure
you create one with those values in order for things work properly.

Here are the values you could use

This will run the project on port 1000 on your localhost. Whatever the dbname, you decide, you have to create a postgres database with thesame name.
PORT=1000
DB_URL="host=localhost user=postgres password=12345 dbname=actorsgo port=5432 sslmode=disable"

#### Installation

After successful cloning this repository, you may need to install the following
- Golang 
- GORM 
    -   go get -u gorm.io/gorm
        go get -u gorm.io/driver/postgres
- GIN
    -   go get -u github.com/gin-gonic/gin
- CompilerDaemon
    -   go get github.com/githubnemo/CompileDaemon
        go install github.com/githubnemo/CompileDaemon
- GotdotEnv
    -   go get github.com/joho/godotenv


## Author Info

- Twitter - [@richmondelaigwu](https://twitter.com/richmondelaigwu)
- Website - [Richmond Elaigwu](https://github.com/Sprof22)

[Back To The Top](#read-me-template)

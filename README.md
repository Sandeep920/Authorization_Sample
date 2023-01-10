# Authentication Sample 

## How to run project
* Create a database having name "auth1" in mysql.

* Create .env file and give appropriate DSN

   DSN= "your_user:your_password@tcp(127.0.0.1:3306)/auth1?charset=utf8mb4&parseTime=True&loc=Local"
  
  & run below command 

```bash 
    $ go run server.go
```

## Tasks and explanations
1) User Registration: User can signUp as given below.

mutation createUser{
  createUser(
    input:{
      name: "Sandeep",
      email: "sand@gmail.com",
      password: "qwert"
    }
  )
  {
    id
    name
  }
}

2) User can login using valid user credential and can generate token as given below.

mutation login{
  login(email: "sand@gmail.com", password: "qwert"){
    token,
    expired_at
  }
}

3)If you want to access users, you have to provide an authorization header(token after login).

query users{
  users {
    name,
    id,
    email,
    created_at,
    updated_at
  } 
}
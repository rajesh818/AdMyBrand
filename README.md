# Rest Apis using golang and Mysql database

Firstly you have install go and mysql database software on your PC.

Then create a database with name as **admybrand**. And create a table using table.sql file that is provided.

## How to connect to local database(MySql)
```
sql.Open("mysql","user:password@tcp(127.0.0.1:port)/databasename")  
usually port = 3306
```

**There are 6 endpoints** 
> **GET**  /get ***to get all users information***

> **GET** /get/{id}  ***to get user data by id***

> **DELETE** /delete ***to delete all users information***

> **DELETE** /delete/{id} ***to delete user data by id***

> **POST** /create ***to create a row in the database for an user***

> **PUT** /update/{id} ***to update the information of user having id as their ID***

Intially the table will not have any rows so we have to use /create to add some rows to the table.

In postman set the method as post and url as "http://localhost:8000/create" and provide user information in the body. Remember that we don't have to provided created at data because by default it will add the createdat info to the row.

*sample user information*

```
{
    "id": 1,
    "name": "Rajesh",
    "dob": "2000-01-01",
    "address": "darmavaram",
    "description": "Final year Student"
}
```

After adding the users data into the tables now we can use other endpoints easily

**/update/{id}** to update the information of particular user.(Provide the updated information in body)

**/get** is used to get all users information 

**/get/{id}** is used to get the information of a particular user. (provide the id)

**/delete** to delete the information of all users

**/delete/{id}** is used to delete the information of a particular user.(provide the id).
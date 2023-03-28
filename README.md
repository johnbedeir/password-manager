# password-manager

<p align="center">
  <img src="imgs/logo.png" style="max-width: 150px">
</p>

This application is a free password manager where you add the `NAME` and the `EMAIL` then a `PASSWORD` will be generated and stored in the `MONGODB` database.

### Steps to run the application:

 1. [Instal MongoDB](https://www.mongodb.com/docs/manual/tutorial/install-mongodb-on-ubuntu/)
 2. [Install Studio3T](https://studio3t.com/knowledge-base/articles/installation/)
 3. Make sure `MongoDB` is installed and running on `localhost:27017`
 4. Open `Studio3T` application and connect to `mongodb://localhost:27017`

 <img src="imgs/1.png" width="400px">

 5. Connect to the database

  <img src="imgs/2.png" width="400px">
  <img src="imgs/3.png" width="400px">

Run the application using the `pass-gen` file:
```
./pass-gen
```

Output:
```
Enter name for the password: gmail
Enter user name: mail@gmail.com
Generated password: k6:9CPgc&\-BH3*L
Password stored in MongoDB database with name gmail in collection passwords
```

<img src="imgs/4.png" width="500px">

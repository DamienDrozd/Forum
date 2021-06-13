# Forum

the link fo our website : https://forum-golang-dacd.netlify.app/

## Objectives :

Forum consists in creating a web forum that allows to :
*    Communication between users.
*    Associating categories to posts.
*    Liking and disliking posts and comments.
*    Filtering posts.

The Forum Project started on Thursday May 20, and the submission date was set for Sunday June 13 at 11:59 pm.


## How to execute the program ?

In order to use our program, you have to type the following on your command prompt : 

```console
go run .
```

Afterwards, the server will start and you will be invited to click on the link below that will take you to the website.

```console 
PS C:\VSCODEProjet\GO\Ynov_B1_GO\ProjetForum\Forum> go run .
Server is starting...

Go on http://localhost:8080/

To shut down the server press CTRL + C
```

### Result 

![image](static/img/presentation/forum1.png)

On your website, you can visit the discussion category and you can register to be able to comment and start discussions.

## Functionnalities :

*   SQLite:
    *   In order to store the data in the forum (like users, posts, comments, etc.) we used the database library SQLite.
*   Authentication:
    *   A user can create an account 
    *   While a user created an account, he can log in the account with his email and his password.
    *   The password is encrypted
*   Communication: 
    *  In order to communiquate, a registered user can create post an can comment posts from other users.
*   Likes and Dislikes: T
    *   he user can like posts and like comments
*   Filter: The post lists can be displays by:
    *   name
    *   categories
    *   created posts
    *   liked posts
*   Docker


## Docker

One of the main goal of the project Forum was to create a Dockerfile, (an executable that contains the files and the dependencies of the program), one image and one container.

### How to do a Docker ?

First of all, you have to create a DockerFile which contains the following :

```code
# The base go-image
FROM golang:1.16

# Create a directory for the app
RUN mkdir /Forum

# Add all files into Forum app 
ADD . /Forum
 
# Set working directory
WORKDIR /Forum

COPY go.mod go.sum ./

RUN go mod download

RUN go build -o main .
 
# Run the server executable
CMD [ "/Forum/main" ]

```

Afterwards, you have to download the Docker sofware on your computer.

Then, you have to run the program ```dockerize.bat``` which contains the following commands :

```bash
@REM ECHO off allows to print only the return of the command
@ECHO OFF
ECHO.
ECHO ---------------------BUILDING IMAGE DOCKER---------------------------
docker build -t forum .
ECHO. 

ECHO ---------------------RUNNING DOCKER ON 8080--------------------------
docker run -d --name Forum -p 8080:8080 forum
ECHO. 
ECHO ---------------------DOCKER IMAGE LIST-------------------------------
docker images
ECHO. 
ECHO ----------------------CONTAINER LIST--------------------------------
docker container ls
```
This program will allow you to create the docker container and run the image. In addition, it will show you the list of your docker images and your container.

<hr>

## Additional Modules


*   #### Moderation:  
    The module is used to add moderation roles in the Forum   
    *   Guests:
        *   These are unregistered-users that can neither post, comment, like or dislike a post. They only have the permission to see those posts, comments, likes or dislikes.
    *   Users: 
        *   These are the users that will be able to create, comment, like or dislike posts.
    *   Moderators:
        *   They should be able to monitor the content in the forum by deleting or reporting post to the admin
        *   To create a moderator the user should request an admin for that role
        *   Approves posted messages before they become publicly visible.
    *   Administrators
        *   Promote or demote a normal user to, or from a moderator user.
        *   Receive reports from moderators. If the admin receives a report from a moderator, he can respond to that report
        *   Delete posts and comments
        *   Manage the categories, by being able to creating and deleting them.

*   #### Advanced - features:
    *   There is an activity page,
        *    A user can see:
            *    The user's posts
            *   Where the user left a like or a dislike
            *   Where yhe user has been commenting
        *   The a section where the user can remove posts and comments.

<hr>

## Sources :

*   SQLite : https://www.youtube.com/watch?v=OHv2K4wL9Yc
*   Encryption : https://gowebexamples.com/password-hashing/
*   Cookies : https://medium.com/wesionary-team/cookies-and-session-management-using-cookies-in-go-7801f935a1c8
*   Docker : 
    *   https://hub.docker.com/
    *   https://hub.docker.com/_/golang
*   Example for the web design : https://cdn.dribbble.com/users/1231043/screenshots/8108762/media/54cf5ab3cc5c44e3e75412357e01cf71.jpg 
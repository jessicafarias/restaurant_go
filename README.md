##### This is a back-end API using golang.

# Restaurant go
Each Restaurant has many comments

[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]

<!-- PROJECT LOGO -->
  <h3 align="center">Track</h3>

  <p align="center">
    This project wirte and read data from a PostgreSQL database<br /></p>

<!-- ABOUT THE PROJECT -->

#### Features:
- Create a new Restaurant
- Create a new Comment for a restaurant
- Get all restaurant comments
- Get all restaurants

### Built With
This project was built using these technologies.
* Golang
* SQL database


### Getting Started
To get a local copy up and running follow these simple example steps.

### Clone
* Clone this repo:
  - Clone with SSH:
  ```
    git@github.com:jessicafarias/restaurant_go.git
  ```
  - Clone with HTTPS
  ```
    https://github.com/jessicafarias/restaurant_go.git
  ```
  - Clone with GitHub CLI
  ```
    gh repo clone jessicafarias/restaurant_go.git
    
 - cd to restaurant_go


### Setup

Start the local webserver:

```$ go run main.go``` will open the project at local webserver at http://localhost:8080/ 

### Using | Requests examples using JS

for all request you will need base URL:
const baseUrl = `http://localhost:8080/`

#### Restaurant fields
- ingeger Id: Not requiered to POST a new restaurant
- string Name: "Restaurant name"
- string Description: "Restaurant description"

#### Comment fields
- ingeger Id: Not requiered to POST a new Comment
- string Username: "User name (wirter)"
- string Body: "Comment body"
- int RestaurantID: Restaurant of id related, the restaurant id should exist


#### Create New Restaurant
```
const url = `${baseUrl}/restaurant`;
const config = `{
    method: 'POST',
    mode: 'cors',
    redirect: 'follow',
    body: JSON.stringify(data),
  }`
```

#### Create New Comment
```
const url = `${baseUrl}/comment`;
const config = `{
    method: 'POST',
    mode: 'cors',
    redirect: 'follow',
    body: JSON.stringify(data),
  }`
```

#### Get all restaurants
```
const url = `${baseUrl}/restaurants`;
const config = `{
    method: 'GET',`
}
```

#### Get all comments by restaurant id
```
const url = `${baseUrl}/comments/{restaurant_id}`;
const config = `{
    method: 'GET',
    `
```

#### Get restaurant by id
```
const url = `${baseUrl}/restaurant/{restaurant_id}`;
const config = `{
    method: 'GET',
    mode: 'cors',`
```

## Author

## 👤 Jessica Michelle Farías Rosado:
Working as a FullStack developer on this project.

 [![Website](https://img.shields.io/badge/-Website-black?style=for-the-badge&logo=Julia&logoColor=white)](https://jessicafarias.github.io/)

 [![LINKEDIN](https://img.shields.io/badge/-LINKEDIN-0077B5?style=for-the-badge&logo=Linkedin&logoColor=white)](https://www.linkedin.com/in/jessica-michelle-farias-rosado/)

 [![EMAIL](https://img.shields.io/badge/-EMAIL-D14836?style=for-the-badge&logo=Mail.Ru&logoColor=white)](mailto:jessica.farias.rosado@gmail.com)
 
 [![TWITTER](https://img.shields.io/badge/-TWITTER-1DA1F2?style=for-the-badge&logo=Twitter&logoColor=white)](https://twitter.com/FariasRosado)


## 🤝 Contributing

Contributions, issues and feature requests are welcome!

Feel free to check the [issues page](https://github.com/jessicafarias/restaurant_go/issues).

## Show your support

Give a :star: if you like this project!



<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[contributors-shield]: https://img.shields.io/github/contributors/jessicafarias/restaurant_go.svg?style=flat-square
[contributors-url]: https://github.com/jessicafarias/restaurant_go/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/jessicafarias/restaurant_go.svg?style=flat-square
[forks-url]: https://github.com/jessicafarias/restaurant_go/network/members
[stars-shield]: https://img.shields.io/github/stars/jessicafarias/restaurant_go.svg?style=flat-square
[stars-url]: https://github.com/jessicafarias/restaurant_go/stargazers
[issues-shield]: https://img.shields.io/github/issues/jessicafarias/restaurant_go.svg?style=flat-square
[issues-url]: https://github.com/jessicafarias/restaurant_go/issues

## 📝 License

This project is [MIT](https://opensource.org/licenses/MIT) licensed.
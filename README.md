<p align="center">
  <a href="" rel="noopener">
 <!-- <img width=200px height=200px src="https://i.imgur.com/FxL5qM0.jpg" alt="Bot logo"></a> -->
</p>

<h3 align="center">Project Title</h3>

<div align="center">

[![Status](https://img.shields.io/badge/status-active-success.svg)]()
[![Platform](https://img.shields.io/badge/platform-golang-blue.svg)]()
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](/LICENSE)

</div>

---


## 📝 Table of Contents

- [About](#about)
- [Demo / Working](#demo)
- [Getting Started](#getting_started)
- [Deployment](#deployment)
- [Usage](#usage)
- [Built Using](#built_using)
- [Authors](#authors)


## 🧐 About <a name = "about"></a>

This project is pure golang.
Gokit is a microservice kit developer.
GoDog is a BBD cucumber for testing purpouses.

## 🎥 Demo / Working <a name = "demo"></a>

![Working](godog.png)

--- 

## 🏁 Getting Started <a name = "getting_started"></a>

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See [deployment](#deployment) for notes on how to deploy the project on a live system.

### Prerequisites

```
GO111MODULE=on go get github.com/cucumber/godog/cmd/godog@v0.9.0 
```
```
go get -u  github.com/go-kit/kit 
```
```
go get github.com/kujtimiihoxha/kit
```

### Installing

A step by step series of examples that tell you how to get a development env running.

Say what the step will be

```
docker-compose up
```

## 🔧 Running the tests <a name = "tests"></a>

Explain how to run the automated tests for this system.

### Break down into end to end tests

This project cool part is bbd for testing , in this case cucumber for go is used

```
go test -v --godog.random
```


## ⛏️ Built Using <a name = "built_using"></a>

- [Golang](https://golang.org/) - Main Framework

## ✍️ Authors <a name = "authors"></a>

- [@miguelmcmc](https://github.com/miguelmcmc) - Idea & Initial work

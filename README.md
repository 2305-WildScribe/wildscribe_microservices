<a id="top"></a>
# WildScribe

[WildScribe Front End Repo](https://github.com/2305-WildScribe/2305-WildScribe_FE)

[Deployment](https://wildscribe.vercel.app/)

## Development Teams

### Front End Team

- Alice Abarca: [GitHub](https://github.com/aliceabarca/) | [LinkedIn](https://www.linkedin.com/in/alice-abarca-431615272/)
- Jocelyn Wensloff: [GitHub](https://github.com/Jwensloff/) | [LinkedIn](https://www.linkedin.com/in/jocelynwensloff/)

### Back End Team

- Parker Boeing: [GitHub](https://github.com/ParkerBoeing) | [LinkedIn](https://www.linkedin.com/in/parker-boeing/)
- Derek Chavez: [GitHub](https://github.com/DChavez18) | [LinkedIn](https://www.linkedin.com/in/derek-chavez/)
- Ian Lyell: [GitHub](https://github.com/ILyell) | [LinkedIn](https://www.linkedin.com/in/ian-lyell/)

---
## Architecture
This backend is designed with a microservice architecture and utilizes GRPC to handle communication between the services. The call from the client hits the main service and is then transfered via GRPC's protocol buffer, which converts the code to binary, where it is received by the microservice and consequently converted back into JSON. 

## Installing Go on macOS

Go (also known as Golang) is a popular programming language developed by Google. This guide will walk you through the steps to install Go on a macOS system.

## Prerequisites

Before you begin, make sure you have the following:

- A macOS computer
- An internet connection

## Installation Steps

Follow these steps to install Go on your macOS system:

1. **Download the Go Installer:**

   Visit the official Golang website to download the installer for macOS. Go to [https://golang.org/dl/](https://golang.org/dl/) and find the macOS version.

2. **Choose the Correct Package:**

   You will see several options for macOS, typically labeled as `goX.Y.darwin-amd64.pkg`, where `X.Y` represents the version number. Click on the latest version to download it.

3. **Install Go:**

   After downloading the package, open the downloaded `.pkg` file by double-clicking it. This will initiate the installation process.

4. **Follow the Installer Instructions:**

   Follow the on-screen instructions to complete the installation. You may need to enter your password to allow the installer to make changes to your system.

5. **Verify the Installation:**

   After the installation is complete, open your terminal. To verify that Go has been installed successfully, you can run the following command:

   ```bash
   go version

---

### How to Install the Project

- Fork and clone this repo
- Navigate to the main folder and run the command ```docker-compose up``` in your terminal
- This should build a docker container for the main service as well as each microservice hosted in the app
- You can build / containerize microservices individually by running ```docker build -t adventure / user```
- If you want to hit individual microservices, GRPC protocols must be used via a tool such as Postman
- Run go test in order to run the test suite

---

# API JSON Contract

## Users

Description of API endpoints for Front End application:

### Getting User

`POST /api/v0/user`

**Request**

```json
{
    "data": {
        "type": "user",
        "attributes": {
            "email": "me@gmail.com",
            "password": "hi"
        }
    }
}
```
**Success Response (200 OK)**:

- **Status**: 200 OK
- **Description**: Successful response with user id
- **Data format**: a hash with a hash of user data

```json
{
    "data": {
        "type": "user",
        "attributes": {
            "name": "Ian",
            "user_id": "652edaa67a75034ea37c6652"
        }
    }
}
```
**Error Response (400 Bad Request)**

```json
{
    "data": {
        "error": "Invalid Request",
        "type": "user"
    }
}
```
---
## Adventures

Description of API endpoints for Front End application:

### Getting Adventures for User

`POST /api/v0/user/adventures`

**Request**

```json
{
    "data":{
        "type": "adventure",
        "attributes":{
            "adventure_id": "652da923ff996de855a6d39d"
        }
    }
}
```
**Success Response (200 OK)**:

- **Status**: 200 OK
- **Description**: Successful response with all adventures associated with user id
- **Data format**: a hash with all adventures, with a hash of adventure data

```json
{
    "data": {
        "type": "adventure",
        "attributes": {
            "user_id": "65299d4ceb708107b33729c6",
            "adventure_id": "652da923ff996de855a6d39d",
            "activity": "Running",
            "date": "10/11/2023",
            "image_url": "https://www.rei.com/dam/parrish_091412_0679_main_lg.jpg",
            "stress_level": "Very High",
            "hours_slept": 8,
            "sleep_stress_notes": "notes about sleep and stress",
            "hydration": "Hydrated",
            "diet": "Good Diet",
            "diet_hydration_notes": "Some Hydraytion",
            "beta_notes": "Running is real hard"
        }
    },
    {
        "type": "adventure",
        "attributes": {
            "user_id": "65299d4ceb708107b33729c6",
            "adventure_id": "652da923ff996de855a6d39d",
            "activity": "Swimming",
            "date": "10/11/2024",
            "image_url": "https://www.rei.com/dam/parrish_091412_0679_main_lg.jpg",
            "stress_level": "High",
            "hours_slept": 9,
            "sleep_stress_notes": "notes about sleep and stress",
            "hydration": "Hydrated",
            "diet": "Good Diet",
            "diet_hydration_notes": "Some Hydraytion",
            "beta_notes": "Swimming is real hard"
        }
    }
}
```

**Error Response (400 Bad Request)**

```json
{
    "data": {
        "error": "Invalid Request"
    }
}
```
---
### Getting An Adventure

`POST /api/v0/user/adventure`

**Request**

```json
{
    "data":{
        "type": "adventure",
        "attributes":{
            "adventure_id": "652ff8c82ed41a2d015d993b"
        }
    }
}
```
**Success Response (200 OK)**:

- **Status**: 200 OK
- **Description**: Successful response with adventure data associated with adventure id
- **Data format**: a hash with adventure, with a hash of adventure data

```json
{
    "data": {
        "type": "adventure",
        "attributes": {
            "user_id": "65299d4ceb708107b33729c6",
            "adventure_id": "652ff8c82ed41a2d015d993b",
            "activity": "Running",
            "date": "10/11/2023",
            "image_url": "https://www.rei.com/dam/parrish_091412_0679_main_lg.jpg",
            "stress_level": "Very High",
            "hours_slept": 8,
            "sleep_stress_notes": "notes about sleep and stress",
            "hydration": "Hydrated",
            "diet": "Good Diet",
            "diet_hydration_notes": "Some Hydraytion",
            "beta_notes": "Running is real hard"
        }
    }
}
```
**Error Response (404 Not Found)**

```json
{
    "data": {
        "error": "Invalid adventure ID",
        "attributes": {
            "adventure_id": [
                {
                    "Key": "_id",
                    "Value": "652ff8c82ed41a2d015d993b"
                }
            ]
        }
    }
}
```
---
### Creating An Adventure

`POST /api/v0/adventure`

**Request**

```json
{
 "data": {
        "type": "adventure",
        "attributes": {
            "user_id": "65299d4ceb708107b33729c6",
            "activity": "Running",
            "date": "10/11/2023",
            "notes": "Running is hard",
            "image_url": "https://www.rei.com/dam/parrish_091412_0679_main_lg.jpg",
            "stress_level": "Very High",
            "hours_slept": 8,
            "sleep_stress_notes": "notes about sleep and stress",
            "hydration": "Hydrated",
            "diet": "Good Diet",
            "diet_hydration_notes": "Some Hydraytion",
            "beta_notes": "Running is real hard"
        }
    }
}
```
**Success Response (201 OK)**:

- **Status**: 201 OK
- **Description**: Successful response with adventure id and success message
- **Data format**: a hash with message, with a hash of new adventure id

```json
{
    "data": {
        "type": "adventure",
        "message": "success",
        "attributes": {
            "adventure_id": "652ff8c82ed41a2d015d993b"
        }
    }
}
```
**Error Response (400 Bad Request)**

```json
{
    "data": {
        "error": "Invalid user ID",
        "attributes": {
            "user_id": "65299d4ceb708107b33729c6"
        }
    }
}
```
---
### Deleting An Adventure

`DELETE /api/v0/adventure`

**Request**

```json
{
    "data": {
        "type": "adventure",
        "attributes": {
            "adventure_id":"6530428eb4e1886116236a8a"
        }
    }
}
```
**Success Response (200 OK)**:

- **Status**: 200 OK
- **Description**: Successful response with success message
- **Data format**: a hash with message and and adventure type

```json
{
    "data": {
        "type": "adventure",
        "message": "success"
    }
}
```

**Error Response (400 Bad Request)**

```json
{
    "data": {
        "error": "Invalid adventure ID",
        "attributes": {
            "adventure_id": "6530428eb4e1886116236a8"
        }
    }
}
```
---
### Updating An Adventure

`PUT /api/v0/adventure`

**Request**

```json
 {
    "data": {
        "type": "adventure",
        "attributes": {
            "user_id": "652ed3250b59c18916efde3f",
            "adventure_id": "652eda24dc59c7aa766a309b",
            "activity": "Walking",
            "date": "10/11/2023",
            "notes": "Running is hard",
            "image_url": "https://www.rei.com/dam/parrish_091412_0679_main_lg.jpg",
            "stress_level": "Very High",
            "hours_slept": 8,
            "sleep_stress_notes": "notes about sleep and stress",
            "hydration": "Hydrated",
            "diet": "Good Diet",
            "diet_hydration_notes": "Some Hydraytion",
            "beta_notes": "Running is real hard"
        }
    }
}
```
**Success Response (200 OK)**:

- **Status**: 200 OK
- **Description**: Successful response with success message
- **Data format**: a hash with message and and adventure type

```json
{
    "data": {
        "type": "adventure",
        "message": "success"
    }
}
```

**Error Response (400 Bad Request)**

```json
{
    "data": {
        "error": "Invalid Request"
    }
}
```
---
# Developed With
<img src="https://raw.githubusercontent.com/devicons/devicon/55609aa5bd817ff167afce0d965585c92040787a/icons/heroku/heroku-original-wordmark.svg" width="50" alt="heroku Logo"><img src="https://raw.githubusercontent.com/devicons/devicon/55609aa5bd817ff167afce0d965585c92040787a/icons/github/github-original.svg" width="50" alt="github Logo"><img src="https://camo.githubusercontent.com/d4de39c8b497d0e00bf90a543e9e43a30e87a057445832e4865197d002392538/68747470733a2f2f7261772e6769746875622e636f6d2f436972636c6543492d5075626c69632f63696d672d676f2f6d61696e2f696d672f636972636c652d676f2e7376673f73616e6974697a653d74727565" width="50"><img src="https://assets.stickpng.com/images/6299f743b04c5ae587c4119d.png" width="50"><img src="https://seeklogo.com/images/M/mongodb-logo-D13D67C930-seeklogo.com.png" width="50"><img src= "https://d3r49iyjzglexf.cloudfront.net/circleci-logo-stacked-fb-657e221fda1646a7e652c09c9fbfb2b0feb5d710089bb4d8e8c759d37a832694.png" width="50"><img src="https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQB--kQ3qZkabOBf1-f51nWAzJYbggP-bvkDBwwk7ZPUSPUqm2hM6L3H9fNgKE3gGyPido&usqp=CAU" width = "50"><img src = "data:image/png;base64,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" width = "50"><img src = "https://d15shllkswkct0.cloudfront.net/wp-content/blogs.dir/1/files/2022/05/Azure.png" width = "50">
---
<a href="#top">Back to Top</a>
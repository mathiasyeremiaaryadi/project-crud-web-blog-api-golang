# Golang CRUD Web Blog Post API

This is a CLI application for expense tracker to manage your expenses. 

Project from: https://roadmap.sh/projects/blogging-platform-api


## Features

- Create Blog Post
- Update Blog Post
- Delete Blog Post
- Get All Blog Posts
- Get Blog Post


## Documentation

### Create Blog Post
Endpoint
```
POST /posts
```

Request
```
{
  "title": "My First Blog Post",
  "content": "This is the content of my first blog post.",
  "category": "Technology",
  "tags": ["Tech", "Programming"]
}
```

Response
```
HTTP 201/Created
{
    "status": "success",
    "data": {
      "id": 1,
      "title": "My First Blog Post",
      "content": "This is the content of my first blog post.",
      "category": "Technology",
      "tags": ["Tech", "Programming"],
      "createdAt": "2021-09-01T12:00:00Z",
      "updatedAt": "2021-09-01T12:00:00Z"
    },
    "message": "new post created"
}
```

### Update Blog Post
Endpoint
```
PUT /posts/1
```

Request
```
{
  "title": "My Updated Blog Post",
  "content": "This is the updated content of my first blog post.",
  "category": "Technology",
  "tags": ["Tech", "Programming"]
}
```

Response
```
HTTP 200/OK
{
    "status": "success",
    "data": {
      "id": 1,
      "title": "My Updated Blog Post",
      "content": "This is the updated content of my first blog post.",
      "category": "Technology",
      "tags": ["Tech", "Programming"],
      "createdAt": "2021-09-01T12:00:00Z",
      "updatedAt": "2021-09-01T12:30:00Z"
    },
    "message": "new post updated"
}
```

### Delete Blog Post
Endpoint
```
DELETE /posts/1
```

Request
```
-
```

Response
```
HTTP 204/No Content
-
```

### Get Blog Post
Endpoint
```
GET /posts/1
```

Request
```
-
```

Response
```
HTTP 200/OK
{
    "status": "success",
    "data": {
      "id": 1,
      "title": "My First Blog Post",
      "content": "This is the content of my first blog post.",
      "category": "Technology",
      "tags": ["Tech", "Programming"],
      "createdAt": "2021-09-01T12:00:00Z",
      "updatedAt": "2021-09-01T12:00:00Z"
    },
    "message": "get post success"
}
```

## Get All Blog Posts
Endpoint
```
GET /posts
```

Request
```
-
```

Response
```
HTTP 200/OK
{
    "status": "success",
    "data": [
      {
        "id": 1,
        "title": "My First Blog Post",
        "content": "This is the content of my first blog post.",
        "category": "Technology",
        "tags": ["Tech", "Programming"],
        "createdAt": "2021-09-01T12:00:00Z",
        "updatedAt": "2021-09-01T12:00:00Z"
      },
      {
        "id": 2,
        "title": "My Second Blog Post",
        "content": "This is the content of my second blog post.",
        "category": "Technology",
        "tags": ["Tech", "Programming"],
        "createdAt": "2021-09-01T12:30:00Z",
        "updatedAt": "2021-09-01T12:30:00Z"
      }
    ],
    "message": "get all post success"
}
```

## Get All Blog Posts With Filter
Endpoint
```
GET /posts?term=first
```

Request
```
-
```

Response
```
HTTP 200/OK
{
    "status": "success",
    "data": [
      {
        "id": 1,
        "title": "My First Blog Post",
        "content": "This is the content of my first blog post.",
        "category": "Technology",
        "tags": ["Tech", "Programming"],
        "createdAt": "2021-09-01T12:00:00Z",
        "updatedAt": "2021-09-01T12:00:00Z"
      }
    ],
    "message": "get all post success"
}
```


## Clone the project

```bash
git clone https://github.com/mathiasyeremiaaryadi/project-crud-web-blog-api-golang.git
```

Go to the project directory

```bash
cd project-crud-web-blog-api-golang
```

Install dependencies

```bash
go build
```

Start the server

```bash
./web-blog-api
```

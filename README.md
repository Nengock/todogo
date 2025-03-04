# Command-Line To-Do List Manager

## Running the Application

### Build the Application
```sh
go build -o todo-app
```

### Add a Task
```sh
./todo-app add "Walk the dog"
```

### List Tasks
```sh
./todo-app list
```

### Complete a Task
```sh
./todo-app complete 1
```

### Delete a Task
```sh
./todo-app delete 2
```

### Running with Docker

#### Build the Docker Image
```sh
docker build -t todo-app .
```

#### Run the Docker Container
```sh
docker run --rm -it -v $(pwd)/tasks.json:/app/tasks.json todo-app add "Walk the dog"
docker run --rm -it -v $(pwd)/tasks.json:/app/tasks.json todo-app list
docker run --rm -it -v $(pwd)/tasks.json:/app/tasks.json todo-app complete 1
docker run --rm -it -v $(pwd)/tasks.json:/app/tasks.json todo-app delete 2
```

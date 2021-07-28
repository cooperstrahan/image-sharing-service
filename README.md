## Testing
Tests are currently dependent on the entries in the database, Tests will pass as long as the current entries in the database are maintained but will fail after 
images are deleted / changed -- suggestion to run tests before messing with the application.



# Image sharing
![alt text](https://github.com/cooperstrahan/image-sharing-service/blob/main/imagesharing.png)

### System Requirements
Go installation
yarn/npm installation

# Set up front end -- image-sharing

## Project setup
```
cd image-sharing
```
```
yarn install
```

### Compiles and hot-reloads for development
```
yarn serve
```

### Compiles and minifies for production
```
yarn build
```

### Lints and fixes files
```
yarn lint
```


# Set up back end -- api
```
cd api/cmd/main
```

## Import project dependencies
```
go build
```
## Run backend
```
./main
```

### TODO:
1. Update User Interface
2. Fix error checking
3. Filter by tags
4. Testing for the front end
5. Update testing for the back end to work with Interfaces in order to work with mocked cases
6. Search
7. Dockerize project

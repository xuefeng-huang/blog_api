# Project setup
1. `git clone https://github.com/xuefeng-huang/blog_api.git`
2. `cd blog_api/ && docker-compose up`
3. server will be up and listening on port 8080

# Runing test
When runing the test, the applition uses a test database in the db container by passing a different db name while initailising. My container for app service named `blog_api_app_1`, so I run tests by this command `docker exec -it blog_api_app_1 bash -c "cd /app/test ; go test -vet=off"`-vet=off flag is to get around an issue which `go test` does not run in golang Alpine containers.

# Side note
TestCreateArticleBad function failed as I could not understand why the binding validation function `ShouldBindJSON` does not work when it is in the test mode. I tried various things to get it work but no good solution found yet. Joy of learning new framework I guess.
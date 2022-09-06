# Server for CheckInBoard



Written in Go. Database Postgresql. Docker-compose

## Connect to the local postgres instance on docker-compose

docker exec -it postgres-database-1 psql CheckInBoard michzuerch 


## goreleaser

git tag -a v0.1.0 -m "First release"
git push origin v0.1.0


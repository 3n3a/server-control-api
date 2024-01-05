# docker build -t server-control-api .
docker run -d --name sca -p 3000:3000 -e "ENVIRONMENT=prod" -e "DEFAULT_API_KEY=1234" -v /var/run/docker.sock:/var/run/docker.sock server-control-api

echo "Starting Tests" && echo ""

echo "should both print a 500 image could not be pulled error"
curl -X POST -H 'Authorization: Bearer 1234' -d "image=docker.io/library/asldfjlkdsfdsjfldjf" http://localhost:3000/docker/images/pull && echo ""
curl -X POST -H 'Authorization: 1234' -d "image=docker.io/library/asldfjlkdsfdsjfldjf" http://localhost:3000/docker/images/pull && echo "" && echo ""

echo "should both print a 200 image successfully pulled message"
curl -X POST -H 'Authorization: Bearer 1234' -d "image=docker.io/library/busybox" http://localhost:3000/docker/images/pull && echo ""
curl -X POST -H 'Authorization: 1234' -d "image=docker.io/library/busybox" http://localhost:3000/docker/images/pull && echo "" && echo ""

echo "should both print a 200 image successfully pulled message"
curl -X POST -H 'Authorization: Bearer 1234' -d "image=docker.io/library/busybox:1.24" http://localhost:3000/docker/images/pull && echo ""
curl -X POST -H 'Authorization: 1234' -d "image=docker.io/library/busybox:1.24" http://localhost:3000/docker/images/pull && echo "" && echo ""

echo "Docker images (busybox latest 1.24)"
docker images | grep busybox

docker stop sca
docker rm sca
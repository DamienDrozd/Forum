echo "### Building image Docker"
docker build -t forum .
echo
echo "### Running Docker on port 8080"
docker run -d -p 8080:8080 forum
echo
echo "### Images list"
docker images
echo
echo "### Containers list"
docker container ls

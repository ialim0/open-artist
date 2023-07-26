echo "Entrez le nom de l'image"
read image
docker image build -f Dockerfile -t $image .
docker images
echo "Donnez le nom du container"
read container
echo "Donnez le port"
read port
docker container run -p $port:$port --detach --name $container $image
docker ps -a
docker exec -it $container /bin/bash
ls -l
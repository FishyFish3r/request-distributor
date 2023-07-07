# request-distributor
cd path/request-distributor

docker-compose up

docker-compose stop servX(example serv1) 
docker-compose start servX(example serv1) 

docker exec -it request-distributor bash
(in bash)#cat logs.txt



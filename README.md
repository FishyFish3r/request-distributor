# request-distributor
//docker start

cd path/request-distributor

docker-compose up

//docker test servsers

docker-compose stop servX(example serv1) 

docker-compose start servX(example serv1) 

//to view logs

docker exec -it request-distributor-dist-1 bash
(in bash)#cat logs.txt



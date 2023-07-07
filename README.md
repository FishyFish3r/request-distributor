# request-distributor
cd path/request-distributor

docker-compose up
___________test_______________
docker-compose stop servX(example serv1) 
docker-compose start servX(example serv1) 
___________logs view__________
docker exec -it request-distributor bash
(in bash)#cat logs.txt



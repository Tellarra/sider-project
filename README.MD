## US 3: Making tasks accessible to siders

To ensure that tasks are available to temporary workers in their mobile application as soon as they are created by companies, the following functional system or workflow can be considered:

1. Store data in the database immediately upon task creation.
2. Configure an Elasticsearch instance to store the new tasks.
3. Develop a program that can populate the Elasticsearch index with data.
4. Host the program in a Kubernetes environment and integrate it with a cloud platform like GCP or AWS.
5. Set up a cron job to run every 10 minutes throughout the day.
6. The cron job triggers the program to check for any new data to add to the Elasticsearch index.
7. If new data is found, perform a partial indexation to store it in Elasticsearch.
8. After approximately 20 minutes, the data becomes available for everyone to see when the `/tasks` endpoint is called.

Implementing this solution may require some effort, but it reduces the number of database calls compared to querying the database for each API call. Additionally, Elasticsearch provides easy manipulation and updating of data.


## What I would have done better

I would have used channels with my goroutines to feed my index, because for now it's too much for the elastic search to handle and index every datas. 
That is why I put a const FeedSize to stock indexation of entities at 2000 for now
I would have made a better organisation of my functions to fit a bit better to the hexagonal architecture

## To RUN project 
First do a 
```
docker-compose up --build
```

When your ES and Kibana is setup you can do

```
go build
```

Or you can do 

```
./siderproject
```

Or you can do 
```
go run main.go
```

If there is no datas in the ES indexes, you need to launch a feed on /feed

## Endpoints

GIN-debug] GET    /health                    See if app is up
[GIN-debug] GET    /tasks                    To get a list of tasks available in the elasticsearch
[GIN-debug] GET    /feed                     To feed your datas to the elasticSearch container
[GIN-debug] PATCH  /tasks/:id                To update an assignee to a task


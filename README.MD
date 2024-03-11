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



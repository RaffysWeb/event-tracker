Install dependencies
```
go install github.com/db-journey/journey/v2
```

Create keyspace
```
cqlsh
create keyspace mykeyspace WITH REPLICATION = { 
   'class' : 'SimpleStrategy', 
   'replication_factor' : 1 
};
  ```

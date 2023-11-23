ps -ax | grep mysql | grep -v grep

# OR for only the running `PID`

ps -ef | grep mysql | grep -v grep | awk '{print $2}'

# OR this If you have this on your machine, I recommend using

pgrep -f mysql

## Postress
If you created the table product with an id column, then the sequence is not simply called product, 
but rather product_id_seq (that is, ${table}_${column}_seq).

This is the ALTER SEQUENCE command you need:
```
ALTER SEQUENCE public.student_id_seq RESTART WITH 3000
```

## MYSQL

```
ALTER TABLE student_tracker.student AUTO_INCREMENT=3000
```


SETUP auto increment id is 1 then

```
TRUNCATE student_tracker.student ;
```
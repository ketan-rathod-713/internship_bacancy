# Cron Job In Unix

Cron commands


```
crontab -l 
```

To show all current jobs


```
crontab -e
```

To edit or add new cron jobs


cron job format

```
***** cd myfolder &  ./create_file.sh
```

***** is called as schedular expression.
minutes, hour, day of the month, month, day of the week

* * * * * or * * or * * * means run job every minute. As cron schedular check every minute that is there any job to run or not. Hence by default if nothing is passed then job will be run every minute.

0 * * * * means every hour -> minutes vali field 
0 0 * * * means every day at 12:00 AM
0 0 * * FRI means at 12:00 AM only on FRIDAY
0 0 1 * * means every 1st day of month at 12:00 AM

0 0 1 * FRI means every 1st FRIDAY of month at 12:00 AM -- WRONG
        At 12:00 AM, on day 1 of the month, and on Friday -- CORRECT

## Terminologies

Minute :- how far past the top of the hour your job runs. For eg. if hour is 12:00 AM and minute is 0 then job will run at 12:00 AM

Hour :- how far past midnight your job runes. For eg. if midnight is 0 then noon is 12 and so on.

Day :- it represents the day of the month.

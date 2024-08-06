# Improvements:

## Task Status table ?

# report API: domain = report

## Team

GET SSE (loading)

- /v1/admin/report/room/:roomID
  for all schedules in the room
- task /week for range duration ex: (1 Jan - 3 May) / person
- average task duration
- median task duration
- task status changes overtime (done % each time)
- person with most task
- person with most task duration
- sort person from most % done

- compare report with prev report
- report table? roomID and duration

## Person

GET SSE (loading)

- local first on FE
- /v1/admin/report/member/:memberID
- num of tasks /week
- average task duration
- median task duration
- task status changes overtime (done % each time

(low priority)

- /v1/admin/report/member/:memberID/pdf
- same but pdf
- generate pdf

# random mock task generator for seeder

1. faker: https://github.com/go-faker/faker/blob/main/example_with_tags_test.go
2. use repository to create data on DB (try this)

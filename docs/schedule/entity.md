```ts
type DurationDay = number;
 
type Schedule = {
    tasks: Array<Task>;
}

type Task = {
    id: string;
    details: TaskDetails;
    dependencies?: Array<Task['id']>;
    startTime: Date;
    duration: DurationDay;
}

type TaskDetails = {
    name: string;
    owner: string;
    assignee: string;
}

```


```go


```
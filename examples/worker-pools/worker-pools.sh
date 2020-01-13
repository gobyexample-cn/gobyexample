# 运行程序，显示 5 个任务被多个 worker 执行。
# 尽管所有的工作总共要花费 5 秒钟，但该程序只花了 2 秒钟，
# 因为 3 个 worker 是并行的。
$ time go run worker-pools.go 
worker 1 started  job 1
worker 2 started  job 2
worker 3 started  job 3
worker 1 finished job 1
worker 1 started  job 4
worker 2 finished job 2
worker 2 started  job 5
worker 3 finished job 3
worker 1 finished job 4
worker 2 finished job 5

real	0m2.358s

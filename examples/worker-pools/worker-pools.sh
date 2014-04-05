# 执行这个程序，显示 9 个任务被多个 worker 执行。整个程序
# 处理所有的任务仅执行了 3s 而不是 9s，是因为 3 个 worker
# 是并行的。
$ time go run worker-pools.go 
worker 1 processing job 1
worker 2 processing job 2
worker 3 processing job 3
worker 1 processing job 4
worker 2 processing job 5
worker 3 processing job 6
worker 1 processing job 7
worker 2 processing job 8
worker 3 processing job 9

real	0m3.149s

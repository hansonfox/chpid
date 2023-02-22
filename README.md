# chpid

A chinese pid generator and validation tool

build on Go Cobra

CN身份证 生成和校验工具

基于go cobra 创建

* 查看帮助
 + `chpid -h `
* sub command example:
+ `parse -p  pid_string           # parse pid,list simply information `
+ `valid -v  pid_string           # valid a single pid `
+ `valid -f  pidfiles.txt         # valid a txt file contain pids `
+ `rand  -r 100 -o output.txt     # random generate pids,and write into a file `

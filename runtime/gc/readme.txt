prof文件是内存分析文件，可以使用pprof来查看分析结果

go tool pprof memprofile.prof
进入到交互命令
查看分析结果

>> top
flat 表示函数自身的内存使用量。
flat% 表示函数自身内存使用量占总内存使用量的百分比。
cum 表示函数及其调用的所有函数的内存使用总量。
cum% 表示函数及其调用的所有函数的内存使用总量占总内存使用量的百分比。

>> list

>> web
web可以得到网页版


火焰图查看
go tool pprof -http=:8080 memprofile.prof
(先要安装 Graphviz)

这样看既可以看火焰图，也可以看箭头图
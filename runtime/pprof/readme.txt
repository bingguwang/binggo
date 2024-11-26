import	"net/http"
import	_ "net/http/pprof"
go func() {
	http.ListenAndServe("0.0.0.0:8222", nil)
}()


访问: http://192.168.2.141:8222/debug/pprof/heap?debug=1
可以看到

# runtime.MemStats
# Alloc = 3080920
# TotalAlloc = 27897960
# Sys = 18010776    进程从系统获得的内存空间，虚拟地址空间。
# Lookups = 0
# Mallocs = 449302
# Frees = 419264
# HeapAlloc = 3080920   进程堆内存分配使用的空间，通常是用户new出来的堆对象，包含未被gc掉的。
# HeapSys = 7831552     进程从系统获得的堆内存，因为golang底层使用TCmalloc机制，会缓存一部分堆内存，虚拟地址空间。
# HeapIdle = 3047424
# HeapInuse = 4784128
# HeapReleased = 1990656
# HeapObjects = 30038
# Stack = 557056 / 557056
# MSpan = 88640 / 114240
# MCache = 2336 / 16352
# BuckHashSys = 1458054
# GCSys = 7408288
# OtherSys = 625234
# NextGC = 4347552
# LastGC = 1721701961368527200
# PauseNs = [0 0 0 0 0 0 0 0 0 0 0 587200 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
记录每次gc暂停的时间(纳秒)，最多记录256个最新记录。
# PauseEnd = [1721701505980590100 1721701521107899000 1721701545266892700 1721701576201153500 1721701614559995900 1721701653732021700 1721701692952325700 1721701732148721500 1721701768330408100 1721701804527830300 1721701843763110400 1721701882963726200 1721701922182434900 1721701961368527200 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
# NumGC = 14    记录gc发生的次数。
# NumForcedGC = 0
# GCCPUFraction = 1.8095937518613244e-05
# DebugGC = false
# MaxRSS = 20221952


上面的信息很全，但是看起来不是很易懂，借助一下工具来看看
go tool pprof -inuse_space http://192.168.2.141:8222/debug/pprof/heap命令连接到进程中
查看正在使用的一些内存相关信息，此时我们得到一个可以交互的命令行。

在交互命令行里

比如；

比如我输入一个top，可以看数据top10来查看正在使用的对象较多的10个函数入口
(pprof) top
Showing nodes accounting for 2703.34kB, 100% of 2703.34kB total
Showing top 10 nodes out of 13
      flat  flat%   sum%        cum   cum%
 1541.26kB 57.01% 57.01%  1541.26kB 57.01%  regexp/syntax.(*compiler).inst (inline)
  600.58kB 22.22% 79.23%  1628.53kB 60.24%  github.com/go-playground/validator/v10.init
  561.50kB 20.77%   100%   561.50kB 20.77%  golang.org/x/net/html.init
         0     0%   100%   513.31kB 18.99%  internal/profile.init
         0     0%   100%  1541.26kB 57.01%  regexp.Compile (inline)
         0     0%   100%  1541.26kB 57.01%  regexp.MustCompile
         0     0%   100%  1541.26kB 57.01%  regexp.compile
         0     0%   100%   513.31kB 18.99%  regexp/syntax.(*compiler).cap (inline)
         0     0%   100%  1026.63kB 37.98%  regexp/syntax.(*compiler).compile
         0     0%   100%   513.31kB 18.99%  regexp/syntax.(*compiler).rune


Flat：函数在样本中处于运行状态的次数或者时间。简单来说就是函数出现在栈顶的次数，而函数在栈顶则意味着它在使用CPU。
Flat%：Flat / Total。
Sum%：自己以及所有前面的Flat%的累积值。解读方式：表中第3行Sum% 32.4%，意思是前3个函数（运行状态）的计数占了总样本数的32.4%
Cum：函数在样本中出现的次数。只要这个函数出现在栈中那么就算进去，这个和Flat不同（必须是栈顶才能算进去）。也可以解读为这个函数的调用次数。
Cum%：Cum / Total

可以看到一些累积分配的内存较多的函数， 此时我们就可以review代码，如何减少这些相关的调用，或者优化相关代码逻辑。


当我们不明确这些调用时是被哪些函数引起的时，我们可以输入top -cum来查找
-cum的意思就是，将函数调用关系 中的数据进行累积，
比如A函数调用的B函数，则B函数中的内存分配量也会累积到A上面，这样就可以很容易的找出调用链。
top -cum
(pprof) top -cum
Showing nodes accounting for 2703.34kB, 100% of 2703.34kB total
Showing top 10 nodes out of 13
      flat  flat%   sum%        cum   cum%
         0     0%     0%  2703.34kB   100%  runtime.doInit
         0     0%     0%  2703.34kB   100%  runtime.main
  600.58kB 22.22% 22.22%  1628.53kB 60.24%  github.com/go-playground/validator/v10.init
         0     0% 22.22%  1541.26kB 57.01%  regexp.Compile (inline)
         0     0% 22.22%  1541.26kB 57.01%  regexp.MustCompile
         0     0% 22.22%  1541.26kB 57.01%  regexp.compile
 1541.26kB 57.01% 79.23%  1541.26kB 57.01%  regexp/syntax.(*compiler).inst (inline)
         0     0% 79.23%  1541.26kB 57.01%  regexp/syntax.Compile
         0     0% 79.23%  1026.63kB 37.98%  regexp/syntax.(*compiler).compile
  561.50kB 20.77%   100%   561.50kB 20.77%  golang.org/x/net/html.init

输出结果还是不太直观，但是还是可以看出一点东西的，如果知道调用关系，就知道在每个函数里占了多少内存了

（字段的意思可以参考 https://dbwu.tech/posts/golang_pprof/）

为了更直观，可以直接看调用图：
************************************************************************************************
go tool pprof -alloc_space -cum -svg http://192.168.2.141:8222/debug/pprof/heap > heap.svg
************************************************************************************************
会生成一个调用图

从方块指出的 箭头上的数字为该函数调用的其他函数累积分配的内存
比如 A--->10000---->B
就是A要调B，而B上已累计内存10000


命令的参数有两种
--inuse/alloc_space
--inuse/alloc_objects
通常情况下：
用--alloc_space来分析程序常驻内存的占用情况;
用--alloc_objects来分析内存的临时分配情况，可以提高程序的运行速度。


生成调用图还有一些其他的命令:
curl -o heap.out http://192.168.33.57:8080/debug/pprof/heap
go tool pprof heap.out
然后输入 png
用浏览器打开png可以看到调用图

我想下载分析报告怎么下载？
下载heap报告:
curl -o a.prof http://192.168.33.57:8080/debug/pprof/heap

怎么查看下载的分析报告？
go tool pprof -http=:8080 a.prof
打开浏览器8080端口即可看到


如果想知道需要优化的是代码的哪一部分，可以在页面查看peek
输出结果会显示调用函数和被调用函数，调用顺序从上到下。
也可以直接查看source的结果，会把源代码里相关部分的分析结果展示给你


对于cpu分析，有没有可能直接获取到分析结果？
go tool pprof -http 127.0.0.1:8848 http://192.168.33.57:8080/debug/pprof/profile?seconds=30



# 排查锁的争用
go tool pprof http://192.168.33.57:8080/debug/pprof/mutex
# 排查阻塞操作
go tool pprof http://192.168.33.57:8080/debug/pprof/block



######################################### go-torch #########################################

火焰图是一种比较直观的分析方式

- 安装
cd /mnt/d/dev/php/magook/trunk/server
git clone http://github.com/brendangregg/FlameGraph.git
cd FlameGraph
cp flamegraph.pl /usr/local/bin
flamegraph.pl -h

USAGE: ./flamegraph.pl [options] infile > outfile.svg
        比如
        ./flamegraph.pl --title="Flame Graph: malloc()" trace.txt > graph.svg
安装torch
go get -u github.com/uber/go-torch



go-torch -alloc_space http://192.168.2.141:8222/debug/pprof/heap --colors=mem
go-torch -inuse_space http://192.168.2.141:8222/debug/pprof/heap --colors=mem


得到的图，就像一个山脉的截面图，从下而上是每个函数的调用栈，因此山的高度跟函数 调用的深度正相关，
而山的宽度跟使用或分配内存的数量成正比。

我们需要留意那些宽而平的山顶，这些部分通常是我们 需要优化的地方。

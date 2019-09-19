#Golang CLI selpg

>__参考：__
>- [开发 Linux 命令行实用程序](https://www.ibm.com/developerworks/cn/linux/shell/clutil/index.html)
>- https://github.com/cyulei/golang_selpg
>- [CLI 命令行实用程序开发基础](https://pmlpml.github.io/ServiceComputingOnCloud/ex-cli-basic)
> <br/>
##测试：
使用的测试文件test.txt的内容为：
![](srcshot/测试文件.jpg)
#### selpg -s1 -e1 test.txt
测试基本参数
![](srcshot/1.jpg)
#### selpg -s1 -e1 < test.txt
测试<符
![](srcshot/2.jpg)
#### cat test.txt | selpg -s1 -e1
测试管道输入到selpg程序
![](srcshot/3.jpg)
#### selpg -s1 -e2 -l5 test.txt >output_file.txt
测试>符，将程序输出输出到output_file.txt文件中
![](srcshot/4.1.jpg)
![](srcshot/4.2.jpg)
output_file.txt
#### selpg -s10 -x > output_file.txt 2>error_file.txt
测试>符，将错误输出到error文件
![](srcshot/6.1.jpg)
![](srcshot/6.2.jpg)
error_file.txt
#### selpg -s1 -e1 test.txt | cat
测试管道传输到下一个CLI应用
![](srcshot/9.jpg)
#### selpg -s1 -e1 -f test.txt
测试-f
![](srcshot/换页符.jpg)
#### selpg -s1 -e1 test.txt -dcat
测试-d（类似管道），不用lp测试改用cat测试
![](srcshot/_d.jpg)


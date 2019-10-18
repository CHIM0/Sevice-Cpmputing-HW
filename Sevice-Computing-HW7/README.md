# Agenda

## 概览
使用cobra作为框架实现的会议管理命令行程序(阉割), 实现的功能:
- 用户注册
- 用户登陆
- 用户登出    
## 使用说明
1. 将仓库`clone`到本地,使用`go build`安装依赖(用go mod 管理环境)
2. 使用`go build -o Agenda.exe` 编译命令行程序Agenda

## 可用指令
  - 获得帮助: `./Agenda help`
  - 注册: `./Agenda register --user=USERNAME --password=PASSWORD --email=EMAIL --telephone=TELEPHONE` 
  - 登录: `./Agenda signin --user=USERNAME --password=PASSWORD`
  - 登出: `./Agenda signout `

## 用例
### 注册

<center><img src="screenshot/register.jpg"/><br>初次注册</center><br>
<center><img src="screenshot/registererr.jpg"/><br>注册信息缺失</center><br>
<center><img src="screenshot/registerexited.jpg"/><br>用户已经存在</center><br>

### 登录
<center><img src="screenshot/signin.jpg"/><br>登录成功</center><br>
<center><img src="screenshot/signinerr.jpg"/><br>登录信息缺失</center><br>
<center><img src="screenshot/signinrepeat.jpg"/><br>重复登录当前用户</center><br>
<center><img src="screenshot/signinrepeat2.jpg"/><br>当前已有其他用户登录</center><br>
<center><img src="screenshot/signinwrong.jpg"/>用户名或密码错误<br></center><br>

### 登出
<center><img src="screenshot/signout.jpg"/>登出</center>
<center><img src="screenshot/signouterr.jpg"/>当前无用户登入，无法登出</center>

REM 使用方法：
REM gen_model.bat usercenter user
REM gen_model.bat usercenter user_auth
REM 再将 genModel 下的文件剪切到对应服务的 model 目录里面，记得改 package

REM 生成的表名
@echo off
chcp 65001 > nul
set tables=%2
REM 表生成的 genmodel 目录
set modeldir=.\%3\model

REM 数据库配置
set host=8.222.131.221
set port=3306
set dbname=buzzbox_%1
set username=root
set passwd=akita

echo 开始创建库：%dbname% 的表：%tables%
goctl model mysql datasource -url="%username%:%passwd%@tcp(%host%:%port%)/%dbname%" -table="%tables%" -dir="%modeldir%" -cache=true --style=goZero
pause
REM 使用方法：
REM gen_model.bat  数据库名 表名 生成目录
REM 例如：gen_model.bat  user user user
REM 生成的表名

@echo off
chcp 65001 > nul
set tables=%2

rem 设置颜色为黄色
color 09

REM 表生成的 genmodel 目录
set modeldir=D:\AkitaCode\Go\project\BuzzBox\service\%3\model

REM 数据库配置
set host=127.0.0.1
set port=3306
set dbname=buzzbox_%1
set username=root
set passwd=root

echo 开始创建库：%dbname% 的表：%tables%

if not "%2"=="" (
    goctl model mysql datasource -url="%username%:%passwd%@tcp(%host%:%port%)/%dbname%" -table="%tables%" -dir="%modeldir%" -cache=true --style=goZero
    echo "生成%2 model成功"
) else (
    echo "参数错误，请按照以下格式输入："
    echo "用法示例："
    echo "gen_model.bat 数据库名 表名 生成目录"
    echo "例如：gen_model.bat  user user user"
    pause
)
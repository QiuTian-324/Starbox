@echo off
chcp 65001 > nul

rem 设置颜色为白色
color 09

rem 清屏
cls

echo ********************************************
echo *                                            *
echo *                BuzzBox                      *
echo *                                            *
echo ********************************************
echo.

echo 正在生成 API 文件...

rem 检查是否传递了参数，如果没有传递则提示并退出
if "%1"=="" (
    echo 请指定要生成的API路径作为参数（例如：用户API传入user）作为参数。
    pause
    exit /b 1
)

rem 切换到指定目录
cd /d D:\AkitaCode\Go\project\BuzzBox\service\%1

for %%i in (api\*.api) do (
    echo 正在处理 %%i...
    echo 处理 %%i...
    goctl api go -api %%i -dir .\ --style=goZero

    rem 检查命令执行的返回值
    if errorlevel 1 (
        echo 生成 %%i 时出错！
        rem 返回根目录
        cd /d D:\AkitaCode\Go\project\BuzzBox
        pause
        exit /b 1
    )
)

echo API 文件生成完成。

rem 返回根目录
cd /d D:\AkitaCode\Go\project\BuzzBox

pause
exit /b 0

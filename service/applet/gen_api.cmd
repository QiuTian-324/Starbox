@echo off
chcp 65001 > nul

rem 设置颜色
color 09

echo 正在生成 API 文件...

rem 切换到指定目录
cd /d D:\AkitaCode\Go\project\BuzzBox\service\applet

for %%i in (api\*.api) do (
    echo 正在处理 %%i...
    echo 处理 %%i...
    goctl api go -api %%i -dir .\ --style=goZero

    rem 检查命令执行的返回值
    if errorlevel 1 (
        echo 生成 %%i 时出错！
        pause
        exit /b 1
    )
)

echo API 文件生成完成。
pause
exit /b 0

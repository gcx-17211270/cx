@REM 用于windows上执行，复制文件到linux机器上

@echo "scp code to target machine"
@set TARGET_MACHINE=192.168.88.137
@set PASSWORD=17211270
@set USER=root
@set DIRECTORY=/root/code/cx
@set PARENT=/root/code

@REM 先配置免密登录 参考 https://zhuanlan.zhihu.com/p/401327519

@cd %cd%
ssh %USER%@%TARGET_MACHINE% 'rm -rf %DIRECTORY%'
@REM @if %errorlevel% neq 0 (
@REM     @echo "rm file failure"
@REM     exit 1
@REM )
@scp -r %cd%/../../cx %USER%@%TARGET_MACHINE%:%PARENT%
@if %errorlevel% neq 0 (
    @echo "scp file failure"
    exit 1
)
@echo "prepare copy code success"

ssh %USER%@%TARGET_MACHINE% 'bash %DIRECTORY%/sh/2.sh'

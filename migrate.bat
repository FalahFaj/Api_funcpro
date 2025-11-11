@echo off
setlocal enabledelayedexpansion

if not exist ".env" (
    echo Error: File .env tidak ditemukan!
    exit /b 1
)

REM
for /f "usebackq tokens=1,2 delims==" %%a in (".env") do (
    set "%%a=%%b"
)

set DSN=host=%DB_HOST% port=%DB_PORT% user=%DB_USER% password=%DB_PASSWORD% dbname=%DB_NAME% sslmode=%DB_SSL_MODE%

goose -dir migrations postgres "%DSN%" %*
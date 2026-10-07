param (
    [string]$Version = "v1.0.0"
)

$ErrorActionPreference = "Stop"
$BuildDate = (Get-Date).ToString("yyyy-MM-ddTHH:mm:ssZ")
$GitCommit = (git rev-parse --short HEAD 2>$null)
if (-not $GitCommit) { $GitCommit = "release" }

$OutputDir = "dist"
if (Test-Path $OutputDir) { Remove-Item -Recurse -Force $OutputDir }
New-Item -ItemType Directory -Path $OutputDir | Out-Null

$LdFlags = "-s -w -X main.Version=$Version -X main.BuildDate=$BuildDate -X main.GitCommit=$GitCommit"

Write-Host "Building go-insta-cli releases ($Version)..." -ForegroundColor Cyan

# Windows amd64
$env:GOOS = "windows"
$env:GOARCH = "amd64"
Write-Host "Compiling Windows (amd64)..."
go build -ldflags $LdFlags -o "$OutputDir/insta.exe" .

# Linux amd64
$env:GOOS = "linux"
$env:GOARCH = "amd64"
Write-Host "Compiling Linux (amd64)..."
go build -ldflags $LdFlags -o "$OutputDir/insta-linux" .

# macOS amd64 (Intel)
$env:GOOS = "darwin"
$env:GOARCH = "amd64"
Write-Host "Compiling macOS (amd64)..."
go build -ldflags $LdFlags -o "$OutputDir/insta-darwin-amd64" .

# macOS arm64 (Apple Silicon M1/M2/M3)
$env:GOOS = "darwin"
$env:GOARCH = "arm64"
Write-Host "Compiling macOS (arm64)..."
go build -ldflags $LdFlags -o "$OutputDir/insta-darwin-arm64" .

# Reset env vars
$env:GOOS = ""
$env:GOARCH = ""

# Generate SHA-256 Checksums
Write-Host "Generating SHA-256 Checksums..." -ForegroundColor Yellow
$ChecksumFile = "$OutputDir/checksums.txt"
Get-ChildItem -Path $OutputDir -File | Where-Object { $_.Name -ne "checksums.txt" } | ForEach-Object {
    $hash = (Get-FileHash -Path $_.FullName -Algorithm SHA256).Hash.ToLower()
    "$hash  $($_.Name)" | Out-File -FilePath $ChecksumFile -Append -Encoding utf8
}

Write-Host "Build complete! Artifacts stored in '$OutputDir/'" -ForegroundColor Green

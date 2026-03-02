$scriptDir = Split-Path -Parent $MyInvocation.MyCommand.Path
Set-Location $scriptDir

function Invoke-GoBenchmark {
    param(
        [Parameter(Mandatory = $true)]
        [string]$Name,
        [Parameter(Mandatory = $true)]
        [string]$Pattern
    )

    Write-Host "`n=== $Name ===" -ForegroundColor Cyan
    $output = go test ./boatrace -bench $Pattern -benchmem -benchtime=1x -run '^$' 2>&1
    $exitCode = $LASTEXITCODE

    $output | ForEach-Object { Write-Host $_ }

    $metricsLine = $output | Where-Object { $_ -match '^Benchmark\S+' } | Select-Object -Last 1
    $metrics = $null

    if ($metricsLine -and $metricsLine -match '^(Benchmark\S+)\s+\d+\s+([0-9\.]+)\s+ns/op\s+([0-9\.]+)\s+B/op\s+([0-9\.]+)\s+allocs/op') {
        $metrics = [PSCustomObject]@{
            Benchmark   = $matches[1]
            NsPerOp     = [double]$matches[2]
            BytesPerOp  = [double]$matches[3]
            AllocsPerOp = [double]$matches[4]
        }
    }

    return [PSCustomObject]@{
        Name     = $Name
        Pattern  = $Pattern
        ExitCode = $exitCode
        Metrics  = $metrics
    }
}

$results = @()
$results += Invoke-GoBenchmark -Name 'NumberOfWaysToWin (small race)' -Pattern 'BenchmarkNumberOfWaysToWin_SmallRace$'
$results += Invoke-GoBenchmark -Name 'RecordBreakingsProducts (test input)' -Pattern 'BenchmarkRecordBreakingsProducts_InputTest$'
$results += Invoke-GoBenchmark -Name 'RecordBreaking (test input)' -Pattern 'BenchmarkRecordBreaking_InputTest$'

Write-Host "`n=== Summary ===" -ForegroundColor Yellow
foreach ($result in $results) {
    if ($result.ExitCode -eq 0 -and $null -ne $result.Metrics) {
        Write-Host ("{0}: {1} ns/op, {2} B/op, {3} allocs/op" -f $result.Name, $result.Metrics.NsPerOp, $result.Metrics.BytesPerOp, $result.Metrics.AllocsPerOp)
    }
    elseif ($result.ExitCode -eq 0) {
        Write-Host ("{0}: completed, metrics not parsed" -f $result.Name)
    }
    else {
        Write-Host ("{0}: failed (exit code {1})" -f $result.Name, $result.ExitCode) -ForegroundColor Red
    }
}

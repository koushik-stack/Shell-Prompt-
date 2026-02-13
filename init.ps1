# ShellPrompt initialization for PowerShell

$PromptBinDir = Split-Path -Parent $PSCommandPath
$PromptBin = Join-Path $PromptBinDir "prompt"

function prompt {
    & $PromptBin pwsh
}

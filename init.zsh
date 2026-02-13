#!/bin/zsh
# ShellPrompt initialization for zsh

# Get the directory where this script is located
PROMPT_BIN_DIR="$(cd "$(dirname "${(%):-%x}")" && pwd)"
PROMPT_BIN="$PROMPT_BIN_DIR/prompt"

# Set the PS1 prompt
setopt PROMPT_SUBST
export PS1='$("$PROMPT_BIN" zsh)'

# Optional: Set secondary prompts
export PS2='> '
export PS3='select: '
export PS4='+ '

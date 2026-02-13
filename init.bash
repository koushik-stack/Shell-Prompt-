#!/bin/bash
# ShellPrompt initialization for bash

# Get the directory where this script is located
PROMPT_BIN_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROMPT_BIN="$PROMPT_BIN_DIR/prompt"

# Set the PS1 prompt
export PS1='$("$PROMPT_BIN" bash)'

# Optional: Set secondary prompts
export PS2='> '
export PS3='select: '
export PS4='+ '

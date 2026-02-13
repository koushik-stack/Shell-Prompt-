#!/usr/bin/env fish
# ShellPrompt initialization for fish

# Get the directory where this script is located
set PROMPT_BIN_DIR (dirname (status filename))
set PROMPT_BIN "$PROMPT_BIN_DIR/prompt"

# Set the PS1 prompt
function fish_prompt
    "$PROMPT_BIN" fish
end

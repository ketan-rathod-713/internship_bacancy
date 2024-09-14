#!/bin/bash

SESSION_NAME="mygamesession"
STARTUP_COMMAND="echo 'startup command'"

# Check if the tmux session is running
tmux has-session -t $SESSION_NAME 2>/dev/null

if [ $? -eq 0 ]; then
    echo "Tmux session '$SESSION_NAME' is running."

    # Check if a Java process is running inside the session
    if tmux list-panes -t $SESSION_NAME -F '#{pane_pid}' | xargs ps -p | grep -q java; then
        echo "Java process found in tmux session. Attaching..."
        tmux attach-session -t $SESSION_NAME
    else
        echo "No Java process found. Creating a new session..."
        tmux kill-session -t $SESSION_NAME
        tmux new-session -d -s $SESSION_NAME "$STARTUP_COMMAND"
        tmux attach-session -t $SESSION_NAME
    fi
else
    echo "No tmux session found. Creating a new one..."
    tmux new-session -d -s $SESSION_NAME "$STARTUP_COMMAND"
    tmux attach-session -t $SESSION_NAME
fi

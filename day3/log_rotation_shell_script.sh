#!/bin/bash
# ============================================================
# Author: Parmeet Singh
# Date: 10-jan-2026
# Demo Real-Time Log Rotation & Cleanup Scheduler
# Generates fake logs and rotates them quickly for demonstration.
# ============================================================

set -x

LOG_DIR="$HOME/demo_logs"
ACTIVE_LOG="app.log"
MAX_LOG_SIZE_MB=1       # rotate after ~1MB
RETENTION_DAYS=1        # cleanup logs older than 1 day
CHECK_INTERVAL=2        # check every 2 seconds for fast demo

# -------------------------
# INIT
# -------------------------
mkdir -p "$LOG_DIR"
touch "$LOG_DIR/$ACTIVE_LOG"

echo "Starting Demo Log Rotation Scheduler..."
echo "Logs will be generated in $LOG_DIR every second."
echo "Monitoring $LOG_DIR every $CHECK_INTERVAL seconds."
echo

# -------------------------
# FUNCTIONS
# -------------------------

rotate_log() {
    TIMESTAMP=$(date +"%Y%m%d_%H%M%S")
    ROTATED="$LOG_DIR/$ACTIVE_LOG.$TIMESTAMP"
    echo "[$(date)] Rotating log: $ACTIVE_LOG → $ROTATED"
    mv "$LOG_DIR/$ACTIVE_LOG" "$ROTATED"
    touch "$LOG_DIR/$ACTIVE_LOG"
    gzip "$ROTATED"
}

cleanup_logs() {
    find "$LOG_DIR" -name "*.gz" -mtime +"$RETENTION_DAYS" -print -delete
}

# -------------------------
# DEMO LOOP
# -------------------------
while true; do
    # Append fake log lines
    for i in {1..100}; do
        echo "This is a demo log line $(date)" >> "$LOG_DIR/$ACTIVE_LOG"
    done

    # Rotate if log exceeds MAX_LOG_SIZE_MB
    LOG_SIZE_MB=$(du -m "$LOG_DIR/$ACTIVE_LOG" | cut -f1)
    if [ "$LOG_SIZE_MB" -ge "$MAX_LOG_SIZE_MB" ]; then
        rotate_log
    fi

    # Cleanup old logs
    cleanup_logs

    # Show current files
    echo "Current files in $LOG_DIR:"
    ls -lh "$LOG_DIR"
    echo "--------------------------------------------"

    sleep "$CHECK_INTERVAL"
done

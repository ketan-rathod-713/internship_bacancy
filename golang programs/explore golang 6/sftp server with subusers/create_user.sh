#!/bin/bash

if [ "$#" -ne 3 ]; then
  echo "Usage: $0 <username> <password> <directory_name>"
  exit 1
fi

USERNAME=$1
PASSWORD=$2
DIRECTORY_NAME=$3

# Check if the shared directory exists
if [ ! -d "/data/incoming" ]; then
  echo "Directory /data/incoming does not exist."
  exit 1
fi

echo "Given directory exists: /data/incoming"

# Create the user without a home directory and with no shell access
useradd -M -s /sbin/nologin $USERNAME
echo "$USERNAME:$PASSWORD" | chpasswd

echo "User added for SFTP: $USERNAME with password $PASSWORD"

# Ensure the shared directory has the correct permissions
chown root:root "/data/incoming"
chmod 755 "/data/incoming"

echo "incoming directory created at /data/incoming and ownership granted to $USERNAME"

# Add an entry in /etc/ssh/sshd_config for ChrootDirectory
cat <<EOL >> /etc/ssh/sshd_config

Match User $USERNAME
    ChrootDirectory /data/incoming
    ForceCommand internal-sftp
    AllowTcpForwarding no
    X11Forwarding no
EOL

# # Restart SSH service to apply changes
systemctl restart sshd

echo "User $USERNAME configured with SFTP access, chrooted to /data/incoming"
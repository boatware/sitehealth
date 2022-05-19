#!/bin/bash

if [[ "$EUID" -ne 0 ]]; then
  echo "Please run as root"
  exit 1
fi

if [[ ! -f "./sitehealth" ]]; then
  echo "Please build the app first."
  exit 1
fi

if [[ -f "/usr/bin/sitehealth" ]]; then
  rm /usr/bin/sitehealth
fi

cp ./sitehealth /usr/bin/sitehealth && echo "Installed sitehealth successfully"
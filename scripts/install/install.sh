#!/bin/bash

# ==========================================================================
# インストールスクリプト
# ==========================================================================
echo "[INFO] Installing packages..."

# パッケージのインストール
sudo apt-get update
sudo apt-get install -y \
  docker \
  docker-compose \
  cron

# crontabの編集
echo "0 1 2 * * /home/appuser/projects/update_cert.sh" 2>&1 \
  > /var/spool/cron/crontabs/appuser

#   サービスの起動
sudo systemctl start cron
sudo systemctl enable cron

echo "[INFO] Installation completed."
#!/bin/bash

# ==========================================================================
# Let's Encrypt証明書更新スクリプト
# ==========================================================================

# 設定
CERTBOT_VOLUME_CERTBOT="./nginx/certbot"      # Let's Encrypt用ディレクトリ
CERTBOT_VOLUME_SSL="./nginx/ssl"              # SSL証明書ディレクトリ
NGINX_CONTAINER_NAME="nginx"                  # Nginxコンテナ名
CERTBOT_IMAGE="certbot/certbot"               # CertbotのDockerイメージ

# ==========================================================================
# 証明書更新
# ==========================================================================

echo "[INFO] Updating Let's Encrypt certificates..."

docker run --rm \
    -v "$CERTBOT_VOLUME_CERTBOT:/var/www/certbot" \
    -v "$CERTBOT_VOLUME_SSL:/etc/letsencrypt" \
    "$CERTBOT_IMAGE" renew

if [ $? -eq 0 ]; then
    echo "[INFO] Certificate renewal succeeded."

    # Nginxのリロード
    echo "[INFO] Reloading Nginx..."
    docker exec -it "$NGINX_CONTAINER_NAME" nginx -s reload

    if [ $? -eq 0 ]; then
        echo "[INFO] Nginx reloaded successfully."
    else
        echo "[ERROR] Failed to reload Nginx."
        exit 2
    fi
else
    echo "[ERROR] Certificate renewal failed."
    exit 1
fi

# 完了
echo "[INFO] Let's Encrypt certificate renewal process completed."
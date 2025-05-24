#!/bin/sh

PASSFILE="/mosquitto/config/passwordfile"
USER=${MQTT_USER:-admin}
PASS=${MQTT_PASSWORD:-admin}

if [ ! -f "$PASSFILE" ]; then
  echo "🛠 Генерация passwordfile..."
  mosquitto_passwd -b -c "$PASSFILE" "$USER" "$PASS"
else
  echo "✅ passwordfile уже существует, пропускаем генерацию."
fi

chmod 644 "$PASSFILE"
chown mosquitto:mosquitto "$PASSFILE" || true

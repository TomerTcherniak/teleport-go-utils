#!/usr/bin/env bash
cp teleport-go-utils.service /etc/systemd/system/teleport-go-utils.service
systemctl daemon-reload
systemctl enable teleport-go-utils
systemctl start teleport-go-utils

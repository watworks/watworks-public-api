#!/bin/sh

set -e

# check for config file if specified
if [ -z "$APP_CONF_PATH"]; then
	# if the file path was specified, but not found - exit immediately, this
	# suggests that something is misconfigured
	if [ ! -f "$APP_CONF_PATH"]; then
		echo "App config path ($APP_CONF_PATH): file not found"
		exit 1;
	fi
# otherwise generate the config file from env vars
else
	cat <<EOF > /app/app.config.yml
counterServiceUrl: "${COUNTER_SERVICE_URL:='foo-bar-baz'}"
EOF

	# set the config filepath, which the app will try to read
	APP_CONF_PATH=/app/app.config.yml
fi

exec /app/watworks-public-api
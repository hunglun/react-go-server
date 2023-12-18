screen -XS https_webserver quit
screen -XS http_webserver quit

screen -dmS https_webserver sudo PORT=443 REACT=../build ./bin/x86_webserver
screen -dmS http_webserver sudo PORT=80 REACT=../build ./bin/x86_webserver

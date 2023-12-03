screen -XS webserver quit
screen -dmS webserver sudo PORT=80 REACT=../build ./bin/x86_webserver

#!/usr/bin/bash
# Get table of known ports from wikipedia via
# https://www.wikitable2json.com/#/API/GetByPage

curl -X 'GET' \
  'https://www.wikitable2json.com/api/List_of_TCP_and_UDP_port_numbers%23Well-known_ports?table=1&lang=en&cleanRef=false' \
  -H 'accept: application/json' \
  > known_ports.json

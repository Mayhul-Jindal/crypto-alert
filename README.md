- [Crypto Alert System](#crypto-alert-system)
  - [Setup](#setup)
  - [System design](#system-design)
- [todos](#todos)
- [dev notes](#dev-notes)

# Crypto Alert System


## Setup

## System design


# todos
- [ ] API --> alert/create read(status, paginated, filter) update delete 
- [ ] Binance websocket go
- [ ] Email using SMTP
- [ ] Redis for cache(fetch all alert)
- [ ] JWT token
  - [ ] Add validation
- [ ] Postgres as a database
- [ ] Kafka as a messqge broker for the task to send email
- [ ] docker-compose

users alerts are accepted even if cryptoWatcher is down
# dev notes
- [X] priority queue ??!
- [X] btrees ?!?
- [X] sorted set redis
- [ ] creating a bucket by rounding vs simple query (test btrees capabilities)
- [ ] kafka
  - [ ] `topic` stores messages
  - [ ] topic can be devided into `partitions` to increase availability and throughput
  - [ ] replication factor decided wether the message in topic is replicated to multiple partitions or not




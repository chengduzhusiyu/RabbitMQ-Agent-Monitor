# RabbitMQ-Agent-Monitor

[![Go Report Card](https://goreportcard.com/badge/github.com/chengduzhusiyu/RabbitMQ-Agent-Monitor)](https://goreportcard.com/report/github.com/chengduzhusiyu/RabbitMQ-Agent-Monitor)
![GoCI](http://goci.ele.me/na/goci/eleme/goci/badge?type=job)
[![Build Status](https://travis-ci.org/chengduzhusiyu/RabbitMQ-Agent-Monitor.svg?branch=master)](https://travis-ci.org/chengduzhusiyu/RabbitMQ-Agent-Monitor)
[![Apache 2 licensed](https://img.shields.io/badge/license-Apache2-blue.svg)](https://raw.githubusercontent.com/oklog/run/master/LICENSE)

RabbitMQ-Agent-Monitor is an agent that used for [open-falcon](http://open-falcon.org/) to monitoring [RabbitMQ](https://www.rabbitmq.com/).

## Arch Requirement
Linux

## Build

```bash
$make build
```

## Agent launch

```bash
$/bin/bash control.sh start|stop|restart
```
It will create a temporary directory `var` in your current path.

## Metrics

***overview metrics*
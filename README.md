# Show And Tell Scheduler

[![Build Status][badge_build_status]][link_build_status]
[![Release Status][badge_release_status]][link_build_status]
[![Repo size][badge_repo_size]][link_repo]

## Overview
A simple Go webapp for demonstrating Kubernetes features.

## Features
- CRUD operations for sessions (time, team, topic, presenter)
- PostgreSQL database backend
- Docker Compose setup with migrations
- Kubernetes manifests with CloudNativePG PostgreSQL database
- 'panic' and 'memory' allocation endpoints (to trigger auto-scaling/ self-healing)

## Endpoints

- `/` - Demo application
- `/health` - Health check endpoint for liveness/readiness probes
- `/demo/panic` - Triggers a panic to demonstrate pod failover
- `/demo/memory` - Allocates memory incrementally until OOM occurs (refresh repeatedly to trigger autoscaling)

## Running Locally

```bash
docker-compose up -d
```

Access at http://localhost:8080


[link_issues]:https://github.com/braveokafor/show-and-tell/issues
[link_pulls]:https://github.com/braveokafor/show-and-tell/pulls
[link_build_status]:https://github.com/braveokafor/show-and-tell/actions/workflows/ci.yaml
[link_build_status]:https://github.com/braveokafor/show-and-tell/actions/workflows/release.yaml
[link_repo]:https://github.com/braveokafor/show-and-tell

[badge_issues]:https://img.shields.io/github/issues-raw/braveokafor/show-and-tell?style=flat-square&logo=GitHub
[badge_pulls]:https://img.shields.io/github/issues-pr/braveokafor/show-and-tell?style=flat-square&logo=GitHub
[badge_build_status]:https://img.shields.io/github/actions/workflow/status/braveokafor/show-and-tell/ci.yaml?style=flat-square&logo=GitHub&label=build
[badge_release_status]:https://img.shields.io/github/actions/workflow/status/braveokafor/show-and-tell/release.yaml?style=flat-square&logo=GitHub&label=release
[badge_repo_size]:https://img.shields.io/github/repo-size/braveokafor/show-and-tell?style=flat-square&logo=GitHub

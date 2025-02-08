# Template repository for deploy

[![Release](https://github.com/variegate-app/deploy_tpl/actions/workflows/release.yml/badge.svg)](https://github.com/variegate-app/deploy_tpl/actions/workflows/release.yml)
[![Lint & test](https://github.com/variegate-app/deploy_tpl/actions/workflows/lint.yml/badge.svg)](https://github.com/variegate-app/deploy_tpl/actions/workflows/lint.yml)
[![CI](https://github.com/variegate-app/deploy_tpl/actions/workflows/ci.yml/badge.svg)](https://github.com/variegate-app/deploy_tpl/actions/workflows/ci.yml)

build & deploy go applications via github actions & docker-compose context

Create secrets
- secrets.CI_TOKEN
- secrets.SSH_USER
- secrets.SSH_PRIVATE_KEY
- secrets.SSH_HOST
---
title: "5-Minute Quickstart"
date: 2026-08-11
prev_title: "Installation Track"
prev_link: "/docs/installation/"
next_title: "Manifest Specification"
next_link: "/docs/manifest/"
---

## Managing Your First Binary

Lown installs toolchains directly from GitHub shorthand or custom Git repository URLs.

### 1. Install Revoq, Muth, and Trig

Install the core ecosystem using short aliases:

```bash
lown install revoq
lown install muth
lown install trig
```

### 2. Verify Installed Binaries

Check your local environment using `lown list` and `lown doctor`:

```bash
lown list
lown doctor
```

### 3. Sync Toolchains Across Projects

Pull updates for all managed tools:

```bash
lown sync
```

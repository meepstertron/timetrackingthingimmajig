# How it works/Mechanism ![](https://img.shields.io/badge/Docs-FFF.svg)
*Last updated 19.11.2026 @ 10:05PM CET*

This document is meant to outline the basic definition on how the time tracking tool developed in this repository works at a high level.

> [!NOTE]
> This file is still actively changing!

<hr>

## Capture multiple signals not just one

| Signal  | What it looks at  | Source  | Trust  |
|---|---|---|---|
|   Filesystem|   Changes to files in a designated project (ignoring deps and `.gitignore`)|   OS File watcher|   Low, anything can trip this|
|   Presence|   Keystroke count, KPS/burstiness, focused window|   Global Keyboard Events, window focus|   |
|   Harness Events|   |   |   |

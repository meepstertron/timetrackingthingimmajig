# Spec V2 ![](https://img.shields.io/badge/Docs-FFF.svg)
*Last updated 19.11.2026 @ 9:00PM CET*

This document is meant to outline the basic definition on how the time tracking tool developed in this repository works at a high level.

> [!NOTE] 
> This file is still actively changing! 

<hr>

## Definitons

 - **Session**: A Single, Contiguous block of time spent working on a project 
 - **Weighted Project (aka WP)**: A Weighted Project  is 10 approved hours of "productive struggle" done by a Member of hackclub on a technical project which ended up in the YSWS DB.
 - **Productive Struggle (aka PS)**: Time spent working outside your comfort zone. Time where you're actually learning. (TOO AMBIGUOUS)
 - **Submitted Hour**: One Hour calculated by the time tracking software submitted to a YSWS program
 - **YSWS**: You Ship We Ship, a program that measures (PS) hours spent on a relevant topic and issues prizes based on WG
 - **Weighted Grants (WG)**: Older definition of 10H of shipped time. Has moved onto Weighted project. One WG is worth roughly 85USD in program budget.
 - **Project**: 
 - **Hack Club**: Non-Profit based in Burlington, VT focused on Incentivising Teenagers (13-19) to code.

## Requirements

 - Needs to run on Less than 500USD/mo compute budget
 - handle ~ 15k MAU
 - Respect user's privacy by collecting as little personal information as possible
 - Gain 1000 weighted projects (which is around 10k Hours if we assume each project is 10H min)

| Which Requirement  | What are we doing?  | implemented?  |
|---|---|---|
|   1. Open Source|  Is licensed under AGPLv3|   yes|
|   2. Accurately track Hours done with  AI|   |   |
|   3. Needs to be Motivating|   |   |
|   4. Needs to be Passive|   |   |
|   5. Centralized server /w accounts|   |   |
|   6. Privacy Friendly|   |   |
|   7. Support Major AI Harnesses|   |   |
|   8. Support all 3 OSsses|   Use Go to cross compile the client and use OS agnostic Libraries/Implement behavior for the big 3|   Partially|
|   9. Trustworthy / Accurately track PS|   |   |
|   10. Needs to be Auditable|   |   |
|   11. Fit on existing server.  |   |   |

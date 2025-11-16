# Welcome to NULL SKY, adventurer!
This is my most recent attempt at creating a usable client for the [SpaceTraders API](https://spacetraders.io/) MMO RPG. My goal with this project is simply to have a good time learning a new language and DB utility that I have zero prior experience with.

If you stumbled across this project looking for some good advice on how to write in Golang or to steal my trade secrets in-game you've come to the wrong place friend! I'm doing this mostly for the data collection and TUI building experience, I'd be lucky if my agent completes even a single contract per bi-weekly reset!

# What this project is:
- A passion project & academic exercise for me to learn Golang and PostgreSQL.
- A fun way for me to both justify playing a fun game and also build skills/tools at the same time.

# What this project is NOT:
- Intended for actual distribution...This is just a fun project for me to learn and grow.
- A point of reference for how programming should be done. 
- While I've been writing PHP and building tools for several years I'm still VERY new to these topics.

# In theory:
This project should be entirely portable and agent agnostic. TLDR: it CAN tie into anyone's agent data by just importing your own personal token during setup and off you go....in theory.

I wrote this on a RHEL based system with no intention of making it interoperable with other distros or kernels since this is really just for me.

# Milestones
- [ ] Setup
    - [-] Register a New Agent
    - [-] Simplify and/or automate new user setup
    - [x] Login with a Preexisting Token

- [ ] Game Server
    - [x] Capture current game server state _**(Stats/Resets/Leaderboards)**_
    - [-] Display game state + leaderboards

- [ ] Contracts
    - [x] View Contracts
    - [x] Accept Contracts
    - [x] Negotiate New Contracts
    - [ ] Deliver and Fulfil Contracts

- [ ] Fleet
    - [-] Display ship stats/status
    - [ ] Manage Ship Inventory  _**(Transfer/Jettison/Extract/Refine)**_
    - [ ] Manage Ship Status     _**(Dock/Orbit/Refuel)**_
    - [ ] Manage Ship Navigation _**(Navigate/Jump/Warp/Flight Mode)**_
    - [ ] Manage Ship Hardware   _**(Install/Remove Mounts)**_

- [ ] Explore
    - [ ] View Systems and Waypoints
    - [ ] Use Jump Gates

- [ ] Economy
    - [ ] View Markets & Shipyards
    - [ ] Trade Goods & cargo
    - [ ] Purchase Ships
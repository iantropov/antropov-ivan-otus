# Architectural Kata

## Описание

https://nealford.com/katas/kata?id=GoingGoingGone

**Going...Going...Gone!**

An auction company wants to take their auctions online to a nationwide scale--customers choose the auction to participate in, wait until the auction begins, then bid as if they were there in the room, with the auctioneer

**Users:**

User scale up to hundreds of participants (per auction), potentially up to thousands of participants, and as many simultaneous auctions as possible

**Requirements:**

- auctions must be categorized and 'discoverable'
- auctions must be as real-time as possible
- auctions must be mixed live and online
- video stream of the action after the fact
- handle money exchange
- participants must be tracked via a reputation index

**Additional Context:**

- auction company is expanding aggressively by merging with smaller competitors
- if nationwide auction is a success, replicate the model overseas
- budget is not constrained--this is a strategic direction
- company just exited a lawsuit where they settled a suit alleging fraud

## Решение

https://stackoverflow.com/questions/17448061/how-many-system-resources-will-be-held-for-keeping-1-000-000-websocket-open
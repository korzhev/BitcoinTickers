# Demo Bitcoin
I've decided that User want to buy Bictoins using Eur, so Exchange course with
 MINIMAL value is shown.
 
# Go version
**In progress**
## Requirements
`cd ./go/`
`GOPATH = $GOPATH +./` 
Tested only on Go 1.8

# Node.js version
## Requirements
`cd ./node/`
node > 7.6

### Install
`npm i`

### Run
`npm start`

Data will be refreshed after each response from tickers.
`Ctrl+C` to terminate process, but it has "close" function in code too.
- [x] **Aggregate unlimited number of feeds.** Just add new to `config.json` 
and create specific parser
- [x] **Given every feed data expiration time, the solution must not use 
obsolete data in comparisons and calculations.** Set in `config.json`

Hints:
- **How responsive is the solution to the new price available in the feed?** - New data will be shown after getting new response from bitcoin ticker or 
exchange ticker.
- **Is it possible to attach multiple tick viewers?** - Yes
- **How does one of the broken feed affect ticker behaviour?** - It is shown 
that it is less active tickers than it could be. 
- **Do you understand the concept of Bid/Ask for the price?** -Yes
- **How old data is you final ticker BTC/EUR giving out to end user?** - It 
depends on intervals, how often is response got from ticker api. Lates 
available data is shown.
- **What is the theoretical limit of feed sources in your solution?** - Limited
 by OS max connections, hardware and Math.min function in js(~100K)
- **Is it possible to test the solution without real access to the feed 
sources?** - Yes
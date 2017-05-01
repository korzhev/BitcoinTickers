/**
 * update data in terminal
 * @param {string} str
 */
function print(str) {
    process.stdout.clearLine();
    process.stdout.cursorTo(0);
    process.stdout.write(str);
    // console.info(str);
}

/**
 * filter tickers with not expired time
 * @param {Array} tickers
 * @returns {Array}
 */
function useNotExpired(tickers) {
    const now = new Date();
    return tickers.filter(t => (now - t.lastSuccessRequestTime) < t.expireTime);
}

/**
 * get Ticker with lowest rate
 * @param {Array} BitcoinTickers
 * @returns {*}
 */
function getMinimumBitcoins(BitcoinTickers) {
    const bt = useNotExpired(BitcoinTickers);
    return bt.reduce((prev, current) => {
        const currentValue = current.lastData.eur;
        if (typeof currentValue !== 'number') {
            return prev;
        }
        return (prev.lastData.eur < current.lastData.eur) ? prev : current;
    }, { lastData: { eur: undefined } });
}

/**
 * get minimal rate
 * @param ExchangeTickers
 * @returns {Object}
 */
function getMinimumExchage(ExchangeTickers) {
    const et = useNotExpired(ExchangeTickers);
    return {
        'eur/usd': Math.min.apply(null, et.map(t => (
            typeof t.lastData !== 'number' ? 0 : t.lastData))),
    };
}

/**
 * require all tickers and start to watch
 * @param {Object} tickerOptions
 * @param {string} base
 * @param {function} onError
 * @param {function} onData
 * @returns {Object}
 */
function setTickers(tickerOptions, base, onError, onData) {
    const name = tickerOptions.name;
    /* eslint-disable */
    const Ticker = require(`./${base}/${name}`);
    const t = new Ticker(tickerOptions);
    /* eslint-enable */
    t.on('error', onError)
        .on('data', onData)
        .start();
    return t;
}

/**
 * stop tickers from list
 * @param {Array} tickerList
 */
function stopTickers(tickerList) {
    tickerList.forEach((t) => {
        t.stop();
    });
}

/**
 * is there any changes in state?
 * @param {Object} state
 * @param {Object} data
 * @returns {boolean}
 */
function checkState(state, data) {
    const aProps = Object.getOwnPropertyNames(state);
    for (let i = 0; i < aProps.length; i++) {
        const propName = aProps[i];
        if (state[propName] !== data[propName]) {
            return false;
        }
    }
    return true;
}

module.exports = {
    stopTickers,
    setTickers,
    getMinimumExchage,
    getMinimumBitcoins,
    print,
    checkState,
};

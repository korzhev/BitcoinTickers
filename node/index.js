'use strict';

const CONF = require('./config.json');
const {
    // stopTickers,
    // checkState,
    setTickers,
    getMinimumExchage,
    getMinimumBitcoins,
    print,
} = require('./helpers');

const BT_NUMBER = CONF.bitcoinTickers.length;
const ET_NUMBER = CONF.exchangeTickers.length;

let state = {
    usd: 0,
    eur: 0,
    'eur/usd': 0,
    bitcoinTickersNumber: 0,
    exchangeTickersNumber: 0,
};

/**
 * set new state and show result
 * @param {Object} data
 * @param {Object} tickerNumber
 */
function mergeData(data, tickerNumber) {
    const newData = Object.assign({}, state, data, tickerNumber);
    // if (!checkState(state, newData)) {
    state = Object.assign({}, state, newData);
    print(`BTC/USD: ${newData.usd.toFixed(2)} ` +
        `EUR/USD: ${newData['eur/usd'].toFixed(2)} ` +
        `BTC/EUR: ${newData.eur.toFixed(2)} ` +
        `Active sources: BTC/USD (${newData.bitcoinTickersNumber} of ` +
        `${BT_NUMBER}) EUR/USD (${newData.exchangeTickersNumber} of ` +
        `${ET_NUMBER}) Time: ${(new Date()).toLocaleString()}`);
    // }
}

/**
 * callback on 'data' event from bitcoin ticker
 * @param {Array} BitcoinTickers
 */
function onDataBitcoin(BitcoinTickers) {
    const data = getMinimumBitcoins(BitcoinTickers);
    const tickerNumber = {
        bitcoinTickersNumber: BitcoinTickers.filter(t => t.active).length,
    };
    mergeData(data.lastData, tickerNumber);
}

/**
 * callback on 'data' event from exchange ticker
 * @param {Array} Exchange
 */
function onDataExchange(Exchange) {
    const data = getMinimumExchage(Exchange);
    const tickerNumber = {
        exchangeTickersNumber: Exchange.filter(t => t.active).length,
    };
    mergeData(data, tickerNumber);
}

/**
 * callback on 'error' event from bitcoin ticker
 * @param {Array} BitcoinTickers
 */
function onErrorBitcoin(BitcoinTickers) {
    const tickerNumber = {
        bitcoinTickersNumber: BitcoinTickers.filter(t => t.active).length,
    };
    mergeData({}, tickerNumber);
}

/**
 * callback on 'error' event from exchange ticker
 * @param Exchange
 */
function onErrorExchange(Exchange) {
    const tickerNumber = {
        exchangeTickersNumber: Exchange.filter(t => t.active).length,
    };
    mergeData({}, tickerNumber);
}

// START PROGRAM
const BitcoinTickers = CONF.bitcoinTickers.map(tickerOptions => setTickers(
    tickerOptions,
    'bitcoin_tickers',
    () => onErrorBitcoin(BitcoinTickers),
    () => onDataBitcoin(BitcoinTickers)));

const Exchange = CONF.exchangeTickers.map(tickerOptions => setTickers(
    tickerOptions,
    'currency_exchange',
    () => onErrorExchange(Exchange),
    () => onDataExchange(Exchange)));

const EE = require('events');
const request = require('request-promise-native');

/**
 * Bitcoin Ticker, emits data and error events
 * @type {AbstractTicker}
 */
module.exports = class AbstractTicker extends EE {
    /**
     * extends event emitterF
     * @param {{
     *      requestOptions: {Object},
     *      name: {string},
     *      interval: {number},
     *      expireTime: {number}
     * }} config
     */
    constructor(config) {
        super();
        this._requestInterval = config.interval || 1000;
        this._interval = null;
        this._requestOptions = config.requestOptions;
        this.name = config.name;
        this.expireTime = config.expireTime || 10000;
        this.lastSuccessRequestTime = new Date();
        this.lastData = {};
        this.active = false;
    }
    /**
     * make async request using request options
     * @returns {Promise.<void>}
     */
    async makeRequest() {
        try {
            const res = await request(this._requestOptions);
            this.lastSuccessRequestTime = new Date();
            this._parse(res);
            this.active = true;
            // console.info('data', this.lastData);
            this.emit('data', this.lastData);
        } catch (e) {
            // console.warn(e);
            this.active = false;
            this.emit('error', e);
        }
    }
    /**
     * parse response data
     * @param {*} res
     * @private
     */
    _parse() {
        throw new Error(`${this.name}.parse() not emplemented`);
    }
    /**
     * start view for changes
     */
    async start() {
        this.active = true;
        try {
            await this.makeRequest();
        } catch (e) {
            console.warn(e);
        }
        this._interval = setInterval(
            () => this.makeRequest(),
            this._requestInterval);
    }
    /**
     * stop ticker
     */
    stop() {
        this.active = false;
        clearInterval(this._interval);
        this.removeAllListeners('data');
        this.removeAllListeners('error');
    }
};

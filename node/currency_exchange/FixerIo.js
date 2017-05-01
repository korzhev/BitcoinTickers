const AT = require('../AbstractTicker');

module.exports = class FixerIoExchange extends AT {
    _parse(res) {
        this.lastData = res.rates.USD;
    }
};

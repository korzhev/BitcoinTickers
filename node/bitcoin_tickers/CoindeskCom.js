const AT = require('../AbstractTicker');

module.exports = class CoindeskComTicker extends AT {
    _parse(res) {
        this.lastData = {
            usd: res.bpi.USD.rate_float,
            eur: res.bpi.EUR.rate_float,
        };
    }
};

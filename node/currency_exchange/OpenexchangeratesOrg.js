const AT = require('../AbstractTicker');

module.exports = class OpenexchangeratesOrgExchange extends AT {
    _parse(res) {
        this.lastData = 1 / res.rates.EUR;
    }
};

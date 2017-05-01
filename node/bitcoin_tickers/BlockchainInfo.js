const AT = require('../AbstractTicker');

module.exports = class BlockchainInfoTicker extends AT {
    _parse(res) {
        this.lastData = {
            usd: res.USD.sell,
            eur: res.EUR.sell,
        };
    }
};

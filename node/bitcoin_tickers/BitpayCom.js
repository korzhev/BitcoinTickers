const AT = require('../AbstractTicker');

module.exports = class BitpayComTicker extends AT {
    _parse(res) {
        this.lastData = {};
        const length = res.length;
        for (let i = 0; i < length; i++) {
            const code = res[i].code;
            if (code === 'EUR' || code === 'USD') {
                this.lastData[code.toLowerCase()] = res[i].rate;
            }
        }
    }
};

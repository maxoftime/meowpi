const axios = require('axios');

module.exports = async function() {
    try {
        let cats = await axios.get('http://localhost:3333/api/v1/all');
        return cats.data.cats;
    } catch(e) {
        return e;
    }
};
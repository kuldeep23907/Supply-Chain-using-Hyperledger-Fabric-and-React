const network = require('../fabric/network.js');
const apiResponse = require('../utils/apiResponse.js');

exports.createProduct = async information => {
    const { name, manufacturerId, price} = information;

    const networkObj = await network.connect(true, false, false, id);
    const contractRes = await network.invoke(networkObj, 'createProduct', name, manufacturerId, price);

    const error = networkObj.error || contractRes.error;
    if (error) {
        const status = networkObj.status || contractRes.status;
        return apiResponse.createModelRes(status, error);
    }

    return apiResponse.createModelRes(200, 'Success', contractRes);
};
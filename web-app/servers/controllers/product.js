const entityModel = require('../models/product.js');
const apiResponse = require('../utils/apiResponse.js');

exports.createProduct = async (req, res) => {
    const { name, manufacturerId, price , userType } = req.body;
    console.log('1');

    if (!name || !manufacturerId || !price || !usertype) {
        return apiResponse.badRequest(res);
    }
    console.log('2');

    if (!userType == 'manufacturer' ) {
        return apiResponse.badRequest(res);
    }
    console.log('3');

    const modelRes = await productModel.createProduct({ name, manufacturerId, price });
    return apiResponse.send(res, modelRes);
};